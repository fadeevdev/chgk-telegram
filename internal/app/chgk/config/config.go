package config

import "gopkg.in/yaml.v2"

type configFile struct {
	ApiKeys struct {
		Telegram string `yaml:"telegram"`
	} `yaml:"apiKeys"`
}

func ParseConfig(fileBytes []byte) (*Config, error) {
	cf := configFile{}

	err := yaml.Unmarshal(fileBytes, &cf)
	if err != nil {
		return nil, err
	}

	c := Config{}

	c.ApiKeys.Telegram = cf.ApiKeys.Telegram

	return &c, nil
}
