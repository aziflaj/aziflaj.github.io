---
title: "Rust Ownership Shenanigans"
pubDatetime: 2026-06-01
description: "Learning Rust ownership the hard way: moves, borrows, mutable references, and a compiler that got hands."
slug: rust-ownership-shenanigans
tags: ["rust", "ownership", "borrow checker", "move semantics", "references", "mutability"]
---

I gave Rust a shot a few years ago, when I wanted to write a concurrent Ruby web server called Serrano. I failed, obviously, otherwise everyone would have stopped using Puma or WEBrick and switched to Serrano already.

This weekend, I thought of going back at it and instead of implementing something as ballsy as a web server, I decided to do the mature thing and actually learn the language first.

Knowing enough C to get by and enough Go to teach you things you think you know, the main mental challenge I had so far was wrapping my mind around Rust's ownership model. I am not proud to admit that I had to fight the compiler to get my code to work, that borrow checker thingy does have some very quirky requirements sometimes... if you don't know the language, that is.

![](/assets/images/20260601/rustgothands.png)

This post is me documenting the punches.

The TL;DR version of it, for people who know the fancy words:

- All<sup>1</sup> Rust assignments are **move semantics**. Assigning doesn't create an "alias", it actually moves data from one label (variable name) to the other.
- Passing by value also falls under the same category: values are **moved** when passed to functions
- Passing by reference (i.e. passing pointers) works the same as in all sane languages that support pointers; no data is moved, just pointers are passed around.<sup>2</sup>
- Actions on references can only change the underlying data when they are marked as mutable.

***
<sup>1</sup>I say _"all"_, I mean _"all except for types that implement `Copy`"_. But that's a detail, and the point is that the default is move semantics, not copy semantics.

<sup>2</sup>In Rustafarian lingo, this is called a "borrow".
***

If you're more of a visual learner, like me, here's what that looks like in code. I won't focus much on the _why_, more on the _what_. There are plenty of "why" blog posts on the interwebz, but nobody shows the what in the same style as I do.

## Assignments move data around

Consider the following code:

```rust
#![allow(unused)]

#[derive(Debug)]
struct Coord(f64, f64);

fn main() {
  let your_house = Coord(45.463504, 6.576340);
  let your_mums_house = your_house;

  // println!("{:?}", your_house); // -> this won't compile

  println!("{:?}", your_mums_house);
}
```

Since your grown ass still lives with your mum, surely both of you have the same location. In a lot of languages, assigning a variable to another variable feels like saying: _Here's another name for the same thing._ An **alias**, if you will. But Rust is special.

In Rust, this assignment:

```rust
let your_mums_house = your_house;
```

does not mean _"now we have two names for the same value."_ It means _"move the value from `your_house` into `your_mums_house`."_

So, what this code shows us is that `your_house` is not actually yours. It's `your_mums_house`.

You, my friend, are homeless.

## Pass-by-value is a move

Now consider this function:

```rust
#![allow(unused)]

#[derive(Debug)]
struct Coord(f64, f64);

fn approx_distance_from_the_sun_in_au(c: Coord) -> f64 {
  // by definition, we're all 1AU away from the sun (give or take)
  1f64
}

fn main() {
  let your_house = Coord(45.463504, 6.576340);
  let your_mums_house = your_house;

  let d = approx_distance_from_the_sun_in_au(your_mums_house);
  println!("{:?}AUs away from the Sun", d);

  // println!("{:?}", your_mums_house); // -> this won't compile
}
```

This looks innocent. We're not assigning `your_mums_house` to another variable directly. We're just passing it to a function.

But somewhere in the execution of this code, Rust sees something morally equivalent to:

```rust
let c = your_mums_house;
```

And that is a move. After passing the value to the function, the function becomes the new owner of that value. When the function ends, `c` goes out of scope. Which means your mum's house is gone. Evicted by function call.

A good solution would be to pass a copy of `your_mums_house` to the function. In that case, the struct would need to implement `Copy`:

```rust
#[derive(Debug, Clone, Copy)]
struct Coord(f64, f64);
```

For tiny things like two `f64`s, that's probably fine.

But not everything should be copied. Big values can be expensive to duplicate, and sometimes you don't want to give ownership away just because a function wants to look at something.

## Pass-by-reference doesn't move data

In languages that support references or pointers, having a reference to a thing usually means you can access the thing without copying it. Rust also lets you do that.

But Rust is special, and it has special rules. The compiler didn't come here to make friends, it came here to throw hands.

```rust
#![allow(unused)]

#[derive(Debug)]
struct Coord(f64, f64);

fn relocate(c: &Coord) {
  println!("Relocating from {:?} to Null Island", c);

  // *c = Coord(0f64, 0f64); // -> this won't compile
}

fn main() {
  let your_house = Coord(45.463504, 6.576340);
  let your_mums_house = your_house;

  relocate(&your_mums_house);

  println!("{:?}", your_mums_house); // now this compiles
}
```

With `&your_mums_house`, we no longer move the data around. We only lend the function access to it.

And by _"access"_, I mean _"look, do not touch."_

The function can read the value, print it, gossip about it, whatever. But it cannot change it, because Rust's values are immutable by default. You have to explicitly say that you want to mutate something.

## Actions on mutable references

Now, say we actually want to pass a reference to a function and let that function change the value.

A good real-world example would be an online chess game. You have a list of moves as the game state. On every move you validate whether the move is legal, so that a 200 Elo player can't check with the king. And if the move is legal, you append it to the game state.

You don't want to copy the entire list of moves every time someone moves a pawn two squares and immediately blunders a bishop. You want to mutate the existing state.

But let us not stray too far from the path. Back to `your_mums_house`. Let's actually relocate this time.

```rust
#![allow(unused)]

#[derive(Debug)]
struct Coord(f64, f64);

fn relocate(c: &mut Coord) {
  println!("Relocating from {:?} to Null Island", c);

  *c = Coord(0f64, 0f64);
}

fn main() {
  let your_house = Coord(45.463504, 6.576340);
  let mut your_mums_house = your_house;

  println!("Before: {:?}", your_mums_house);

  relocate(&mut your_mums_house);

  println!("After: {:?}", your_mums_house);
}
```

This is Rust making mutation very, very explicit.

You cannot accidentally mutate through an immutable reference. You cannot pass a mutable reference to something that was not declared mutable. And you cannot have random parts of the code poking the same value while mutation is happening.

The thing I like about these mutable references is their exclusivity. In other languages I know, having a pointer means you can edit the underlying data. Pass that pointer to concurrent functions, and you have a race condition (if you forget to use a mutex). Pass that pointer to a function that might panic, and you have a dangling pointer. Pass that pointer to a function that might return early, and you have a memory leak.

With these ownership shenanigans, borrow checker and whatnot, Rust protects you from all of that. You cannot have two mutable references to the same value at the same time. You cannot have a mutable reference while an immutable reference exists. You cannot accidentally have a mutable reference because you forgot to mark it as mutable.

At first, this feels annoying. You write code that feels reasonable. The compiler says no. You add `mut`. The compiler says no again. You add `&mut`. The compiler says no in a different font. Damn compiler got hands.

But it's all for a good cause.
