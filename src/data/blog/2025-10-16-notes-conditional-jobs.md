---
title: "Runbook Notes 007: Conditional Jobs & Dynamic Logic"
pubDatetime: 2025-10-16
description: "Runbook can now make runtime decisions with boolean expressions, environment variables, and manual approvals. And it's powered by a language I wrote a year ago, called PinguLang"
slug: notes-conditional-jobs
tags: [
"runbook-notes",
"workflow-engine",
]
---

> #### Previous notes:
> - _001_ - [Designing my Workflow Engine](/posts/notes-designing-workflow-engine)
> - _002_ - [Kubernetized](/posts/notes-kubernetized)
> - _003_ - [New UI, Deadlocks, and AI Overengineering](/posts/notes-ui-deadlocks-overengineering)
> - _004_ - [Cloud-Native Graceful Shutdowns](/posts/notes-cloud-native-graceful-shutdowns)
> - _005_ - [Quotas and Limits](/posts/notes-quotas-and-limits)
> - _006_ - [Github Integration](/posts/notes-github-integration)


Earlier this week I was contemplating on how to add support for **Conditional Jobs** to Runbook. I already have some sort of conditional execution; you can specify a list of steps upon the current step `depends_on`, and the current step will only get executed after the steps it depends on have finished successfully. In case one of them fails, the workflow fails as well, pretty straight forward.

But the `depends_on` is a sort of "static condition": it only sets the order of execution, and the execution can't be stopped or started based on some arbitrary checks set by the user. What I wanted with this Conditional execution was something smarter: a way for workflows to make decisions **at runtime**, not just based on whether the previous step passed or failed.

***

## A CI/CD Example
Consider a CI/CD pipeline where you don’t always want to deploy automatically. Sometimes deployment should only happen:

- After someone manually approves it, **AND**
- When it’s triggered from specific branches

For example:

```yml file=runbook.yml
name: "Deploy to Vercel"

steps:
  # ...
  - name: "Deploy"
    depends_on:
      - "Run Tests"
    if: '(branch == "main" or branch == "qa") and approved_by(@aziflaj)'
    command: pnpm exec vercel --logs --target=$ENV
```

The `"Deploy"` step always happens if the `"Run Tests"` step succeeds, but before it runs it has to check some conditions:

1. Is the triggering branch `main` or `qa`?
2. Has `@aziflaj` approved this step?

If either condition fails, the step is skipped (or stopped for approval)

***

## Evaluating booleans

I considered 3 different approaches for evaluating these expressions:

1. **Embedd a JS interpreter**, something like [robertkrimen/otto](https://github.com/robertkrimen/otto) and evaluate expressions written in a language the user (most likely) already knows
2. **Write my own boolean expression parser** using [goyacc](https://pkg.go.dev/golang.org/x/tools/cmd/goyacc) and some reading on [Backus-Naur form](https://en.wikipedia.org/wiki/Backus%E2%80%93Naur_form)
3. Remember I've already done this last year, and PinguLang can already evaluate boolean expressions, among other things

So I decided to add my old project as a dependency of my current project:

![](/assets/images/20251016/pingul-runbook.png)

PinguLang lets me inject variables and intrinsic functions in the scope of the execution, like this:

```go file=cond.go
// source -> lexer -> tokens -> parser -> AST -> eval
lxr := lexer.New(string(step.If))
prs := parser.New(lxr)
program := prs.ParseProgram()

scope := object.NewScope()
scope.Set("branch", &object.String(Value: []rune("qa"))) // the triggering branch
scope.Set(
  "approved_by",
  &object.IntrinsicFunc(Value: func(args ...object.Object) object.Object {
    // [redacted implementation]
    return &object.Boolean(Value: true)
  }),
)

result := eval.Eval(scope, program)
result.IsTruthy() // the evaluated boolean we expect
```

***

## Putting It to the Test

After wiring this up, I ran a worƒlow that has both passing and failing conditions:

![](/assets/images/20251016/condrun.png)

That `"Step 2 Cond F"` is grayed out because its condition is evaluated to false. And how it looks like in the yml is:

```yml file=runbook.yml
name: "Testing Docker & Compose"

global:
  env:
    TEXT: "Hello World"
    BOOL_TRUE: true
    TRUTHY: true

steps:
  - name: "Step 1"
    command: printenv

  - name: "Step 2 Cond T"
    depends_on:
      - "Step 1"
    if: "$BOOL_TRUE == $TRUTHY" # evaluates to true
    env:
      TEXT: "Running Step 2 Cond T"
    command: printenv

  - name: "Step 2 Cond F"
    depends_on:
      - "Step 1"
    env:
      TEXT: "Running Step 2 Cond F"
    if: "$BOOL_TRUE != $TRUTHY" # evaluates to false
    command: printenv

  - name: "Run via Docker Compose"
    depends_on:
      - "Step 1"
    env:
      BUILDKIT_PROGRESS: "plain"
    command: |
      repo_cloner

      cd runbook-activated
      ls -la

      docker compose build --quiet
      docker compose up -d redis
      docker compose run --rm app
      docker compose down
```

***

## Global Env Vars
As a side quest, I also added support for globally defined environment variables, which can be overridden per step as needed. These globals are also exposed to the `if` conditions as `$ENVVARS`, so you can reference and compare them dynamically within your expressions.

Every week, we get one step closer to the first closed beta testers for Runbook.
