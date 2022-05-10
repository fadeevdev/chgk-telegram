package telegram

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	apiKey string
}

func New(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (c *Client) SendMessage(chatID uint64, message string) error {
	cl := http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.apiKey)

	payload := fmt.Sprintf(`
		{
			"chat_id": %d,
			"text": %q
		}
	`, chatID, message)

	resp, err := cl.Post(url, "application/json", bytes.NewBuffer([]byte(payload)))

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return err
}
