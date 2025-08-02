---

title: Testing Go with Ruby
pubDatetime: 2021-03-28
description: Tearing down language barriers, using RSpec tests for a Git-like project built in Golang
comments: true
tags:
  - go
  - golang
  - ruby
  - rspec
  - testing

slug: testing-go-with-ruby
---

Earlier this month I started building a [Git clone in Go](https://github.com/aziflaj/gogot) as a way to learn more about the language and see how it feels working with it. All was going well and untested, until I faced some weird issue in the way blobs were created. Of course, I didn't know what the issue was back then, I just knew my time machine command was kinda broken. So I thought writing tests would be a good way to catch these errors in time. I tried writing some tests in Go for a couple of core files, but to be honest I didn't really think it would be a viable approach for my Gogot. So I decided to use RSpec instead of the default Go's testing library. Why RSpec? Glad I asked:

1. De facto standard-setter in behavior-driven testing. The describe/context/it blocks have been copied by libraries in languages other than Ruby, because it's easy to understand and to describe different scenarios with it. Go (afaik) lacks something similar, and I highly doubt it will get something similar in the foreseeable future.

2. I know Rspec better than [insert testing framework here]. It is a spoon that I have been sharpening for 5 years now, and it doesn't seem like the right time to get rid of it yet.

3. I can execute shell commands in Ruby using backticks, and it makes writing tests for Gogot as a shell utility really straightforward.

So to make it a bit fancier, I'm using a Dockerfile to build Gogot inside a Go container, and then I'm "installing" the compiled 'gogot' onto a new Ruby container which contains only the test code, where I'm finally running all the tests through the 'rspec' command. I added a shell script that will build the container and will run the tests, and if everything goes right the whole output should fit inside a terminal's screen :D

![light-side-better-than-dark-side](/assets/images/20210328/tests.png)


For a closer look at the tests, here's writing the scenario "Adding files without a path prints usage information" in code, following the good practices of [Better Specs](https://www.betterspecs.org/). The commands inside the backticks are executed "for real", and I have to make sure to clean up after every example (test) that I run. Since this is executed inside a container, the world outside it in not affected.

![testing](/assets/images/20210328/spec.png)


Besides all that testing suite and connecting it to Travis CI, I have been working on a 'status' command that looks like this (ignore the "nothing to commit message" at the end, it's still work in progress):

![gogot status](/assets/images/20210328/status.jpeg)


Looking at the languages section of the Github repo, I'm thinking this is one language away of becoming a chimera of bad code in more languages than you need, but I feel this approach is better for testing it as a real tool. Writing the tests this way helped me find more than enough issues in the code and also refactor here and there.
