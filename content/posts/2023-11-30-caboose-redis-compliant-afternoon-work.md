---
layout: post
title: "Caboose: a Redis-compliant Afternoon's Work of Genius"
date: '2023-11-30'
---

TL;DR: I built an **In-Memory, Redis-compliant** (_to some extent, don't get your hopes high_) **Key-Value Data Store**.
It took me a bit over 2 hours, but it works and it sucks at the same time.
You can find it in my GitHub: [aziflaj/caboose](https://github.com/aziflaj/caboose)

## Obligatory Backstory
My experience with badly written databases goes way back to 2012, when I wanted
to develop a social network in C++ (however crazy that sounds).
I had no idea SQL was a thing, so I implemented the persistence layer myself
in binary files and spaghetti code -- you might even call it an **"Object-Oriented 
DataBase Management System"**, an OODBMS.

A few years after that, around 2017,  a friend and I had this idea of putting
Albania in the map: by developing a key-value database, with in-memory as well as 
disk persistence support, named [Balo DB](https://github.com/balodb) (TM Pending).
It seems like every 5 years I have to consider writing a database system.

During my summer workation of 2023, when I was [backpacking across Eastern Europe](https://www.youtube.com/watch?v=OWxrV0-ObCY),
I had this weird idea (again) to write a Proof of Concept, Redis-compliant database.
At this point I've been using Redis in almost every configuration possible:
as a cache ([like this](https://aws.amazon.com/elasticache/redis/)),
as a websocket adapter via Pub/Sub ([similar to this](https://socket.io/docs/v4/redis-adapter/)),
as a queue for async processing ([via Sidekiq](https://github.com/sidekiq/sidekiq)),
even as the primary DB of an application.

## What really is Redis?

Recently I had to debug some Rails sessions stored in Redis; it was a red herring
but I got more exposure to the internals. I had already read about this thing called RESP
(more on that later) and in my mind, a "Redis-compliant Key-Value Store" is as easy as 1-2-3:

1. Listen for TCP requests on a given port
2. Parse an incoming request and pass it to a Hash Table
3. Respond back to the requester with the result of the previous operation

But as usual, the devil is in the details. This "compliancy" concept revolves around
something called REdis Serialization Protocol -- RESP. The requests a Redis Client sends
to a Redis Server have to be serialized in a specific format, described in details in
the [Redis docs](https://redis.io/docs/reference/protocol-spec/) and in a high level lazy approach on my README.
There are a few versions of RESP, I think I was too lazy to go through the whole doc but from my
couple of hours long experience, the ones you truly _**need**_ are:

- Bulk Strings
- Errors
- Arrays -- Very important, almost all requests are arrays of strings
- Integers, and you can hack your way into using `1` and `0` as Booleans

So the first thing I did was [this `sarge` module](https://github.com/aziflaj/caboose/tree/main/sarge) (almost sounds like [`serde`](https://serde.rs/))
which handles Serialization and Deserialization following RESP specifications. And once the
binding agent is ready, we can move onto listening for requests and manipulating the Hash table,
which is handled by [the `vic` module](https://github.com/aziflaj/caboose/tree/main/vic).

## Thinking about concurrency

If you didn't know, [Redis is (mostly) single-threaded](https://redis.io/docs/management/optimization/latency/#single-threaded-nature-of-redis).
Similar to your run-o-the-mill Node.js application, it does some async I/O wizardry and
some event loop shenanigans, but I don't feel competent enough to talk about it.

My Caboose DB uses goroutines to handle each request on its own lightweight, "green thread" (not an actual OS-level thread).
I also "protected" my almost-global Key Value Store with mutexes, both on table level and on record level, as such:

```go
type KVStore struct {
  mu      sync.RWMutex
  data    map[string]string
  mutexes map[string]*sync.Mutex
}
```

The table-level mutex is locked before each operation, and record-level mutexes are locked before a value is set or deleted.
The good thing about the table-level [`sync.RWMutex`](https://pkg.go.dev/sync#RWMutex) is that it can get locked by multiple readers,
but only one writer. The record-level mutex is a simple `sync.Mutex`, and two requests can't write or delete the same record at the same time.

In simple laymen terms, what these mutexes achieve are:

- allowing multiple requests to **set/update different records** at the same time,
- allowing multiple requests to **read the same record** at the same time,
- denying multiple requests to **set/update the same record** at the same time

That's acceptable and more than enough when you "hack into it" for a couple of hours and release a Proof of Concept that's not gonna be used by anyone, right?

## So... what's the point?

It sounds impressive to other nerds `¯\_(ツ)_/¯`

It also addresses a couple of steps from [John Crickett's Redis challenge](https://codingchallenges.fyi/challenges/challenge-redis),
and I _might_ use it in the future as a playground for data structures used by In-Memory Data Stores,
or idealistically _inspire someone to follow my footsteps_ (knowing myself, I will just forget I did this).
