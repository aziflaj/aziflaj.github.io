---
title: "Runbook Notes 004: Cloud-Native Graceful Shutdowns"
pubDatetime: 2025-09-04
description: "Runbook finally supports graceful shutdowns. In this update: fixing the lost context cancellation from Note 002, using Go’s errgroup to propagate errors, and a neat state-management pattern with first-class functions. No more zombie pods when workflows fail or get cancelled."
slug: notes-cloud-native-graceful-shutdowns
tags: [
"runbook-notes",
"workflow-engine",
"golang",
"kubernetes",
"cloud-native",
"concurrency",
"graceful-shutdown"
]
---

> #### Previous notes:
> - _001_ - [Designing my Workflow Engine](/posts/notes-designing-workflow-engine)
> - _002_ - [Kubernetized](/posts/notes-kubernetized)
> - _003_ - [New UI, Deadlocks, and AI Overengineering](/posts/notes-ui-deadlocks-overengineering)

After taking some time off for strolls in Andalusia and dips in Costa del Sol, I'm back to staring at colorful text in a terminal emulator. I found the strength to get back at Runbook, and address an issue I introduced two notes ago. If you can recall, one of the losses in the [Gains & Losses](/posts/notes-kubernetized/#gains--losses) section of Note 002, was _"[...losing] the support for Go contexts: a context getting cancelled doesn’t translate into the pod shutting down. At least not auto**magically**."_

This wasn't a big deal for a while, since I had more important stuff to worry about, but while performing some tests that were _expected_ to fail, being able to stop workflows mid-run was becoming sort of a need. So, after more than a month, I finally addressed graceful shutdowns head-on. Here is a sample run of a workflow that gets triggered by an HTTP event, runs for a while, then it has a sudden stop due to one of its step failing (as expected).

![](/assets/images/20250905/cancelflow.gif)

It works in a similar way when a user manually cancells a running workflow, but the video can't show it because I haven't added the button in the UI yet. I also can't show you how I'm cancelling workflows, due to the _"My Code My Rules"_ Policy. What I can show you here, is some niceties in Go.

***

## Error Groups

An error group (`errgroup`) is basically a context-aware `sync.WaitGroup`. Like a `WaitGroup`, it lets you spawn multiple goroutines and wait for them to finish. The difference is that `errgroup` adds two key features:

- Error propagation – if any goroutine returns an error, the entire group fails.
- Context cancellation – once one goroutine fails, the context passed to the group is cancelled, signaling the rest to stop.

That means you can spin up a bunch of concurrent jobs, and if one of them dies, the rest know they should clean up and exit instead of running blindly or ending up with zombie jobs.

Here’s a simplified example of how I use it:

```go
func SpawnJobs(ctx context.Context, workflow *Workflow) error {
	g, gCtx := errgroup.WithContext(ctx)
	for _, step := range workflow.Steps {
		g.Go(func() error {
			job, err := SpawnJobForStep(gCtx, step)
			if err != nil {
				return fmt.Errorf("something went wrong: %w", err)
			}

      // block until done
			return WaitForStepToFinish(gCtx, job)
		})
	}

	return g.Wait()
}
```

If something goes wrong while `SpawnJobForStep()` or `WaitForStepToFinish()`, the `gCtx.Done()` channel will get closed, so somewhere inside those functions is some code like:

```go
select {
  case <- gCtx.Done():
    // cleanup
  default:
    // do stuff
}
```

This pattern is very useful in my orchestrator: workflows can have many steps, and error groups make sure a failure in one propagates to all.

***

## First-class functions

You can pass functions as params to other functions in JavaScript, since JS is a Functional Programming language. Well, you can do the same in Go. You can also do it in C, but let's not start calling C a FP language now.

> You can definitely do FP in C, but y'all ain't ready for that conversation

I found I had the need to do some arbitrary changes to the global run state: the Redis I've been trying to avoid since the beginning, the culprit of the [self-deadlock](/posts/notes-ui-deadlocks-overengineering/#the-self-deadlock) from the previous note. I could put all these changes in functions like this:

```go
func HandleUseCaseA() {
	currentState, err := GetState(key)
	if err != nil {
		return err
	}

  globalState.m.Lock()
  defer globalState.m.Unlock()

  // Handle A
}

func HandleUseCaseB() {
	currentState, err := GetState(key)
	if err != nil {
		return err
	}

  globalState.m.Lock()
  defer globalState.m.Unlock()

  // Handle B
}
```

And risk reaching `HandleUseCaseZZ`... But I decided to better "expose" the lock in this way:

```go
func WithLock(key string, fn func(*RunState) error) error {
	currentState, err := GetState(key)
	if err != nil {
		return err
	}

  globalState.m.Lock()
  defer globalState.m.Unlock()

	return fn(currentState)
}
```

So, let the caller define how the use case is handled, and the state exposes the lock/unlock logic in a secure way.

In a similar fashion, while this piece of code might not seem related to the workflow cancellation, it is very crucial when it comes to state management. And state management helps me spawn and shut down jobs correctly, so there's more here than it meets the eye.

***

Runbook can now shut down gracefully. No more orphaned pods left running when workflows fail or get cancelled. No more wasted resources. Runbook is more reliable than ever.
