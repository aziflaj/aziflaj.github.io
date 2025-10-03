---
title: "Runbook Notes 006: Github Integration"
pubDatetime: 2025-10-03
description: "This update covers Runbook’s GitHub integration — from OAuth logins to pulling repos and running workflows. After a week of battling JWT tokens, it finally works!"
slug: notes-github-integration
tags: [
"runbook-notes",
"workflow-engine",
"github",
"gitlab",
"integration"
]
---

> #### Previous notes:
> - _001_ - [Designing my Workflow Engine](/posts/notes-designing-workflow-engine)
> - _002_ - [Kubernetized](/posts/notes-kubernetized)
> - _003_ - [New UI, Deadlocks, and AI Overengineering](/posts/notes-ui-deadlocks-overengineering)
> - _004_ - [Cloud-Native Graceful Shutdowns](/posts/notes-cloud-native-graceful-shutdowns)
> - _005_ - [Quotas and Limits](/posts/notes-quotas-and-limits)

After a week, I can proudly announce I can run Runbook workflows from Github!
I can't find the words to express how happy I am to have implemented this feature, especially
after a whole week of consistent failures of debugging attempts on Github's JWT tokens...
so here's a full blown blogpost to show what and how and why.

***

## Social OAuth

The first time I was introduced to GitHub, I was told "It's a social platform for programmers" as if it's a Facebook for coders... But I do have some followers there, and I have followed some people over the years.

After a few weeks of trying to remember my password for Runbook, I decided to add OAuth to it. I'm currently supporting
GitHub and GitLab login, along with the good ol' email/username and password, but that's something you can see from the video.

I first started with [markbates/goth](https://github.com/markbates/goth/), but then I saw how much it depends on session, and my app was using an auth token/refresh token approach, and I was too lazy to read through the documentation on how to make my approach goth-compatible, so I went **I'll just implement OAuth myself, with blackjack, and [explicit]!**. So I just rolled out my own OAuth handling logic based on [x/oauth2](https://cs.opensource.google/go/x/oauth2). I even added a contract to follow for future integrations:

```go
type ForgeService interface {
	AuthCodeURL(state string) string
	GetUserInfo(ctx context.Context, code string) (*models.OAuthUser, error)
	UpsertUser(ctx context.Context, oauthUser *models.OAuthUser) (*models.User, error)
}
```

And it works well for both GitHub and GitLab! In order to add more OAuth Providers in the future, all I need to do now is just implement these 3 functions and everything works like a charm.

***

## Integration is not OAuth

When signing in with GitHub, you get back a couple of user tokens to use when you want to act on behalf of the signed in user. These are what I'm using to get stuff like the user's emails and other data. But when it comes to actually using the user's repos, that's a totally different story.

See, signing in with GitHub is an apple, and cloning a repo from the user's account is not even an orange; it's an orangutan that needs taming. Since Runbook is registered as a Github Application which you can install on your Individual or Organization's settings, it needs a different path of authorization and execution.

For every user or organization that installs your application, GitHub gives you an installation ID. This ID can be traded for an installation token, which allows your app to act on behalf of the user or organization.

I started my journey by using the Google endorsed [bradleyfalzon/ghinstallation](https://github.com/bradleyfalzon/ghinstallation) package, but whatever I did I always got the same error message from GitHub:

```
Failed to fetch projects
Failed to get installation <installation-id> details (may be revoked or suspended)
GET https://api.github.com/app/installations/<installation-id>
401 A JSON web token could not be decoded []
```

**Why couldn't the damn JSon web token be decoded?!** I have the damn `.pem` file in place. I have the right App ID, trust me. I created and recreated the GitHub App twice, I replaced all the values properly. I quadruple checked the installation ID for my user and I hardcoded it in all the right places. _THIS THING SHOULD WORK!_

Everything pointed to the JWT tokens not being signed correctly. After spending 2-3 afternoons testing and failing repeatedly, I gave up. I thought **I will sign the tokens myself, with blackjack, and [explicit]!**

So I rolled out a simple JWT signer, as per the GitHub documentation, based on [golang-jwt/jwt](https://github.com/golang-jwt/jwt/). And it worked... but why?! Why would my approach work and the ghinstallation wouldn't? According to the AI shenanigans, chats and whatnot, it was a conflict between the jwt/v4 that ghinstallation uses internally, and the jwt/v5 that I was using. After that many afternoons of undecodable JWT tokens, I coulnd't even care for it anymore. If it works, it works. Ship it first, then ask questions later.

I did some extra testing, it seemed like my JWT signing was working properly as per the GitHub responses; I committed, and called it a day.

***

### Pulling Repos from GitHub

After I got the Installation Tokens working properly, the next logical step would be to

1. Pull Runbook Definitions from the GitHub repo
2. Schedule the workflows for execution
3. (hidden step) Clone the repo because the workflow depends on it

I almost forgot about that 3rd step, but we'll get back at it. Luckily, the first one was simple: the Go GitHub SDK already provides me with a [GetContents](https://pkg.go.dev/github.com/google/go-github/v75/github#RepositoriesService.GetContents) function to pull the Runbook definition, a single `.runbook.yml` in the root of the project. That's all I need to know in order to schedule a workflow for execution...

_"But wait," - you say, - "I want to run some npm install and some npm run tests on my workflow!"_ Pulling a single file is not enough, I needed that 3rd hidden step.

So I created another executable in the same [BORE](https://aziflaj.github.io/posts/notes-kubernetized/#bore-build-once-run-everywhere) repo I have for Runbook, called `repo_cloner`. What it does, is that it uses the same JWT signer I was blabbering about in the previous section to clone the repository on behalf of the user who set up the Runbook installation. The recording you saw at the beginning of this blog was triggered by a barebones repository with a single `echo` bash script in it, and this `.runbook.yml` file:

```yml
name: "Default Runbook Workflow"

steps:
  - name: "Greeting"
    env:
      TEXT: "Hello, World! I am"
    command: |
      USER=$(whoami)
      echo "${TEXT} ${USER}!"

      repo_cloner
      ls -la
      cd runbook-activated
      chmod +x ./run.sh
      ./run.sh
```

Seeing Runbook be able to run workflows from GitHub was an unimaginable relief after a whole week of battling unsigned tokens. I think I will take the weekend off, go celebrate with some beers, and worry about GitLab on Monday. Things are coming together really nicely.
