# Usecases

A lot of these overlap with other existing services (Discord, Twitch, GitHub, IRC, YouTube etc.) but this expands on the coworking nature of the community.

## Verbs

* `who`  - output short summary about a person suitable for all Twitch chat to see
* `stat` - give the over time statistics about a community member
* `ping` - get a real-time status about a person, their mood, are they online, busy, etc.
* `work` - current project summary of what a person is actively working on
* `todo` - public todo list with progress and percentages
* `pomo` - control one's own pomo timer or use the default for the community
* `mood` - what is the person's current mood, this is popular on GitHub as well
* `week` - upcoming public calendar for the coming week (only), optionally recurring
* `note` - send a short offline note to someone (or oneself), with delivery date
* `post` - post a temporary micro-blog message including links to other longer posts
* `time` - start one of several personal timers
* `poke` - setup an alert for oneself sort of like an alarm at a given time, possibly recurring
* `poll` - activate a poll to expire within a given time frame
* `vote` - vote for a currently active poll

## Nouns

* User
  * Member - member of the community
  * Moderator - mid-level perms
  * Owner - community senior management

* Timer
  * Pomo - pomodoro style timer
    * Interval - length of each interval
    * IntervalCount - number of short intervals before long break
  * Countdown - countdown timer leading up to events

* Event - something that happens once at a specific moment in time

## Member

* Member of a community wants to register with the central coworking server.
* CRUD todo item.
* CRUD current todo item.
* CRUD personal timer (including pomo).
* CRUD plan.

## Moderator

* Set a timer (including pomo) for the whole community.
* Post a note for everyone.

## Administrator

* CRUD Moderators
