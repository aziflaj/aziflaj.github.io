---
layout:     post
title:      "Week #7: 7 questions answered"
date:       2015-11-28
summary:    "This is the 7th week of my challenge, and it has been good so far. I have been learning a lot and I intend to learn a lot more, maybe even get back to the same technologies and explore them more deeply. I learned a new JVM language called Kotlin, I learned provisioning with Chef, I did some Node.js development using Sails.js and I also made a simplified IRC in Java. I decided to take this week off from developing. Instead, I picked 7 interesting questions over time and I intend to answer them as good as I can."
tags:       [challenge, questions, c#, java, python, orm, javascript, nodejs, artificial, intelligence, neural, networks]
---

This is the 7th week of [my challenge](https://github.com/aziflaj/52), and it has been good so far. I have been learning a lot and I intend to learn a lot more, maybe even get back to the same technologies and explore them more deeply. I learned [a new JVM language called Kotlin](https://github.com/aziflaj/ToDo-kotlin/), I learned [provisioning with Chef](https://github.com/aziflaj/vagrant-chef), I did some [Node.js development using Sails.js](https://github.com/aziflaj/Sailor) and I also made [a simplified IRC in Java](https://github.com/aziflaj/IRC). I decided to take this week off from developing. Instead, I picked 7 interesting questions over time and I intend to answer them as good as I can.

The only issue I've had so far with the challenge is that I decided what technologies, methodologies and tools I will use during the challenge, but not what exactly to build. I don't want to build the same thing over and over again, changing only the programming language (_e.g._ build the same web application in Rails and Scala). Maybe I buy some time finding some nice ideas while I answer these questions :-)

## #1: C# vs Java
This is a question I've asked myself a lot before deciding to focus in Java. Recently, a professor asked me this question and I have also seen the same question asked in many developer groups and forums.

The syntax of both C# and Java is derived from C and C++, so developers who were familiar with those languages could adapt to them quickly. They both are Object Oriented, requiring every program to be wrapped in a class. They both use a Garbage Collector to manage memory allocation and deallocation. Programs written in C# and Java are portable among different machine architectures: C# uses the Common Language Infrastructure (CLI) and Java uses the Java Virtual Machine (JVM). Other similar features include generics, lambdas and anonymous functions, exceptions, single inheritance (multiple inheritance with interfaces), etc.

Without going too deep into the core of each language, there is one thing I like which is implemented in C# and not (yet) in Java: Language Integrated Query a.k.a LINQ. You can think of LINQ as a SQL-like syntax for collections of objects. The closest thing to LINQ in Java is the Stream API added in Java 8, but it is nothing close to LINQ (no SQL-like syntax). Also, C# allows the developer to use pointers in code blocks marked as `unsafe`, while Java doesn't have this feature. C# allows operators to overload, Java doesn't. In fact, some operators are pre-overloaded in Java, like the `+` operator to concatenate Strings, but the programmer can't overload more. This is one of the features that C# inherits from C++ and Java doesn't.

Both languages can be used for Mobile, Web and Desktop development. The leading mobile operating system, Android, is developed in C, C++ and Java. All native Android applications will have to use Java somehow. All hybrid apps, on the other hand, are wrapped in a `WebView` so the developer is allowed to use HTML, CSS and JavaScript to create applications (see [Cordova](https://cordova.apache.org/)). Before Android, some kinds of mobile applications were created using [Java ME](https://en.wikipedia.org/wiki/Java_Platform,_Micro_Edition), a flavor of Java for low-resource environments. 
In the web, there are plenty of ways how to use Java. [Java EE](https://en.wikipedia.org/wiki/Java_Platform,_Enterprise_Edition) is a flavor of Java designed for large distributed enterprise and internet environments. You can use Java in the server-side, using [Servlets](https://en.wikipedia.org/wiki/Java_Servlet) or [JavaServer Pages](https://en.wikipedia.org/wiki/JavaServer_Pages) (JSP). A Servlet is a back-end service that provides some response when accessed. This may or may not be HTML, regarding to the application developed. JSPs use Java code embedded in an HTML page and they are compiled into Servlets the first time they're accessed. You can also use it for front-end web development, where a Java Applet is embedded into a web browser. For both front- and back-end, there are plenty of frameworks that simplify the work of the developer. You can use Google Web Toolkit (GWT) for front-end development (Gmail is developed using GWT) or frameworks like [Spring Framework](https://spring.io/) for back-end development.

The same can be done using C# and .NET Framework, which actually adds most of its power to C#. It gives support for web application development, database connectivity, network communication, mobile development for Windows phones, etc. 

Basically, the languages really look alike and are used in similar domains. Deciding which one to use comes to the terms whether you like Microsoft or not. While you can use C# in most platforms, you can get the best out of it only in Microsoft platforms. A web development company could use .NET technologies because of Microsoft support about the Development (all the tips and tricks about C# and .NET) and the Operations (everything regarding to Windows Server and IIS). Java on the other hand is cross platform and you can run it in every machine, whether is Unix, Linux, or Windows.

## #2: Active Record or Data Mapper?
These are both different types of Object Relational Mappers (ORM). They represent a coding pattern used to map an object (in the OOP sense) into a relational entity, _e.g._ a record in a table of a relational database. ORMs are used in most (if not all) MVC frameworks, usually for creating Models and wrapping the business logic. 

If you have used Ruby on Rails, you've probably seen this:

{% highlight ruby %}
class User < ActiveRecord::Base
  # The rest of your model
end
{% endhighlight %}

or the same thing in Laravel:

{% highlight php startinline=true %}
use Illuminate\Database\Eloquent\Model;

class User extends Model 
{
  // The rest of your model
}
{% endhighlight %}

This is the example of an Active Record. The class you're mapping in the DB is the one that handles persistence. If you want to store a new user, you'd go like:

{% highlight ruby %}
user = User.new
user.last_name = 'Bond'
user.first_name = 'James'
user.save
{% endhighlight %}

You don't have to write getters or setters; the Active Record pattern doesn't require them. The `save` method is inherited from the `ActiveRecord::Base` class, which gives the `User` class the power to persist its objects into the DB. 

The other camp, Data Mapper, can be represented by other powerful ORMs such as Hibernate for Java and Doctrine for PHP. In this pattern, the model classes are just plain classes that don't have to inherit or extend anything. The persistence is then handled by a separate mechanism. When using Doctrine ORM in a Symfony2 application, you create a simple class with the table columns as fields:

{% highlight php startinline=true %}
// The User model
class User
{
  private $firstName;
  private $lastName;
  
  public function getFirstName()
  {
    return $this->firstName;
  }
  
  public function getLastName()
  {
    return $this->lastName;
  }

  public function setFirstName($firstName)
  {
    $this->firstName = $firstName;
  }
  
  public function setLastName($lastName)
  {
    $this->lastName = $lastName;
  }
}
{% endhighlight %}

Storing a user will be like:

{% highlight php startinline=true %}
// Create a user
$user = new User();
$user->setLastName("Bond");
$user->setFirstName("James");

// store a user
$entityManager = $this->getDoctrine()->getManager();
$entityManager->persist($user);
$entityManager->flush();
{% endhighlight %}

The `User` class is nothing more than a plain old PHP class with getters and setters for the fields (like a JavaBean for PHP). As you can see, the Entity Manager bundled with Symfony2 handles the persistence.

The main difference (and probably the most important one) is which _"thing"_ handles persistence. In the Active Record pattern, the object can store itself using the `save` method, while in the Data Mapper pattern the persistence is handled by another mechanism called Entity Mapper.

The `ActiveRecord` pattern is simpler to use and easier to learn, while the `DataMapper` is a bit hard to set up. In some cases, you have to configure the Entity Manager by yourself (most frameworks have already done this for you). On the other hand, `DataMapper`s are more flexible in the sense of executing raw SQL queries when implementing the Entity Manager. This can add a lot of power and use the lower-level database more efficiently. 

The "Which one to use?" question doesn't have a proper answer. You have to decide which one has the features that you need in your application. When using a framework, you can choose to use the bundled ORM, include another ORM you want, or use raw SQL queries if you should.

## #3: How is it possible to run JavaScript in a server?
Simple, in the same way you run it in a browser.

JavaScript (JS) was designed from Netscape to make it possible to run small applications in a web browser, and so to make websites more interactive with their users. The standardized version of JavaScript is called ECMAScript (ES for short). What you probably know as JS (if you come from the front-end development world), is in fact ES with added DOM (Document Object Model) and some other browser-dependent APIs (Application Program Interface). ECMAScript is the structure and syntax of the language. When you write loops or functions, when you use variables and objects, you're using ES.

Before analyzing how to use JS (or ES) outside the browser, let's take a look at how browsers work. A browser is a conglomerate of many components. Some of the components are:

- the UI layer, the one you see and interact with
- the rendering engine, which parses the HTML or XML files and styles them using CSS
- a networking layer to communicate with a back-end service
- a JS interpreter which (surprisingly) interprets the JavaScript code
- a data storage layer, such as IndexedDB or localStorage
- etc.

The JS interpreter, or the JS Engine, is that piece of the browser that executes JavaScript. If you divide that part of the browser from the rest of it, you can still execute JavaScript **without** all the helper APIs like the DOM or `XMLHttpRequest`. There are many different JS engines out there: Firefox uses SpiderMonkey, Java uses Rhino (until Java7) and Nashorn (since Java8), Chrome uses V8 engine, etc.

That should actually answer the question in the first place: **to use JS beyond the web in general, and in the server-side in particular, all we need to use is a JS engine.**

I have used Nashorn as a JavaScript terminal, to execute JS scripts in the shell. Java comes with a Command-Line utility called `jjs` which you can use in the same way as a Python or Ruby terminal. Here's one of my articles [explaining how to use Nashorn](http://www.sitepoint.com/introducing-nashorn-javascript-engine/) to execute JavaScript in the terminal and also call JS from Java code (and vice-versa).

Node.js is a way of executing JS applications in the server, using V8 as the JS engine. Combined with other technologies like MongoDB, it allows usage of JavaScript in the front-end, in the back-end and also in the database, providing a full-stack JavaScript solution. Node.js is based on an event-driven, non-blocking I/O model that allows it to be fast and achieve high performance. Projects like WordPress.com [have moved away of PHP and MySQL](https://developer.wordpress.com/calypso), building a REST API in Node.js which is then consumed by a Single Page Application (SPA) in the front-end. If you want to give it a try, you can read my [Introduction to the MEAN stack](http://www.sitepoint.com/introduction-to-mean-stack/).

ES is evolving and [its new standardization](http://www.ecma-international.org/ecma-262/6.0/) is called ES2015 or ES6. This new version introduces some more facilities that make ES development easier and more familiar to other developers coming from other languages like Java. ES6 will be implemented soon in V8 Engine, and so will be usable with Node.js. Nevertheless, there are some transpilers, like [Babel.js](https://babeljs.io/), that allow you write ES6 code for every browser and then compile it into ES5, the version supported by all modern browsers.

## #4: Why certain big companies (including Facebook) decide to use a languages like C/C++ along PHP?
The main reason is performance: compiled languages like C and C++ are faster in the terms of running time.

Initially, Facebook was written in PHP. The language is simple and very easy to get started with and is designed to build dynamic websites. This gives PHP a slight advantage. But there are some drawbacks too, mostly related to the fact that PHP is an _interpreted language_. Every interpreted language should have a runtime environment upon which the program is executed. On the other hand, _compiled languages_ do not have this runtime environment (Java is an exception, but it compensates with JIT compilation); they are executed by the machine, the bare metal. Having that added layer (the runtime environment), makes interpreted languages slower than compiled ones. This might not be a big issue in the beginning and you might think that scaling up will make the application faster even when using PHP, but scaling up means spending more money for additional hardware. In the beginning, that's what Facebook did. Then, they decided to try the other approach: compile some parts of their platform. The legend says that up to 75% of their servers were released from executing PHP code, since the compiled code required less resources and executed faster.

> I don't have any resource saying that the 75% thing is real, except what I've heard in a presentation in a PHP User Group

Now, Facebook uses not only C and C++ (through [warp](https://code.facebook.com/posts/476987592402291/under-the-hood-warp-a-fast-c-and-c-preprocessor/)), but even Java, Python, Erlang, Ruby, Python and other languages, some of them for internal usage and not for the products we use everyday. The mobile applications are built in Java for Android and Objective-C for iOS. Also Facebook gave us a new language called [Hack](http://hacklang.org/), which is more like PHP on steroids; they gave us [HHVM](https://code.facebook.com/projects/564433143613123/hhvm/) for faster PHP execution (although the PHP7 engine, called PHPNG might surpass that); [React](https://code.facebook.com/projects/176988925806765/react/) and [React Native](https://code.facebook.com/projects/450791118411445/react-native/), and [plenty more tools  for developers](https://code.facebook.com/projects). All those are used also by Facebook in their products.

If you're interested in what and how Facebook uses everything below the surface, read this Quora entry: [What is Facebook's architecture?](https://www.quora.com/What-is-Facebooks-architecture-6)


## #5: How does Siri work?
I will have two approaches here, one for the common iOS users and the other one for the rest of us.

### Siri for the common iOS user
Siri is one of the smartest things you know! It knows what you want to do and does it better than you. You can ask it everything and it will always give you the correct answer. But be warned, Siri is connected over the web. Even Pentagon is connected over the web. And probably, even The Red Button (yes, THAT red button) is connected with a computer over the web. So next time talking to Siri, think about Skynet and what it did in the end of Terminator 3 before making it angry.

![Siri = Hal](https://raw.githubusercontent.com/aziflaj/aziflaj.github.io/master/images/52-projects/week7/sirii-hal.jpg)

### The real face of Siri
Siri is a Personalized Assistant that Learns (PAL) which is not created by Apple. Siri is an acronym for _Speech Interpretation and Recognition Interface_. The guys behind SIRI ([SRI International](https://www.sri.com/)) wanted to port the application in other platforms like Android and Blackberry but after the acquisition by Apple, that didn't happen. 

When you say something to Siri, this is what happens under the cover:
- Siri uses ASR (Automatic Speech Recognition) to translate what you say into text
- That text is parsed using NLP (Natural Language Processing) techniques so the device can _"understand"_ what you're saying
- After understanding that, if what you're saying is a command (_e.g._ "Call Mom"), Siri performs the command required
- That step can be integrated with 3rd party services (like external APIs) if the command required can't be executed by the device only
- The output from the 3rd party can be transformed into a human readable response and using TTS (Text To Speech), it is read by Siri

Speech recognition is not a simple task, neither is natural language processing. It is easy for humans to recognize speech and process spoken languages, but that's because we have a brain (probably not true, but let's agree for the sake of the argument). In fact, one of the ways to make a computer recognize speech and process natural language is to give it a brain, with neurons and all the stuff. This is the job of Artificial Intelligence. And the AI scientists use Artificial Neural Networks (ANNs) to give machines a brain, teach them how to learn and act like humans.

Now, Siri is not one of a kind. It's not black magic or something, and other companies have done more or less the same thing. Microsoft has Cortana in its devices and Google has Google Now integrated in some of its services (like Chrome).

## #6: How to hack?
Don't!

Most of people asking this question don't really know the true meaning of hacking. If you want to learn how to attack the security of some kind of service, I'm not the right guy to ask. But that's not the meaning of hacking and a hacker is not someone who breaks into the database of a bank and steals $10M. I'm against that kind of "hacking" and I'd suggest everyone not to do it. It's illegal, it's criminal, and you'll probably get caught. 

Hacking (the good thing) started in the 60s in the MIT with the communities of computer enthusiasts and hobbyists, focusing on both hardware and software. It is the culture of individuals who enjoy the intellectual challenge of creatively overcoming and circumventing limitations of systems to achieve novel and clever outcomes. Some hackers are also programmers, but programming is not what divides hackers from the rest. Instead, it is the way of doing things, the playfulness during the process and the added value in the end of the process that makes a hacker different. Richard Stallman, the guy behind GNU project, said this about the hackers who program:

_What they had in common was mainly love of excellence and programming. They wanted to make their programs that they used be as good as they could. They also wanted to make them do neat things. They wanted to be able to do something in a more exciting way than anyone believed possible and show "Look how wonderful this is. I bet you didn't believe this could be done."_

Some well known hackers, who also happen to be programmers are:

- Linus Torvalds, the guy who gave you Linux and Git
- Dennis Ritchie and Ken Thompson, the creators of C 
- Steve Wozniak, the creator of Apple I and II and my favorite Steve of Apple
- Donald Knuth, the author of "The Art of Computer Programming"

Hacking is a culture that identifies all those who push the limits of what they know. Being the son of a farmer in the 1600s and discovering that bodies interact without any physical contact, that's hacking (yes, Newton was a hacker). 

When talking about hackers that deal with security breaching or testing, I choose to call the first kind "crackers" and the other "pentester" (short for "penetration tester"). They are also called black- and white-hat hacker (black-hat being the bad guy). The black-hat do mainly bad things to your system, like Distributed Denial of Service (DDoS), try to inject malicious code in your website with techniques called Cross-Site Scripting (XSS) or SQL Injection (SQLi), they might use different P2P networks (like torrent services) to install Remote Administration Tools (RATs) in your PC, etc. The other group, white-hats, do everything they can to protect you from these attacks. Being a white-hat means having all the skills of a black-hat (and probably even more) and also the ethical training not to use that skill for malicious purposes. They are like the cops of the internet, equipped with a Gatling gun, so try not to make them angry :-)

## #7: What should I learn next?
Being a software developer or any other kind of professional in the IT community means you should always learn. The technology changes day by day and you have to follow it. 

Learning everything is not bad, but it is not simple either. You may want to be a full-stack developer, writing code for the server and the browser and after that, going into mobile development. But the languages and the technologies are endless. There are thousands of programming languages today and I don't believe anyone who says "I know all the programming languages" (yes, someone told me that). I also don't take seriously anyone who wants to learn all the programming languages (I've been told that too). What you should do is start learning new things at your own pace. You can try new languages or explore new aspects of the languages you know, just like I'm doing with this challenge. This is a list of some languages, technologies or IT fields I'd recommend you to try. If you don't feel you should try them, don't. It's just a matter of opinion.

I'll start with **data analysis**. People like seeing infographics with numbers and results of different statistics and other similar things. The basics of data analysis: take a big dump of data and analyze it in order to extract useful information. Data analyst is one of the most required job positions in Silicon Valley, so if you're into statistics or mathematics or anything similar, you'll probably like it. Data analysis is closely related to **Big Data**. A good application of big data analysis would be a recommendation mechanism integrated in eBay, something that would recommend you to buy an iPhone protector case after you buy an iPhone. The applications are endless. Try learning Python, R or Matlab if you're thinking of a career in this field.

If you have been developing back-end applications as a web developer, it is now a good time to check out JavaScript. [Node.js](https://nodejs.org/en/) and [Meteor](https://www.meteor.com/) seem good alternatives. I've mentioned a couple of things about Node.js above in question #3. Meteor on the other hand allows you to run the same code both in the client and in the server. Using the same source code, you can run your app in a cloud service, in a web browser and in a mobile device. It comes with a command-line interface that helps you get started and manage the source code, and just like Node.js, it is reactive (asynchronous). For the UI, you can use [its rendering framework called Blaze](https://www.meteor.com/tutorials/blaze/creating-an-app), or integrate it with other frameworks like [AngularJS](https://www.meteor.com/tutorials/angular/creating-an-app) or [ReactJS](https://www.meteor.com/tutorials/react/creating-an-app).

More about JavaScript. If you haven't adapted a MVC framework yet, give AngularJS or EmberJS a look. They are stable, they have a massive community online and you can find help for everything you need. Also, both of them can easily be integrated with other back-end technologies, like Rails, Java or PHP. But if you are already familiar with any of them, take a look at ReactJS. It is very fast compared with other frameworks (including Angular). You can use it as the single framework in your next project, or integrate it with any other MV* framework like Angular. I'd recommend to try using React instead of Angular 1.x directives if your application will have to frequently update the DOM or you have to show very long lists (or similar) that will need some time to load. You can see here [some benchmarks with Angular and React](http://www.williambrownstreet.net/blog/2014/04/faster-angularjs-rendering-angularjs-and-reactjs/) and then decide when (not if) you will use them both.

Another milestone I'd like to have and probably you'd like that too, is learning a **functional programming (FP) language**. While all programming languages allow the developer use functions (or methods, procedures, subroutines, etc.), some of them don't support functional programming. Some FP concepts are introduced to Java 8 (with lambdas or anonymous functions), but yet the language is not functional. _Functional languages treat functions as first-class citizens._ You can have functions that take functions as arguments and/or return functions as results. These functions are called higher-order functions and if you have some mathematical background you'll now derivative and integral to be higher-order functions. FP languages also support pure functional functions, which comes with some interesting and useful properties. If some functions don't have any data dependency between each other, they can be performed in parallel without any issue. Recursion is another important topic when talking about FP languages. Some languages (like LISP) don't have support for loops like other languages, so you have to use _tail recursion_. 
If you want to learn any FP language, JavaScript can be the first language if you haven't already used it (also other benefits if you take it up now; read above). Other languages are LISP, Haskell or Erlang (Whatsapp and Facebook Messenger's back-end infrastructure is built upon Erlang). An interesting case would be Scala, which supports both Object-Oriented Programming and FP.

After talking about Siri and how it works, I don't think I could go on without mentioning a couple of words about Artificial Intelligence (AI). The days when AI was only for PhD research purposes are over. Nowadays, more and more businesses are using AI systems. The recommendation mechanism mentioned at the data analysis paragraph is built upon a concept called Machine Learning (ML), a branch of AI that teaches machines how to learn and make predictions. Siri, Google Now, Cortana are AI programs. Pattern recognition is used in Optical Character Recognition (OCR), in medicine and Computer-Aided Diagnosis. Now that Google open-sourced [TensorFlow](http://www.tensorflow.org/), I expect more and more people will go into the world of Artificial Intelligence. So if you are one of those people and want to know what languages you need for that, I'd recommend Python. TensorFlow is written in Python and the language is simple and general-purpose. LISP has been used for AI for a long time, alongside Prolog. But actually you can use every programming language you know to implement AI. Even C and C++ can be used, and Matlab has a module for Neural Networks which you can find interesting to work with.

And that's for the 7th week. I hope I taught you anything interesting. If you have any other question, please add it below in the comment section and I will try to answer as good as I can, just like I did with the questions above.

