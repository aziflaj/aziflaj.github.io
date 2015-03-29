---
layout: post
title: Old programming jokes and what is wrong with them
summary: 
tags: [reveal, js, presentation]
---

<p>
While browsing my really old pictures on my computer and some old posts on social media, I (rastisa) in some funny used-to-be stuff that I used to say, write and believe. Well, not all of them are true, my opinions about them has changed during time, and I don't believe or support them anymore. Some of them are n00b thinkings (don't judge, we all start as n00bs). Here I am giving a list of some of them, hopefully any young programmer doesn't do the same mistakes.
</p>


<h2>C++ is better than Java</h2>
![C++ better than Java]({{ site.url }}/images/programming-jokes/Java.jpg)

First of all, there is no language better than another! Except Visual Basic; all languages are better than Visual Basic. 
I started programming with C++, and it was feeling nice. The way the computer was doing what I wanted it to do was amaising! If it blew up, it was my fault, not the computer's. I wanted to learn every small component of C++, and I loved every tiny bit of it. But I knew that C++ was really old, most of people didn't use it on an everyday basis. Instead of C++, Java was (and still is) the most used programming language by the developers in the world, with a community of over 10 million of developers. So I took a look at Java (a really fast look, I'd say), trying to find things to hate it. I managed to find a couple actually:

<ul>
	<li>The absence of pointers</li>
	<li>Forcing to use Object Oriented Programming (_hence_ no procedural code)</li>
	<li>Single inheritance (C++ has multiple inheritance: one child class may have more than one parent)</li>
	<li>No inline assembler (C++ inherited this from its old parent: C)</li>
	<li>etc.</li>
</ul>

I remember calling _JavaFag_ everyone who loved Java and thought it was better than every other language. Now, I proudly say that <strong>I was wrong</strong>. Why? Well, there is a reason for every one of the "reasons" above.

<h4>The absence of pointers</h4>
Whoever wrote C code knows the importance of pointers. C++ inherits this importance from the good ol' C and of course makes it better. So the first thing you think when you learn that Java lacks pointers is "WTF?!" (that stands for **W**hy **T**his **F**eature). But whait: Java is implemented above a C/C++ layer. This means that, underline, it **should use pointers**. And it actually does. And also you as a Java developer use pointers almost every line of code you write. How? Every time you use an object, you actually are using **a pointer to a class instance**. That's why objects are passed by reference on methods. Also, when you write non-static methods of a class, you use the <code>this</code> keyword. Yep, another pointer, just like in C++. 

<h4>OOP - only</h4>
Consider C++ as a linking bridge between C and Java. C is mostly procedural (other paradigms also), while Java is mostly OOP (other paradigms also). C++ is both procedual and Object Oriented. Is this a good thing? It depends on your angle. But there is a reason why Java supports OOP so much: **code modulasization**. Good programmers write code that humans can understand, and modular code is better to read and to understand. No more words needed!

<h4>Single Inheritance</h4>
There is a reason why a child class shouldn't have more than one parent. Consider a class called <code>Child</code> which extends a <code>Mother</code> class and a <code>Father</code> class. Now both parents have a method (function) that is called <code>speak()</code>. If the <code>Child</code> class implements that method is OK, but what if it doesn't? Will there be called the <code>Mother</code>'s method, or the <code>Father</code>'s one? This is what they call <a href="http://en.wikipedia.org/wiki/Multiple_inheritance#The_diamond_problem">the diamond problem</a>. Java tries to solve this problem by simply not allowing multiple inheritance.

<h4>No inline assembler</h4>
I have one question: Why do you want to write assembly code in a Java program?! Well, let's supose you have a strong reason (still, why?!). There is a way to write code in C, C++, Assembly and other native languages by using  <a href="http://docs.oracle.com/javase/7/docs/technotes/guides/jni/">Java Native Interface</a>. It is not the best thing to do, since it would remove the WORA of Java and make the code platform-dependent.


<h2>C# - May the hate be with you</h2>
