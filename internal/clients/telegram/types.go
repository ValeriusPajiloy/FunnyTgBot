package telegram

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}
type ChatResponse struct {
	Ok     bool `json:"ok"`
	Result Chat `json:"result"`
}
type ChatMemberResponse struct {
	Ok     bool         `json:"ok"`
	Result []ChatMember `json:"result"`
}
type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type From struct {
	Username string `json:"username"`
}

type Chat struct {
	ID              int      `json:"id"`
	ActiveUsernames []string `json:"active_usernames"`
}

type ChatMember struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
