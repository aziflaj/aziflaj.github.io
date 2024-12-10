---
layout: post
title: Designing a Better User Experience for My Fingers
date: '2021-12-28'
comments: true
---

Those of you who know me must be surprised I'm even considering writing about UX. For those who don't know me, I believe an introduction is due. I am one of those people who stare at the terminal for hours every day, either fixing bugs or writing them. Sometimes I have to do some clicking around to manually test the code I wrote, but most of my time is spent either executing commands in a terminal, or writing code in a [terminal-based editor](https://neovim.io/). People don't expect me, someone who's willing to start a debate on why Zsh is the superior shell and how your mouse-dependent development process is slowing you down, to start talking about designing user interfaces and optimizing user experience. And I don't blame them, but we are living in strange times...

If you are in the same field as me, you most likely heard people throwing terms like UI and UX when it comes to designing graphical interfaces and defining users interactions with the product they're building. I want to extend the interface definition beyond graphical, the interaction definition beyond clicking on a screen, and the user definition to include code monkeys, programming primates and software simians (so, most of my audience). When talking about you as a user of a computer, your interface is your keyboard (and your mouse, if you're one of those people), and it's designed by someone who never saw human hands.

![staggered kb]((https://github.com/aziflaj/aziflaj.github.io/blob/main/assets/images/20211228/staggered.jpeg?raw=true)

Normal (a.k.a. staggered) keyboard designers had the impression people have diagonally moving fingers of the same length but also, they thought the fingers on the right hand are longer than the ones on the left. That's why the distance between Y and J is almost double the distance between T and F.

When it comes to designing interfaces, one of the first steps is to gather information on how potential users will want to use it; what problem will the system solve, how the users will interact with it, etc. And as someone who gets paid to sit in front of a computer and type on a keyboard, I have a thing or three to say when it comes to interacting with it. If you're also in the same situation, sitting in front of a computer having to write code for hours every day, I believe the following assumptions will apply to you:

1. You press the spacebar either with your left, or your right thumb, which means...
2. One of your thumbs is underutilized
3. Most of the time, when you reach for the number row, you're holding the SHIFT key down
4. Sometimes you press on the caps lock by mistake, and end up aCCIDENTALLY YELLING MID-SENTENCE

So, after long consideration, I came to the conclusion that the keyboard layout most of us are using was not designed with developers in mind. Consider the following, pretty basic example in PHP, which I got from [Lumen's landing page](https://lumen.laravel.com/) (I swear I'm not a PHP developer!):

```php
<?php

$app->get("user/{id}", function($id) {
    return User::findOrFail($id);
})
```

Imagine having to reach for SHIFT every time you use a variable, or when you declare or call a function, or access something in a different namespace... A good UX designer would put these frequently used symbols only one key away. Not two keys away, not a key combo looking like some weird jazz chord, just one key away (in their defense, they're designing one ring, err... keyboard to rule them all).

That's why I decided to build my own keyboard, with blackjack and hoo... I mean, with a better layout customized to my needs. Due to an old wordplay, I tend to name my projects alphabetically in an Adjective Animal format, both words starting with the same character. So now I'd like to introduce you to [Nasty Narwhal](https://github.com/aziflaj/qmk_firmware), my first mechanical keyboard build:



![nasty narwhal]((https://github.com/aziflaj/aziflaj.github.io/blob/main/assets/images/20211228/keymap.png?raw=true)

Aight, this is the keymap's base layer , not the keyboard itself; the keyboard is the one on the pic on top. It's a 4-row, orthogonal, split keyboard with a few changes when compared to your normal one. Firstly, the number row is what I like to call "shift-inverted" (or schwifted): pressing the keys will give you the symbols, and holding on to "shift" will give you the numbers. Along with that, I swapped the behaviour of the minus, quotes and semicolon keys since I use underscores, double quotes and colons (the punctuation, getyomindoutofthegutter) more than minuses, single quotes and semicolons. The long spacebar is no longer that long, and it's only on the left side. The arrow keys aren't there anymore but they can be accessed via RAISE + H/J/K/L (courtesy of Vim). I also have a rotary encoder on the right side of the keyboard which I can use as a mouse scroll wheel and to control the volume.

What I don't like about this? Great question. I made it "harder" to type [brackets] and {braces}; they're still hidden under a layer, respectively LOWER + O/P and RAISE + O/P. Also, I'm now a slower typist, rocking it at 20 words per minute instead of my usual 90. But on the other hand, after using it for a couple of days now, I feel my pinkie finger being more useful when typing. Isn't a productive pinkie what we all need in life?
