---
title: "Runbook Notes 002: Kubernetized"
pubDatetime: 2025-08-01
description: "In this note, I dig into what changed when Runbook went Cloud-Native™: YAML parsers betrayed me, Go contexts stopped working, logs got messy, and yet... orchestration got real. Also, I'm pattenting the BORE acronym, and I still don’t have a database."
slug: notes-kubernetized
tags: [
  "kubernetes",
  "workflow-engine",
  "cloud-native",
  "go",
  "build-in-public"
]
---

> #### Previous notes:
> - _001_ - [Designing my Workflow Engine](/posts/notes-designing-workflow-engine)

After merging last week’s Pull Requests, written half-drunk on my patented Ironwort tea cocktail, I'm proud to announce:

Runbook is now a **Cloud-Native**, **Event-Driven**, and **Infrastructure-Agnostic** orchestrator for **Ephemeral**, **Distributed Workloads**.

🧠 Built for multi-tenancy, fault-tolerance, and high-concurrency execution. _IT SCALES!_ </br>
🛡️ Security by design: containerized, sandboxed, demure </br>
🗒️ Designed around declarative YAML workflows </br>
🚀 Zero lock-in. Born to be ~~wild~~ run anywhere

_It’s not just orchestration. It’s orchestration done right™️_

Of course, there's some jokes there, and a lot of flashy words that mean nothing to anyone with a gram of working brain. But they're definitely true. I did achieve all those, through meticulously planned and purposefully deferred decisions I mentioned in my previous note. And I learned a few things along the way.


<details>
<summary>Anyway here's a sample workflow</summary>

```
          ╭───╮
╭───╮     │ C │
│ A │ ╭───┼● ●┼╮
│  ●┼─╯   ╰───╯│  ╭───╮
│  ●┼─────╮    │  │ D │
╰───╯     │    ╰──┼● ●┼╮
╭───╮     ╰───────┼●  ││╭───╮
│ B │             ╰───╯││ E │
│  ●┼────────╮         ╰┼●  │
╰───╯        ╰──────────┼●  │
                        ╰───╯
```
</details>

<details>
<summary>And here's a sample run of that workflow in K8s</summary>

![](/assets/images/20250801/runbook.gif)

</details>



***

## Kubernetes broke my Orchestration

I added a YAML parser to Runbook since day one. When I started working on K8s, I noticed the K8s Go client also has a YAML parser included. So I forgot the good ol' _"Don't fix what's not broken"_ mantra and decided to use this new parser instead of the old one. Because why have 2 thing when 1 thing do trick.

For reference, here's how a very basic Runbook YAML definition looks like:

```yml
name: "Default Runbook Workflow"

steps:
  - name: "Step A"
    env:
      SOME_ENVVAR: "value"
    command: printenv

  - name: "B"
    command: |
      echo "Look busy..."
      sleep 50

  - name: "C"
    depends_on:
      - "Step A"
      - B
    command: echo "Step with dependencies"
```

And here's the code before and after I switched from my previous YAML parser to the K8s one:

```go
// Old Code
import "github.com/goccy/go-yaml"
// ...
var runbook types.Runbook
if err := yaml.Unmarshal(ymlFileBytes, &runbook); err != nil {
  return fmt.Errorf("%w: %v", types.ErrUnmarshalRunbook, err)
}

// New Code
import "k8s.io/apimachinery/pkg/util/yaml"
// ...
var runbook types.Runbook
decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(ymlFileBytes), 4096)
if err := decoder.Decode(&runbook); err != nil {
  return fmt.Errorf("%w: %v", types.ErrUnmarshalRunbook, err)
}
```

The next day, while doing some routine tests, I noticed all steps were kicking off at once: the step dependencies weren't being taken into consideration. **The workflow orchestrator was broken.** Kind of a big deal, when all you have is a workflow orchestrator, right? So I started reading the code again, putting logs and breakpoints for everything.

After spending almost half an hour debugging, I decided to log the steps themselves as they were parsed before starting their execution. And lo and behold, here's what I see:

``` go
{ Name: "Step A", DependsOn: [], Command: "printenv" }
{ Name: "B", DependsOn: [], Command: "echo \"Look busy...\"\nsleep 50" }
{ Name: "C", DependsOn: [], Command: "echo \"Step with dependencies\"" }
```

No dependencies on step C?! Who'd've thunk the K8s YAML parser would mess up parsing arrays... I reverted to the original parser, ran everything again, and I saw my dependencies being loaded correctly. Great success!

Now I'm using 2 parsers: one for parsing my own Runbook definitions, the K8s one to parse K8s YAMLs, and praying the K8s parser can at least parse its own configs properly...

***

## Gains & Losses

When I wrote the previous entry of these Runbook Notes, workflows were executed as Goroutines. It's easy controlling everything when nothing leaves the machine. But now that we're cloud native, properly distributed, and loaded with buzzwords... you gain some and you lose some.

Moving to K8s Jobs I lost the support for Go contexts: when a context is cancelled, which no longer translates into the pod shutting down. At least not auto<strong><em>magically</em></strong>.

Another nice feature I lost was log streaming. With steps as goroutines, I could easily set up [`io.Pipe`](https://pkg.go.dev/io#Pipe)s and stream `$stdout` and `$stderr` independently. With K8s jobs, I can only read pod logs as a stream of text, and they don't come with `$stdout`/`$stderr` flavors. I have to build my own tooling for that sort of distinction.

But I'm fine with losing this kind of functionality. You gain some, you lose some. Manual shutdowns and log sorting are the price to pay for a cloud-native, event-driven, and infrastructure-agnostic orchestrator for ephemeral, distributed workloads. The main product right now is the Orchestrator. For everything else, there's a GitHub issue and a work plan.

***

## BORE: Build Once, Run Everywhere

Reading this blog is a bore, and that's the same method I'm using to build my too many components of Runbook. What started as a single binary has become 3 binaries in a trench coat, with probably more to come. So instead of having multiple projects and losing track of what is implemented where, I took the "monorepo" approach and pushed it one step further: "monobinary"! My 3 binaries in a trench coat are actually a single large(r) binary, with different modes of execution.

```bash
# Instead of:
$ mybin1
$ mybin2

# I do:
$ mybin --runmode mode1
$ mybin --runmode mode2
```

I only have a single build step, and I put the build artifact into different container images for different purposes. **Build Once, Run Everywhere**. Or to put it in an advertisement-friendly way:

_Dare To Be BORE-ing!_
