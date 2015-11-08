---
layout:     post
title:      "Week #4 challenge: Server-side JavaScript"
date:       2015-11-08
summary:    "This week on my 52 weeks, 52 projects challenge I tried a back-end JavaScript framework called Sails.js. Sails is an MVC framework based on Express.js and socket.io that resembles the MVC behavior of frameworks like Ruby on Rails. I also used MongoDB, a NoSQL database, to store the data."
tags:       [challenge, node, javascript, mvc, framework, mongo, mongodb]
---

This week on my **52 weeks, 52 projects** challenge I tried a back-end JavaScript framework called [Sails.js](http://sailsjs.org/). Sails is an MVC framework based on [Express.js](http://expressjs.com/) and [socket.io](http://socket.io/) that resembles the MVC behavior of frameworks like [Ruby on Rails](http://rubyonrails.org/). I also used MongoDB, a NoSQL database, to store the data.

You've probably heard before that it is possible to use JavaScript on the server-side just like you use Rails, [Spring](http://spring.io/), PHP, etc. This is possible because there are projects like [Node.js](https://nodejs.org/en/) that allow the developer use JavaScript to run in the backend of a web service. Node.js uses a non-blocking, event-driven I/O model that makes your web application/service lightweight and efficient.

Using Sails it's not the only way to develop Node.js applications. In fact, there's a full-stack JavaScript solution called **MEAN**, which stands for Mongo, Express.js, [Angular.js](http://angularjs.org/) and Node.js. Since MongoDB actually stores data in the BSON (Binary JSON) format, you can say that the MEAN stack is a way of using JavaScript for (almost) everything in your web applictation. This week's project is not a MEAN application: it doesn't have the Angular part (yet). But since Sails uses Express.js under the cover and I used MongoDB, we can agree that it is a MEN application (I hope there isn't any feminist unfollowing me for that).

During this week I created a demo API using Sails. It is not something very special to be done during a week, but I didn't really work too much on it. During the past weeks I used to work 3 hours/day on the projects, but this week wasn't the best. One or two days I couldn't work at all on this. Anyway, it gives the basic idea of building a simple API in Sails. It handles five routes, all sharing the base `api/v1/ship` route:

| Method | Route             | Description  |
|:------ |:------------------| :-----|
| GET    | api/v1/ship/      | Get the list of all ships stored |
| GET    | api/v1/ship/:id   | Get the data for a single ship with the ID of `id` |
| POST   | api/v1/ship/      | Insert a ship in the database |
| PUT    | api/v1/ship/:id   | Update data for a ship with the given ID of `id` |
| DELETE | api/v1/ship/:id   | Delete the ship with ID of `id` |

The data is stored in a Mongo database. Sails comes with a simple-to-use ORM/ODM called [Waterline](http://sailsjs.org/documentation/concepts/models-and-orm), which makes it very simple to use any kind of database management system, both SQL and NoSQL. If you need to change the DB management system, all you have to do is install the proper database driver, change the configurations file, and that's it! The interface that Waterline gives to you is the same and DB-independent, so the rest of the code remains the same.

Unfortunately, I didn't add any kind of validation or authentication/authorization mechanism, which is not cool. Next week, I intend to add such things, both a validation mechanism (so the data posted is the right one) and an authentication/authorization mechanism (so not all users can delete and/or update data). Maybe in a later time, I'll use the same application for playing with some Android development, probably using Reactive Extensions for Java.

This is for the 4th week. If you want to check the application, [you can find it on GitHub](https://github.com/aziflaj/Sailor) with other information on how to run the application. I included the exported data from MongoDB, so you can use the same data if you want. Also I added some shell scripts that I used to test `POST`-ing, `PUT`-ing and `DELETE`-ing.

If you're interested in learning more about the MEAN stack, you can read this [Introduction to the MEAN Stack](http://www.sitepoint.com/introduction-to-mean-stack/) or check out my [Intro to MEAN](https://aziflaj.github.io/presentations/intro-to-mean/#/) presentation.
