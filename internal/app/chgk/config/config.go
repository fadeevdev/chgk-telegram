package config

import (
	"os"
)

func ParseConfig() (*Config, error) {
	c := Config{}

	c.Port = os.Getenv("PORT")
	c.ApiKeys.Telegram = os.Getenv("TELEGRAM_BOT_TOKEN")
	c.Postgres.Host = os.Getenv("POSTGRES_HOST")
	c.Postgres.Port = os.Getenv("POSTGRES_PORT")
	c.Postgres.DbName = os.Getenv("POSTGRES_DBNAME")
	c.Postgres.DbUser = os.Getenv("POSTGRES_USER")
	c.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")

	return &c, nil
}
