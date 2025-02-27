---

title: Ruby and card tricks
comments: true
description: Magic is not real, but card tricks are...
pubDatetime: 2015-08-18
tags:
    - ruby
    - magic
    - cards
    - math

slug: ruby-card-tricks
---

Most of the card tricks I know are based on doing something without the other guy watching, such as manipulating the cards, seeing them in one way or another, etc. But I know a card trick that is highly mathematical rather than a real _'trick'_. Hence, it would work with any kind of object that is [stackable](https://www.google.com/search?q=stackable&ie=utf-8&oe=utf-8#q=define+stackable) on other peers.

The trick is rather simple. It is done using 21 playing cards. With the cards facing down, you ask the other guy to pick one, see it, and put it with the other cards. Then, you start putting the cards in 3 stacks, one after the other, and in the end ask the other guy to point to the stack that contains the card (without pointing the card, of course). You put this stack between the other two, and repeat this process another 2 times (3 times in total). In the end, the other guy's card will be the 11th card in the big pile of cards.

Simple, _"magical"_ and mathematical. I couldn't stop myself from picking a pen and a paper and asking myself _"Why does this happen?"_. Some basic mathematical calculations later, it was obvious that after doing that kind of shuffling 3 times, the only possible position for the card is the 11th one. Think of the cards as an array of 21 cards, from 0 to 20. They are divided in 3 groups, named **G0**, **G1** and **G2**. A random card with a given index `i` will end up in the group with index `i mod 3` and its index inside that group will be `floor(i/3)`. In every shuffle, 7 is added to the group index, since that group goes in between two other groups. Finally, a (not so) complicated, three-floor equation will be raised:
![three-floor equation](/assets/images/20150818/equation.png)

> The result is 10 and not 11 because I am using 0-based counting

After figuring that out, I thought of wrapping it in a Ruby program. Why Ruby? First of all, because I don't know Ruby. It is one of the most used programming languages in the world, and I don't know how to write Ruby code. Secondly, I think Ruby on Rails (RoR) might be a pretty good framework to experiment sometimes, and why not pick it up as the framework of choice.

The full source code is this:
```ruby
require 'set'

def divide_cards(card_set)
  card_array = card_set.to_a
  group0 = Array.new
  group1 = Array.new
  group2 = Array.new

  for index in 0 ... card_array.size do
    if index % 3 == 0
      group0.push(card_array[index])
    elsif index % 3 == 1
      group1.push(card_array[index])
    else
      group2.push(card_array[index])
    end
  end

  return group0, group1, group2
end

suits = %w(Clubs Diamonds Hearts Spades)
cards = %w(A 2 3 4 5 6 7 8 9 10 J Q K)

trick_cards = Set.new # the list of 21 cards to do the trick

# Generate all 21 cards for playing
until trick_cards.size == 21
  rnd_card = cards.sample.to_s
  rnd_suit = suits.sample.to_s
  trick_cards.add("#{rnd_card} of #{rnd_suit}")
end

3.times do
  # divide the cards into 3 groups
  group0, group1, group2 = divide_cards(trick_cards)

  # show the cards
  printf '%-17s', 'Group 0'
  printf '%-17s', 'Group 1'
  printf "%-17s\n", 'Group 2'

  7.times do |i|
    printf '%-17s', group0[i]
    printf '%-17s', group1[i]
    printf "%-17s\n", group2[i]
  end

  # read the group
  group_index = -1 # stub value
  puts 'In which group is your card located? '
  group_index = gets.to_i until group_index >=0 && group_index <= 2

  # put the selected group in the middle
  trick_cards.subtract(trick_cards)
  if group_index == 0
    trick_cards.merge(group1)
    trick_cards.merge(group0) # right in the middle
    trick_cards.merge(group2)
  elsif group_index == 1
    trick_cards.merge(group0)
    trick_cards.merge(group1) # right in the middle
    trick_cards.merge(group2)
  else # group_index == 2
    trick_cards.merge(group0)
    trick_cards.merge(group2) # right in the middle
    trick_cards.merge(group1)
  end
end

puts "Your card should be #{trick_cards.to_a[10]}"
```

I noticed some really nice features of Ruby.
First of all, its loops are great! I just love that `7.times do |i|` loop, with the 7 iterations and the piping operator. The fun thing is that long time ago, when I was thinking of creating a programming language with an Albanian syntax, I was considering ways of creating loops in the same way, by using:

```python
5.repeat {
  # block of code
}
```

Also, the first function I wrote on Ruby could return multiple values. I've seen something similar in Scala but this is actually the first time I use it.

Overall, Ruby was nice, this magic trick was _demagicated_ (the process of removing the magic out of something), and who knows what's next.
