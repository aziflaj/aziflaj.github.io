---
layout: post
title: "The Apple Farm Parable: Understanding Monitoring and Observability"
date: '2025-01-27'
image:
  path: "https://github.com/aziflaj/aziflaj.github.io/blob/main/assets/images/20250127/applefarm.png?raw=true"
  alt: "The Apple Farm Parable"
---

In informal Albanian slang there's a saying. Whenever you want someone to explain a topic in the most simplistic terms, you ask them to _"explain it using apples"_. We had a similar case in my Software Engineering class at [UNYT](https://unyt.edu.al) last week, where students didn't totally comprehend the difference between monitoring and observability, not knowing where one ends and where the other begins. So as a good storyteller, if I do say so myself, I came up with the following parable.

Imagine we own an apple farm. We have planted different apple trees over a vast field and when the fruits are ripe, we send the right specialists to harvest them, sort them in different crates based on quality and species, load them in trucks and deliver them to either grocery stores or apple jam factories.

Given we are good engineers, - _and also very lazy,_ - we have automated a lot of the processes in our farm. We don't have to check for ripe fruit regularly, we use a network of IoT sensors backed by a ML system that based on the measured events, it notifies the right people to take action. For example, when time comes, harvesters get a push notification that _"Tree#1337 is ready to harvest"_; they go and gather the fruits, and pass the batch through an automated apple sorting machine. When this machine fills a crate, it is weighted, registered with a barcode for tracking purposes, and loaded in the back of a truck destined to an apple jam factory or local grocery. At this point, a designated driver gets notified to show up and drive the truck to the right destination.

At the end of each month, we run some financial and productivity checks, e.g. whether we have yielded the expected amount of jam from what we harvested. It turns out, we didn't! And our Certified Applefarm Auditor starts investigating on the reasons why. They try to answer questions like:

- Did we send all crates to the right destination?
- Were the crates delivered intact?

Because we have tracking barcodes for each crate, and because we were smart enough to log them through each event that affects them, the Auditor is able to pinpoint the issue: a certain crate did not reach the destination. Maybe apples fell off the truck, or maybe the driver was [hungry for apples](https://www.youtube.com/watch?v=-VKDlEhFJm8). We can't know for sure with what data this system provides, but we can close in on the scope of the underlying issues.

***

In more technical terms, measuring the state of the apples and trees through our IoT mesh and notifying the right people when the time is right, is equivalent to setting up metrics and alerts, two basic concepts of Monitoring. Giving each crate a tracking barcode and logging their state as they pass through our Apple'n'Jam delivery system is equivalent to request tracing through trace IDs and logging, two basic concepts of Observability.

Monitoring answers the question _"What?"_, as in _"What is the current state of the system?"_, while Observability answers the question _"Why?"_, as in _"Why is the system in the current state?"_.
