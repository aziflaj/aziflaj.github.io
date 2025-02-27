---

title: "LeMuR: a way to \"ship your machine to customers\""
pubDatetime: 2024-06-20
image:
  path: https://github.com/aziflaj/aziflaj.github.io/blob/main/assets/images/20240620/lemur.png?raw=true
  alt: "LeMuR"
description: When someone says "It works on my machine", the consensus achieved by the Internet Hive Mind is to reply with "we won't ship your machine to the client"... Ugh, computer nerds think they have a sense of humor. I am here to tell you that yes, we can ship your machine to the client. We've had that technology for years! And no, I'm not talking about containers, this tech is even older!

slug: lemur-ship-your-machine-to-customers
---

When someone says _"It works on my machine"_, the consensus achieved by the
Internet Hive Mind is to reply with _"we won't ship your machine to the
client"_... Ugh, computer nerds think they have a sense of humor.

I am here to tell you that yes, we can ship your machine to the client. We've
had that technology for years! And no, I'm not talking about containers, this
tech is even older!

![](/assets/images/20240620/lemur.png)

## The (not-so) good ol' times

Before the Cloud and DevOps as we know it, we the back-end generalists had to
ask for SSH access to a (remote or local) server and provision it ourselves. I
remember the procedure being as follows:

1. Install Nginx 👨🏻‍💻
2. Uninstall Apache HTTP Server, we're modern like that 😎
3. Install and configure MySQL 👨🏻‍💻
4. Translate Apache's `.htaccess` into Nginx `conf` files 🫣
5. Use FTP to copy files to the server 🚀
6. Cross your fingers 🤞 and hope you don't see an error page

It worked well, despite the fact that our local machines were part of the CI/CD
system, minus the continuous part; everything was manual. And there was always
the risk of missing some native dependencies when you deployed, because the
server and your local machine were running different operating systems. I remember
running Fedora on my machine and CentOS on my servers, simply to face less "off
by one dependency" errors.

Then everything changed with Docker!

## No, Docker didn't change *everything*

