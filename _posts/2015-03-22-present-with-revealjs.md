---
layout: post
title: Present with Reveal.JS
comments:   true
summary: If you think using HTML, CSS and Javascript can't be used for making presentations better than PowerPoint, think again! Reveal.JS is the framework that makes PowerPoint irrelevant.
category: 
    - JavaScript
    - Reveal.js
---

<p>
Everybody uses presentations. Professors at our school use them to make lectures less boring than they actually are. We at <a href="http://www.slideshare.net/thingslab/" target="_blank">Things Lab</a> use them to show what we are talking about, and we make sure the presentations are as verbose as they should be so the reader can understand what it is about even without being present at the Lab during the presentation. You would normally use Microsoft Office's PowerPoint, or <a href="http://www.libreoffice.org/discover/impress/" target="_blank">LibreOffice's Impress</a> if you are an Open Source Evangelist. But of course, there is another way, which is actually better in my opinion. If you can use <a href="https://github.com/showcases/javascript-game-engines" target="_blank">Javascript to write games</a>, hell yes you can use it to make presentations!
</p>

For the last presentation at Things Lab, I used an HTML Presentation Framework called <a href="https://github.com/hakimel/reveal.js" target="_blank">Reveal.JS</a>. It is a really neat presentation-making tool which uses HTML, CSS and Javascript to make your presentations look good in the browser. For me, it is easier to write HTML than use a WYSIWYG interface such as Powerpoint, so I really liked Reveal.JS.

You start by cloning the github repo of Reveal.JS or by downloading the files. You will probably not have to use all the files. In fact, I removed most of them. All I kept were <code>index.html</code> and the folders <code>css</code>, <code>js</code>, <code>lib</code> and <code>plugin</code>. All the rest, such as <code>test</code>, <code>.travis.yml</code>, <code>Gruntfile.js</code>, etc, were not needed for my presentation.

Every slide of a Reveal.JS presentation is a <code>section</code>. You can place more than one <code>section</code> inside a <code>section</code>, and this would make your slides show one below the other (using the up-down arrow to show them). It is rich with transition animations and themes. You can also use Markdown syntax to write your slides. If you feel you want to upload the presentation into Slideshare, Reveal's got you covered. By adding <code>/print-pdf</code> at the URL, the presentation becomes print-friendly, and you can save it as PDF.

If you are not a coder and/or HTML is not your thing, you can still use Reveal.JS. There is a visual editor, called <a href="http://slides.com/">slides.com</a>, which could help you make your presentations in a simple, web-based service.

You can find the presentation <a href="http://aziflaj.github.io/real-world-webapp/" target="_blank">here</a>.

<i>How about you? Have you ever used Reveal.JS or any other alternative? Share it in the comment section.</i>
