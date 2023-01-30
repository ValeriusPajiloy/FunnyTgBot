package telegram

import (
	"fmt"
	"strconv"
	"strings"
	"tgbot/internal/events"
)

// TODO: declare command here
const (
	//cmd
	HelpCmd = "/help"
	//game
	DotaCmd    = "/dota"
	MainCmd    = "/main"
	ValheimCmd = "/valheim"
	//notifications
	SpellCmd = "/колдую"
)

const (
	//key
	RealKey = "real"
)

func (w *Worker) doCommand(event events.Event, meta Meta) error {
	text := strings.TrimSpace(event.Text)
	textSlice := strings.Split(text, " ")
	if text != "" {

		command := text[0] == '/'
		if command {
			cmd := textSlice[0]

			hasKeys := len(textSlice) > 1
			keys := make([]string, len(textSlice)-1)
			for i := 0; i < len(textSlice)-1; i++ {
				keys[i] = textSlice[i+1]
			}

			//TODO: Add commands here
			switch cmd {
			case HelpCmd:
				if hasKeys {
					if keys[0] == RealKey {
						w.tg.SendMessage(meta.ChatId, msgRealHelp)
					}
				} else {
					w.tg.SendMessage(meta.ChatId, msgHelp)
				}

			case SpellCmd:
				if hasKeys {
					if len(keys) < 3 {
						if keys[0] == "help" {
							w.tg.SendMessage(meta.ChatId, msgHelpSpell)
						} else {
							w.tg.SendMessage(meta.ChatId, msgNotFullSpell)
						}
					} else {
						spell := keys[0]
						timeMin, err := strconv.Atoi(keys[1])
						target := keys[2]
						if err != nil {
							w.tg.SendMessage(meta.ChatId, msgNotFullSpell)
						} else {
							w.tg.SendMessage(meta.ChatId, fmt.Sprintf(msgSpell, timeMin, target, spell))
							w.notify.CreateNotifies(spell, timeMin, target, meta.ChatId)
						}
					}
				} else {
					w.tg.SendMessage(meta.ChatId, msgEmptySpell)
				}
			// case DotaCmd:
			// 	tegAll := ""
			// 	//TODO: think about select all users
			// 	w.tg.SendMessage(meta.ChatId, fmt.Sprintf("%s %s", tegAll, msgGoDota))
			// case MainCmd:
			// 	tegAll := ""
			// 	//TODO: think about select all users
			// 	w.tg.SendMessage(meta.ChatId, fmt.Sprintf("%s %s", tegAll, msgGoMain))
			// case ValheimCmd:
			// 	tegAll := ""
			// 	//TODO: think about select all users
			// 	w.tg.SendMessage(meta.ChatId, fmt.Sprintf("%s %s", tegAll, msgGoValheim))
			default:
				w.tg.SendMessage(meta.ChatId, msgUnknown)
			}
		}
	}
	return nil
}
