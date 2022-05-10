package config

import (
	"os"
)

func ParseConfig() (*Config, error) {

	c := Config{}

	c.GrpcServiceAddress = os.Getenv("GRPC_SERVICE_ADDRESS")
	c.Port = os.Getenv("PORT")

	return &c, nil
}
