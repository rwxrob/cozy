# ☕️ Cozy Community Builder

***Work in progress. Watch for v0.1+ before attempting to use for anything.***

Cozy is a tool designed to bring people together in real-time to promote progress for everyone in a collaborative community by sharing what people are doing, what they want to do, and discovering others they can help. Members can be working on the same project for the same organization or entirely different ones. Cozy aspires to be more that just another social media platform. It's the missing part of most social media platforms that is required to help each other know follow and assist others with their specific objectives *in real-time* in a way that promotes casual, organic help and mentoring when the need arises simply by helping everyone know what everyone else is doing. In fact, cozy was inspired by multi-user UNIX systems where having a `.project` and `.plan` made following what others where doing a normal thing. Cozy restores this part of what social media has seem to have forgotten about the successful collaborative systems of the past.

## Sharing...

People work better when they can share what they are doing easily. "Coworking" has emerged as the way to do this. Whether it is a group of people working for different companies all at the same coffee shop or people at the same company joined by (unproductive) cubicles, the idea is the same. Work *near* others (if not *with* them). Remote work has motivated the creation of several approaches to this idea but the original digital coworking platform has existed since the 70s: UNIX.

People have been `finger`-ing one another, chatting over `ntalk` and `irc`, and sharing `.plan` and `.project` files since before today's social media billionaires were even born. One could argue that all modern social media platforms are derivatives of these original multi-user UNIX host systems and tools. You would think people designing modern social media would actually learn from these original designs, but sadly they have not. Cozy aspires to address these omissions in a modern way:

* REST OpenAPI (leveraging industry grade TLS encryption)
* Static files with no database dependency
* Builtin Twitch, GitHub, Slack, and Discord integration

See the `cozy help` page for more information.

## Simple HTTP server

A Cozy server is just a plain old HTTP server with the content organized in a predictable way (nginx, Apache, `net/http`, etc.). That's it. This means that people can cache, snapshot, and exchange entire community servers with a single `wget` command. The magic as to the presentation and searching and notifications from this content happens in any of the different clients that store and index this data in different ways.

The magic that enables this to work is simple: limit the amount and time-to-live of all the data. Don't update your data within a give time frame and it gets cleaned up automatically releasing any unique ID you might have been squatting on.

## *Everything* is public

That's right. There are no pretense to privacy here, no "sliding into someone's DMs." This platform is 100% public. Sure you can leave a message for someone with `tell` but everyone else will see it as well. After all, this is a coworking space. If you do want to get all whispery with someone from the community find another corner of the Internet to cozy up (you'll have no problem finding a place, if you haven't already). Keeping everything public also allows the community to police itself and quickly shut down harassment from anonymous trolls.

## Authenticated external account required

The Internet has already figured out how to ensure people are authenticated properly. Pick one of the many OIDC providers on the Internet or allow people to login using any one of them. Here's a short list of the most common ones:

* GitHub
* Google
* YouTube
* Twitch
* Facebook

Requiring a person have one of these accounts eliminates most of the risk of harassment and abuse, but not all. The rest can be policed by the community, or the community moderators.

## What about chat?

There is no chat. That has been covered well by Slack, Discord, and IRC. But we do plan to create as many integrations into existing chat and social media platforms as possible, which isn't hard because everyone in a community can easily map their socials to their unique ID enabling things like `!who rwxrob` and `!work rwxrob` from any system.

## What about decentralized social media?

First of all, this is a coworking sharing platform, which qualifies as social media, but only barely. This isn't the next Slack, or Facebook, or MS Teams. It's a simple way to share small bits of on-demand information with one another using a opt-in, event-driven subscriber model design. Subscriber models fundamentally depend on a hub of some kind. The alternative, forcing every single person using the tool to also be a distributor, is draconian at best. In fact, most peer-to-peer designs forget about these fundamental design requirements for any such system.

Besides, centralization isn't automatically evil. It's *how* stuff is centralized and what is the scope of the information. Since Cozy assumes every single bit of data placed onto it is meant to be public there is no need for concern about monitoring, in fact, the more monitoring the better. And control? Well, because this project is open source and meant to help communities build their own silos there is nothing stopping servers from federating the content, like Mastodon, but with a coworking focus.
