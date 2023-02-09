package clients

type Client interface {
	SendMessage(chatID int, text string) error
}
