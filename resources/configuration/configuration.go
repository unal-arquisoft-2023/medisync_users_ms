package configuration

import (
	"errors"
	"fmt"

	"github.com/joho/godotenv"
)

type ConfigurationRepository interface {
	Get(key string) (string, error)
}

type ConfigurationGodotEnv struct {
	config map[string]string
	path   string
}

func (c *ConfigurationGodotEnv) Get(key string) (string, error) {
	if val, ok := c.config[key]; ok {
		return val, nil
	}
	return "", errors.New(fmt.Sprintf("Key %s not found, using file %s", key, c.path))
}

func NewConfigurationGodotEnv(path string) (ConfigurationRepository, error) {

	config, err := godotenv.Read(path)

	if err != nil {
		return nil, err
	}

	return &ConfigurationGodotEnv{
		config: config,
		path:   path,
	}, nil
}
