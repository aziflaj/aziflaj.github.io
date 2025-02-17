---
layout: post
title: "Computers can't store numbers"
date: '2025-02-17'
math: true
---

> I have been blogging here for 10 years now. It all started with [V for Vagrant](/posts/v-for-vagrant/).

It's safe to assume you, the reader of this blog, have taken a
programming course at some point in your life. And if you have, you
know about `int`s and `float`s and `double`s and why or how they are used.
If you took a Introduction to Computer Science course, or maybe a 
Computer Architecture course, you probably know how these different numbers are stored
in binary. If not, a brief intro would be:
 - `int`s are stored as binary numbers, with a fixed number of bits. Sometimes
 they are signed (e.g. `-3`) so we reserve one bit for the sign. We call it [2's complement](https://en.wikipedia.org/wiki/Two%27s_complement).
 - `float`s are stored in a format called [IEEE-754](https://en.wikipedia.org/wiki/IEEE_754),
 and it represents each number as a **power of 2** times a **fraction**, both of which 
 are stored in binary. The only difference between `float` and `double` is the number of bits
 used to represent the exponent from the power of 2 and the fraction.

Given that, and knowing a bit of math (onto which we might go any moment now), I'm making this
bold statement: Computers can't store most numbers. And by "most", I mean _almost all_ numbers.
So to make it even more engaging and/or enraging: **Computers can't store numbers**.

## What are you talking about?

Consider this simple C program:

```c
#include <stdio.h>

int main() {
    int population = 8206358989;
    printf("%d people in the world", population);

    return 0;
}
```

If you ignore the compiler warnings and run this program, you will get something like:

```
-383575603 people in the world
```

That's what happens when you need more bits than what you have, an overflow occurs.
Since the above program was compiled on a 32-bit machine, the `int` type is 32 bits long.
Using 32 bits you can only store up to 2^32 - 1 numbers, and that's usually split in half
for both positive and negative numbers. The world's population needs a few more bits than that.
Though this exact overflow won't happen in a 64-bit machine, the overflow can
still happen if you try to store a number larger than 2^64 - 1.

The solution? Use a BigInteger library, which stores numbers as arrays of digits and performs
arithmetic operations on them. Since now you're working with individual digits, you can store
an arbitrarily large number as an array of its digits. Some operations might be slower, but that's the trade-off.
And using this digit representation, you can store all the integers you want.
Unluckily for us, mathmagicians couldn't just settle for integers,
they had to invent more numbers...

## Real numbers
Real numbers are split into two categories: rational and irrational numbers.
Rationals, which also encompass integers, are numbers that can be expressed
as fractions. We already have a way to represent them with IEEE-754, as long as we
can represent the fraction as $$ \left( 2 + \sum_{i=1}^{23}\frac{b_i}{2^i} \right) \cdot 2^{e} $$,
where $$ b_i \in \left\{ 0, 1 \right\} $$ is the _i_-th bit, and
$$ e \in \mathrm{N},  -126 \leqslant e\leqslant 127 $$ is the exponent.

Now, there are 2 problems here. Firstly, IEEE-754 represents the same amount of numbers between two power of two's.
It is "denser" in numbers between two small powers of 2 and "sparser" between two large powers of 2, in the sense that
the represented numbers are closer together in the former case and farther apart in the latter.

![](https://cdn.masto.host/socialjvnsca/media_attachments/files/109/841/543/718/147/188/original/50ef83b5011afb0f.png)

This density and sparsity can manifest itself in very interesting ways. Example, this program won't print what you expect:

```c
#include <stdio.h>

int main() {
    float x = 16777217.0f;
    // this prints 16777216.000000
    printf("%f\n", x);
    return 0;
}
```

Similarly to the overflow mentioned above, this one is called *Underflow*.

Secondly, there's an infinite amount of numbers between two power of 2's. We can only represent a finite amount of them.
Strictly speaking, we can't represent most real numbers. Putting them in a balance, on one side all the numbers we can represent
and on the other side all those we can't represent, the latter side is much heavier. Infinitely heavier, [set-cardinality-wise](https://en.wikipedia.org/wiki/Cardinality).

To reiterate: **Computers can't store numbers**.

## What can we do?
Well, we can store _all_ integers by using BigInteger digit representation. We can store _some_ real numbers by using IEEE-754,
but it comes with the trade off of not being able to [add 0.1 to 0.2](https://0.30000000000000004.com/). This is why you should
never use floats when you're representing money, lest you want to lose a few cents here and there.

So what we can do is more trade-offs. Similar to BigIntegers, i.e. representing numbers as arrays of their digits,
we can represent some real numbers as fractions, where both the nominator and the denominator are BigIntegers. This way,
0.1 and 0.2 would be stored as `{ nominator: 1, denominator: 10 }` and `{ nominator: 2, denominator: 10 }`, which
would make their sum as `{ nominator: 3, denominator: 10 }` without any loss of precision. This is how Python's `fractions` module works.

Using this nominator/denominator representation, we just made it possible to represent infinitely many real numbers, _all_ the rational numbers.
But then again, we can't represent irrational numbers, a set of uncountably infinite size.
And uncountable infinities are a lot heavier than countable infinities.

**Computers can't store numbers**.

When it comes to irrationals, like $$ \pi $$ or $$ e $$ or $$ \sqrt{2} $$, we can only store approximations. We can store a 
rational number that is close to 10e-100 to $$ \pi $$, but it won't be $$ \pi $$. We know more than 100 trillion digits of $$ \pi $$,
but for most practical purposes, `3.14159` is enough. It's said that NASA only uses the first 15 digits of $$ \pi $$ for their calculations.
And that's the trade-off we have to make: precision for performance.

As with everything in computer science, it's all about trade-offs. **Computers can't store _most_ numbers**, but they can store enough numbers to be useful, to some extent.
