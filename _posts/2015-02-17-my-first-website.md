---
layout:     post
title:      Bulding my first website <br> with HTML and CSS
date:       2015-02-17
summary:    Yesterday I finished a short course at Learnable titled Build Your First Website - HTML & CSS
tags:       [HTML, CSS, Learnable, website]
---

<p>
Yesterday I finished a <a href="http://goo.gl/6PuJD4" target="_blank">short course</a> at Learnable titled "Build Your First Website: HTML & CSS". I wasn't a complete noob on HTML and CSS but my CSS skills were not good. I had to check StackOverflow for any simple line of CSS. So I started the course at Learnable, hoping it would be the last "Beginner CSS" course.
</p>

What I liked most about this course was its practicality: you would have to build a website just like <a href="http://www.buildyourfirst.website/#" target="_blank">this one</a>. I learned a lot of stuff during that course. I finally learned why people use `meta` tags on HTML and why this piece of code is everwhere:

{% highlight html %}
<!-- [if il IE 9]>
  <script src="assets/js/html5.js"></script>
  <script src="assets/js/respond.js"></script>
<![endif]-->
{% endhighlight %}

Also I learned what <a href="https://smacss.com/" target="_blank">SMACSS</a> (Scalable and Modular Architecture for CSS) says about <a href="https://smacss.com/book/categorizing" target="_blank">categorizing CSS rules</a> in five different categories:

<ul>
<li>Base - <code>base.css</code> is used for the default, base rules of CSS.</li>
<li>Layout - <code>layout.css</code> is used to divide the view into sections, such as rows and/or columns.</li>
<li>Modules - <code>modules.css</code> holds the CSS for the modules of the view, such as features of a product, social media links, etc.</li>
<li>State - <code>state.css</code> defines how the modules or layouts will look when in a particular state.</li>
<li>Theme - <code>theme.css</code> defines how modules or layouts might look</li>
</ul>

Though, only 3 of them were used in this course: Base, Layout and Modules. Also, there is a 4th CSS file, called `normalize.css`. <a href="http://necolas.github.io/normalize.css/" target="_blank">Normalize.css</a> makes browsers render all elements more consistently and in line with modern standards. It precisely targets only the styles that need normalizing.

Another nice thing I learned was the usage of CSS sprites - merging some images into only one file, <a href="https://github.com/aziflaj/first-website-html-css/blob/master/assets/img/icon-sprite.png" target="_blank">like this</a>. I have heard of them before, I was aware of their existence, but this was the first time I used them. Even though merging some of your images into a single image file results in a file size bigger than the sum of the sizes of smaller images, it also results in less HTTP requests to download the images, thus in a webpage that loads faster.

In general, it was fun learning CSS through this course. I can't say that I'm a CSS master now. I know more than before, but I still qualify myself as a CSS noob. Thankfully, I pushed the code on <a href="https://github.com/aziflaj/first-website-html-css" target="_blank">my GitHub account</a> so I can later use it as a reference, a very helpful one for me. I am now waiting for the next course, "Build Your First Website: JavaScript". I hope not to forget what I learned about CSS until then.
