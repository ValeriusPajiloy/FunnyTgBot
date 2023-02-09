# Funny Telegram Bot
A simple pet project for learning the Go language. 
Telegram bot with rofl functionality
## -------------------------------------------------
### READY:
### ---------------------
#### Notification module.
Funny notification image like magik spells.
Use
> /колдую SPELL TIME TARGET 

After the specified TIME of minutes bot write:
> Сейчас у TARGET SPELL

Also bot will be notify about SPELL on 30 min, 15 min, 5 min and every 5 last min before spell time like:
> SPELL у TARGET через 30 минут

##### Example
Use:
> /колдую понос 45 @Valera

Result after 15min
> понос у @Valera через 30 минут

Result after 45min
> Сейчас у @Valera понос

### ---------------------
### IN WORK:
### ---------------------
#### Tag all to game module.
Usage:

For tag all in group:
> /go *group*

Add new group:
> /go addgroup *group*
	
Get all groups for chat
> /go groups

Add users in group:
> /go add *group* *userTag*(maybe more one)`

Example:

Your message:
> /go dota

Answer:
> го в dota @Valera @Max @Egor @Kolya
### ---------------------
### IN PLANS:
### ---------------------
#### Chinese social ranking

The idea is to store a rating for each user.

When a user writes messages that combine good words and the word China, the party - the rating goes up.
When a user writes a message that combines bad words and the words China, the party, the rating goes down.

on the /rating command, the bot displays the user's rating.

## -------------------------------------------------
Many thanks to Nikolai Tuzov, I learned a lot from his "Telegram Bot in Golang" playlist, examples of links below

https://www.youtube.com/playlist?list=PLFAQFisfyqlWDwouVTUztKX2wUjYQ4T3l

https://github.com/GolangLessons/Read-Adviser-Bot/tree/lessons

Ideas, suggestions and advice read here: https://t.me/PajiloyValera

