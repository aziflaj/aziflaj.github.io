---

title: "Meet PinguL: A Programming Language I created"
pubDatetime: 2024-04-29
description: "I took a Compiler Design course during my Masters. We were tasked to create a compiler for a subset of Java, they called it "MiniJava". It's somewhat of a "rite of passage" and funnily enough, for a lot of my peers, the course was their introduction to Finite State Automata and Regular Expressions. I was already a skilled RegEx craftsman at that point, and reading RegEx on a whiteboard was never something I appreciated... Suffice to say I skipped most of the classes. I remember understanding the idea of a Lexer, then dosing off during the Parsing lectures, and never being present for the AST part. So I did what the student version of me did best: I traded some other course's project for the Compiler Design project, never learned a thing, got a passing grade, and moved on with my life."
image:
  path: https://github.com/aziflaj/aziflaj.github.io/blob/main/assets/images/20240429/pingul.jpg?raw=true
  alt: "PinguL"

slug: meet-pingul-a-programming-lang-i-created
---


I took a Compiler Design course during my Masters. We were tasked to create a compiler for a subset of Java, they called it "MiniJava". It's somewhat of a "rite of passage"
and funnily enough, for a lot of my peers, the course was their introduction to Finite State Automata and Regular Expressions. I was already a skilled RegEx craftsman at that point,
and reading RegEx on a whiteboard was never something I appreciated... Suffice to say I skipped most of the classes. I remember understanding the idea of a Lexer,
then dosing off during the Parsing lectures, and never being present for the AST part. So I did what the student version of me did best:
I traded some other course's project for the Compiler Design project, never learned a thing, got a passing grade, and moved on with my life.

![](/assets/images/20240429/pingul.jpg)

That was a bit of a lie, mostly for shock value. Though I really found the course extremely dull, I remember reading about the Ruby VM and how the AST is built and evaluated. I even tried to read the Ruby source code, but I understood nothing. Either way, stealing bits and pieces here and there from lectures and the internet,
I built a mental image of the whole Lexer -> Parser -> AST pipeline. And last week, I finished writing a programming language that makes JavaScript look like
it was written by a genius. And I didn't use any lexer or parser generator. I could have; after all lexing and parsing are both two very studied and very much solved problems by now. But, where's the fun in that? No, this was all from scratch, madman style.

## Wat is this abomination?

The language is called PinguL and it's [freely available on GitHub](https://github.com/aziflaj/pingul) for everyone to ridicule. It is a Dynamically Typed, Garbage-Collected, somewhat Functional, Interpreted programming language. Its syntax is inspired by JavaScript,
with some hints of Python and Ruby, and some C-style boolean evaluations. Unlike JavaScript, it wasn't written in 10 days, but it has a few of the same quirks -- _more on that later_. It has a REPL, it supports
Integers, Booleans, Strings and Lists, but not Floating Point numbers. It has support for first-class functions, meaning you can pass functions as arguments to other functions.
It has a Garbage Collector _by accident_; since the language is written in Go and it is interpreted by the Go runtime, it incidentally profits from Go's GC.

But this article is not about the language itself, you can read more about it in the [GitHub repository](https://github.com/aziflaj/pingul). We are here to talk shit about it. And shit we shall talk...

## Shit Talking Session

PinguL has some limitations, as it usually is the case with languages you write for fun. It doesn't have an extensive standard library, it has like 6 ~~built-in~~ intrinsic functions.
It doesn't have a module system, you can't import code from other files. It doesn't even have support for loops. But what it has, is a lot of heart. And a handful of issues...

### Issue #1: Boolean Tomfoolery

Check this out:

![](/assets/images/20240429/bools.png)

There are a few values that PinguL evaluates to `true` in some cases, even though they aren't booleans. A non-empty string, a non-empty list, and a non-zero integer are all evaluated to `true` when used as conditionals. So
the following code, for example, will print `STRING(true)` for all three cases:

```javascript
if ("non-empty string") {
    print("true");
}

if ([1, 2, 3]) {
    print("true");
}

if (42) {
    print("true");
}
```

Similarly, an empty string, an empty list, and `0` are all evaluated to `false`, and these print `STRING(false)`:

```javascript
if ("") {
    print("true");
} else {
    print("false");
}

if ([]) {
    print("true");
} else {
    print("false");
}

if (0) {
    print("true");
} else {
    print("false");
}
```

If you're wondering what's that "string" thing you're seeing in `STRING(true)`, that's the data type. Because that's how I decided to stringify these expressions for debugging purposes, deal with it 🤷‍♂️

### Issue #2: Parentheses break Math

I don't know where I went wrong with parentheses, but they don't work as expected. I only found out about this when I was writing the readme, and never looked into it. I'm sure it's an easy fix that will never happen.

Parentheses work fine when you write if statements and when you write or call functions, but when you
try to do math with them, they break. Check this out:

![](/assets/images/20240429/parens.png)

It doesn't seem to be an issue with parentheses in general, `if`s and `func`s seem to work. Nor is it an issue with negative numbers, `-4 * 3` is evaluated properly. But when you
add parentheses to the mix, it all goes sideways. Seems like everything before the last `)` is ignored, as both these expressions get evaluated to `6`:

```javascript
(pingul)>> (1 + 2 + (3 * 4) + 5) + 6
INT(6)

(pingul)>> (1 + 2 + (3 * 4) + 5) * 6
INT(6)
```

Is it fixable? Yeah, I'll have to double check how these expressions are parsed and if that's fine, how they are evaluated.

Will I fix it? :slightly_smiling_face: It's not a bug, it's a feature to stop you from using PinguL in production.

### Issue #3: Functions are... quirky

There are a few interesting things about functions in PinguL. Firstly, they're first-class citizens. You can pass them as arguments to other functions, you can return them from functions,
you can even do IIFE (Immediately Invoked Function Expressions) with them. This is valid PinguL code:

```javascript
var square = func(x) {
  return x * x;
};

var sq3 = func(f, val) {
  return f(val);
}(square, 3);

print(sq3);
```

This will print `INT(9)`, as expected. IIFEs and function passing work as expected, just like in JavaScript.

The issue arises when you try and use function-based encapsulation. You can define a function inside another function, but the inner function doesn't have access to the outer function's scope.
So the following code will not work as expected:

```javascript
var outer = func() {
  var x = 42;

  return func() {
    return x;
  };
};

var i = outer();
print(i());
```

It will only print `NIL`, since the inner (returned) function doesn't have access to the outer-scoped `x` variable.

### Issue #4: No Loops

What the title says. PinguL doesn't have loops. You can't write `for` loops, you can't do `while` loops, you can't even `do-while` loops.

But you can recurse, and you can Map and Reduce over lists, so why even bother with loops? I'm not gonna show you how to do that,
but there's a whole section in [the README](https://github.com/aziflaj/pingul/?tab=readme-ov-file#loops) about it.

## Conclusion

Now that I've written a programming language, I can finally say I've done it all.

I've written a Version Control System, [Gogot](https://github.com/aziflaj/gogot).
I've written a Database, [Caboose](https://github.com/aziflaj/caboose/).
And now, I've written a Programming Language, [PinguL](https://github.com/aziflaj/pingul/).

Bold claims and a lot of asterisks.
