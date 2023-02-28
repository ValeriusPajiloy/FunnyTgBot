[ru](README_RU.md)
# Funny Telegram Bot
A simple pet project for learning the Go language. 
Telegram bot with funny functionality


Realized
1. [X] [Timer notification.](#notification)  
2. [X] [Tag all group.](#tagAllGroup) 
3. [ ] [Chinese social ranking.](#ChineseSocialRanking)  
4. [ ] [Random picture.](#RandomPicture)
___
## Modules:
<a name="notification"><h3>1. Timer notification.</h3></a>
Funny notifications that look like magic spells.

You can set a timer for a spell by specifying a target

[More](docs/en/NOTIFICATION.md)
___
<a name="tagAllGroup"><h3>2. Tag all groupe.</h3></a>
There is no way to use @all in telegram.

This module allows you to create groups, add users there and tag the entire group with one message

The module was made in order to collect a group in the game, so the syntax looks like this:
`/go dota`

[More](docs/en/TAG_ALL_GROUP.md)
___
<a name="ChineseSocialRanking"><h3>3. Chinese social ranking</h3></a>
The idea is to store a rating for each user.

When a user writes messages that combine good words and the word China, the party - the rating goes up.
When a user writes a message that combines bad words and the words China, the party, the rating goes down.

on the `/rank` command, the bot displays the user's rating.

[More](docs/en/SOCIAL_RANKING.md)
___
<a name="RandomPicture"><h3>4. Random picture</h3></a>
The module allows you to create "themes", and add pictures to these themes.
After that, when using `/theme_name`, the bot will send a random picture added to this topic

___
Many thanks to Nikolai Tuzov, I learned a lot from his "Telegram Bot in Golang" playlist, examples of links below

https://www.youtube.com/playlist?list=PLFAQFisfyqlWDwouVTUztKX2wUjYQ4T3l

https://github.com/GolangLessons/Read-Adviser-Bot/tree/lessons

Ideas, suggestions and advice write here: https://t.me/PajiloyValera

