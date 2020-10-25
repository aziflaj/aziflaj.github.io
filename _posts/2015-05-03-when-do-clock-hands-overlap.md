---
layout: post
title: When do the minute and hour hands of a clock overlap?
comments:   true
summary: When? When?!
category: 
    - python
    - clock
---

A really nice logical question is _"How many times a day do the minute and hour hands of a clock overlap?"_. Well that's easy, 22 times. During the first half of the day, they overlap 11 times: 00:00, ~01:05, ~02:10, and so on, but not ~11:55 (since they overlap at 00:00). In 24 hours, that would be 22 times. I'm not trying to find out **how many times**, but **when** does this happen.

> Caution! If you think you will lose your time while I solve this task just for the fun of solving it, please don't go any further.

Let's think of the hour hand. In one hour, it makes `1/12` of a full circle, which means `1/12` of 360 degrees. So the [angular velocity](http://en.wikipedia.org/wiki/Angular_velocity) of the hour hand is **30 degrees/hour** or **0.5 degrees/min** or **1 degree every 2 minutes**.

What about the minute hand? It makes a full circle in one hour, which means **6 degrees/min**. That's basic maths, every 6th grader can calculate these. Also, every 6th grader can calculate when do the clock hands overlap, but I'm going to use [Python](http://python.org/) for this task.

Now we should consider that if we are calculating the time when the hands overlap, for example, when the hour is 1, it is better to start checking after 01:05, or in general `5 * h` minutes when the hour is `h`. I am now writing a simple function to calculate the angle of each hand for a given hour:

{% highlight python %}
# When do the minute and hour hands of a clock overlap?

def hourly_overlap(hour):
	mins = hour * 5
	mins_initial_position = mins * 6
	hour_initial_position = hour * 30 + 0.5 * mins

	print (" hours: {0}\n mins: {1}\n hours angle: {2}\n mins angle: {3}".format(hour, mins, hour_initial_position, mins_initial_position) )


#main
hourly_overlap(1)
{% endhighlight %}

This would output:
{% highlight bash %}
hours: 1
mins: 5
hours angle: 32.5
mins angle: 30
{% endhighlight %}

So when the time is 01:05, the hour hand is shifted 32.5 degrees from the base position (which is hour 12) while the minutes hand is shifted 30 degrees from the base position. 

Now we are going to "move" the minutes hand by one minute (6 degrees). If we find a time when the minutes and the hour hands are shifted with the same angle, that is when the overlap happens. If the minutes hand passes the hour hand, we found an interval when the minutes hand overlaps the hour hand.

To do this, I changed my program to this:

{% highlight python %}
# When do the minute and hour hands of a clock overlap?

h_velocity = 0.5 # degrees/minute
min_velocity = 6 # degrees/minute


def hourly_overlap(hour):
	mins = hour * 5
	mins_position = mins * 6
	hour_position = hour * 30 + 0.5 * mins

	while True:
		# travel one minute ahead in time
		mins += 1
		if (mins >= 60):
			mins = 0
		mins_position += min_velocity
		hour_position += h_velocity

		# check if the angles are the same
		if mins_position == hour_position:
			print("{0:02d} : {1:02d}\n".format(hour, mins))
			break

		# if the mins angle passes the hour angle, we found an interval
		elif mins_position > hour_position:
			print("The hands overlap between {0:02d}:{1:02d} and {0:02d}:{2:02d}\n".format(hour, mins-1, mins))
			break

		# otherwise loop

#main
hourly_overlap(1)
{% endhighlight %}

It outputs this: **The hands overlap between 01:05 and 01:06**. And it's `true`! When it's 01:05 the hands are respectively at 32.5 degrees and 30 degrees, while at 01:06 they are at 33 and 36 degrees. 

Now what we have to do, is add a loop to calculate when does the overlapping happen for any hour from 00:00 to 11:00. Before that, I'd like to add a use case which we already know the result. If the hour is 11 or 12 or 00, we know that the overlap happens at 00:00. so the final code would be:

{% highlight python %}
# When do the minute and hour hands of a clock overlap?

h_velocity = 0.5 # degrees/minute
min_velocity = 6 # degrees/minute


def hourly_overlap(hour):
	if hour == 11 or hour == 12 or hour == 0:
		print("At {0:02d} o'clock, the overlap happens at 00:00".format(hour))
		return

	mins = hour * 5
	mins_position = mins * 6
	hour_position = hour * 30 + 0.5 * mins

	while True:
		# travel one minute ahead in time
		mins += 1
		if (mins >= 60):
			mins = 0
		mins_position += min_velocity
		hour_position += h_velocity

		# check if the angles are the same
		if mins_position == hour_position:
			print("At {0:02d} o'clock, the overlap happens at {0:02d}:{1:02d}".format(hour, mins))
			break

		# if the mins angle passes the hour angle, we found an interval
		elif mins_position > hour_position:
			print("At {0:02d} o'clock, the overlap happens between {0:02d}:{1:02d} and {0:02d}:{2:02d}".format(hour, mins-1, mins))
			break

		# otherwise loop

#main
i = 0

while (i<12):
	hourly_overlap(i)
	i += 1
{% endhighlight %}

And the output is:
{% highlight bash %}
At 00 o'clock, the overlap happens at 00:00
At 01 o'clock, the overlap happens between 01:05 and 01:06
At 02 o'clock, the overlap happens between 02:10 and 02:11
At 03 o'clock, the overlap happens between 03:16 and 03:17
At 04 o'clock, the overlap happens between 04:21 and 04:22
At 05 o'clock, the overlap happens between 05:27 and 05:28
At 06 o'clock, the overlap happens between 06:32 and 06:33
At 07 o'clock, the overlap happens between 07:38 and 07:39
At 08 o'clock, the overlap happens between 08:43 and 08:44
At 09 o'clock, the overlap happens between 09:49 and 09:50
At 10 o'clock, the overlap happens between 10:54 and 10:55
At 11 o'clock, the overlap happens at 00:00
{% endhighlight %}

I wouldn't like to go ahead and find the second when the overlap happens, but it is possible and it's not very hard to do it. If you would like to do it, _please comment below with a [gist](https://gist.github.com/) of the code._
