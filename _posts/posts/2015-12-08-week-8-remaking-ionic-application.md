---
layout:     post
title:      "Week #8: Remaking an Ionic Application"
date:       2015-12-08
comments:   true
summary:    This week I went back to one of the frameworks I like in the hybrid app development "field", the Ionic Framework, and built a logo game using it"
category:
    - project52
    - javascript
    - angularjs
---

Native mobile development is an in-demand skill you should have, but it has its drawbacks as everything else. Of course, the performance is as good as possible, but you can't simply port the application from one platform to the other. This is where hybrid mobile apps come into play. They probably don't have the same high-performance as native apps, but using the same code-base you can compile applications for ALL the platforms. This week, I went back into one of the frameworks I like in the hybrid app development "field", the [Ionic Framework](http://ionicframework.com/).

> I won't talk much about how native development differs from hybrid development so if you are interested in that, please read this other article: [http://www.sitepoint.com/native-vs-hybrid-app-development/](http://www.sitepoint.com/native-vs-hybrid-app-development/)

Ionic is a framework that allows the developer build Android and iOS applications using AngularJS, or at least this was in the beginning. Now it goes really beyond that and gives the developer way more power than any other similar solution. The framework is based on [Cordova](http://cordova.apache.org/) so basically the applications are simply(?) Cordova applications (notice the question mark). Over tiame, Ionic became the #1 framework and platform for hybrid mobile applications, providing tools that every developer would love. The platform now provides analytics and push notifications. You can build applications by simply dragging and dropping UI widgets around in the [Ionic Creator](http://ionic.io/products/creator), and you can manage building and emulation in the [Ionic Lab](http://lab.ionic.io/). If you need to use device resources, like the embedded SQLite database, the camera or geolocation, [ngCordova](http://ngcordova.com/) will be of help. These are Cordova plugins wrapped in Angular services and factories that you can easily use in your application.

When I first got my hands dirty with Ionic, it didn't have all these great tools. Still, it was simple to learn and understand. I used it once to show how a logo game can be build for Firefox OS, and the article is on SitePoint ([Part 1](http://www.sitepoint.com/firefox-os-game-development-ionic-framework/) and [Part 2](http://www.sitepoint.com/firefox-os-game-development-ionic-framework-part-2/)). Now I decided to re-write the same application for some reasons:

- I wanted to try some of the new tools, like the Ionic Creator and the Ionic Lab
- Angular is getting its major update into v2, and [so is Ionic](http://blog.ionic.io/announcing-ionic-2-0-alpha/). So I will write the same application once again, to compare v2 and v1.
- I've never used ngCordova, so now I have a playground for it

I started using Creator to build the UI, and most of the final UI is the one built with this tool. It is very easy to use, just drag a widget and drop it where you need it. This is a great help for people like me who still suck at CSS and styling things. After building most of the UI, you can download the project and start making necessary addition and changes in the code. The generated code will include some modules: the main `app` module, a module for controllers, a module for directives, another one for services and factories, and a final one for routes. Every module is a single file, but I didn't like this structure so I changed it a bit, following the [Angular style guidelines](https://github.com/johnpapa/angular-styleguide) by [@john_papa](https://twitter.com/john_papa): one file per component, components wrapped in IIFEs, naming conventions, etc.

For the data, I kept the same technique as in the original version (which you can find [in GitHhub (aziflaj/yalg)](http://github.com/aziflaj/yalg)) and stored them into a JSON file. Also the structure of this file is not very different from the first version.

The game, which you can [find it here](https://github.com/aziflaj/LogoGame/), is not finished. I mean, it is more or less functional, but it is not production-ready and I don't intend to upload it on Play Store or App Store. This is just a playground for me and anyone else who wants a place to test their skills with the Ionic platform. If you want to make it production ready, you should firstly add more images and probably even change the ones which are already there. A good game would require more levels, scoring and sharing scores in social media, etc. I will probably add some of these features while still keeping this as a proof of concept.

And that's for the 8th week! I'm a bit late with this because during this week I've been working on some side projects, including Vagrant, Chef and Laravel, but maybe one day I will have a weekly challenge including those too. Until then, stay classy San Diego :)
