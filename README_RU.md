[en](README.md)
# Funny Telegram Bot
Простой проект для изучения языка Go.
Telegram-бот с забавным функционалом.

Реализовано
1. [X] [Timer notification.](#notification)  
2. [X] [Tag all group.](#tagAllGroup) 
3. [X] [Chinese social ranking.](#ChineseSocialRanking)  
4. [ ] [Random picture.](#RandomPicture)

___
## Модули:
<a name="notification"><h3>1. Timer notification.</h3></a>
Забавные уведомления, похожие на магические заклинания.

Вы можете установить таймер для заклинания, указав цель

[Подробнее](ru/NOTIFICATION.md)
___
<a name="tagAllGroup"><h3>2. Tag all groupe.</h3></a>
В телеграмме нельзя использовать @all.

Этот модуль позволяет создавать группы, добавлять туда пользователей и тегать всю группу одним сообщением.

Модуль сделан для того, чтобы собрать группу в игре, поэтому синтаксис выглядит так:
`/go dota`

[Подробнее](ru/TAG_ALL_GROUP.md)
___
<a name="ChineseSocialRanking"><h3>3. Chinese social ranking</h3></a>
Идея состоит в том, чтобы хранить рейтинг для каждого пользователя.

Когда пользователь пишет сообщения, в которых сочетаются хорошие слова и слова Китай, партия - рейтинг поднимается.
Когда пользователь пишет сообщение, в котором сочетаются плохие слова и слова Китай, партия, рейтинг падает.

По команде `/rank` бот отображает рейтинг пользователя.

[Подробнее](ru/SOCIAL_RANKING.md)
___
<a name="RandomPicture"><h3>4. Random picture</h3></a>
Модуль позволяет создавать «темы» и добавлять в эти темы картинки.
После этого при использовании `/theme_name` бот будет присылать случайную картинку, добавленную в эту тему

___
Большое спасибо Николаю Тузову, многому научился из его плейлиста "Telegram Bot in Golang", примеры ссылок ниже

https://www.youtube.com/playlist?list=PLFAQFisfyqlWDwouVTUztKX2wUjYQ4T3l

https://github.com/GolangLessons/Read-Adviser-Bot/tree/lessons

Идеи, предложения и советы пишите сюда: https://t.me/PajiloyValera

