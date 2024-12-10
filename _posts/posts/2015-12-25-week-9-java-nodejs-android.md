---
layout:     post
title:      "Week #9: Java, Node.js and Android"
date:       2015-12-25
comments:   true
summary:    For the last 2-3 weeks I've been working with Node.js and Java. The projects are still work-in-progress, but I feel obliged to write the blog post and describe what I've been working on.
category:
    - project52
    - java
    - nodejs
    - mongo
    - android
---

It has been a while since my last blog post and the last challenge. But I have been working on a couple of projects in parallel, including school projects and personal projects so I didn't have enough time to write a blog post about the progress of the challenge. Anyway, these last 2-3 weeks I've been working with Node.js and Java. Even though the projects aren't yet finished, I feel obliged to write the blog post and describe what I've been working on.

As you can guess from the title, I've been spending my time writing a Java and Javascript, partly because of homeworks and partly because of side projects. One of the homeworks (more like home projects) was to develop an application; it doesn't matter what technologies or languages you use (it couldn't be more underspecified). The other one was to develop a Java game that is similar to [Yahtzee](https://en.wikipedia.org/wiki/Yahtzee) or [Balut](https://en.wikipedia.org/wiki/Balut_%28game%29) (I never played any of them ever so I might be wrong). 

## The Node.js Application
The underspecified project was actually a good idea; it was given in the Software Engineering course and its purpose is to check the ability to work in groups and to pick the right tool for the job. So we decided to create an application called **Restauranteers**. Restauranteers is an application that allows users (i.e. customers) to find the best places to hang out, based on reviews and ranking. We decided to build this application using Node.js for the back-end, MongoDB as the database, and possibly later build any mobile apps, whether native or hybrid (using AngularJS and Ionic Framework).

We chose [MongoDB](https://www.mongodb.org/) because of the simplicity of it: the application isn't designed to be the foundation of a future business (but you never know), so we didn't want to care that much about the database schema and normalization. Using MongoDB, we can store records in the same way as we can structure JSON objects. We also decided to use [Sails.js](http://sailsjs.org/), a MVC framework for Node.js. The front-end will initially be in jQuery and Bootstrap, but maybe we can switch to AngularJS. 

The application will have 4 types of users: 
- Customers, who can write reviews about different places and also give a rating
- Managers, who can create and update information about their restaurants (or clubs, bars, etc.), like the description, the address, their menu, etc.
- Administrators, who can create new restaurants and assign their unique manager (after their request)
- Developers, who need a Customer account and can be equipped with an API key, which then can be used to access the RESTful API of the application.

As I said before, the application is not yet finished. We are still working on it, and we expect that in the end of January most of it will be finished. The source code can be found on Github ([aziflaj/restauranteers](https://github.com/aziflaj/restauranteers)) and we will also publish it on Heroku (you can find the link on the repo).

## The Java Game
So this year we have an Object-Oriented Programming course, where we use Java to learn OOP concepts. Although, the course seems more like a Java course than an OOP course. In my opinion, an OOP course should teach the base concepts of OOP (encapsulation, inheritance and polymorphism) which are language-independent. When you're learning the Java Swing API, while being an application of OOP, you are learning more Java than OOP. Anyway, the project we had in this course was to develop a game that is similar to Yahtzee or Balut. I learned the name of the game after reading an article on Wikipedia about [all the dice games](https://en.wikipedia.org/wiki/List_of_dice_games). Balut was the first on the list, so I wrote that as the name of the application, the main package and the Github repository ([aziflaj/Balut](https://github.com/aziflaj/Balut)), but after reading about [Yahtzee on Wikipedia](https://en.wikipedia.org/wiki/Yahtzee), I realized I named things wrong (something you'd expect by someone who never played these games). 

I also found out that the same game is given as a homework in Stanford. But they provide the UI to the students and some other parts of the application as _jar_ archives, while we were required to build everything by ourselves. Anyway, I decided to try and structure the application in a Model-View-Presenter model. I put the models (`Dice.java` and `Player.java`) in a `model` package; I created an interface and an interface implementation for all 5 dice and the player and put them in a `presenter` package; I divided the UI in panels and put each panel in a different package, and all these UI-related packages in a `view` package. I also added a `utils` package with utility classes, like the `ScoreHelper.java` which calculates the points of each dice roll, or the `ResourceHelper.java` which gets the dice faces from the `resources` folder and returns their URL so they can be rendered in the UI. Another package in `utils`, the `db` package, holds classes for connecting with a SQLite database to store the result of each game and also some information about the players. In the assignment, we are required to store information about each player: their name and age. The `DatabaseOpenHelper.java` is a singleton class that in the first call creates the database and the necessary tables if they aren't already created.

![Did I do it OK?](https://i.imgflip.com/wbvw9.jpg)

After a while, it became a bit hard thinking about adding other features or tracking bugs. When I decided to add a `GameController.java`, bugs started dancing around the application like moths around a light. I started adding public methods to each view to get their presenters, so then they can be required at the upmost class that makes the supervision of the game (that is, `GameController`). I think maybe I didn't apply MVP correctly, or maybe using a Dependency Injector since the beginning would allow me to require presenters in a simpler way rather than adding getter methods at each view. Even though the game is almost finished and I'm not thinking of rewriting it using a proper MVP or DI, I'd like to know your opinion about what I did and also any suggestion for future reference.

## Android as a Side Project
Since last month, I started following an Android course by [Rob Percival](https://twitter.com/techedrob). This week I developed four Android apps (one every night) called "Basic Phrases", "Egg Timer", "Brain Trainer" and "Guess the Celebrity", which you can all find on my Github account ([aziflaj/AndroidCourse](https://github.com/aziflaj/AndroidCourse)). Basic Phrases speaks out some phrases in French when you tap a button. Egg Timer is a countdown timer up to 10 minutes, where you can use a seekbar to speficy the countdown time; it will play a sound when the time is up. Brain Trainer checks how many additions you can solve in 30 seconds. Guess the Celebrity fetches a list of 100 celebrities and their photo from a website and asks the user to find the name of the celebrity. What I don't like about this app is that it downloads the HTML of the website and then parses it to find the celebrities and their photo, so in case the HTML of the website changes, the application will have to change too. A better idea would be to use some kind of 3rd-party API that returns XML or JSON responses.

![Android apps](https://github.com/aziflaj/aziflaj.github.io/blob/main/assets/images/20151225/android.png?raw=true)

The apps are simple, only one Activity per each of them. That's because I finished only 36% of the course, and the best apps are yet to come. If you're curious, some of the applications that I'll develop in the course are some clone apps for Twitter, Instagram, Uber, Flappy Birds and Snapchat, and of course other simple applications using maps or third-party APIs. Even though this is not the first time I develop Android application and I have some experience writing Java code, the course is structured in a way that you don't need previous experience in Java or Android to follow it.

And that's what I've been up to, in case you were wondering. Finally, thank you for reading my blog on this Christmas day! I wish you a Merry Christmas and thank you for reading my blog. 
