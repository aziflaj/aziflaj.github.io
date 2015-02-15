---
layout:     post
title:      V for Vagrant 
date:       2015-02-15
summary:    Vagrant is a very nice provision tool that helps me with web development. I don't want to install every tool on my machine; I want my environment as clean as possible. Vagrant helps me create virtual machines with all the tools required for any project.
tags:       [Vagrant, Provision, Puppet, PuPHPet]
---

<p>
It has been almost one year since I firstly heard of <a href="https://www.vagrantup.com/">Vagrant</a>, and maybe 7-8 months since I firstly used it. Vagrant is a provision tool; it provides you a full development environment with all the tools you require. 
</p>

You can find it on <a href="https://github.com/mitchellh/vagrant">GitHub</a> and after reading through its readme, you will figure out its really simple to set it up. After installing Vagrant, all you have to do is execute this at your tty:

{% highlight bash %}
vagrant init hashicorp/precise32  # or any other virtual system you want
vagrant up
{% endhighlight %}

## Why should you use Vagrant?
Next time you develop a web application, probably your development environment (your machine) and the production environment (the deployment server) won't be the same. You can still develop your application and test it in your environment, but some things might not go as well when you deploy on the production server. You may try to mimic the production environment by changing your development environment, but are you willing to do that for each project? 

This is where Vagrant comes into play. It allows you to create a virtual machine that will be the same as the production environment. If you are working on a team and everyone has a different environment, you can all sync into the same environment as the production environment, and be sure that the code **will** work. You can finally get rid of the most famous excuse in the developer's vocabulary: _"It works on my machine..."_

## How to use Vagrant
Vagrant comes with a simple-to-use CLI and using that you can create a new environment, start working into the environment, change its features if you have to, destroy the environment, and so on.

The first thing you have to do is to <a href="https://docs.vagrantup.com/v2/vagrantfile/">create a Vagrantfile</a>, which is everything Vagrant needs to know about the environment that is about to create. Next, you have to execute
{% highlight bash %}
vagrant up
{% endhighlight %}

This will start to create the environment. 
<p>
<i>Keep in mind that this environment that is being created is nothing more than a virtual machine, so it needs a software such as <a href="https://www.virtualbox.org/">VirtualBox</a> or any other similar.</i>
</p>
