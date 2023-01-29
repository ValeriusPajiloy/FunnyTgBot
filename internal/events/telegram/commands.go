package telegram

import (
	"strings"
)

// TODO: declare command here
const (
	//cmd
	HelpCmd = "/help"

	//key
	RealKey = "real"
)

func (w *Worker) doCommand(text string, chatId int, username string) error {
	text = strings.TrimSpace(text)
	textSlice := strings.Split(text, " ")

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
					w.tg.SendMessage(chatId, msgRealHelp)
				}
			} else {
				w.tg.SendMessage(chatId, msgHelp)
			}
		default:
			w.tg.SendMessage(chatId, msgUnknown)
		}
	}

	return nil
}
