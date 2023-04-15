---
layout: post
title: "The Curious Case of Port 6000: A Comedy of Errors and Chromium's Shenanigans"
date: '2023-04-15'
---

Some time ago I started working on a side project, mainly to experiment with some
technologies I don't use on a daily basis. I had been working on a REST API for
a while, had it running on a container on port `3000`, as you usually do;
there's a nice ring to 3000. It works for [Andre 3000](https://open.spotify.com/artist/74V3dE1a51skRkdII8y2C6?autoplay=true),
why wouldn't it work for me? I also had it tested, both with automated testing and some manual testing via Insomnia, and when
things started looking good, I moved to the frontend.

At this point, all I know is React and jQuery.
I don't know how the cool kids are building frontend stuff now but my go-to tool has
always been React (for the past 7 years anyway), so I again went with React. For the first time decided to use a
framework, because I was too lazy to start installing way too many React libraries.
So I started with a new Dockerfile, installed this newfangled thing called [Remix](https://remix.run/), set it up to
run on port `6000` (because 3000 times 2 ... dunno, reasons), and then ran all my containers
to see whether they started on the right order.

They did. _"Good"_ I thought ... but it was all but good. The container was running, `docker ps`
showed it there occupying the right port, along with the rest. I visit http://localhost:6000, I see nothing ...

_"Odd"_, - I thought, - _"is my server down?"_ It wasn't. I hit refresh, then hard refresh,
then desperately opened incognito, maybe some weird caching issue ...

Nothing changed.

{{< image src="/images/20230415/unreachable.png" position="center" style="border-radius: 8px;" >}}

At that point I just didn't give it much thought. It was late anyway, I was tired.

***

2 or 3 months later, I found some free time to get back to the same project. I spawned the 
containers up again, visited the same localhost:6000 and ... nothing.

_"Is the frontend process even running?"_ I thought. So I get into the running container,
and indeed the `npm run dev` was there, running. Remix was working just fine, but I didn't see it in the browser.

_"Maybe some weird host resolution issue?"_ I check my `/etc/hosts` but again, nothing interesting. 

_"Maybe another service running at the same port?"_ ... yeah, dumb thought, how can 2 processes bind on the same port? 
Either way, I did `lsof -i :6000` and ... it was the frontend container.

But then, a sign of hope. Innocently, I ran the poor man's browser:

```bash
$ curl localhost:6000
```

And lo and behold, HTML in my terminal!

_"Something fucky is at play here!"_

No obvious reason for this to happen. So I Googled `curl works but browser does not` and started checking
different StackExchange results, but didn't find anything that would explain this. I even found someone that blamed it on [DNS](https://superuser.com/questions/924950/site-displays-with-curl-but-not-within-the-browser).
But we all know it is never DNS (unless it's DNS, but this wasn't the case). Or even [CORS](https://stackoverflow.com/questions/38689350/for-what-reason-i-can-access-the-resources-by-curl-but-not-in-the-browser),
but no, this was surely not a CORS issue. Out of desperation, I switched to port 3001 and ... it worked. It finally worked!

But that wasn't important anymore.

I didn't care that it was fixed, I needed to know why it was broken!
Why would it not work for port 6000? I knew that ports up to 1023 are system ports and a no-go,
so nothing should stop me from using port 6000 ... unless Chromium has other plans.
And I accidentally found a file on Chromium's codebase that explains the fuckiness at play.

https://chromium.googlesource.com/chromium/src.git/+/refs/heads/master/net/base/port_util.cc#99

```cpp
// The general list of blocked ports. Will be blocked unless a specific
// protocol overrides it. (Ex: ftp can use port 21)
// When adding a port to the list, consider also adding it to kAllowablePorts,
// below.
const int kRestrictedPorts[] = {
    1,      // tcpmux
    7,      // echo
    9,      // discard
    11,     // systat
    13,     // daytime
    15,     // netstat
    // ...
    6000,   // X11
    // ...
};
```

Damn X11! Who'd've thunk?! Apparently, browsers restrict access on some ports due to security concerns and potential vulnerabilities.
Chromium, and later I learned [Firefox](https://www.reddit.com/r/firefox/comments/ttms50/since_when_was_this_a_thing_in_firefox_trying_to/)
as well, blocks access to some ports, 6000 included, and unless you specifically tinker with the settings to allow access on these ports, the browser
just won't let you.

And now I can sleep in peace that the issue knowing that it wasn't me being too dumb to set up a simple
container running a simple command building a simple frontend app. The issue was me being ignorant about
X11 reserved ports, a window system that [hit a two decade low on development pace during 2022](https://www.phoronix.com/news/XServer-2022-Development-Pace) ...

