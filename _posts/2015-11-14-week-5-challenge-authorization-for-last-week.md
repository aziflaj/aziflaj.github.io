---
layout:     post
title:      "Week #5 challenge: Authorization for the Last Week Project"
date:       2015-11-14
comments:   true
summary:    "Fifth week already! If your remember the project I developed last week, you probably recall that I didn't add some things to it. For example, everyone could just delete everything from the database, just like everyone could store anything. This week, I added a simple authorization mechanism, based on API keys, so not everyone could have access to everything."
tags:       [challenge, node, javascript, mvc, framework, mongo, mongodb, validation, authorization, authentication]
---

Fifth week already! If your remember [the project I developed last week](https://aziflaj.github.io/week-4-challenge-server-side-javascript/), you probably recall that I didn't add some things to it. For example, everyone could just delete everything from the database, just like everyone could store anything. This week, I added a simple authorization mechanism, based on API keys, so not everyone could have access to everything. 

I started the week by adding the Validator service. What this validator does is a simple check if all the required parameters are set for the ship to store in the database. If any of the parameters isn't set, it doesn't allow the application to store anything in the DB. I also added two users, one authorized as admin and the other one as a simple editor. Only the admin has the access to deleting records, both the admin and the editor can create and update records, and everyone (even unauthorized users) can see the records. 

The whole authorization is based into a simple API key added in the URL when the HTTP request is sent to the server. So the whole HTTP request should be sent to `api/v1/ships?api_key=<YOUR-API-KEY-HERE>`. I generated the API keys by md5-ing the concatenated username and password of the user. E.g. the admin has the username `admin` and the password `admin`, so the API key is `md5(adminadmin)`. These API keys are stored into the database, so everytime a request is sent, this API key is checked in the DB and the application checks if the request should be executed or not. 

I also changed some shell scripts and divided them in two folders, `v1` (the scripts writen in the 4th week) and `v2` (the scripts writen in the 5th) week. The scripts from the 4th week are changed in the meaning that now they require the API key to be passed as command-line parameter. Also now you don't have to `POST` the data into the DB, but you can directly import them into Mongo by executing the `shells/v2/mongoseed.sh` script. 

There are also more things that may be added to the same application. I think I should add some UI or a front-end client, maybe after I learn some Ember.js. Also, it would be a good idea managing the ships from a UI, using sessions (and cookies?) instead of API keys. While API keys might be a good idea when building mobile apps and other API clients, sessions are simpler for most users. You have to login and that it, no API keys to memorize.

And this is for the 5th week. You can find the source code in [the same Github repository](https://github.com/aziflaj/Sailor) as the previous week. I added a branch for the 4th week so you can still access the source code of the last week.
