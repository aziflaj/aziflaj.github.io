---
layout:     post
title:      V for Vagrant
pubDatetime:       2015-02-15
comments:   true
description:    Vagrant is a very nice provision tool that helps me with web development. I don't want to install every tool on my machine; I want my environment as clean as possible. Vagrant helps me create virtual machines with all the necessary tools for any project.
tags:
    - vagrant
    - provision
    - puppet

slug: v-for-vagrant
---

<p>
It has been almost one year since I first heard of <a href="https://www.vagrantup.com/">Vagrant</a>, and maybe 7-8 months since I first used it. Vagrant is a provision tool; it provides you a full development environment with all the tools required.
</p>

You can find it on <a href="https://github.com/mitchellh/vagrant">GitHub</a> and after reading through its readme, you will figure out it's really simple to set up. After installing Vagrant, all you have to do is execute this at your command-line:

```bash
$ vagrant init hashicorp/precise32    # or any other virtual system you want
$ vagrant up
```

## Why should you use Vagrant?
Next time you develop a web application, probably your development environment (your machine) and the production environment (the deployment server) won't be the same. You can still develop your application and test it in your environment, but some things might not go very well when you deploy to the production server. You may try to mimic the production environment by changing your development environment, but are you willing to do that for each project? What if you are working on multiple projects at the same time?

This is where Vagrant comes into play. It allows you to create a virtual machine that will be the same as the production environment. If you are working on a team and everyone has a different environment, you can all sync into the same environment as the production environment, and be sure that the code **will** work. You can finally get rid of the most famous excuse in the developer's vocabulary: _"It works on my machine..."_

## How to use Vagrant
Vagrant comes with a simple-to-use CLI which helps you create a new environment, start working into the environment, change its features if you have to, destroy the environment and so on.

The first thing you have to do is to <a href="https://docs.vagrantup.com/v2/vagrantfile/">create a Vagrantfile</a>, which tells Vagrant every detail about the new environment that it is about to create. Next, you have to execute
```bash
$ vagrant up
```

This will make Vagrant start creating the environment.
<p>
<i>Keep in mind that this environment that is being created is nothing more than a virtual machine, so it needs a software such as <a href="https://www.virtualbox.org/">VirtualBox</a> or any other similar to work.</i>
</p>

When you run `vagrant up` for the first time, it will download the Operating System for the virtual machine, so be patient. This will happen only once, so this means next time you go on working on the project or even use the same operating system with Vagrant, it won't have to re-download it.

After the process, you can log into the "newborn" system by executing:
```bash
$ vagrant ssh
```

<p>
<i>On Windows, to execute SSH you need to have an SSH client installed, such as PuTTY. Also, you can use the SSH service that comes with Git command-line.</i>
</p>

After SSH-ing, you are free to use the development environment as you wish. Normally, you would create a shared folder, called `synced folder` between your host system and the virtual one. Usually, this folder is the same folder as `Vagrantfile` is, but you can change it as well. This means that the source code is accessed by the host and the guest system **at the same time**, and you can still use your favorite IDE/text editor.

When you end your work, all you have to do is execute:

```bash
$ exit              # to exit from the VM
$ vagrant halt      # to stop the VM
```

Also if you finished with the project and want to get rid of the whole Vagrant files, all you have to do is add `vagrant destroy`. It will delete all Vagrant-related files and leave only the source code behind.

## Last words
You can choose from a long list of virtual machines at <a href="https://vagrantcloud.com/">Vagrantcloud.com</a> and pick the one that suits you best. If you need help configuring your Vagrant system, there are many GUI configuration tools online which you can use for free. I wrote an article called <a href="http://www.sitepoint.com/5-easy-ways-getting-started-php-vagrant/">5 Easy Ways to Get Started with PHP on Vagrant</a> which you may find useful if you are thinking of using Vagrant for PHP development.
