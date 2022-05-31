package chgk_api_client

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	Host string
}

func New(host string) *Client {
	return &Client{
		host,
	}
}

func (c *Client) GetRandomQuestion() (*Question, error) {
	cl := http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("%sxml/random/", c.Host)

	resp, err := cl.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %v", err)
	}
	var res Search
	if err := xml.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	if len(res.QuestionList) > 1 {
		return &res.QuestionList[0], nil
	}
	return nil, errors.New("no random questions")
}
