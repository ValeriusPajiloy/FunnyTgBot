package telegram

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"tgbot/lib/e"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

// Client work with TelegramApi
type Client struct {
	host     string
	basePath string
	client   http.Client
}

func NewClient(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

//Send message to tgApi
func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)
	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("cant SendMessage ", err)
	}
	return nil
}

//Get updates from tgApi
func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, e.Wrap("cant Updates ", err)
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, e.Wrap("cant Updates ", err)
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
		return nil, e.Wrap("cant do request ", err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, e.Wrap("cant do request ", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, e.Wrap("cant do request ", err)
	}

	return body, nil
}
