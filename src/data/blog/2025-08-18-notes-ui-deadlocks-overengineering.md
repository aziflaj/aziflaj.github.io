---
title: "Runbook Notes 003: New UI, Deadlocks, and AI Overengineering"
pubDatetime: 2025-08-18
description: "Goodbye juggling multiple terminals, hello frontend! In this update: workflow visualizations, over-engineered log deduplication, and the self-deadlock that stole an afternoon from me."
slug: notes-ui-deadlocks-overengineering
tags: [
"runbook",
"workflow-engine",
"frontend",
"golang",
"debugging",
"build-in-public",
"ai-fails"
]
---

> #### Previous notes:
> - _001_ - [Designing my Workflow Engine](/posts/notes-designing-workflow-engine)
> - _002_ - [Kubernetized](/posts/notes-kubernetized)

Last time you read one of these notes, Runbook had to be demoed via a bunch of terminals running at the same time: one for orchestration logs, one for K8s pods, a third one for job logs, and a fourth one to trigger workflows and run commands.

Though all that can be used to prove the workflows run as expected, it doesn't make for a sellable product that appeals to CEOs and CTOs who prefer a clickable UI.

## Goodbye CLI, Hello UI

One month after starting to work on Runbook, we now have a passable, working prototype of a User Interface. A Front End if you will. Here's a GIF showing workflows executing in real time with logs streaming in.

![](/assets/images/20250818/frontend.gif)

> Pardon the low quality, the original screen record was a few MBs and I had to gif-ify it and compress it to be usable in this blog. You'll see a better quality soon, hopefully running YOUR workflow CI/CD pipeline

I'm not much of a front-end developer; I prefer to address myself as a Craftsman of Software Solutions (CSS for short, no relation to CSS) and when people can't figure out my interfaces, it's obviously their fault. But my lack of web development know-how doesn't translate to a lack of vision: I know how a workflow should look like.

Instead of losing time and putting too much effort in researching JS libraries, learning how Vite works (I guess Webpack is not a thing anymore), picking between TypeScript ~~compilers~~ transpilers written in Rust and whatnot, I hacked together a frontend fast enough to demo, ugly enough to make it obvious it’s still early. The result is a work of art that mostly works, crafted with the help of a language model that adds as many bugs as it fixes.

I am using [dagre](https://github.com/dagrejs/dagre) to calculate the positioning of the workflow graph nodes, and React Flow to make said nodes look useful. Each node represents a step of the workflow, with loading indicators to show progress, green and red colors to show successful/failure states, and when clicked, you are taken to a view which shows you the logs of the step you just clicked. These logs are polled in Near Real Time, because WebSockets are an overkill sometimes.

I even had to rewrite one of my Go components into a web service. Maybe one of these days I'll publish a guide into developing web services in Go, since I already have [a guide on how to deploy them](https://aziflaj.github.io/posts/deploy-go-kamal-gh-actions/), but let's see what the future holds. Either way, the side effects of fast prototyping are interesting bugs which you can read about, one scroll of your mouse below.

***

### AI-Assisted Over Engineering

I had an interesting issue with logs. The logs get pulled in batches as they get generated in the execution of each workflow step. It can happen that 2 consecutive batches include the same log line twice: e.g. batch-N pulls lines 10-15, and batch-(N+1) pulls lines 15-20, meaning the log line 15 is duplicated. Since the FE is vibe coded and I was too lazy to update how batches work, I asked my Artificially Intelligent assistant (emphasis on ass) to deduplicate these lines. And the AI assistant happily spat out the most over-engineered code imaginable, worthy of a PhD thesis on HashSets.

For context, my Log Entries are stored as (among other fields): `{ line_no: 42, message: "something happened" }`. And to deduplicate, the AI assistant decided to firstly write a function that builds a hash from a given batch of log entries, where the `line_no` acts as the hash key. Then, when a second batch is received from the backend, it runs a hash-merge function. The hash-merge logic was the over-engineered part here. It first creates a set with the keys from both hashes, since sets can't have duplicated values. Then, it converts the set into an array and sorts it (log lines should be printed in order). With the hash keys in order, it iterates through all of them and pulls the values from each hash; if the key is found in both hashes it pulls the value from the latest batch, in order to overwrite potentially old values. Sounds technically correct, right?

But, line numbers are... numbers. I don't start counting log lines from 0 but from 1, but that doesn't mean I can't use an array to store all these log lines. Instead of using a Set + Hash + whatever logic that was, I can simply use a dynamic array, use `line_no--` as the array index and store the `message` as the array value. And when a new batch of log lines is pulled, I just resize the array and add the new values where needed. No overcomplicated logic. Maybe our AI assistants aren't there yet and our jobs are safe.

Either way, this wasn't a bug, but rather a case of reasoning language models not reasoning properly. What was a bug, was this other thing that took me a full afternoon to debug, and it was me who wrote it, not AI.

***

## The Self-Deadlock

On the first of these notes, I wrote how I went through hoops to avoid adding Redis to my infra stack. I am simply using an in-memory in-process hash map to store state, and rebuild it from somewhere _(you don't need to know where from)_ in case my process gets killed and restarted. The problem is that this hash map is accessed by multiple goroutines, so it needs to be protected with locks. And here's my implementation of it (obviously changed, for obvious reasons). See if you can find the issue:

```go file=state.go
type state struct {
  m sync.Mutex
  runStates map[string]*RunState
}

var globalState state

// ...
func GetState(key string) (*RunState, error) {
  globalState.m.Lock()
  defer globalState.m.Unlock()

  //... do stuff
  return runState, nil
}

func DoStuff(key string, event Event) error {
  globalState.m.Lock()
  defer globalState.m.Unlock()

  currentState, _ := GetState(key)

  // ... update state if necessary
  return nil
}
```

If you saw it, you have better eyes than I do. It really took me half a day to figure out I have a self-deadlock here:
- DoStuff is called, it acquires a lock
- GetState is called, it tries to acquire lock; can't
- GetState blocked, waiting for DoStuff to release the lock
- DoStuff blocked, waiting for GetState to finish

There are 2 ways to work around this issue. The Java way is to use a [Reentrant mutex](https://en.wikipedia.org/wiki/Reentrant_mutex) (also known as a recursive mutex). But they can be tricky, and fortunately Go doesn't include a Reentrant mutex in the stdlib. So, the good samaritan approach is not to implement a Reentrant mutex by yourself (it's easy, try it), but to rewrite the code as this:

```go file=state.go
func DoStuff(key string, event Event) error {
  currentState, _ := GetState(key) // [!code ++]

  globalState.m.Lock()
  defer globalState.m.Unlock()

  currentState, _ := GetState(key) // [!code --]

  // ... update state if necessary
  return nil
}
```

Just move the GetState call out of the region. And everyone lives happily ever after.

***

Runbook now has a usable UI and improved responsiveness to events. For now, it runs. And now, you can see it. But there's still more to do in order to start selling it as a product. Maybe Note #004 will come with something YOU can use. But let's see.
