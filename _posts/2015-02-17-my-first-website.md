---
layout:     post
title:      Bulding my first website with HTML and CSS
date:       2015-02-17
summary:    Yesterday I finished a short course on Learnable titled Build Your First Website-HTML & CSS
tags:       [HTML, CSS, Learnable, website]
---

<p>
Yesterday I finished a <a href="http://goo.gl/6PuJD4">short course</a> on Learnable titled "Build Your First Website: HTML & CSS". I wasn't a complete noob on HTML and CSS but I knew (and I still do) that I suck at CSS. I always forget what any line of CSS does (yeah, programming is way simpler). So I started the course at Learnable, hoping it would be my last CSS course I follow.
</p>

The best thing with this course was its practicality: in the end you would have to build a website just like <a href="http://www.buildyourfirst.website/#">this one</a>. I finally learned why people use `meta` tags on HTML and why this piece of code is everwhere:

{% highlight html %}
<!-- [if il IE 9]>
  <script src="assets/js/html5.js"></script>
  <script src="assets/js/respond.js"></script>
<![endif]-->
{% endhighlight %}

Also I learned what <a href="https://smacss.com/">SMACSS</a> (Scalable and Modular Architecture for CSS) says about <a href="https://smacss.com/book/categorizing">categorizing CSS rules</a> in five different categories, even though for this project were used only 3 of them. 

Another nice thing I learned was the usage of CSS sprites. I have heard of it before, I was aware of its existence, but this was the first time I used them. Even though merging all your images into a single image file results in a bigger file than the sum of the smaller images divided, it also results in less HTTP requests to download all the images, thus in a faster loading of the webpage.

In general, it was fun to learn CSS using this course. I can't say that I'm a CSS master now. I feel a bit less noob than before, but I still qualify myself as a CSS noob. Thankfully, I pushed the code on <a href="https://github.com/aziflaj/first-website-html-css">my GitHub account</a> so I can later use it as a reference, a very helpful one for me. I am now waiting for the next course, "Build Your First Website: JavaScript". Hopefully I don't forget what I learned about CSS until then.
