---
layout: post
title: Old programming jokes and what is wrong with them
comments:   true
summary: While browsing some old pictures on my computer and some old posts on social media, I stumbled upon some used-to-be-funny stuff that I used to say, write and believe. Well, not all of them are true, my opinions about them have changed with the time, and I don't believe or support them anymore. Some of them are n00b thoughts (don't judge, we all start as n00bs). Here I am giving a list of some of them, hopefully any young programmer doesn't make the same mistakes.
tags: [c++, java, c#, javascript, programming, jokes, visual basic]
---

<p>
While browsing some old pictures on my computer and some old posts on social media, I stumbled upon some used-to-be-funny stuff that I used to say, write and believe. Well, not all of them are true, my opinions about them have changed with the time, and I don't believe or support them anymore. Some of them are n00b thoughts (don't judge, we all start as n00bs). Here I am giving a list of some of them, hopefully any young programmer doesn't make the same mistakes.
</p>


<h2>C++ is better than Java</h2>
![C++ better than Java]({{ site.url }}/images/programming-jokes/Java.jpg)

First of all, there is no language better than another! Except Visual Basic; all languages are better than Visual Basic. 
I started programming with C++, and it feelt nice. The way the computer was doing what I wanted it to do was amazing! If it blew up, it was my fault, not the computer's. I wanted to learn every small component of C++, and I loved every tiny bit of it. But I knew that C++ was really old, most of people didn't use it on an daily basis. Instead of C++, Java was (and still is) the most used programming language by the developers in the world, with a community of over 10 million of developers. So I took a look at Java (a really fast look, I'd say), trying to find reasons to hate it. I managed to find a couple actually:

<ul>
	<li>The absence of pointers</li>
	<li>Forcing to use Object Oriented Programming (<i>hence</i> no procedural code)</li>
	<li>Single inheritance (C++ has multiple inheritance: one child class may have more than one parent)</li>
	<li>No inline assembler (C++ inherited this from its old parent: C)</li>
	<li>etc.</li>
</ul>

I remember calling _JavaFag_ everyone who loved Java and thought it was better than every other language. Now, I proudly say that <strong>I was wrong</strong>. Why? Well, there is a reason for each of the "reasons" above.

<h4>The absence of pointers</h4>
Whoever wrote C code knows the importance of pointers. C++ inherits this importance from the good ol' C and of course makes it better. So the first thing you think when you learn that Java lacks pointers is "WTF?!" (that stands for <strong>W</strong>hy <strong>T</strong>his <strong>F</strong>eature). But wait: Java is implemented above a C/C++ layer. This means that, underline, it <strong>should use pointers</strong>. And it actually does. And also you as a Java developer use pointers almost every line of code you write. How? Every time you use an object, you actually are using <strong>a pointer to a class instance</strong>. That's why objects are passed by reference on methods. Also, when you write non-static methods of a class, you use the <code>this</code> keyword. Yep, another pointer, just like in C++. 

<h4>OOP - only</h4>
Consider C++ as a bridge between C and Java. C is mostly procedural (other paradigms also), while Java is mostly OOP (other paradigms also). C++ is both procedual and Object Oriented. Is this a good thing? It depends on your point of view. But there is a reason why Java supports OOP so much: <strong>code modularisation</strong>. Good programmers write code that humans can understand, and modular code is better to read and to understand. No more words needed!

<h4>Single Inheritance</h4>
There is a reason why a child class shouldn't have more than one parent. Consider a class called <code>Child</code> which extends a <code>Mother</code> class and a <code>Father</code> class. Now both parents have a method (function) that is called <code>speak()</code>. If the <code>Child</code> class implements that method is OK, but what if it doesn't? Will there be called the <code>Mother</code>'s method, or the <code>Father</code>'s one? This is what they call <a href="http://en.wikipedia.org/wiki/Multiple_inheritance#The_diamond_problem" target="_blank">the diamond problem</a>. Java tries to solve this problem by simply not allowing multiple inheritance.

<h4>No inline assembler</h4>
I have one question: Why do you want to write assembly code in a Java program?! Well, let's supose you have a strong reason (still, why?!). There is a way to write code in C, C++, Assembly and other native languages by using  <a href="http://docs.oracle.com/javase/7/docs/technotes/guides/jni/" target="_blank">Java Native Interface</a>. It is not the best thing to do, since it would remove the WORA of Java and make the code platform-dependent.


<h2>C# - May the hate be with you</h2>
![C# - May the hate be with you]({{ site.url }}/images/programming-jokes/Csharp.jpg)

After throwing hate to Java, I'd go on with the sentence: <i>"And there is C#, the closed-source Microsoft copy of Java"</i>. I actually hated C# more than Java. There were reasons after that. Firstly, I don't like to name my methods using <code>PascalCase</code> just like my classes (at least Java uses <code>camelCase</code>). Remove that, and C# looks just like Java (the syntax, at least). Keep in mind that C# was released 5 years later than Java (conspirative programmer is conspirative). Well, I didn't write too much C# code but I know there are some features of C# that Java doesn't have, or that are introduced much later. The best feature of C# is <a href="http://en.wikipedia.org/wiki/Language_Integrated_Query" target="_blank">LINQ</a> (Language Integrated Queries), a way of executing SQL-like queries for object collections. The closest thing that Java has to that is the <a href="http://javadocs.techempower.com/jdk18/api/java/util/stream/package-summary.html" target="_blank">Stream API</a> added with Java 8, but that is nowhere close to LINQ (there are some <a href="http://en.wikipedia.org/wiki/Language_Integrated_Query#Implementations_in_other_languages" target="_blank">LINQ implementation</a> though).

Another issue, C# is strongly related to Microsoft. Microsoft's .NET framework exists only for Microsoft products, and this makes C# less cross-platform (0 times more than C++). There was always the <a href="http://www.mono-project.com/" target="_blank">Mono Project</a>, but it was really slower than JVM on every machine. But finally, Microsoft <a href="http://blogs.msdn.com/b/dotnet/archive/2014/11/12/net-core-is-open-source.aspx" target="_blank">decided to open-source it</a>. Later, we probably will see a cross-platform .NET, with a higher performance and usable in Mac OS, iOS, Linux and Android.


<h2>Javascript - its not Java script</h2>
![Javascript - its not Java script]({{ site.url }}/images/programming-jokes/Javascript.jpg)

Javascript has it drawbacks. First of all, it sucks!. 
<ul>
	<li>Not standard for all the browsers (don't come up with jQuery yet)</li>
	<li>Braces don't guarantee scope; everything exists in the global scope (WTF - Why This Feature)</li>
	<li>Semicolons are optional (Come on! Use them or don't use them! Make up your mind, woman!)</li>
	<li>Object Oriented (you'd wish!)</li>
	<li>Weak typing (prefix all the variables!)</li>
	<li>Debugging is a nightmare</li>
</ul>

OK, let's not be harsh on the kid. It may be not the best programming language, but it's not the worst (I'm looking at you, Visual Basic!). But after using it for a while, I started to like it. Well, not Javascript, but <a href="http://angularjs.org/" target="_blank">AngularJS</a> (my favorite JS framework) and <a href="http://ionicframework.com/" target="_blank">Ionic Framework</a>. 


Let's be honest, every language has its drawbacks. No programming language is perfect, and it's their right not to be (Visual Basic is abusing with this right). But thankfully, there isn't only one programming language. You are free to use the one which has the features you require for the project you have. There are two kind of programming lanugages: the ones you love, and the ones you don't know. So don't hate any of them, don't set high standards for the programming languages. All of them are good for something (except one... but now you know).

_How about you? Did you hate any programming language? Write it below on the comment section._
