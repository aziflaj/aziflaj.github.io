---
layout: post
title: We really started from the bottom...
date: '2021-03-13'
summary: A random internet person had commented on a C++ calculator Gist I had published around 5 years ago and the memories cued in
comments: true
category:
  - c++
---

I think I wrote my first program at some point between 8th and 9th grade, and it was one of those Windows BAT scripts you write to mess up with your friends computer. Fast forward to high school, I was creating "Mozilla Firefox" shortcuts on the desktop that shut your PC down. So many frustrated kids in the CS lab, and a couple of laughers in the corner... Or that time I had to disable the antivirus because it was deleting a (broken) keylogger I wrote the previous night. Ah, good times.

This morning I woke up by an email; a random internet person had commented on a forgotten Gist I had published some 5 years ago. It was a calculator I wrote in C++. The date said "01-01-13 12:01" but I highly doubt I did this on New Year's Eve. I believe I wrote this somewhere around October or November 2012, back when I was learning C++ but before I learned anything about OOP. My set up at the time was a 2002-built Compaq Evo laptop running the latest Windows XP (32bit) on a Pentium 4, rocking on 512MB of RAM. I had tried different code editors, from Notepad to Notepad++, but I had decided to use a real IDE and not some lame text pad. Meet Turbo-C++, the blue kid on the block:

![blue-dabade-dabada](https://github.com/aziflaj/aziflaj.github.io/blob/main/assets/images/20210313/turboc.png?raw=true)

Probably the reason I have to wear glasses now, I can't remember how many hours I spent in front of this blue screen (there was only a full-screen option), stroking my then-yet-to-grow beard in attempts of debugging, squeezing my brainjuice and thinking "I now can speak to machines".

But let's get on with the Gist and review it:

```cpp
/*
  Name: CALCULATOR v1.0
  Author: Aldo Ziflaj
  Date: 01-01-13 12:01
  Description: A simple calculator that can add, subtract, multiply,
               divide, raise to n-th integer power and
               calculate the factorial
*/


#include <iostream>
#include <cstdio>
#include <cmath>
#include <cstring>
using namespace std;
```

Look at that, already using comments when there was nobody to read the code. Smart kid, huh. I remember being too annoyed with all the includes, so I once created a header file called "common.h" and inside I put all the content from "iostream.h", "stdio.h" and "string.h". Instead of including 3 different files, I could do a simple include "common.h" (I would take the smart kid comment back; but I only did this once).

```cpp
unsigned int factorial(int a) {
  int r=1;
  for (int i=a;i>0;i--) r*=i;
  return r;
}
```

Why didn't I use recursion here? Good question, I **had to** think about overflowing stack. Would an overflow be a big issue for a command-line calculator? Yes! This is meticulous design, I wasn't doing things just because back then.

It just goes to shit from there anyway...

```cpp
int main () {
  double a,b,r,memory=0;
  char op, choice[10];


  cout<<"+ to add, - to subtract, * to multiply, / to divide, \
	^ to power, ! to factorial\n\n";

  start: // THE BEGINING


  cin>>a; // 1ST NUMBER
  reused: //
  cin>>op; // OPERATOR
  // WHAT TO DO
  if (op=='!') r=factorial(a);
  else {
        cin>>b;
        if (op=='+') r=a+b;
    	if (op=='-') r=a-b;
	    if (op=='*') r=a*b;
    	if (op=='/') r=a/b;
	    if (op=='^') r=pow(a,b);
	    }
  // COUT THE RESULT
  cout<<"="<<r<<endl;

```

Ignore _start:_ and _reused:_, we'll come to those later. Speaking of meticulous design, the way we'd build calculators 1-2 years later (in college) is by asking for the numbers and the operator separately, e.g. like this:

```bash
Enter 1st number: 5
Enter 2nd number: 8
Enter the operator: +
Your result is: 13
```

I was so ahead of my time, my calculator would do:

```bash
+ to add, - to subtract, * to multiply, / to divide, ^ to power, ! to factorial
5+8 (<<< the user input)
=13
```

I really thought doing it inline was like some secret grail that nobody was searching for but it'd be fun if they found it. Continuing strong with the comments, I don't know if I'd understand what's going on here without the comment:

```cpp
// COUT THE RESULT
cout<<"="<<r<<endl;
```

For some time, I was actually thinking of cout (/kəʊt/) as a verb meaning "spit out" as in "spit it out, machine!", a verb which Urban Dictionary ruined for me.

Next on, the memory:

```cpp
// CREATING THE MEMORY
cout<<"\nType 'mi' to insert the number into memory, or 'mc' to clear memory\n";
cin>>choice;
if (!strcmp(choice,"mi")) memory=r; // INSERT THE MEMORY
else if (!strcmp(choice,"mc")) memory=0; // CLEAR THE MEMORY
else cout<<"command unknown, program will go on\n"; // DEFAULT
```

If you recall the old calculators had [these MC/MR/MS/... buttons](https://www.quora.com/What-do-MC-MR-MS-M+-and-M-in-calculators-do), used for storing some result to the memory and using in at a later time. It's not difficult to implement, and maybe some scientific calculators still use these, but I haven't seen them implemented anymore. A relic of the past, if you will. I used MI (Memory Insert) instead of MS (Memory Store) because... reasons. But at least the comments are useful this time.

And finally, you reach the spaghettiest of the codes, the _poo de la poo_, the "base for the loop" (yeah I have no idea what that means either):

```cpp
// CREATING THE BASE FOR THE LOOP
cout<<"\nType 'restart' to start again from the beggining, \
	'reuse' to use the result, \n'mr' to reuse the number \
	in the memory, or 'quit' to quit: ";

cin>>choice; // TELL THE PROGRAM WHAT TO DO
if (!strcmp(choice,"restart")) goto start; // START FROM THE BEGINING
if (!strcmp(choice,"reuse")) { a=r; cout<<a; } // USE THE RESULT
if (!strcmp(choice,"reuse")) goto reused; // START WITH THE RESULT
if (!strcmp(choice,"quit")) goto end; // QUIT THE PROGRAM
if (!strcmp(choice,"mr")) { a=memory; cout<<a; } // USE THE MEMORY
if (!strcmp(choice,"mr")) goto reused; // START WITH THE MEMORY
getchar();
end: // TO QUIT THE PROGRAM
return 0;
```

So after you get your result, you get asked if you want to calculate another thing (restart), use the result (the ANS button in calculators), use a stored result or quit. So many goto calls, no proper indentation, useless comments that don't explain anything and for some reason, there are two if clauses that check whether the choice is "reuse" or "mr". Why is that, past me? Couldn't you do the assignment and the goto jump in the same if body?

Funny how they never mentioned goto calls in college, it really made it easier thinking in JMP terms when they were teaching us assembly.

Anyway, it still works. Here's the link if you wanna try for yourself:

https://gist.github.com/aziflaj/68813816ef298a03cd4b

![white-iterm-dont-care](https://github.com/aziflaj/aziflaj.github.io/blob/main/assets/images/20210313/output.png?raw=true)

Try calculating 11 factorial and see if you can fix it :)
