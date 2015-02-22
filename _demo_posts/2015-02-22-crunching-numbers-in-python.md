---
layout:     post
title:      Crunching numbers in Python
date:       2015-02-22
summary:    When I finished high school and I was spiritually preparing to go to an Engineering College, I knew that too many calculations were ahead. During the first year, I had to follow 5 mathematical courses (or should I say curses?). That is when I discovered the simplicity of Python, and using it to solve math problems seemed the simplest thing to do.
tags:       [Numberoid, Python, Unit, Test]
---

<p>
When I finished high school and I was spiritually preparing to go to an Engineering College, I knew that too many calculations were ahead. During the first year, I had to follow 5 Math-related courses (or should I say <strong>curses<strong>?). That's why when I discovered the simplicity of Python, I started using it to solve math problems.
</p>

Now probably the ones who use Python as a general-purpose programming language every day are cursing me and saying things such as _"Python is way more than a mathematical tool for you!"_ or _"Go on with Matlab, fool! Python is for other things!"_. I know that, but back then I had never heard of MatLab or any other math-related program. Before, I tried to use Java for solving some mathematical problems, but for small problems, writing Java is boring when compared to Python. So what I loved about Python was its simplicity in both writing and reading the code, something I still consider as one of the best features of Python. But still, Python is not my programming language of choice, except when solving math problems (<a href="http://cdn2.cdnme.se/cdn/8-1/2159399/images/2011/u_mad_bro_picture_challenge_3-s469x428-160564-535_161009110.jpg" target="_blank">UMAD, Python-lover?</a>).

One of the most programming-related and still **not** programming-related mathematical course I followed was Numerical Analysis. I say it is programming-related because all of the mathematical solutions can be simply adapted into code, and yet the subject doesn't require you to do so (Computer Engineering... yeah!). But still, I learned Python for a reason, so I started to write some code in order to solve some of the exercises. I started with matrices since they are simple to implement (already implemented as lists of lists). So I was writing some functions to transpose matrices, find their determinants, inverting them, and also to solve systems of linear equations, something we had to do really **a lot**.

And I was thinking, why don't I go on with this? I know, my simple code will never be as expert as Matlab; comparing it with such a powerful tool like Matlab is like comparing the cutting ability of a katana and a toothpick (Matlab = katana if I got you confused there). But still, let's do it! So I created the repository, called <a href="https://github.com/aziflaj/numberoid/" target="_blank">Numberoid</a>, and pushed the code online. Only 6 commits at the time of writting this blogpost, but who cares? We will say "It's under active development!".

While writing the code, I learned some nice things about writing Python code. Firstly, <a href="http://docs.python-guide.org/en/latest/" target="_blank">The Hitchhiker’s Guide to Python</a> is a good place where you can learn a lot of things about writing good Python code. I learned that in order to create a Python module, all you have to do is create a folder with a file called `__init__.py`, which might even be empty. It acts more or less like a marker of a folder where you place your packages (_I need someone to make this clear for me: are Python packages just like Java packages or namespaces?_).

Another nice thing I learned was Unit Testing using a module called `unittest`, so I put all the tests at a folder called just `/test`. For now, there is only one test case, the one for <a href="https://github.com/aziflaj/numberoid/blob/master/test/testmatrix.py" target="_blank">testing matrices<a/>. Not all the tests cases are implemented. I'm still waiting for a helper with some free time. No Python experience required, just free time, eagerness to learn and internet connection. Also, most (if not all) of the mathematical methods which are going to get implemented can be found on Wikipedia, so all we have to do is find the methods, implement them into functions and merge them into the rest of the code. 

So what I require for now is:
<ol>

<li>A helper (not a <a href="https://s-media-cache-ak0.pinimg.com/736x/94/6c/c7/946cc7383075dc6f3e645a5e27b0d794.jpg" target="_blank">minion<a/>, even though I love minions) to write Python.</li>

<li>Someone who has a bit (or even a <i>byte</i>) of experience with numerical analysis and nymerical methods, but if the not-minion helper has this kind of skill it'd be great.</li>

<li>I don't know what else, but I'll figure out with time</li>

</ol>

So if anyone is interested to help, the repository is online, so feel free to <a href="https://github.com/aziflaj/numberoid/fork" target="_blank">fork it</a>, or if you have any suggestion you can <a href="https://github.com/aziflaj/numberoid/issues" target="_blank">file an issue</a>, or even leave it as a comment in the comment section below.
