---

title: Converting Markdown to HTML using Ruby
pubDatetime: 2016-07-26
description: I wrote a Ruby module to translate MD files to HTML
comments: true
icon: icons/project52.png
tags:
  - project52

slug: converting-markdown-to-html-using-ruby
---

Since the first time I used Markdown, I have wanted to use it for everything I write. So far, I have used Markdown on my Github-hosted blog, for my [SitePoint articles](https://www.sitepoint.com/author/aldoziflaj/), and I have integrated it in a couple of CMS-like Rails applications (including my yet-to-be-published portfolio).

Markdown is a lightweight markup language that can be easily converted to HTML. It was created by [John Gruber](https://en.wikipedia.org/wiki/John_Gruber) (no relation to [Hans Gruber](https://diehard.fandom.com/wiki/Hans_Gruber)) and [Aaron Swartz](https://en.wikipedia.org/wiki/Aaron_Swartz) (a.k.a. [The Internet’s Own Boy](https://www.imdb.com/title/tt3268458/)) “to write using an easy-to-read, easy-to-write plain text format, and optionally convert it to structurally valid XHTML (or HTML)”. You can read more about the Markdown format on the [Daring Fireball](https://daringfireball.net/projects/markdown/) website.

On the 25th week of my **Project52**, I wrote this small Ruby class that converts Markdown into HTML. My goal was to learn more about Regular Expressions, and how to use them to parse complex texts. In this project, I used regex to find special characters in the Markdown format and change them into standard HTML tags.

![Mastering Regex](https://miro.medium.com/max/1058/1*go5dPUOtrJ36TP7RjvqYWg.jpeg)

The most important code is a single class called `MdHtml`, with one `#to_html` public method and some other private methods that handle the conversion. Also I wrote a simple script to call this class on a given markdown and convert it into an HTML file.

The source code is on my GitHub account ([aziflaj/md2html](https://github.com/aziflaj/md2html)), and you can fork it anytime to make it better than it is. The whole purpose was to get a grasp on regular expressions, which have been really black magic to me so far… not that I am a regex master now, but at least I know more about how to use them in Ruby and how the whole Markdown to HTML conversion works. Maybe this is not the best way to make the conversion, but who cares! YOLRO (You Only Learn Regex Once)… JK you never learn regex.

_Previously posted on [my Medium blog](https://medium.com/@aziflaj/converting-markdown-to-html-using-ruby-c9c92b45aeb1)_
