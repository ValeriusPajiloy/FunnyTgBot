package telegram

import (
	"fmt"
	tgclient "tgbot/internal/clients/telegram"
	"time"
)

type Notification struct {
	tg *tgclient.Client
}

func NewNotification(tg *tgclient.Client) *Notification {
	return &Notification{
		tg: tg,
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
		textMsg = fmt.Sprintf(msgSpellThirtyMinutsPrefix, target, spell)
	case 15:
		textMsg = fmt.Sprintf(msgSpellFiveteenMinutsPrefix, target, spell)
	case 5:
		textMsg = fmt.Sprintf(msgSpellSomeMinutsPrefix, five, target, spell)
	case 4:
		textMsg = fmt.Sprintf(msgSpellSomeMinutsPrefix, four, target, spell)
	case 3:
		textMsg = fmt.Sprintf(msgSpellSomeMinutsPrefix, three, target, spell)
	case 2:
		textMsg = fmt.Sprintf(msgSpellSomeMinutsPrefix, two, target, spell)
	case 1:
		textMsg = fmt.Sprintf(msgSpellSomeMinutsPrefix, one, target, spell)
	case 0:
		textMsg = fmt.Sprintf(msgSpellNowPrefix, target, spell)
	}
	timeMin := countMin - until
	time.AfterFunc(time.Minute*time.Duration(timeMin), func() {
		n.tg.SendMessage(chatId, textMsg)
	})
}
