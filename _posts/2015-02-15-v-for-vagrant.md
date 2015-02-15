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

```bash
vagrant init hashicorp/precise32    # or any other virtual system you want
vagrant up
```

##Why should you use Vagrant?
Next time you develop a web application, probably your development environment (your machine) and the production environment (the deployment server) won't be the same. You can still develop your application and test it in your environment, but some things might not go as well when you deploy on the production server. You may try to mimic the production environment by changing your development environment, but are you willing to do that for each project? 

This is where Vagrant comes in. It allows you to create a virtual machine that will be the same as the production environment. If you are working on a team and everyone has a different environment, you can all sync into the same environment as the production environment, and be sure that the code **will** work.
