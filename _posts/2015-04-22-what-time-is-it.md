---
layout: post
title: What time is it?
summary: "A couple of months ago, I saw a nice Javascript project: a background changing clock. Based on the Hex color codes, it could change the background based on the time, using hours, minutes and seconds as hashcode generator. I liked it and decided to write it by myself."
tags: [clock, time, javascript]
---

A couple of months ago, I saw a nice Javascript project: a background changing clock. Based on the [Hex color codes](http://www.color-hex.com/), it could change the background based on the time, using hours, minutes and seconds as hashcode generator. I liked it and decided to write it by myself.

I started with a simple HTML file, called `index.html` with this content:

{% highlight html %}
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>jsClock</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>

    <div id="clock"></div>

    <script src="clock.js"></script>
</body>
</html>
{% endhighlight %}

It's nothing unusual actually, just the basic HTML5 template that IntelliJ IDEA provides, plus a `div` with `id` set to `clock`, a Javascript file called `clock.js` and a CSS stylesheet.

The most important part of this clock is actually that Javascript file:

{% highlight javascript %}
var clock = function() {
    var date = new Date();

    var hours = (date.getHours() >= 10 ? date.getHours() : '0' + date.getHours());
    var mins = (date.getMinutes() >= 10 ? date.getMinutes() : '0' + date.getMinutes());
    var secs = (date.getSeconds() >= 10 ? date.getSeconds() : '0' + date.getSeconds());

    var hexColor = "#" + hours + mins + secs;

    document.getElementById("clock").innerHTML = hexColor;
    document.bgColor = hexColor;

    setTimeout(clock,1000);
};

clock();
{% endhighlight %}

This simple script gets the hours, minutes and seconds, adds them a trailing 0 if the number is only with one digit and sets the background color to this value every second. Nothing special, is it?

![clockjs in action]({{ site.url }}/images/clockjs/clock.gif)

You can find the source code on [this GitHub repository](https://github.com/aziflaj/jsClock). Feel free to fork it and change it as you wish.
