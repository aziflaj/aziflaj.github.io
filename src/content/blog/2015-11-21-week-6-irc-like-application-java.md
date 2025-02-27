---
layout:     post
title:      "Week #6: An IRC-like Application in Java"
pubDatetime:       2015-11-21
comments:   true
description:    As a challenge, I built a client-server application that allows multiple clients to communicate with each others using java Threads and Sockets
tags:
    - project52
    - java
    - socket

slug: week-6-irc-like-application-java
---

I started my 6th week of the challenge very late. In fact, I changed plans during the week, and decided to go with this application on Tuesday.

## Why did I change the project middle-week?
A couple of times during the Software Engineering class, the professor has been saying "Don't think Apache2 is Rocket Science! Even you can build your own server". You don't have to get that for granted, but keep in mind that the Internet Protocol (and also UDP and TCP) is built upon network sockets, so building a server basically means using sockets everywhere. That being said, he started showing how to use sockets with a simple PHP program, explaining what sockets are, how do they work, how PHP actually uses C in the background to achieve InterProcess Communication (IPC), etc. One of the things he mentioned was that of using multiple threads for serving more than one clients, and I actually took that as a challenge; to use Java Threads and [Server]Sockets in order to create a client-server application that allows clients communicate with each other.

## What did I build
The application is very simple in concept and you're probably familiar with the idea. It is an application that resembles an Internet Relay Chat, or IRC. IRCs were a way of group chatting over TCP that dates back in the late 80s. It is still used actually, just not that actively. You've probably heard of groups that use IRC:

- [Anonymous uses IRC](https://www.anonops.com/) over TSL
- [Ruby community](https://www.ruby-lang.org/en/community/) uses IRC in the #ruby channel
- [Mozilla uses IRC](http://irc.lc/mozilla/developers/) to connect all the developers

One of the most known IRC networks known is [Freenode](https://freenode.net/), with more than 90 000 users in the world. Of course it's not the only one, and some (so-called) modern group chat networks like [Slack](https://slack.com/) are built upon the same idea and model (with some added confidentiality).

What I did this week (actually, only in 2 days) was a simple application that resembles an IRC node, _i.e._ a server connected to multiple clients that allows different clients communicate with each other in a **public chat**. It should be noted that it is not a fully-featured IRC:

- Usernames might conflict
- You can't ping another user
- You have to manually create a different server process in order to create a different message thread
- message threads are not stored or logged

Nevertheless, it can be a good Proof of Concept for different applications and examples. It uses [Sockets](https://docs.oracle.com/javase/7/docs/api/java/net/Socket.html) and [ServerSocket](https://docs.oracle.com/javase/7/docs/api/java/net/ServerSocket.html), different running threads for different clients, [Swing API](https://docs.oracle.com/javase/7/docs/api/javax/swing/package-summary.html) for the User Interface, and uses a single `.properties` file to store some necessary information to connect a client with the server.

## How to set it up

Firstly, you'll need to clone the repository of the project on GitHub:

http://github.com/aziflaj/IRC

There are two folders, `Client` and `Server` and also two shell scripts for compiling and running both the server and the client.

> When I was in the first year of Computer Engineering, there was this professor teaching Computer Science 101 explaining the structures in C. He started his explanation saying: "Structures are... structures". That being said, it should be obvious which one is the Client and which one is the Server.

The client comes with a `client.properties.example` file which you have to rename to `client.properties` and also put some useful information inside. Property files are containers of key-value pairs, and this file contains information about the server you want to connect (the hostname and the port) and also about you (the username). These values are then used to connect you to the server. Unfortunately, since most programmers are lazy, the server is not that customizable. The server runs on port 4444 (hard-coded) and you'll have to change that if you want to bind the process to another port.

You can run the server and the client by executing:

```bash
#
# This is bash
# Sorry, Windows user :-(
#

# Run the server
$ chmod +x server
$ ./server -r

# Run the client
$ chmod +x client
$ ./client -r
```

Then you can use the UI client to chat all day long.

And that's for the 6th week! If you want to make any changes, feel free to fork me on GitHub and fix anything you think should be fixed.
