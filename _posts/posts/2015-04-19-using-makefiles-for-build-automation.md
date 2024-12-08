---
layout: post
title: Using Makefiles for build automation
comments:   true
summary: How did they do it before pipelined Continuous Integration were a thing?
category: 
    - makefile
    - C/C++
---

In my [last presentation](http://aziflaj.github.io/real-world-webapp/) at Things Lab, I mentioned **build automation** as one of the best practices used in Continuous Integration (CI). But build automation goes even beyond CI; you don't neccessarily have to use CI to use automated builds. Actually, build automation dates even before CI was mentioned for the first time.

There are different utilities that help you automate the build process. If you ever programmed in Java you probably have heard of tools such as [Apache Ant](http://ant.apache.org/), [Maven](https://maven.apache.org/) or [Gradle](https://gradle.org/). All three of them are build tools that help you automate the compiling and building process when you develop a software. And that's not all. These tools also help you deploy the WAR (**W**eb **AR**chive) into the development/production server; or obtain all the dependencies of your application; or building the APK (Android Package) file for different Android versions; etc. But compared to a utility called [`make`](https://www.gnu.org/software/make/manual/make.html#Introduction), these tools are new and with a lot more features.

You might be asking yourself "Why do I need to use any of these tools when my IDE takes care of the building process?". The answer is: "You don't!". It is true that most (if not all) modern IDEs take care of the building process with much less effort than writing the configuration files for the build automation. But basically what these IDEs are doing, is writing customized build configurations for each project, so they "know" what to do and when to do it, where to get the source files, etc. So if you want to know more of how the IDEs do this, you might want to learn how to use any build tool. Also, if you have to use a machine with limited resources which can't support any powerful IDE, you'd have to write the build configuration by yourself.

The first build automation tool I've been using is probably the oldest I've been using: the `make` utility. This utility allows you to specify a list of files and their dependencies in a special file known as a `Makefile`. Based on this `Makefile`, the `make` program builds your project by automatically recompiling **only when it's necessary**. How? By firstly checking if your source files are changed after the creation of the executable(s). 

So far, I've been using `Makefile`-s for C and C++ programs (see [here](https://github.com/fzaninotto/Faker/blob/master/Makefile) a `Makefile` used for PHP). When you are working on large projects, it is a good idea to use header files. What you do is group all commonly used definitions inside one header file. These definitions might be structures and functions used to manipulate these structures (in both C & C++ ), definitions of one or more classes (in C++ using Object Orientation), global constants or macros definitions (like `EOF` in `<stdio.h>`), etc. Then, you write the `Makefile` with different (or only one) build targets. I am giving here the last `Makefile` I wrote for a small homework:

```Makefile
CC = gcc 
FLAGS = -std=c99 -Ilibs -Isrc
MAIN_FILES = src/app.c src/text.c
TEST_FILES = tests/runtests.c

all:
	$(RM) tests/*.exe
	$(RM) libs/*.exe
	$(RM) src/*.exe
	$(RM) *.exe
	
	$(CC) $(FLAGS) $(MAIN_FILES) -o app

	printf "\n"	
	printf "*********************\n"
	printf "*** READY TO EXEC ***\n"
	printf "*********************"
	printf "\n"


test:
	$(CC) $(FLAGS) $(TEST_FILES) -o test
	
	printf "\n"
	printf "*********************\n"
	printf "*** READY TO TEST ***\n"
	printf "*********************"
	printf "\n"


clean:
	$(RM) tests/*.exe
	$(RM) libs/*.exe
	$(RM) src/*.exe
	$(RM) *.exe

.SILENT: all test clean
.PHONY: all test clean
```

I'm trying to explain some of it, even though I'm not a `Makefile` master. It begins with some macro definitions such as:

 - `CC` for the C compiler 
 - `FLAGS` for the different flags passed to the compiler
 - `MAIN_FILES` to define which are the files to get compiled for the main bundle
 - `TEST_FILES` to define which are the files to get compiled for testing the app
 
After those macros, there are three build targets defined: `all`, `test` and `clean`. By default, the `make` utility calls the first build target defined if no target is passed as parameter (which in our case, is `all`).

To create a build target, all you have to do is firstly define the name of the build target (such as `all` or `test`) followed by `:` and then by all the commands that would build the application based on that build target. Make sure to add a `tab` before each command, or else `make` won't like it at all. In your build targets, you can call the defined macros by using the `$(MACRO)` syntax. 

A couple of targets you would want to use are the `clean` and the `install` target. Check [this `Makefile`](https://github.com/aziflaj/ffos-cli/blob/master/Makefile) I wrote for FFOS-CLI. The `clean` target is used to clean the project directory from different files that might be generated while running the application or testing it. By calling:

```bash
make clean
```

you make sure to remove all these files, since they are not necessary for the project itself. So you run that target anytime you want to share all your code directory with the others. The `install` target does what you are thinking of; installs the application. Referring to the FFOS-CLI's `Makefile`, you install it by executing:

```bash
make
sudo make install
```

It firstly compiles the application based on the first target (which in that case is `ffos`) and then adds it into the `/usr/local/bin` directory so it can be called at any other directory by that user.

Even though `make` is old and probably not used that much today, there are other build tools and dependency managers that are a must for every developer. I can mention Gradle, which now is the default build tool for Android development; Ant and Maven for every old-school Java developer; [Composer](https://getcomposer.org/) for PHP, which is not a build tool but a dependency manager, very useful for all PHP developers; [Bower](http://bower.io/) for Javascrip developers, which does almost the same thing as Composer for PHP, etc. Using these build tools is a must for every developer, and I would recommend all to learn one of these.
