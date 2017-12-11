package slack

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const ENDPOINT = "https://slack.com/api/"

type Client struct {
	token string
}

type Message struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func NewClient(token string) *Client {
	return &Client{token}
}

func (client *Client) Send(channel string, message []Message) error {
	attachments, err := json.Marshal(message)

	if err != nil {
		return err
	}

	_, err = http.PostForm(ENDPOINT+"chat.postMessage", url.Values{
		"token":       {client.token},
		"channel":     {channel},
		"as_user":     {"true"},
		"attachments": {string(attachments)},
	})

	return err
}
