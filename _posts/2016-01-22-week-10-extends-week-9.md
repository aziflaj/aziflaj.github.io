---
layout:     post
title:      "Week #10 extends Week #9"
date:       2016-01-22
comments:   true
summary:    It has been almost a month since the last blog post. School requires a lot of my time so I haven't been playing around starting from December 20. Now that I finished all of my school projects (and I'm getting ready for the exams season) I think it is better to write a blog post about what I've been doing lately.
category:
    - project52
    - java
    - php
    - android
---

It has been almost a month since the last blog post. School requires a lot of my time so I haven't been playing around  starting from December 20. Now that I finished all of my school projects (and I'm getting ready for the exams season) I think it is better to write a blog post about what I've been doing lately.

Firstly, I'm proud to say that I've changed my Github longest streak from 10 days to 17 days. It is not something really big actually, considering that I have more or less one week that I don't `commit` anything. But it means that I've been busy, and the whole purpose of my **Project 52** is to keep me busy programming and learning. I also had the chance to finally learn more Bash scripting. It isn't that hard; it is very useful when you're using Linux. And if you want to know more about my 10th week (which is actually the whole January) go on reading below.

## Model View Presenter implemented in Balut
If you can recall [from my last blogpost](http://aziflaj.github.io/week-9-java-nodejs-android/), I had to build a Yahtzee game in Java. Since I didn't know the game was actually called Yahtzee, I named the repository Balut ([aziflaj/Balut](http://github.com/aziflaj/Balut)). You can find how to build and run the game using Gradle in the readme of the repository (please read it, it's not that hard).

![Balut](https://raw.githubusercontent.com/aziflaj/aziflaj.github.io/master/images/52-projects/week10/balut.png)

The task was just to develop the game but I thought the game itself wasn't hard enough, so I decided to experiment with implementing a Model-View-Presenter pattern. The MVP pattern is a derivative of Model-View-Controller, where the Presenter layer defines the presentation logic of the application (e.g. what happens when a button is clicked). The game is now finished, using SQLite for storing all the records of the game and also uses JUnit for Unit Testing the scoring class. But there is a small issue now: I don't like the MVP I implemented, I could do it better:

- I think it would be a good idea creating three interfaces called `MvpModel`, `MvpView` and `MvpPresenter` and using them to create better wiring between the views and their presenters. 
- Using a Data Manager to connect with the SQLite database and also store the state of the game would also be a good idea. I could use this Data Manager to write and retrieve POJOs from the DB. Of course, the Data Manager would retrieve SQL records from the SQLite and then wrap them in Java objects (so it's not a real ORM).
- The `GameController` class is suposed to be a Singleton (and it is), but it is built poorly. Firstly, it has a `start()` method that actually calls the `getInstance()` method. I don't know why I did this but I don't like it anymore. I also don't want to change it because... reasons.

The MVP architecural pattern is very popular right now as the pattern of choice of Android applications. Android doesn't have a clear, defined architecture to follow when developing apps in order to divide the logic from the view. That's why some developers use this pattern to structure their application code.

## Procedural PHP, Wordpress Style
We had to develop a job searching application using PHP and MySQL in a subject called "Web Development". The subject covers the basics of PHP, like form processing, database connection other related stuff while using the XAMPP server. The issue is that we're using a really old book! It is a Deitel book published in 2008, when PHP was in its 4th version. Things were a lot different back then! Anyway, we agreed to use PHP 5.3 in the application and **no OOP at all**. Also, we had to use `mysql_` functions: the old, procedural, not safe way of connecting to PHP. 

Of course I can't agree to this! If you're going to use procedural PHP, at least do it right. I used the better, improved API for MySQL (`mysqli_`) and also used a Vagrant machine to host the project while development. If you want to know more about Vagrant, you can read my old blogpost about it, [V for Vagrant](https://aziflaj.github.io/v-for-vagrant/). The project is now on Github ([aziflaj/upt-web-dev](https://github.com/aziflaj/upt-web-dev)), with the whole files and the database schema.

This is also the first time I use [MySQL Workbench](https://www.mysql.com/products/workbench/). Workbench is a GUI tool for executing SQL queries in local or remote MySQL databases, create databases using Enhanced Entity-Relationship (EER) diagrams, etc. You can find the SQL script generated by Workbench in the repository, along with the EER diagram of the database.

As for Vagrant, I didn't build one of my own. Instead I used [Scotchbox](http://box.scotch.io/), the Vagrant config developed by [Scotch.io](http://scotch.io/). I realized I'm not a sysadmin and probably will never really need to configure Vagrant or other server machines, so I don't need to learn how to configure Vagrant and/or Docker. I didn't drop interest in these tools, I still think they're really good, just that I don't need to know how to configure them (more Dev, less Ops).

The application works, more or less. I tried not to make anything different from what the course covers. So every form in the application is built inside a table and that is probably the ugliest thing I've done in this application. We couldn't use Bootstrap for better styling, but I tried to use the SMACSS model of structuring the CSS and it helped me a lot while writing CSS (I still suck at CSS though). The code is so spaghetti, an italian chef wouldn't make a difference between it and [spaghetti alla puttanesca](https://en.wikipedia.org/wiki/Spaghetti_alla_puttanesca). Anyway, I tried to build something good (with all the restraints) and focus on structuring the application as better as I could. 

If you want to check it, head over at Github and let me know what you think about it ([aziflaj/upt-web-dev](https://github.com/aziflaj/upt-web-dev)).

## Parse on Android
In the last blogpost, I said I had built 4 Android apps in the course I'm following. Until now, I've build 6 more and the 7th is under development. The last application I developed is an Instagram clone which I called [Exchangeagram](https://github.com/aziflaj/AndroidCourse/tree/master/Exchangeagram) (The Internship strikes back). The interesting thing about this application was that I had to use the Parse API for Android to make it. Parse is a very interesting backend service which you can use as cloud storage on your web and mobile application. Building the backend of Instagram is not what I needed to do to develop the application, that's why we used Parse.

Parse allows you to communicate with the backend in a background thread without blocking the Main thread. It also uses callbacks in most of its API as Single Abstract Methods (SAM), like this:

{% highlight java %}
ParseUser.logInInBackground(username, password, new LogInCallback() {
    @Override
    public void done(ParseUser user, ParseException e) {
        if (user != null) {
            Toast.makeText(getApplicationContext(), "Logged in", Toast.LENGTH_SHORT).show();
            startActivity(new Intent(MainActivity.this, UserListActivity.class));
            finish();
        } else {
            Toast.makeText(getApplicationContext(), e.getMessage(), Toast.LENGTH_SHORT).show();
        }
    }
});
{% endhighlight %}

![Exchangeagram](https://raw.githubusercontent.com/aziflaj/aziflaj.github.io/master/images/52-projects/week10/exchangeagram.jpg)

After this one, the next app is an Uber clone, where I will have to use Parse SDK and its Location-oriented functionality. 

Finally, I also have been playing with some Bash scripting, learning the basics of awk usage and I've been liking it so far. In case you don't know, awk is a CLI utility in Linux which is used for data extraction from text messages. I normally use it when I kill processes, like Firefox when it stops responding:

{% highlight bash %}
$ kill `ps ax | grep firefox | grep -v grep | awk '{ print $1 }'`
{% endhighlight %}

And more or less that's it for this month. The next month will also be very silent for me. I will not go on with the challenge because of the finals. I will resume the challenge in March, as week 11, probably going on with more Android projects. Until then, stay classy.

