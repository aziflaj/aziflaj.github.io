---
layout: post
title: Using Gradle for Build Automation
comments:   true
summary: Makefiles might be good enough for C/C++ projects, but no Java developer is using them. In this blogpost, I'm going to show you the basics of Gradle usage.
category: 
    - java
    - gradle
    - build tools
    - automation
---

In a previous blogpost I wrote about [using Makefiles for build automation](http://aziflaj.github.io/using-makefiles-for-build-automation/). Makefiles might be good enough for C and C++ projects, but no Java developer is using them ([there are reasons](http://stackoverflow.com/questions/2209827/why-is-no-one-using-make-for-java)). Instead, there are better build tools for Java that are far better than `make`, such as Ant, Maven and Gradle. In this blogpost, I'm going to show you the basics of Gradle usage.

## Why using Gradle
Gradle is a real [polyglot](http://gradle.org/why/polyglot-builds/): Java is not the only language it speaks. You can use it to build C/C++/Objective-C projects, or for other JVM languages such as Scala or Groovy (Gradle is build with Groovy), and even for [web-oriented build automation](https://github.com/filipblondeel/gradle-gulp-plugin)!

When using Gradle, you can choose what to compile. This means that you aren't obligated to build all your project, but even [small fractions of it](https://docs.gradle.org/current/userguide/multi_project_builds.html#sec:execution_rules_for_multi_project_builds). Also, since Gradle can cache things such as test results and artifacts, it saves time by not doing work that is already done.

While Ant and Maven write their logic in XML [configuration] files, Gradle uses **real source code** to express its logic. This _code-over-configuration_ philosophy makes it easier to write concise and powerful build files. And yet, it is extendable by a great amount of plugins.

## Build with Gradle
I used Gradle to build a small, "Hello World" application. It is made of two source files (check the [01 branch](https://github.com/aziflaj/gradle-basics/tree/01)):

```java
// src/main/java/demo/App.java
package demo;

public class App {

  public static void main(String... args) {
    Greeter greeter = new Greeter();
    greeter.sayHello();
  }
}



// src/main/java/demo/Greeter.java
package demo;

public class Greeter {
  private String name;

  public Greeter(String name) {
    this.name = name;
  }

  public Greeter() {
    this("World");
  }

  public void sayHello() {
    String greeting = "Hello, %s!";
    System.out.printf(greeting, this.name);
  }

}
```

Gradle requires a project strucure like:

```bash
[project-name]
  +----src
  |     +---- main
  |     |      `---- java
  |     |
  |     `----tests
  |            `---- java
  |
  `---- build.gradle
```

The `src/main/java` directory is the one where you put the source files. So there I put the `demo` package with the code.

In the `build.gradle` file, you can simply put:

```groovy
// compile java
apply plugin: 'java'

// make the runnable application
apply plugin: 'application'

mainClassName = "demo.App"
```

This one says to Gradle to build the application and to use the `App.java` file as the main file to run the executable. When you run `gradle build` (make sure to install Gradle first), you will see a long list of reports, finishing with "BUILD SUCCESSFUL". It will create a new `build` directory, with the compiled and packaged files in there. You can run the executable by executing `gradle run`.

## Resolve dependencies
If your app is using third-party libraries (a.k.a dependencies), you can get them using Gradle (check the [02 branch](https://github.com/aziflaj/gradle-basics/tree/02)). I am appending these lines in the `build.gradle` file:

```groovy
repositories {
  mavenCentral()
}

dependencies {
  compile "joda-time:joda-time:2.2"
}
```

You can now use the Joda Time library in your app. By adding `mavenCentral()` at `repositories`, we say to Gradle that it can obtain necessary files from the Maven repository, and then we declare which one we need to obtain at `dependencies`.

## Create the wrapper
when you are working on a team, maybe not all of the members will have Gradle installed, and not all of them must have it installed in order to use it. Instead, one of the members can create a Gradle wrapper which can then be used as a legit Gradle build tool. Append this task in the `build.gradle` file:

```groovy
task wrapper(type: Wrapper) {
    gradleVersion = '2.3'
}
```

When you run `gradle wrapper`, some files will be generated to allow others use gradle (check the [03 branch](https://github.com/aziflaj/gradle-basics/tree/03)), such as gradlew.sh for Linux/OSX users and gradlew.bat for Windows users. Make sure to include the added files in Version Control!

So that was all. I hope you enjoyed trying Gradle and probably are thinking of using it as the build tool of choice. You can check out the repository [here](https://github.com/aziflaj/gradle-basics)
