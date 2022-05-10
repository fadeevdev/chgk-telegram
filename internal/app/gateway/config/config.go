package config

import "gopkg.in/yaml.v2"

type configFile struct {
	GrpcServiceAddress string `yaml:"grpc_service_address"`
	Port               string `yaml:"port"`
}

func ParseConfig(fileBytes []byte) (*Config, error) {
	cf := configFile{}

	err := yaml.Unmarshal(fileBytes, &cf)
	if err != nil {
		return nil, err
	}

	c := Config{}

	c.GrpcServiceAddress = cf.GrpcServiceAddress
	c.Port = cf.Port

	return &c, nil
}
