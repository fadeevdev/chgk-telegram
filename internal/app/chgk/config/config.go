package config

import (
	"os"
)

func ParseConfig() (*Config, error) {
	c := Config{}

	c.Port = os.Getenv("PORT")
	c.ApiKeys.Telegram = os.Getenv("TELEGRAM_BOT_TOKEN")

	return &c, nil
}
