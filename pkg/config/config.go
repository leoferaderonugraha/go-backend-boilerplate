package config

import (
    "io/ioutil"
    "encoding/json"
)

type Config struct {
	DatabaseURL string `json:"database_url"`
}

func LoadConfig(path string) (*Config, error) {
    // Load json based configuration

    file, err := ioutil.ReadFile(path)

    if err != nil {
        return nil, err
    }

    config := &Config{}
    if err := json.Unmarshal(file, config); err != nil {
        return nil, err
    }

	return config, nil
}

