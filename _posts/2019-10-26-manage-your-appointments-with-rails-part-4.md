---
layout: post
title: Manage your appointments with Rails, Part 4
date: '2019-10-26'
summary: Adding Stimulus to our Rails application
icon: icons/rails.png
comments: true
category:
  - rails
---

If you do a Google search for the best front-end framework, you'll get a list of big names such as React, Vue, Angular, etc. Frameworks like Backbone and Ember, which were among the first front-end frameworks, are now eclipsed by React and other similar JS packages. Funnily enough, React isn't a framework :)

![Best front-end framework 2019](https://cdn-images-1.medium.com/max/1600/1*n0QLVdzXzBZYw_TaLRO3cg.png)

The JS community has gone through a handful of build tools, starting with [grunt.js](https://gruntjs.com/), but it seems as the de-facto standard now is [Webpack](https://webpack.js.org/). Rails decided to adapt to this new era of JS application from version 5, and introduced a gem called [Webpacker](https://github.com/rails/webpacker). The framework which came with jQuery out of the box for a really long time, now gave its users the possibility to build more complex applications and to follow the new trends from the JS community.

Some time ago, [Stimulus.js](https://stimulusjs.org/) came by, built by the same people that gave you Rails. It isn't a framework, it's just a JS library that, according to the docs, _"augments your HTML with just enough behavior to make it shine"_. So let's see how both these younglings play together with Rails.

## Install Stimulus

Fire up your terminal and run:

```bash
$ rails webpacker:install:stimulus
```

And that's it. See you on Part #5 :)

Unlike other JS frameworks/libraries, Stimulus isn't something you can use to manage your whole front-end; it makes an automatic connection between a JS class and an HTML element and lets you add some functionality to that HTML element through the class.

![](https://cdn-images-1.medium.com/max/1600/1*V0aPEt6lZ64qP-3lmdjCfQ.jpeg)

I added this controller to show free slots:

```js
// app/javascript/controllers/free_slot_controller.js
import { Controller } from "stimulus"

export default class extends Controller {
  template = `<li>Free slot at $time</li>`;
  dateOptions = {hour: '2-digit', minute: '2-digit', day: '2-digit', month: 'short', year: 'numeric'};

  connect() {
    this.load()
  }

  load() {
    fetch('/slots/free')
      .then(response => response.json())
      .then(data => this.jsonToHtml(data))
      .then(html => {
        this.element.innerHTML = html
      })
  }

  jsonToHtml(data) {
    return data.map(item => {
      const formattedDate = (new Date(item.scheduled_at)).toLocaleString('en-US', this.dateOptions)
      return this.template.replace('$time', formattedDate)
    }).join('');
  }
}
```

In order to bind this controller to an HTML Element, we use the `data-controller` attribute:

```html
<ul class="free-slots" data-controller="free-slots">
</ul>
```

The `connect()` method is a lifecycle callback which is called when the controller is connected to the DOM. In this method we fetch the free slots from the app and format them into HTML to show them to the user.

The same thing goes for the upcomming meetings as well:

```js
// app/javascript/controllers/upcoming_meetings_controller.js
import { Controller } from "stimulus"

export default class extends Controller {
  template = `<li>Meeting with <a href="mailto:$guest_email">$guest_email</a> at $time</li>`;
  dateOptions = {hour: '2-digit', minute: '2-digit', day: '2-digit', month: 'short', year: 'numeric'};

  connect() {
    this.load()
  }

  load() {
    fetch('/slots/upcoming')
      .then(response => response.json())
      .then(data => this.jsonToHtml(data))
      .then(html => {
        this.element.innerHTML = html
      })
  }

  jsonToHtml(data) {
    return data.map(item => {
      const formattedDate = (new Date(item.scheduled_at)).toLocaleString('en-US', this.dateOptions)
      return this.template.replace('$time', formattedDate)
        .replace(/\$guest_email/g, item.guest_email)
    }).join('');
  }
}
```

Since Stimulus.js doesn't support templates (yet), there are two different approaches to using templates as in every other JS frontend:

 - _The Rails Way_: You create HTML partials and respond with these partials instead of JSON objects. The Stimulus controllers will then put the server rendered HTML in the right place inside the application.
 - _The JS Way_: Similar to what I did above; you can create JS string templates, sprinkle some variables on top of it (like `$guest_email` in the above example), and then substitute the variable names with the real values sent by the server.

## That's All Folks

There is much more to Rails that what I wrote in these 4 blog posts. Building applications can sometimes get so complex, you need to break your Rails monolith into smaller, more modular applications that work together for a single quest: To make the customer happy. Going into the monolith breaking topic, you might want to take a look at [The Modular Monolith](https://medium.com/@dan_manges/the-modular-monolith-rails-architecture-fb1023826fc4). 