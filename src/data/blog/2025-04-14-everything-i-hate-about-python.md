---
title: "Everything I hate about Python"
pubDatetime: 2025-04-14
description: "A lighthearted, brutally honest rant about Python from the perspective of an old timer who’s dabbled in too many languages to count. From tuple weirdness to PEP8 gripes, this post pokes fun at Python’s quirks while offering a cautionary tale for those new to the language — or just wondering why their loop variables won’t go out of scope."
slug: everything-i-hate-about-python
tags: ["python", "rant", "language-design", "pep8"]
---

I used too many languages during the last +10 years as a professional software craftsman. A long time ago, I wrote Java. I did some personal projects in C and Elixir, I wrote a lot of Ruby and Node.JS. I taught classes in Go and JavaScript, I even dabbled in Python for number crunching and ML shenanigans. I like to think I'm well versed in building software, regardless of the language. But Python is something I never enjoyed working with; the language seems like a toy, the ecosystem too hectic and immature... and this is my experience from back in the days, when Python folks were still fighting the 2.x VS 3.x holy war.

I recently started properly learning Python, for personal reasons. I did it both the old-school and the new-school way: I downloaded books and went through them, I consulted ChatGPT to transfer my current knowledge to Python, and I even ran some experiments in the Python REPL.

After so many years of considering Python a stupid, silly language, I came to the conclusion I was not wrong. Python truly is a stupid, silly language. It's like that line from Kung Pow: _We purposely designed it wrong, as a joke._

And here I will rant about what I learned and why I think it's stupid, comparing it with other languages... but keep in mind there are no bad languages. _There's languages you hate, and languages you never used_. It's just a rant, I'm nitpicking features, no hard feelings if you like Python (you shouldn't, but you can't teach taste).

![](/assets/images/20250414/python.png)

## Tuples are stupid
Tuples are these immutable ordered list structures that Python uses to... store things in immutable ordered lists I guess? Like an immutable array, but fancier. They're surrounded by parentheses... **usually**. And that "usually" is a very weak restriction, because you can add an accidental, single comma and ruin the life of too many people. Consider the following example:

```py
# This is an `int`, all good
thing = 1     

# This is a `tuple`, notice the parentheses
thing = (1, 2)

# Still an `int`, not a `tuple`
# Understandable, regardless of the parentheses
thing = (1)    

# A `tuple` of a single item. Most likely intentional
thing = (1,)

# A `tuple` of a single item. See the comma?
# Python flipping you the finger
thing = 1,
```

That last one, that no-parentheses, most likely unintentional, single-item tuple. That's the single comma that will ruin the life of too many people if put in the wrong place. 

It should have been considered a Parser bug from the day one. It should have thrown an `Unexpected TOKEN(',') in line 69`. But they decided to go with it. 

## OOP is stupid
I have always worked with OOP languages, and I've even done some functional programming for fun and non-profit. But OOP is something I've always worked with; I know how it works and how to design with Objects. I remember having issues with the multiple inheritance of C++, and I know how good you can have it if you use Composition over Inheritance. Yet Python's OOP feels half baked...

Consider this following class:

```py
class Stupid:
    """
    PEP8 is stupid, but we'll deal with that later
    """
    public_var = "ok"
    _protected_var = "...fine, whatever"
    __private_var = "are you fucking kidding me?!"

    def __init__(self):
        """
        private initializer i assume, right PEP8?!
        """

    def do_stuff(self):
        """
        What kind of derranged programming language
        needs the caller to be passed as a param?
        Well, C does the same thing... but C ain't OOP!
        """
```

The 3 pillars of OOP: Encapsulation, Inheritance, Polymorphism... Best 2 out of three I guess. Or maybe, best 1 out of 3, but let's worry about polymorphism later. 

Whose idea was it to implement OOP and forget about access control?! How do you live in a world where your private variables can be overwritten by anyone?!

That whole `_protected` and `__private` thing is just a convention, because the language itself doesn't support any access levels and pythonistas decided to go "please don't touch my doubly-underscored variables, they're for me only". By the same logic, all the magic methods should be considered private, because they start with a double underscore. But fear not, pythonistas have another convention for `__magic__` methods... It's conventions piled upon conventions because the language is designed wrong.

Also, I now have to pass `self` as the first param to every method? Also also, it is only called `self` by convention?! Because this would work exactly the same way:

```py
class Stupid:
    def __init__(self):
        self._hotel = "trivago"

    def hotel(this):
       print(type(this).__name__ + ": " + this._hotel)

s = Stupid()
s.hotel()
```

This prints `Stupid: trivago` and yes, you can call it `this` instead of `self` and feel normal. Hell, you can call it `dumdum` if you want and it'd still work the same! If you tell me this is a language feature and not a bug, I will personally piss on your grave.

I understand where this comes from though. In good ol' C, we used to write "classes" and "methods" like this:

```c
typedef struct {
    char* _hotel;
} Stupid;

Stupid* stupid_init() {
    Stupid* s = malloc(sizeof(Stupid));
    if (s == NULL) return NULL;

    s->_hotel = strdup("trivago");
    return s;
}

char* stupid_hotel(Stupid* self) {
    return self->_hotel;
}
```

In C, we did prepend all the function names with the struct name, and we did pass the struct as the first param to the function. But C is not Object-Oriented; that was an artifical thing we did for our sanity. This is unacceptable from a "modern OOP language".

## PEP8 is stupid
Those `public`, `_private`, `__protected` shenanigans are a convention, and PEP8 is the recommendation from the Python overlords on how to write code, which describes those shenanigans. It also suggests the following two, which I hate with a passion:

### 79 char line limit
Why not 80, [for historical purposes](https://softwareengineering.stackexchange.com/a/148678)?! It used to be that a billion years ago, when we were writing programs in punch cards, those cards had an 80 character limit. Why go with 79? Why are you taking a character away from it? What goal are you trying to achieve? 

True, most people will go for something longer, maybe 100 or 120 char lines because they can still read it. I prefer shorter lines, 80 seems to work fine for me and I can still open 3 side-by-side windows on my 14" screen and read them just fine. But why 79? So specific...

### 4 Spaces vs 1 tab
We're going into holy wars now! The language doesn't care. If you start indenting with a single space instead of 4, and if you are consistent through the file, Python will still interpret the code correctly. Try it, try using 2 spaces instead of 4. You will have more real-estate on your screen for code, and you might have more code to read on those 79 char long lines...

I personally prefer tabs, maybe because I got too accustomed to Go and `gofmt` uses tabs instead of spaces. A tab is a single character by the way, and in most editors you can configure the visual width of it. I have `tabstop=2` on my Vim config which makes it _visually_ 2 characters wide, but it's still a single character. A single byte in the file's binary.

But y'all folks prefer 4 spaces... We can't have nice things in the Python world.

## Loops are stupid

Or maybe, loopity scopes are stupid. Consider this basic example:a

```py
arr = ["foo", "bar", "baz"]
for i, item in enumerate(arr):
    pass

print(i)    # => prints `2`, even though `i` is out of scope
print(item) # => prints `baz`
```

Why did `i` and `item` leak out of the loop? Because Python is stupid, that's why!

## Jokes aside...
Learning from my experience with [PinguL](https://github.com/aziflaj/pingul/), it's easy to mess things up when writing parsers. So this rant is not for people to start calling Python a bad language, it's more of a "careful out there buddy" kind of blogpost written by someone that doesn't use Python as a daily driver. Take everything with a grain of salt, especially if you're a die-hard Pythonista. But if you're like me, coming from stricter, more opinionated ecosystems,just know you're not crazy for finding Python a little... too loose around the edges.
