---
layout:     post
title:      Building a Slot Game in Java 
date:       2015-03-02
summary:    "A couple of years ago, when I was learning Java programming, I thought of testing myself and my programming skills by writing a game in Java. Now, I'm not going to call it game programming, since game programming is way more that what I did. In fact, what I did was just a test for me. So I decided to write a game I was playing on my old Nokia E50 phone, a Slot Game."
tags:       [Java, Slot, Slots, Game]
---

<p>
A couple of years ago, when I was learning Java programming, I thought of testing myself and my programming skills by writing a game in Java. Now, I'm not going to call it <i>"game programming"</i>, since game programming is way more that what I did. In fact, what I did was just a test for me. So I decided to write a game I was playing on my old Nokia E50 phone, a Slot Game.
</p>

This slot game I was playing on my phone was really simple. It had only 3 slots with different items in each. You had to push the **Spin** button in order to spin the slots, and you won a small amount of coins if two or three slots were alike. Of course, 3 slots were better than 2. It is not really hard to make a game like this, but for a beginner it is good to start with. As I remember, this was probably the first program that I could tell others: <i>"Look at what I just did!"</i>

So I started working on it (I remember using <a href="https://netbeans.org/" target="_blank">NetBeans</a> at that time), firstly as console-only, and then using GUI. The first thing I did, was deciding what kind of images (actually their names, not the images themselves) I would use. I wrote this line of code:

{% highlight java %}
String[] symbols={"Seven","Shamrock","Diamond","3Bar","Star","Bell", "Bar","Orange","Lemon"}; //slot symbols
{% endhighlight %}

I also decided what would be the amount of "money" that the user would win if he matched two or three symbols:

{% highlight java %}
//amount winning
int[] twoMatches={30,16,15,12,11,10,9,7,5};
int[] threeMatches={60,32,30,24,22,20,18,14,10};
{% endhighlight %}

After that, I went on writing the code that was suposed to randomly choose one of the elements in the <code>symbols</code> array. This can be done using the <a href="http://docs.oracle.com/javase/7/docs/api/java/lang/Math.html#random()" target="_blank"><code>Math.random()</code><a/> method, or the calling the <a href="http://docs.oracle.com/javase/7/docs/api/java/util/Random.html#nextInt(int)" target="_blank"><code>nextInt()</code></a> method in a <a href="http://docs.oracle.com/javase/7/docs/api/java/util/Random.html" target="_blank"><code>Random</code></a> instance, or you could use the wrong way I did in the beginning:

{% highlight java %}
// a random from 0 to 9
int randomInt = (int) System.currentTimeMillis() % 10;
{% endhighlight %}

Of course, I soon switched to calling <code>Math.random()</code>, and in order to get a number that I could use as index for my array, I wrote this block of code:

{% highlight java %}
double random=Math.random();
random*=8;
int choice=(int) random;
{% endhighlight %}

So the variable <code>choice</code> was the random index that I could use to get a random item from the array (keep in mind that Math.random() returns a **double** between 0.0 and 1.0)
