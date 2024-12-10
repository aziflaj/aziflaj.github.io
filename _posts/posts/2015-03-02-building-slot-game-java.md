---
layout:     post
title:      Building a Slot Game in Java 
date:       2015-03-02
comments:   true
summary:    The gambler's worst nightmare, a Slot machine where you can't win.
category:
    - java
---

<p>
A couple of years ago, when I was learning Java programming, I thought of testing myself and my programming skills by writing a game in Java. Now, I'm not going to call it <i>"game programming"</i>, since game programming is way more than what I did. In fact, what I did was just a test for me. So I decided to write a game I was playing on my old Nokia E50 phone, a Slot Game.
</p>

This slot game I was playing on my phone was really simple. It had only 3 slots with different items in each. You had to push the **Spin** button in order to spin the slots, and you would win a small amount of coins if two or three slots were alike. Of course, 3 slots were better than 2. It is not really hard to make a game like this, but for a beginner it is good to start with. As I remember, this was probably the first program that I could tell others: <i>"Look at what I just did!"</i>

So I started working on it (I remember using <a href="https://netbeans.org/" target="_blank">NetBeans</a> at that time), firstly as console-only, and then using GUI. The first thing I did, was deciding what kind of images (actually their names, not the images themselves) I would use. I wrote this line of code:

```java
String[] symbols={"Seven","Shamrock","Diamond","3Bar","Star","Bell", "Bar","Orange","Lemon"}; //slot symbols
```

I also decided what would be the amount of "money" that the user would win if he matched two or three symbols:

```java
//amount winning
int[] twoMatches={30,16,15,12,11,10,9,7,5};
int[] threeMatches={60,32,30,24,22,20,18,14,10};
```

After that, I went on writing the code that was suposed to randomly choose one of the elements in the <code>symbols</code> array. This could be done by using the <a href="http://docs.oracle.com/javase/7/docs/api/java/lang/Math.html#random()" target="_blank"><code>Math.random()</code><a/> method, or by calling the <a href="http://docs.oracle.com/javase/7/docs/api/java/util/Random.html#nextInt(int)" target="_blank"><code>nextInt()</code></a> method at a <a href="http://docs.oracle.com/javase/7/docs/api/java/util/Random.html" target="_blank"><code>Random</code></a> instance, or by the wrong way I used to do at the beginning:

```java
// a random from 0 to 9
int randomInt = (int) System.currentTimeMillis() % 10;
```

Of course, I soon switched to calling <code>Math.random()</code>, and in order to get a number that I could use as an index for my array, I wrote this block of code:

```java
double random=Math.random();
random*=8;
int choice=(int) random;
```

So the variable <code>choice</code> was the random index that I could use to get a random item from the array (keep in mind that Math.random() returns a **double** between 0.0 and 1.0)

So after choosing 3 random items, I just printed them out at the console, saying whether there were none, two or three matches, and calculating the amount of money won, if there was any, with the given coefficient. It was a good start; I only had to think of the UI, and I suck at UI design. But for this one, all that I needed was a really simple design which I managed to code as I was planning. 

For the items to show at slots, I just googled them and found 12 of them in a single sprite. I downloaded the sprite and started my old photo editing software which sometimes can really be magical; **Microsoft Paint**! I started cropping images from the sprite, paying attention to their dimensions that should be the same, 122px by 114px. Why these magical values? Just because!

What was left to do, was the UI. I could use the really-helpful drag-and-drop UI builder that ships with NetBeans, but I wanted to do it myself. I had a really hard time figuring out which kind of layout to use, since the only one I really knew was <a href="http://docs.oracle.com/javase/7/docs/api/java/awt/GridLayout.html" target="_blank"><code>GridLayout</code></a>. Anyway beside that, I managed to use <code>FlowLayout</code> and <code>BorderLayout</code>. There is a difference between them, but I'm not really capable of pointing that out, so you can check the online JavaDoc for them.

I managed to build the game, and started to play it. I figured out that the coefficients for multiplying the bet were too damn high, but I didn't care as long as I knew that the game worked. 

![My helpful screenshot](https://github.com/aziflaj/aziflaj.github.io/blob/main/assets/images/slots-screenshots/snapshot.bmp?raw=true)

## My bad practices
As you can see, the <a href="https://github.com/aziflaj/slots" target="_blank">source code</a> of this simple game is in only one file. This is something that I don't like to do anymore. A better way to do it is by making the code as modular as it can be. By building small modular elements, dividing the UI from the logic of the application, you help yourself during the testing and debugging phase. So the first thing that I would like to change, is dividing the whole <code>class Slots extends JFrame</code> from the class that calls it. 

This is done by firstly creating a file called <code>Slots.java</code> that will contain only the code for the UI. Then, creating an <code>ActionListener</code> that will listen to different button clicks (there are 5 different buttons). Finally, creating a class called <code>App.java</code> that will only create a <code>Slots</code> instance and make it run.

Basically, the <code>App.java</code> would look like this:

```java
public class App {
  public static void main(String[] args) {
    javax.swing.SwingUtilities.invokeLater(new Runnable() {
      @Override
      public void run() {
        try {
          Slots gameWindow = new Slots();
          gameWindow.setVisible(true);
        } catch (Exception ex) {
          javax.swing.JOptionPane.showMessageDialog(null,
              ex, 
              "Error", 
              javax.swing.JOptionPane.ERROR_MESSAGE);
          System.exit(1);   
        }
      }
  });
  }
}
```

As I remember, <code>SwingUtilities.invokeLater()</code> is used to divide the UI thread from other threads, so if any UI changes are needed, they won't stall the application.

The Listener class, which might be called something like <code>SlotButtonListener</code>, might be something like this:

```java
public class SlotButtonListener implements ActionListener {
  @Override
  public void actionPerformed(ActionEvent ae) {
    if (ae.getSource() == button1) {
      //do something if button1 is clicked
    } else if (ae.getSource() == button2) {
      //do something if button2 is clicked
    } 
    //...
    
    //and so on
  }
}
```

And finally, in the <code>Slots</code> class there would be only the code for defining the UI of the game. All the buttons would have <code>SlotButtonListener</code> as action listener. 

Anyone who wants to change the code following these advices is free to do it. You can <a href="https://github.com/aziflaj/slots/fork" target="_blank">fork it</a> anytime you want.

_Do you have any Java programming advice for me? Feel free to comment below_
