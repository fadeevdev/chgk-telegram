package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
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

func (c *Client) SendMessage(chatID uint64, message string) (*pb.Message, error) {
	cl := http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.apiKey)

	payload := fmt.Sprintf(`
		{
			"chat_id": %d,
			"text": %q
		}
	`, chatID, message)

	resp, err := cl.Post(url, "application/json", bytes.NewBuffer([]byte(payload)))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	respMessJson := &Message{}
	respMess := &pb.Message{}

	respMess.From = &pb.User{}

	json.NewDecoder(resp.Body).Decode(respMessJson)
	respMess.Id = uint64(respMessJson.Result.From.Id)
	respMess.From.Id = uint64(respMessJson.Result.From.Id)
	respMess.From.IsBot = respMessJson.Result.From.IsBot
	respMess.From.FirstName = respMessJson.Result.From.FirstName
	respMess.From.Username = respMessJson.Result.From.Username
	respMess.Date = uint64(respMessJson.Result.Date)

	return respMess, err
}
