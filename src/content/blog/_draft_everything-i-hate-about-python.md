---
title: "Everything I hate about Python"
pubDatetime: 2025-04-30
draft: true
description: ""
slug: everything-i-hate-py
---

I used to mock Python as a language a lot, until I decided to learn it... and I came to the conclusion I was not wrong. It's like that line from Kung Pow: _We purposely designed it wrong, as a joke._

Language designed by the utterly derranged. 

## Tuples are stupid
```py
thing = 1    # int
thing = (1)  # int, understandable
thing = (1,) # tuple
thing = 1,   # tuple, fuck you
```

## OOP is stupid

```py
class Stupid:
    """
    PEP8 is stupid but okay
    """
    public_var = "ok"
    _protected_var = "...fine, whatever"
    __private_var = "are you fucking kidding me?!"

    def __init__():
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

## PEP8 is stupid
Besides the public, _private, __protected bs...

- 79 char line limit... why not 80, for historical purposes?!
- 4 Spaces vs 1 tab... I have a `tabstop=2` on my Vim config so if I'd use tabs I wouldn't get horizontally inclined lines of code. But we can't have nice things in this world.

## Loops are stupid
Why is scope so dumb?

```py
arr = ["foo", "bar", "baz"]
for i, item in enumerate(arr):
    # do stuff

print(i) # => It print's 2, even though `i` is out of scope
```