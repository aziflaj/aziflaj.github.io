---
layout:     post
title:      "Week #12: Uber and Flappy Bird"
date:       2016-03-05
summary:    "This week I went back into following the Android course by Rob Percival, building an Uber clone and after that, a Flappy Bird clone. During this time, I was working with Parse Geolocation API and libGDX for developing the game. Also I followed an introduction to Android Wear, how to build simple applications for smart watches and test these apps in the emulator."
comments:   true
tags:       [challenge, android, uber, game, libgdx]
---

This week I went back into following the Android course by [Rob Percival](https://twitter.com/techedrob), building an Uber clone and after that, a Flappy Bird clone. During this time, I was working with Parse Geolocation API and libGDX for developing the game. Also I followed an introduction to Android Wear, how to build simple applications for smart watches and test these apps in the emulator.

## Building the Uber clone
I actually started working on Suber, the Uber clone, almost one month ago. Unfortunately, I was forced to stop developing it because of the exams, and now I finally finished it. 

Suber, a wordplay between "super" and "uber", uses Parse as a Backend to allow riders request taxis and drivers pick up these riders. The location of the rider is shown on the map, using Google Map API. If the user is registered as a rider, all he can do is request a taxi or cancel a request by pressing a button. If the user is a driver, then he can see a list of riders distance and can choose one of them to give a ride. 

As for Android Marshmallow, the user can disable some of the permissions of the application, so the developer has to check manually if the permissions are allowed by the user. In the case of the location permissions, you have to add this piece of code anytime you are accessing location services:

```java
if (ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION)
        != PackageManager.PERMISSION_GRANTED

        && ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_COARSE_LOCATION)
        != PackageManager.PERMISSION_GRANTED) {

    Toast.makeText(this, "Don't have location permission", Toast.LENGTH_LONG).show();
    // Do something here...
}
```

The new thing I learned while building Suber was the usage of `ParseGeoPoint` to store the location of the user in Parse, and after to use that location to find the closest riders to a driver. Parse allows you to store the location of a user by using a `ParseGeoPoint` instance and then execute Geo-based queries. I used these queries to fetch a list of riders in a radius of 100 kilometers from the driver:

```java
// mLocationManager and mBestProvider are already initialized
Location last = mLocationManager.getLastKnownLocation(mBestProvider);
if (last != null) {
  ParseGeoPoint driverLocation = new ParseGeoPoint(last.getLatitude(), last.getLongitude());
  // "Requests" is the class (table) that holds the requests of riders for a driver
  ParseQuery<ParseObject> query = ParseQuery.getQuery("Requests"); 
  query.whereWithinKilometers("location", driverLocation, 100); // geo-based query
  query.whereDoesNotExist("driverUsername"); // if the driverUsername exists, the rider already has found a taxi
  query.findInBackground(new FindCallback<ParseObject>() {
    @Override
    public void done(List<ParseObject> objects, ParseException e) {
      if (e == null && objects.size() > 0) {
        for (ParseObject obj : objects) {
          // Use the objects
        }
      }
    }
  });
}
```

You can see that the `FindCallback` is implemented as a SAM (Single Abstract Method) interface. For the apps to cover most of the devices, Android developers should use Java 6 and not benefit from the facilities of newer Java versions, like lambdas of Java 8. Nevertheless, there are (at least) two alternatives which you can use here: Retrolambda ([orfjackal/retrolambda](https://github.com/orfjackal/retrolambda)) and [Kotlin](https://kotlinlang.org/). Retrolambda allows you to use Java 8 lambdas in Android, while Kotlin is a new language that supports lambda and is compiled into Java 6 bytecode, which can be used by Android apps. I call it "Swift for Android" and you can read one of my articles about using Kotlin on Android here: [Streamline Android Java Code with Kotlin](http://www.sitepoint.com/streamline-android-java-code-with-kotlin/).

> That article was written before Kotlin hit beta, so it might be a bit out of date. Anyhow, most of the code should work fine.

Building Suber wasn't that hard. In other applications, I have used the location and Google Map API, so while working on Suber I just had to go check the other apps when I forgot something. The documentation of Parse is a good resource (actually, the best) when it comes to learning the Parse API. Unfortunately, Facebook is closing Parse down, which will become unavailable next in January 2017. This means that after one year, these applications won't work unless I host a Parse Server ([ParsePlatform/parse-server](https://github.com/ParsePlatform/parse-server)) on [Heroku](https://devcenter.heroku.com/articles/deploying-a-parse-server-to-heroku) or any other similar service.

## Building Flappy Bird with libGDX
LibGDX is a Java game development framework for building desktop, mobile and HTML5 games. I used libGDX for building a Flappy birds clone which I called _"Flippy Bird"_. The whole game is developed in a single Java class provided by the libGDX builder. This class is fairly simple:

```java
package com.aziflaj.flippybird;

import com.badlogic.gdx.ApplicationAdapter;

public class FlippyBird extends ApplicationAdapter {
  @Override
  public void create() {
    // set up the game components
  }
  
  @Override
  public void render() {
    // render the game view
  }
}
```

The `create` method is called once in the beginning of the game. In that method I created all the sprites as `Texture` instances and sat up some other components, like the shapes for detecting collisions between objects. The `render` method is called repeatedly to paint the screen for the game, making all the changes. In the `render` method I've put all the logic of the game, like where to draw tubes for the game, randomize the gap location between the tubes (moving it up and down), what happens when the user taps the screen, etc.

![flippy bird](https://raw.githubusercontent.com/aziflaj/aziflaj.github.io/master/images/52-projects/week12/flippy.png)

I really enjoyed building a game using libGDX. The library seems both simple to use and very powerful, but this might be the opinion of a noob game developer after building a simple game using libGDX. 

I also learned how to build simple Android Wear applications. I've actually never used an Android Wear device except the emulator, so I am not to be count as an expert. But I learned how to build simple apps, connect them to an Android device, create notifications and layouts for different screens (round and square), and this was nice for a 12th week of the Project 52.
