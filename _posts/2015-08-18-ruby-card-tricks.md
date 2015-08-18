---
layout: post
title: Ruby and card tricks
summary: "Magic is not real, but card tricks are. Most of the card tricks I know are based on doing something without the other guy watching, such as manipulating the cards, seeing them in one way or another, etc. But I know a card trick that is highly mathematical rather than a real ‘trick’. Hence, it would work with any kind of object that is stackable on other peers."
tags: [ruby, magic, cards, rails, math, mathematics]
---

Magic is not real, but card tricks are. Most of the card tricks I know are based on doing something without the other guy watching, such as manipulating the cards, seeing them in one way or another, etc. But I know a card trick that is highly mathematical rather than a real _'trick'_. Hence, it would work with any kind of object that is [stackable](https://www.google.com/search?q=stackable&ie=utf-8&oe=utf-8#q=define+stackable) on other peers.

The trick is rather simple. It is done using 21 playing cards. With the cards facing down, you ask the other guy to pick one, see it, and put it with the other cards. Then, you start putting the cards in 3 stacks, one after the other, and in the end ask the other guy to point to the stack that contains the card (without pointing the card, of course). You put this stack between the other two, and repeat this process another 2 times (3 times in total). In the end, the other guy's card will be the 11th card in the big pile of cards.

Simple, _"magical"_ and mathematical. I couldn't stop myself from picking a pen and a paper and asking myself _"Why does this happen?"_. Some basic mathematical calculations later, it was obvious that after doing that kind of shuffling 3 times, the only possible position for the card is the 11th one. Think of the cards as an array of 21 cards, from 0 to 20. They are divided in 3 groups, named **G0**, **G1** and **G2**. A random card with a given index `i` will end up in the group with index `i mod 3` and its index inside that group will be `floor(i/3)`. In every shuffle, 7 is added to the group index, since that group goes in between two other groups. Finally, a (not so) complicated, three-floor equation will be raised:
![three-floor equation](/images/magic-tricks/3rd.png)

> The result is 10 and not 11 because I am using 0-based counting


After figuring that out, I thought of wrapping it in a Ruby program. Why Ruby? First of all, because I don't know Ruby. It is one of the most used programming languages in the world, and I don't know how to write Ruby code. Secondly, I think Ruby on Rails (RoR) might be a pretty good framework to experiment sometimes, and why not pick it up as the framework of choice.

The full source code is this:
{% gist aziflaj/e70f7dd2cdf86de6e955 %}

I noticed some  really nice features of Ruby. 
First of all, its loops are great! I just love that `7.times do |i|` loop, with the 7 iterations and the piping operator. The fun thing is that long time ago, when I was thinking of creating a programming language with an Albanian syntax, I was considering ways of creating loops in the same way, by using:

{% highlight python %}
5.repeat {
  # block of code
}
{% endhighlight %}

Also, the first function I wrote on Ruby could return multiple values. I've seen something similar in Scala but this is actually the first time I use it.

Overall, Ruby was nice, this magic trick was _demagicated_ (the process of removing the magic out of something), and who knows what's next.

