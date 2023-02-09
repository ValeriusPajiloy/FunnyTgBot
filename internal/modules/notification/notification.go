package notification

import (
	"fmt"
	"strconv"
	client "tgbot/internal/clients"
	mongodb "tgbot/internal/storage/mongo_db"
	"time"
)

type Notification struct {
	tg   client.Client
	repo *mongodb.Repository
}

func NewNotification(tg client.Client, repo *mongodb.Repository) *Notification {
	return &Notification{
		tg:   tg,
		repo: repo,
	}
}

func (n *Notification) CreateAnswer(hasKeys bool, keys []string, chatId int) {
	if hasKeys {
		if len(keys) < 3 {
			if keys[0] == "help" {
				n.tg.SendMessage(chatId, MSG_HELP_SPELL)
			} else {
				n.tg.SendMessage(chatId, MSG_NOT_FULL_SPELL)
			}
		} else {
			spell := keys[0]
			timeMin, err := strconv.Atoi(keys[1])
			target := keys[2]
			if err != nil {
				n.tg.SendMessage(chatId, MSG_NOT_FULL_SPELL)
			} else {
				n.tg.SendMessage(chatId, fmt.Sprintf(MSG_SPELL, timeMin, target, spell))
				n.CreateNotifies(spell, timeMin, target, chatId)
			}
		}
	} else {
		n.tg.SendMessage(chatId, MSG_EMPTY_SPELL)
	}
}

func (n *Notification) CreateNotifies(spell string, countMin int, target string, chatId int) error {

	if countMin > 30 {
		n.createTimer(spell, countMin, target, chatId, 30)
	}
	if countMin > 15 {
		n.createTimer(spell, countMin, target, chatId, 15)
	}
	for i := 5; i >= 0; i-- {
		if countMin > i {
			n.createTimer(spell, countMin, target, chatId, i)
		}
	}

	return nil
}

func (n *Notification) createTimer(spell string, countMin int, target string, chatId int, until int) {
	var textMsg string
	switch until {
	case 30:
		textMsg = fmt.Sprintf(MSG_SPELL_THIRTY_MINUTS_PREFIX, target, spell)
	case 15:
		textMsg = fmt.Sprintf(MSG_SPELL_FIVETEEN_MINUTS_PREFIX, target, spell)
	case 5:
		textMsg = fmt.Sprintf(MSG_SPELL_SOME_MINUTS_PREFIX, FIVE, target, spell)
	case 4:
		textMsg = fmt.Sprintf(MSG_SPELL_SOME_MINUTS_PREFIX, FOUR, target, spell)
	case 3:
		textMsg = fmt.Sprintf(MSG_SPELL_SOME_MINUTS_PREFIX, THREE, target, spell)
	case 2:
		textMsg = fmt.Sprintf(MSG_SPELL_SOME_MINUTS_PREFIX, TWO, target, spell)
	case 1:
		textMsg = fmt.Sprintf(MSG_SPELL_SOME_MINUTS_PREFIX, ONE, target, spell)
	case 0:
		textMsg = fmt.Sprintf(MSG_SPELL_NOW_PREFIX, target, spell)
	}
	timeMin := countMin - until
	time.AfterFunc(time.Minute*time.Duration(timeMin), func() {
		n.tg.SendMessage(chatId, textMsg)
	})
}