To be honest, Docker -- _and containerization in general_ -- does address the
native dependencies issue. But it wasn't our first solution. A friend and
colleague of ~8 years reminded me about this fancy little tool called
[Vagrant](https://www.vagrantup.com/), and boy did I love it back then. I could
now set up Virtual Machines identical to the production servers so not only I,
but everyone on the team could avoid them pesky native dependency errors. Docker
solved the same problem differently, by using LXC instead of full fledged
Virtual Machines. They eventually won by being lighter and bundling less
dependencies inside the container but to be frank, they weren't a
_"revolutionary solution to an unsolvable problem"_ per se.

Good thing about Docker containers is that you can build Linux containers that
will run on ARM by using a Windows machine running on Intel CPU. Cross OS, cross
CPU, cross platform in general. But boy are them builds Slow as a Snail!

> _They just dropped a new SaaS acronym: Slow as a Snail. Now every software you
> develop is SaaS.._

Some time ago, I had to build Docker images to be used cross platform, and due to
financial and technological limitations we had to use Intel-based CI runners to build
images that would run on x64 and arm64 machines. The x64 builds took a
fair 5 minutes to build, which is acceptable. Arm64 on the other hand... Our CI
runners have a timeout of 60 minutes, and I never saw it finish with a usable
artifact. It always timed out because QEMU emulation is SaaSnail.
Eventually, we decided to just build x64 images
and thank Apple for releasing Rosetta2, so we could use those x64 images on M1
Macs.

But we have a solution to this whole problem, a solution that involves neither
containerization, nor virtualization. No, we're better than that. This solution
involves ***raw skill*** and ***extreme abuse*** of existing tech.

## Introducing LeMuR

A half-assed acronym that got inspired by "**L**ocal, **M**eet **R**emote", and
a solution we came up with in a mere 10 minutes during an ad-hoc call.
The solution could technically work for everyone, but not everyone will be able to pull it off.

LeMuR is perfect for small teams and founders who don't want to spend time on
provisioning overly-complicated infrastructure. It builds upon battle tested techniques
like *Trunk-Based Development*, *CI/CD* and *GitOps*, it supports Feature Branch
Deployments and Staging Environments out of the box and really, as long as you
know what you're doing, you're only limited by your imagination.

It's not a new tool, it's a new
development procedure that literally allows you to ship your development machine
to the customer, by... (drumroll)... ***developing in the production
environment!*** That's right folks, LeMuR is like that "bugfixing in production"
meme but safer™️.

![](https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExdnN1aGl5cXViMTEyYTRkcTB6a2s3dDBvanMwdWRzbHAwMDgwbTJ5cyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/XjlNyeZp5lDri/giphy.webp)


**So how to do this LeMuR thing? Glad I asked.**

Get a VPS from your preferred provider and install Nginx in it. Put your
application code in a folder inside the VPS, e.g. in `/path/to/prod/app`. This
will be your production code. To run it, I'd suggest some tool that can do Hot
Module Reloading when the files change. Rails does this out of the box, if
you're using Node.js you can use pm2 with the `--watch` option, or just be a
_mensch_ and use `inotifywait` like a madman, [here's a
hint](https://stackoverflow.com/questions/8699293/how-to-monitor-a-complete-directory-tree-for-changes-in-linux).

> _If I mention that you could add a `post-update` git hook to restart the app
> server, would it spoil the rest of LeMuR setup for you? No worries if you're
> clueless about it, we'll get back to this._

Once you've done that, start your server and bind it to some port, I hear
`:3000` is a good option. And to finally publish your app for the world, add a
proxy pass to your Nginx config:

```nginx
location / {
  # ...
  proxy_pass    http://127.0.0.1:3000/;
  # ...
}
```

That's all that's needed to publish your Production app, so let's jump into the Dev
environment, where we will blaspheme and insult all the known developer deities.
In the same VPS, add an exact copy of your app to `/path/to/dev/app`; this will
be your development environment. Add a new "remote" to your dev Git config:

```bash
$ git remote add lemurprod /path/to/prod/app
```

> _You see where this is going, right? Does that `post-update` hook make sense now?_

We're gonna need some Git hooks ([more on that
here](https://www.atlassian.com/git/tutorials/git-hooks)) to support all those
CI/CD shenanigans, so technically you'll need to:

- Run any linters in the `pre-commit` hook
- Run tests in the `pre-push` hook, stopping the push if the tests fail.

Also, you will need a development server, so let's add this to the Nginx config:


```nginx
location /lemur-devsrv/ {
  # ...
  proxy_pass    http://127.0.0.1:9000/;
  # ...
}
```

This assumes you're running the dev server on port `:9000` because it has to be
a different port from the prod server. And whatever you do, [don't use port
6000](https://aziflaj.github.io/posts/2023-04-15-curious-case-port-6000/).

Finally, to address editing code. Either be a madman and edit files in the
`path/to/dev/app` folder using Vim with a decent configuration (or Neovim, I
don't judge),
or use [VSCode and Remote development over
SSH](https://code.visualstudio.com/docs/remote/ssh-tutorial) or something
similar for whatever editor/IDE you use. And how does the development flow
works?

1. You branch off `main`, Trunk Based Development with style
2. You edit files in the `/path/to/dev/app` folder, you can check changes from
anywhere in the world since you're exposing the dev app in
`https://{vps-ip}/lemur-devsrv/`
3. You commit your changes in your branch, and it either passes with no linter
errors, or it fails with linter errors which you will have to address (hashtag
bestpractices)
4. You eventually are happy with your branch, so you merge it in your dev app
   `main` branch
5. You do `git push lemurprod main` and it either pushes after all the tests
pass (hashtag continuousintegration) or it fails and you have to fix your broken
test suite
6. Your production server in `/path/to/prod/app` runs its `post-update` hook to
restart the server, or your `inotifywait` triggers a server restart script, and
your app gets updated (hashtag continuousdeployment)


## You promised Feature Branch Deployments and Staging Environments

Yes, I did. I don't recommend them, you should not have long-lived branches and
you should use Feature Flags instead. Anyway, here's how you can add a Staging
Environment to your LeMuR setup:

1. Create a `/path/to/stag/app` to be your staging environment
2. Add a new remote to your Git config: `git remote add lemurstag /path/to/stag/app`
3. Add a new Nginx config for the staging server:
    ```nginx
    location /lemur-stagsrv/ {
      # ...
      proxy_pass    http://127.0.0.1:3001/;
      # ...
    }
    ```
4. Run your staging app on port `:3001`, use either `inotifywait` or the git
   hook to restart the server when you push to the staging branch
5. ???
6. Profit!


And if Staging doesn't cut it for you and you need Feature Branch Deployments
because you're a masochist, save this bash script as a `deplfb` and use it to deploy feature branches at will:

```bash
#!/bin/bash

branch=$1
port=$2

# Which branch bruh?! And which port?!
if [ -z "$branch" ]; then
  echo "Usage: deplfb <branch-name> <port>"
  exit 1
fi

# add a new worktree for the branch
mkdir -p /path/to/feat
git worktree add /path/to/feat/$branch $branch

# Add new Nginx config for the feature branch
cat <<EOF > /etc/nginx/sites-available/your-app
location /lemur-$branch/ {
  # ...
  proxy_pass    http://127.0.0.1:$port/;
  # ...
}
EOF

# Create a new server for the feature branch
cd /path/to/feat/$branch
# TODO: run the dev server with code reloading on $port
# TODO: restart nginx to apply the new config
```

Just address the `TODO`s and you're good to go. I'm skipping the environment
variables, but for those you can either use `.env` files or something like
[direnv](https://direnv.net/).

## Conclusion

Next time your peers feel like "It works on my machine" is not a valid excuse,
tell them about LeMuR. Educate them on what a sprinkle of bash scripting and
bunch of Git hooks can do. Show them the power of developing in production.
Tell them "Yes, we _can_ ship my machine to the client" and watch
their faces as they realize you're not joking.

Embrace the chaos. Live on the edge. Develop in production. Ship your machine!
