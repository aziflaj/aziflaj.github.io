---

title:      "Week #13: Playing Blackjack with Machine Learning"
pubDatetime:       2016-04-11
description:    A (very dumb) probability model deciding whether you should draw another card or not to win at Blackjack.
comments:   true
tags:
    - project52
    - java
    - machine learning

slug: week-13-playing-blackjack-with-machine-learning
---

It has been almost a month since the last blogpost, and I've been too busy to keep it up with my Project 52. This is my last semester as a Computer Engineering undergraduate and usually is the heaviest semester of all. That's why I've been so silent for the last month and didn't even have too much time to focus on this project.

For the last months I've been reading a lot about machine learning, AI and some related topics. I've read online articles, papers and doctorates from well-known (and also from not so well-known) people in the AI field, and I've gathered plenty of information about the field in general. I've grown up to be a great fan of neural nets since the first OCR I trained ([MNIST](http://yann.lecun.com/exdb/mnist/)). Anyway, this "week" I've been working on something I'd like to call "a probability model for playing Blackjack". If that intrigued you, please keep reading.

## What is this probability model?
There are a couple of approaches to Machine Learning (ML), like decision trees, Support Vector Machines (SVM), neural nets, probabilistic models, etc., each of them with their use cases (which sometimes can overlap). A probabilistic model is some kind of model that makes decisions based on a given probability. Examples of these are [Markov Models](https://en.wikipedia.org/wiki/Hidden_Markov_model) used for speech recognition, [Naive Bayes classifier](https://en.wikipedia.org/wiki/Naive_Bayes_classifier) for classifying (e.g. spam or ham), etc.

> I may have been mistaking 'probabilistic model' for 'statistical model'. If so, please let me know in the comment section.

For this challenge, I built a simple probabilistic model that answers a simple question about [Blackjack](http://wizardofodds.com/games/blackjack/basics/): **"What's the probability that by hitting (i.e. drawing another card) the player will still be in the game and won't lose?"** The simulation I ran was a one-player game of Blackjack, which initially played to gain experience and calculated the probability of still being part of the game if the player drawed another card from the deck.

## The simulation
I wrote the simulation program in Java because I am more familiar with the language, and it was simpler for me to create the deck of cards in Java. I started by writing this code, modelling a playing card in Java:

```java
package com.aziflaj.ventuno;

import java.util.Random;

public class Card {
    private Face mFace;
    private Suit mSuit;

    private Card(Face face, Suit suit) {
        mFace = face;
        mSuit = suit;
    }

    public static Card generateRandomCard() {
        Random rnd = new Random();
        Face[] faces = Face.values();
        Suit[] suits = Suit.values();
        return new Card(faces[rnd.nextInt(faces.length)], suits[rnd.nextInt(suits.length)]);
    }

    public int getPoints() {
        return mFace.asInt();
    }

    public Face getFace() {
        return mFace;
    }

    @Override
    public String toString() {
        return String.format("%c, %c", mFace.asChar(), mSuit.asChar());
    }

    @Override
    public boolean equals(Object obj) {
        return obj instanceof Card &&  mFace == ((Card) obj).mFace && mSuit == ((Card) obj).mSuit;
    }

    public enum Face {
        ACE, TWO, THREE, FOUR, FIVE,
        SIX, SEVEN, EIGHT, NINE,
        TEN, JACK, QUEEN, KING;

        public int asInt() {
            switch (this) {
                case ACE:
                    return 1;
                case TWO:
                    return 2;
                case THREE:
                    return 3;
                case FOUR:
                    return 4;
                case FIVE:
                    return 5;
                case SIX:
                    return 6;
                case SEVEN:
                    return 7;
                case EIGHT:
                    return 8;
                case NINE:
                    return 9;
                case TEN:
                    return 10;
                case JACK:
                    return 10;
                case QUEEN:
                    return 10;
                case KING:
                    return 10;
                default:
                    return 0;
            }
        }

        public char asChar() {
            switch (this) {
                case ACE:
                    return 'A';
                case TWO:
                    return '2';
                case THREE:
                    return '3';
                case FOUR:
                    return '4';
                case FIVE:
                    return '5';
                case SIX:
                    return '6';
                case SEVEN:
                    return '7';
                case EIGHT:
                    return '8';
                case NINE:
                    return '9';
                case TEN:
                    return 'T';
                case JACK:
                    return 'J';
                case QUEEN:
                    return 'Q';
                case KING:
                    return 'K';
                default:
                    return '0';
            }
        }
    }

    public enum Suit {
        HEARTS, DIAMONDS, CLUBS, SPADES;

        public char asChar() {
            return name().charAt(0);
        }
    }
}
```

Initially, I thought of building a model that takes in consideration **the cards** the player is holding in the hand instead of the sum of the cards. It would work as a big search tree, with the initial node being the first card drawn and the probability of drawing that; its nodes would be the probabilities of still being in the game if one of the other cards was drawn. By knowning these probabilities, the model would decide if it would draw or stop there. Knowing how many possible combinations of cards are there in the wild, this would be a really big and complicated. But then I figured out (wasn't really that hard) that the suit of the card is not at all important for the game, and not even the order the cards were drawn (call it _"feature engineering"_). That's why that model would be an overkill for the task.

The implementation of the deck is this simple:

```java
package com.aziflaj.ventuno;

import java.util.Stack;

public class CardDeck {
    private Stack<Card> deck;

    public CardDeck() {
        deck = new Stack<>();
        int i = 0;
        while (i < 52) {
            Card card = Card.generateRandomCard();
            if (!cardDuplicate(card)) {
                deck.push(card);
                i++;
            }
        }
    }

    public Card drawCard() {
        return deck.pop();
    }

    private boolean cardDuplicate(Card card) {
        for (Card c : deck) {
            if (c.equals(card)) return true;
        }
        return false;
    }
}
```

As you can see, I'm generating random cards and putting them in the deck (a stack) if the card is not already there. This is a fairly naive implementation, since the algorithm tries more than once to insert some cards because of collisions (duplicated cards).


## The probability map
So I changed the model to consider only the sum of the cards drawn by the player. Initially, with no probability model implemented and a greedy player that wants only to draw cards, the player was winning in 16% - 18% of the games. The probability for each sum was:

```
2 => 1.000000
3 => 1.000000
4 => 1.000000
5 => 1.000000
6 => 1.000000
7 => 1.000000
8 => 1.000000
9 => 1.000000
10 => 1.000000
11 => 1.000000
12 => 0.763199
13 => 0.717735
14 => 0.679474
15 => 0.646142
16 => 0.620540
17 => 0.591418
18 => 0.564090
19 => 0.542397
20 => 0.521169
21 => 1.000000
```

I like to call the above a _probability map_, since it maps every not losing sum to the probability we require.

It is easy to understand why the probability for sums 2 to 11 is 1.0: for whatever card the player draws, the sum will be less than or equal to 21, and the player will still be in the game (or win). The same goes for the sum of 21. Also for each sum between 12 and 20, it is noticeable the decreasing probability. If the sum is 12, the only cards that would kick the player out of the game are the 10 point cards and there are less of those compared to less-than-10 point cards, hence the decreasing probability.

While I agree with the probability calculated for the sums 2-11 and that of 21, I'm not so comfortable with the other probabilities. I'm pretty sure that there is no 50% chance to draw a card and win if the sum of the cards is 20. That's why I decided to "normalize" those probabilities into "better" values, i.e. decrease the winning probability of 20 and increase the winning probability of 12. I applied 3 different formulas for these probabilities, making a strech in that "probability map":

![models](https://raw.githubusercontent.com/aziflaj/aziflaj.github.io/master/images/52-projects/week13/models.png)

The formulas are applied for probabilities of the sums 12 to 20, since the other sums are a sure probability of 1.0. On the left, _P<sub>S</sub>_ is the probability applied at decision making, while _p(s)_ is the probability from the "probability map" above.

I ran the simulation again, with the above models, and the result was this:

```bash
Untrained:
Winning probability: 0.16997

Linear Stretching:
Winning probability: 0.79648

1st Exponential Stretching:
Winning probability: 0.80659

2nd Exponential Stretching:
Winning probability: 0.89401
```

If you draw the above models as curves in a 2D plane, the first and the second models will almost overlap in the (.52, .77) interval; that's why the winningprobability of those models is almost the same. On the other hand, the third model gives a winning probability of almost 90%, which is mostly OK (I guess) for Blackjack players.

So that was the really simple probability model I built for the 13th challenge, which you can find on my Github account ([aziflaj/ventuno](https://github.com/aziflaj/ventuno/)). As I said, I've been a bit busy these last weeks so it will probably take some time until the next coding challenge. Until then, I have another idea. Remember the 7th week? I'm thinking of repeating the same every 7 weeks, so probably the next "challenge" would contain my answers to 7 other interesting questions I'll find online. If you're having any questions lately, please let me know on the comment section so I can find your answer. Until then, stay classy.
