package config

type ApiKeys struct {
	Telegram string
}

type Postgres struct {
	Host     string
	Port     string
	DbName   string
	DbUser   string
	Password string
}

type Config struct {
	ApiKeys  ApiKeys
	Port     string
	Postgres Postgres
}
