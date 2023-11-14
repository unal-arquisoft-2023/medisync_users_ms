package configuration

import (
	"errors"
	"fmt"
	"os"

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
	val := os.Getenv(key)
	if val != "" {
		return val, nil
	}
	return "", errors.New(fmt.Sprintf("Key %s not found, using file %s", key, c.path))
}

func NewConfigurationGodotEnv(path string) (ConfigurationRepository, error) {

	err := godotenv.Load(path)
	// fmt.Println("Loaded config ", config)

	if err != nil {
		return nil, err
	}

	return &ConfigurationGodotEnv{
		path: path,
	}, nil
}
