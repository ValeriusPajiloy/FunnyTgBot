package telegram

import (
	"context"
	"errors"
	"fmt"

	"tgbot/internal/clients/telegram"
	"tgbot/internal/events"
	"tgbot/internal/modules/notification"
	tagall "tgbot/internal/modules/tag_all"
)

var ErrUnknownMessageType = errors.New("UnknownMessageType")
var ErrUnknownMetaType = errors.New("UnknownMetaType")

type Worker struct {
	tg     *telegram.Client
	notify *notification.Notification
	tagall *tagall.TagAll
	offset int
}

type Meta struct {
	ChatId          int
	SenderName      string
	ActiveUsernames []string
}

func NewWorker(tg *telegram.Client, notify *notification.Notification, tagall *tagall.TagAll) *Worker {
	return &Worker{
		tg:     tg,
		notify: notify,
		tagall: tagall,
	}
}

func (w *Worker) Fetch(limit int) ([]events.Event, error) {
	updates, err := w.tg.Updates(w.offset, limit)
	if err != nil {
		return nil, fmt.Errorf("cant Fetch %s", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, update := range updates {
		if update.Message != nil {
			res = append(res, event(update))
		}
	}

	w.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (w *Worker) Process(ctx context.Context, event events.Event) error {
	switch event.Type {
	case events.Message:
		w.processMessage(ctx, event)
	case events.Unknown:
		return fmt.Errorf("cant Process %s", ErrUnknownMessageType)
	}
	return nil
}

func (w *Worker) processMessage(ctx context.Context, event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return fmt.Errorf("cant processMessage %s", err)
	}

	if err := w.doCommand(ctx, event, meta); err != nil {
		return fmt.Errorf("cant processMessage %s", err)
	}
	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, fmt.Errorf("cant meta %s", ErrUnknownMetaType)
	}
	return res, nil
}
func event(update telegram.Update) events.Event {
	fetchType := fetchType(update)

	res := events.Event{
		Type: fetchType,
		Text: fetchMessage(update),
	}

	if fetchType == events.Message {
		res.Meta = Meta{
			ChatId:          update.Message.Chat.ID,
			SenderName:      update.Message.From.Username,
			ActiveUsernames: update.Message.Chat.ActiveUsernames,
		}
	}
	return res
}

func fetchType(update telegram.Update) events.Type {
	if update.Message == nil {
		return events.Unknown
	}
	return events.Message
}
func fetchMessage(update telegram.Update) string {
	if update.Message == nil {
		return ""
	}
	return update.Message.Text
}
