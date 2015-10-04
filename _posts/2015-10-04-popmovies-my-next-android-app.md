---
layout:     post
title:      PopMovies: my next Android app
date:       2015-10-05
summary:    A couple of months ago I took a course in Udacity called "Developing Android Apps". It is developed by Google and it's the first part of the Android Nanodegree at Udacity. The final project was building a cloud-connected app to allow users to discover the most popular movies playing, a really good experience.
tags:       [android, ionic, kotlin, google, udacity]
---


A couple of months ago I took a course in Udacity called "[Developing Android Apps](https://www.udacity.com/course/developing-android-apps--ud853)". It is developed by Google and it's the first part of the [Android Nanodegree](https://www.udacity.com/course/android-developer-nanodegree--nd801) at Udacity. The final project was building a cloud-connected app to allow users to discover the most popular movies playing, a really good experience.

## In the beginning...
... there was Java. You probably don't know, but firstly [I was one of the Java haters](https://aziflaj.github.io/old-programming-jokes-what-is-wrong-with-them). Nevertheless, I gave it a try and now it is my favorite programming language. It's not a surprise now, I wanted to learn Android development the first thing after I finished learning how to build Swing GUIs three years ago. But using an old PC (Windows XP, Pentium 4, 512MB of RAM) with Eclipse and ADT plugin was not the best thing to do. I'm really surprised how that PC was able to run [Netbeans IDE](https://netbeans.org/) ([miminum requirements](https://netbeans.org/community/releases/70/relnotes.html#system_requirements) though). Anyway, I left the Android thing aside and went on with other technologies, like PHP and Javascript. 

After developing [ApPiano](https://marketplace.firefox.com/app/appiano/), I felt the urge to go back into mobile development. That's when I tried [Ionic Framework](http://ionicframework.com/). Now it has become better and more supported than before, since Android's WebView is based on Chromium. It is really nice, and through [ngCordova](http://ngcordova.com/) you can use the native APIs in a really simple way.

Anyway, nothing beats Native! The performance is way better and it has out of the box support for all native APIs. That's why I started the course at Udacity.

> If you want to know more about Hybrid and Native apps, [read my article at SitePoint](http://www.sitepoint.com/native-vs-hybrid-app-development/). 

## What Movie to Watch Next?
> [Hot Fuzz](http://www.imdb.com/title/tt0425112/) is nice

So after building the Sunshine app, I had to build this application that showed the most popular movies playing now. The movies were to be fetched by a cloud service called [The Movie Database (TMDb)](https://www.themoviedb.org/) and showed to the user. Quite a simple task if you think about it. Anyway, it included most of the core Android skills:

- Using XML for layouts
- Fetching data from a cloud service out of the UI thread
- CRUD operations in the embedded SQLite
- Writing `ContentProvider`s
- Unit testing
- Syncing through a thread-safe Service
- Showing notifications
- Writing Adapters (mostly [CursorAdapter](https://developer.android.com/reference/android/widget/CursorAdapter.html)s)

I tried to develop it by using the Android framework as much as I could, and then thinking of using third-party libraries. But sometimes it is better to use third-party libraries: Nobody wants to reinvent the wheel! The Android framework is capable of doing everything you need to do, but sometimes it can be tiring and frustrating. Third-parties simplify the work and increase code readability. That's why I used:

- [Picasso](https://square.github.io/picasso/) for image loading, 
- [Retrofit](https://square.github.io/retrofit/) for HTTP client; it turns an HTTP API into a Java interface
- [Gson](https://google.github.io/gson/apidocs/) for fetching Java Objects from the JSON response from the server
- [CWAC MergeAdapter](https://github.com/commonsguy/cwac-merge) for populating a ListView from different `CursorAdapter`s

Some other third-party libraries I'd like to use are [RxAndroid](https://github.com/ReactiveX/RxAndroid) for Reactive programming and any ORM like [GreenDAO](http://greendao-orm.com/) or [Realm](https://realm.io/docs/java/latest/) (still not stable though). Also, it seems a good idea using [Kotlin](http://kotlinlang.org/docs/tutorials/kotlin-android.html) or [Xtend](https://www.eclipse.org/xtend/), just like Swing for Android. Probably my next Android application will be built upon [AndroidBoilerplate](https://github.com/hitherejoe/Android-Boilerplate).

The source code can be found [on GitHub](https://github.com/aziflaj/PopMovies). I hoe it will serve to me (and probably to anyone else) as a reference in the future when developing Android apps. Also, there are also other resources about Android development I'd like to share. You can read these [5 Best Resources for Android Developers](http://www.sitepoint.com/5-resources-for-android-developers/) on SitePoint.

_If you have anything to add, any suggestion or question, comment below and let me know_
