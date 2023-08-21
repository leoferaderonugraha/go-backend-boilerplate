package config

import (
    "os"
    "encoding/json"
    "fmt"
)

type Config struct {
	DatabaseURL string `json:"database_url"`
    RedisURL string `json:"redis_url"`
    RedisPassword string `json:"redis_password"`
    RedisDB int `json:"redis_db"`
}

var instance *Config

func GetConfig() (*Config, error) {
    // Could only be instantiated once

    if instance == nil {
        fmt.Println("Reading config.json")
        configData, err := os.ReadFile("config.json")

        if err != nil {
            return nil, err
        }

        instance = &Config{}

        if err := json.Unmarshal(configData, instance); err != nil {
            return nil, err
        }
    }

    return instance, nil
}
