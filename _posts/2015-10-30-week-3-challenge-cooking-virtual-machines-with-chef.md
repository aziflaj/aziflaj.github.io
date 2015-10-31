---
layout:     post
title:      "Week #3 challenge: Cooking Virtual Machines with Chef"
date:       2015-10-31
summary:    "A long time ago in a galaxy far, far away (and probably in a parallel universe) Han Solo decided to open a restaurant, becoming Chef Solo and one of the best in the business. His specialty was a dish including Automation, Servers, Data centers, and some \"secret formula\" which we're given to believe its a Ruby DSL... Star Wars jokes aside, this week on my 52 weeks, 52 projects challenge I tried to provision a virtual server using chef solo."
tags:       [challenge, automation, provision, devops, ruby, vagrant, chef, solo]
---

A long time ago in a galaxy far, far away (and probably in a parallel universe) Han Solo decided to open a restaurant, becoming [Chef Solo](https://docs.chef.io/chef_solo.html) and one of the best in the business. His specialty was a dish including Automation, Servers, Data centers, and some "secret formula" which we're given to believe its a Ruby DSL... Star Wars jokes aside, this week on my [**52 weeks, 52 projects**](https://aziflaj.github.io/52-weeks-52-projects/) challenge I tried to provision a virtual server using chef solo.

## Virtual Servers Using Vagrant
[Vagrant](https://www.vagrantup.com) is a tool for creating lightweight virtual machines, without the heavy User Interface. When you need a development environment, may choose to use a \_AMP server, or create a replica of the production environment and be prepared for (almost) everything that may happen. The first one is easy: most (if not all) \_AMP stacks are simple to install and set up and come with a graphical interface that is simple to understand and use (think PhpMyAdmin). But there are some limitations to that, which I've previously written in a blog post titled [V for Vagrant](https://aziflaj.github.io/v-for-vagrant/). 

<iframe width="560" height="315" src="//www.youtube.com/embed/_I94-tJlovg" frameborder="0"> </iframe>

If you follow what the above video says, using a development environment that is a replica of the production one, will help you in many aspects. That's why instead of any \_AMP stack, I choose to use Vagrant to create virtual machines that act as local servers. Of course, Vagrant just creates a plain virtual server machine with only the OS installed, nothing else. To prepare the server for operation you need to **provision** it, _i.e._ install the proper software, make some configurations, etc. This is where Chef comes into play.

## Cooking with Chef
[Chef](https://www.chef.io/chef/) is a configuration management tool, which uses a pure-Ruby Domain Specific Language (DSL) for writing configurations in what are called _"recipes"_. These recipes describe how Chef should provision the VM, like what password to use for the database, what version of software to install, what packages to include, etc. Basically, what Chef does is turn your infrastructure into code, making it more developer-friendly. By turning infrastructure into code, Chef allows you to include it into Version Control, test it, etc. It is used in many big companies, including [Facebook](https://www.chef.io/customers/facebook/), [Rackspace](https://developer.rackspace.com/blog/cooking-with-chef/) and [Airbnb](http://nerds.airbnb.com/making-breakfast-chef-airbnb/).

I created a virtual machine that includes Apache2 webserver, PHP and Composer, Node.js and PostgreSQL. It is not a big thing, but it is a good start for me after using PuPHPet and Laravel's Homestead for a long time without caring about how they're created.

> Both PuPHPet and Homestead are created using Puppet, which is also a very important provisioning system just like Chef. One difference is that Chef uses a Ruby DSL turning infrastructure into **code**, while Puppet turns infrastructure into **configuration**. Think of it as Gradle vs Maven or Gulp vs Grunt, if you're familiar with those build tools.

And that’s it for the third week! You can find the source code in [this GitHub repository](https://github.com/aziflaj/vagrant-chef); you can fork it and change anything if you want to. If you want to ask me anything, leave a comment below.
