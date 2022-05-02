---
layout:     post
title:      "Week #2 report: AngularJS + Gulp"
date:       2015-10-24
comments:   true
summary:    On the second week of my "52 weeks, 52 projects" challenge, I went back to re-learning my favorite JavaScript framework, AngularJS.
category:
    - project52
    - angularjs
    - sass
    - gulp
    - build tools
---

On the second week of my [**52 weeks, 52 projects**](http://aziflaj.github.io/52-weeks-52-projects/) challenge, I went back to re-learning my favorite JavaScript framework, AngularJS.

This week, I thought of making a simple invoice application, that lets the user to create an invoice, add items to the invoice and creates a PDF version of that. It is not something new in the market. In fact, I tried to remake [this invoice generator](http://invoice-generator.com/), trying to make the UI as similar as that. It is not a copycat though; I'm not that good at CSS. Instead, I focused on the Angular part and another tool called **Gulp.js**.

## AngularJS
You probably already know what AngularJS is: a Superheroic JavaScript framework. I learned it more or less one year ago, and since then I've been using it whenever I could instead of vanilla JS and/or jQuery. This time, I tried to follow the [AngularJS style guideline](https://github.com/johnpapa/angular-styleguide) by [John Papa](http://johnpapa.net/). This guideline shows almost everything you need to know and apply when you create an AngularJS application, like:

- How to name files
- How to make Controllers and Directives
- When to use Services and Factories
- How to use IIFEs to wrap every component of the application
- [RTFM for more](https://github.com/johnpapa/angular-styleguide)

It is actually the first time I use IIFEs (Immediately-Invoked Function Expression) in a JS application; I never knew why to use them. If you don't already know, they're used to avoid polluting the global scope; probably there is already a dependency using the same variable name as you.

## Gulp.js
For starters, Gulp is a task automation tool. Just like Maven, Gradle and similar other tools, it makes sure to automate some tasks for the developer, simplifying work and easing the process.

I used it for tasks like minifying and concatenating JavaScript files, and also for compiling, minifying and concatenating SASS files. This is a good web development practice because it makes the assets smaller and the website can be served quicker. In fact, Ruby on Rails uses the same technique for serving websites and web applications faster.

> Also, Rails appends a hash value to the name of the compiled files, which is created upon the content of the files. So the files are cached in the user browser and re-cached when the hash value changes (when the content of the files changes).

```js
var gulp = require('gulp'),
    uglify = require('gulp-uglify'),
    concat = require('gulp-concat'),
    jscs = require('gulp-jscs'),
    jshint = require('gulp-jshint');
    
gulp.task('js', function() {
  gulp.src(['./js/**/*.js'])
      .pipe(jshint())
      .pipe(jscs())
      .pipe(jscs.reporter())
      .pipe(uglify())
      .pipe(concat('all.js'))
      .pipe(gulp.dest('./build'));
});
```

This is one of the tasks I wrote, that helps me follow some [JavaScript Coding Style](http://jscs.info/) (JSCS), minify/uglify the files and then concatenate them into one single, minified file called `build/app.js`. Is this file which is included in the HTML page. The same thing applies to SASS. Also, I used Gulp to create a live-reload service for the application: whenever I changed anything in the JS or SASS files, the page is automatically refreshed and the view is updated. It can be very helpful when developing with more than one screen: one for writing code and the other one for actually seeing the application running.

Probably later I will add a back-end service to this UI. Right now, I'm thinking about using Node.js or maybe Rails (or any other Ruby framework), but even Java might be an alternative, with [Spring Framework](https://spring.io/) or [Spark](http://sparkjava.com/). 

And that’s it for the second week! You can find the source code in [this GitHub repository](https://github.com/aziflaj/simple-invoices-frontend), you can fork it and change anything if you want to. 
If you want to ask me anything, leave a comment below.
