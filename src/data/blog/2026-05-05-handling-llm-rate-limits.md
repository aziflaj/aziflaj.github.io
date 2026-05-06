---
title: "Handling LLM Rate Limits in Go: Circuit Breakers and Durable Queues"
pubDatetime: 2026-05-05
description: "How to build resilient LLM integrations in Go using circuit breakers and durable queues, avoiding retry storms and handling rate limits with backpressure."
slug: handling-llm-rate-limits
tags: [
    "go",
    "llm",
    "rate-limit",
    "circuit-breaker",
    "durable-queue",
]
---

We are currently going through an era of products needing to provide prompt-based interfaces while not having enough resources to self-host LLMs. And since talking to a computer in plain English is way easier _(from a user's perspective)_ than clicking around (YMMV), as product builders and engineers, we will face an issue more often than not:

Company pays for an amount of tokens, which happen to not be unlimited. Product users then consumes all those tokens. All the consecutive requests fail with a _"you're too poor to buy more tokens"_ error. Users start complaining.

There's no point in sending another request (or another token) to that LLM until you either get richer, or until your quota gets reset and you can start prompting again.

History repeats itself, and while LLMs are great and intelligent and all that, for all practical purposes they _can_, and I'd advocate **should**,  be treated as any other dependency over the network:

-   LLMs can be slow, and your requests can timeout
-   LLM responses can fail, not necessarily due to your input
-   The LLM can rate limit and/or throttle you
-   The LLM can go offline for whatever reason, and come back online whenever it feels like

Our goal here, as builders, is to build resilient systems that stay predictable under these circumstances.

## Retries Aren't Silver Bullets

A few days ago, I posted on LinkedIn a [story](https://www.linkedin.com/feed/update/urn:li:activity:7454816827496460288/) of how being rate limited by an API role-playing as an LLM resulted in high error rates.

The naive, perhaps tempting solution might be to retry. But retries, even with backoff, can make things worse when the upstream is rate limiting you. A retry is not always recovery. Sometimes, it is unnecessary pressure.

If every failed request immediately becomes another request, the system does not recover. It amplifies the problem. What it needs instead is **backpressure**.

Luckily, for the use case we had, the service depending on the LLM didn't have a time constraint. Assume that air resistance is negligible and that customers don't care how long the response takes to come back, as long as the result is useful.

To work around these limitations, the solution we can go for is a combination of two patterns, which probably the title gave away too soon:

- **Circuit Breaker**, to stop spamming the LLM when it is clearly pushing back
- **Durable Queue**, to make sure work is not lost and can be retried later

The circuit breaker protects the upstream. The queue protects our work. Together, they give us controlled retries instead of chaos.

## Implementing a Circuit Breaker

A circuit breaker answers a single question: _Should we call the 3rd party ~~API~~ LLM right now?_

As per some paraphrased definition, a service client should invoke a remote service via a **proxy** that functions in a similar fashion to an electrical circuit breaker. When the number of consecutive failures crosses a threshold, the breaker trips; the circuit opens and for the duration of a timeout period all attempts to invoke the remote service will fail immediately. After the timeout expires, the breaker allows a limited number of test requests to pass through; the circuit is now half-open. If those requests succeed, the breaker resumes normal operation and the circuit closes. Otherwise if there is a failure, the timeout period begins again.

Here's a simple implementation of this whole logic in Go:

```go file=circuit_breaker.go
package breaker

var ErrCircuitOpen = errors.New("circuit breaker is open")

type CircuitState string

const (
  StateClosed   CircuitState = "closed"
  StateOpen     CircuitState = "open"
  StateHalfOpen CircuitState = "half_open"
)

type CircuitBreaker struct {
  mu sync.Mutex

  state CircuitState

  failures    int
  maxFailures int

  openedAt time.Time
  cooldown time.Duration

  shouldTrip func(error) bool
}

func NewCircuitBreaker(maxFailures int, cooldown time.Duration, shouldTrip func(error) bool) *CircuitBreaker {
  return &CircuitBreaker{
    state:       StateClosed,
    maxFailures: maxFailures,
    cooldown:    cooldown,
    shouldTrip:  shouldTrip,
  }
}

func (b *CircuitBreaker) Do(ctx context.Context, fn func(context.Context) error) error {
  if err := b.beforeCall(); err != nil {
    return err
  }

  err := fn(ctx)
  if err == nil {
    b.recordSuccess()
    return nil
  }

  if b.shouldTrip == nil || b.shouldTrip(err) {
    b.recordFailure()
  }

  return err
}

func (b *CircuitBreaker) beforeCall() error {
  b.mu.Lock()
  defer b.mu.Unlock()

  if b.state == StateOpen {
    if time.Since(b.openedAt) >= b.cooldown {
      b.state = StateHalfOpen
      return nil
    }

    return ErrCircuitOpen
  }

  return nil
}

func (b *CircuitBreaker) recordSuccess() {
  b.mu.Lock()
  defer b.mu.Unlock()

  b.failures = 0
  b.state = StateClosed
}

func (b *CircuitBreaker) recordFailure() {
  b.mu.Lock()
  defer b.mu.Unlock()

  b.failures++

  if b.state == StateHalfOpen || b.failures >= b.maxFailures {
    b.state = StateOpen
    b.openedAt = time.Now()
  }
}
```

The breaker does not retry anything. We simply ask it to `Do` something for us, and it decides whether we are allowed to hit the upstream dependency or not.
Also, this breaker is in-memory and while it can be used by multiple workers (see the mutex), it does not survive restarts. If that's what you need (and most likely you do), look into persisting breaker's state and using distributed locking instead of mutexes.

## Be lazy, retry later

Following the principle of _"not doing today what can be put off till tomorrow"_, instead of tying the payload's lifecycle to that of an HTTP request, we persist work to retry when it benefits us:

```go file=queue.go
type Job struct {
  ID      string
  Payload Payload
  Attempt int
}

type Payload struct {
  Prompt string
}

type Queue interface {
  Enqueue(ctx context.Context, payload Payload) error
  ClaimNext(ctx context.Context) (*Job, error)
  Complete(ctx context.Context, jobID string) error
  RetryLater(ctx context.Context, jobID string, next time.Time) error
}
```

This `Queue` interface doesn't tie to a specific persistence system. You can implement it to write to a database table, or a message queue, or a stream you can read from at a later time. The implementation doesn't matter. Only the guarantee does: the payload survives failure, and can be retried at a later time.

## Wiring it all together

We use a Worker instance to combine the two patterns. The worker is responsible for processing prompts, and it uses the `CircuitBreaker` to decide whether to call the LLM or not. If the breaker is open, or if the call fails, it enqueues the job to be retried later.

```go
type LLMClient interface {
  ProcessPrompt(ctx context.Context, text string) (LLMResponse, error)
}

type Worker struct {
  queue   Queue
  breaker *CircuitBreaker
  llm     LLMClient
}
```

A worker calls the `LLMClient` when the `CircuitBreaker` is closed, and in case of failure it enqueues the `Job` (along with the `Payload`) to be retried at a future time. We can build a breaker that opens on specific errors, like:

```go
breaker := NewCircuitBreaker(
  5,                 // Max failures threshold
  15 * time.Minute,  // Cooldown period
  func(err error) bool { // trip the breaker for HTTP 429 and 5xx
    var httpErr *HTTPError
    if errors.As(err, &httpErr) {
      return httpErr.StatusCode == http.StatusTooManyRequests ||  // HTTP 429
             httpErr.StatusCode >= http.StatusInternalServerError // HTTP 5xx
    }

    return false
  },
)
```

With the modules we now have, we can implement a `ProcessJob` behavior that enables the worker to process a prompt or enqueue for later:

```go
func (w *Worker) ProcessJob(ctx context.Context) error {
  job, err := w.queue.ClaimNext(ctx)
  if err != nil {
    return err
  }
  if job == nil { // feels weird but I'll allow it
    return nil
  }

  var llmRes LLMResponse
  err = w.breaker.Do(ctx, func(ctx context.Context) error {
    res, err := w.llm.ProcessPrompt(ctx, job.Payload.Prompt)
    if err != nil {
      return err
    }

    llmRes = res
    return nil
  })

  if err != nil {
    return w.retryLater(ctx, job)
  }

  w.queue.Complete(ctx, job.ID)

  // TODO: process LLM Response

  return nil
}

func (w *Worker) retryLater(ctx context.Context, job *Job) error {
  next := time.Now().Add(backoff(job.Attempt)) // implementation left to the reader
  return w.queue.RetryLater(ctx, job.ID, next)
}
```

## Takeaways

If you're not a Go developer, the main takeaway here is to not think of LLMs as special citizens from a systems perspective. They are unreliable network dependencies. If you treat them like magic, your system will break. If you treat them like any other dependency in any distributed system, you can use all the tools you already know to work, and your system will survive.

The prompt is only half of the solution. The execution model is just as important.
