package telegram

import (
	"strings"
	"tgbot/internal/events"
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
			case START_CMD:
				w.tg.SendMessage(meta.ChatId, MSG_HELP)
			case HELP_CMD:
				if hasKeys {
					if keys[0] == REAL_KEY {
						w.tg.SendMessage(meta.ChatId, MSG_REAL_HELP)
					}
				} else {
					w.tg.SendMessage(meta.ChatId, MSG_HELP)
				}

			case SPELL_CMD:
				w.notify.CreateAnswer(hasKeys, keys, meta.ChatId)

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
				w.tg.SendMessage(meta.ChatId, MSG_UNKNOWN)
			}
		}
	}
	return nil
}
