package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	mongodb "tgbot/internal/storage/mongo_db"
)

const (
	getUpdatesMethod            = "getUpdates"
	sendMessageMethod           = "sendMessage"
	getChatMethod               = "getChat"
	getChatAdministratorsMethod = "getChatAdministrators"
)

// Client work with TelegramApi
type Client struct {
	host     string
	basePath string
	client   http.Client
	repo     *mongodb.Repository
}

func NewClient(host string, token string, repo *mongodb.Repository) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
		repo:     repo,
	}
}

// Send message to tgApi
func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)
	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return fmt.Errorf("cant SendMessage %s", err)
	}
	return nil
}

// Get updates from tgApi
func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, fmt.Errorf("cant Updates %s", err)
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("cant Updates %s", err)
	}
	return res.Result, nil
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	url := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("cant do request %s", err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cant do request %s", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cant do request %s", err)
	}

	return body, nil
}
