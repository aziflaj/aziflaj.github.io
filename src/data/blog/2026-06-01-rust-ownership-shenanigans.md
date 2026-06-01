---
title: "Rust Ownership Shenanigans"
pubDatetime: 2026-06-01
description: ""
slug: rust-ownership-shenanigans
tags: [
]
---

I gave rust a shot a few years ago, when I wanted to write a concurrent Ruby web server called Serrano. I failed, obviously, otherwise everyone would stop using Puma or Webrick and would have switched to Serrano already.

This weekend, i thought of going back at it and instead of implementing something as ballsy as a web server, I decided to learn the language. Knowing enough C to get by and enough Go to teach you things you think you know, the main mental challenge I had so far was to wrap my mind around the rust's ownership model.

The TL;DR version of it, for people who know the fancy words:

- Besides primitive types, all Rust assignments are **move semantics**. assigning doesnt create an "alias", it actually moves data from one label to the other.
- Passing by value also falls under the same category: values are **moved** when passed to functions
- Passing by reference (i.e. passing pointers) works the same as in all sane languages that support pointers; no data is moved, just pointers are passed around.
- Actions on references can only change the underlying data when they are marked as mutable.

![](/assets/images/20260601/rustgothands.png)

And if you're more of a visual learner, like me, here's what all that looks like in code. I wont focus much on the why, more on the what. There's plenty of "why" blogposts on the interwebz, but nobody shows the what in the same style as I do.

## Assignments move data around

Consider the following code:

```rs
#![allow(unused)]

#[derive(Debug)]
struct Coord(f64, f64);

fn main() {
  let your_house = Coord(45.463504, 6.576340);
  let your_mums_house = your_house;

  // println!("{:?}", your_house); // -> this wont compile
  println!("{:?}", your_mums_house);
}
```

Since your grown ass still lives with your mum, surely both of you have the same location. Assigning a variable to another variable in all languages would mean "here's another name for this variable, and both names point to the same place in memory". But Rust is special, and assignments mean "move this data from the old name to the new one".

So, what this code shows us, is that `your_house` is not actually yours, it's `your_mums_house`. You... are homeless.

## Pass-by-value is a move

Consider the following function:

```rs
#![allow(unused)]

#[derive(Debug)]
struct Coord(f64, f64);

fn distance_from_the_sun_in_au(c: Coord) -> f64 {
  1f64 // by definition, we're all 1AU away from the sun (give or take)
}

fn main() {
  let your_house = Coord(45.463504, 6.576340);
  let your_mums_house = your_house;

  let d = distance_from_the_sun_in_au(your_mums_house);
  println!("{:?}", d);

  //println!("{:?}", your_mums_house); // -> this wont compile
}
```

When trying to calculate the distance from the Sun, somewhere in the execution of the code there's a `c = your_mums_house` where `c` is the param of that function. That's a move, and as we know, using a variable after its value is moved is a big no-no.

A good solution to this would be to pass a copy of `your_mums_house` to the func, in which case you need the struct to `#[derive(Clone, Copy)]` (or implement them yourself). But big things can be expensive to copy. And in that case, references can be a better workaround

## Pass-by-reference doesn't move data

In all languages that support references, having a reference to a _thing_ (without objectifying things) means you have access to the latest value of that thing (if you're reading it), or changing that value to anything (within constraints, obviously).

But rust is special, and its things are immutable by default, so immutable references mean "you can use this thing without moving it". Consider this:

```rs
#![allow(unused)]

#[derive(Debug)]
struct Coord(f64, f64);

fn distance_from_the_sun_in_au(c: &Coord) -> f64 {
  // c = &Coord(0f64 ,0f64); // -> this won't compile

  1f64 // by definition, we're all 1AU away from the sun (give or take)
}

fn main() {
  let your_house = Coord(45.463504, 6.576340);
  let your_mums_house = your_house;

  let d = distance_from_the_sun_in_au(&your_mums_house);
  println!("{:?}", d);

  println!("{:?}", your_mums_house); // now this compiles
}
```

With the reference of your_mums_house passed to the func, we no longer move data around. We only pass a pointer to where this data lives in memory, and we can access that place directly. And by "access", i mean "read". The data is immutable by default.

### Actions on mutable references

Now, say you wanna pass a reference to a function, and use that function to change the thing you're referring to. A good example would be an online chess game: you use a list of moves to as the game state. On every move you first validate whether the move is valid (in case an Elo 200 player wants to check with the king) and if the move is valid then you append it to the game state. You don't want to copy a long list of moves inside this function, because you care for the memory footprint of your chess engine.

But let us not stray too far from the path, and get back to `your_mums_house`:


```rs
#![allow(unused)]

#[derive(Debug)]
struct Coord(f64, f64);

fn distance_from_the_sun_in_au(c: &mut Coord) -> f64 {
  *c = Coord(0f64, 0f64); // relocate to Null island
  1f64
}

fn main() {
  let your_house = Coord(45.463504, 6.576340);
  let mut your_mums_house = your_house;

  println!("{:?}", your_mums_house);

  let d = distance_from_the_sun_in_au(&mut your_mums_house);
  println!("Distance from the sun: {} AU", d);

  println!("{:?}", your_mums_house);
}
```

The only way to change the reference passed to the function is by explicitly marking it as mutable when writing the func, and also by marking `your_mums_house` as `mut` when we declare it. Otherwise, you'll end up fighting the rust compiler, and that damn compiler has hands.
