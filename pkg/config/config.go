package config

import (
    "os"
    "encoding/json"
)

/* Currently only support for primitive data types */
type JSON_SUPPORTED_TYPES interface {
    string |
    int |
    int32 |
    int64 |
    float32 |
    float64 |
    bool
}

var instance *map[string]any

func loadConfig() (*map[string]any, error) {
    // Could only be instantiated once

    if instance == nil {
        configData, err := os.ReadFile("config.json")

        if err != nil {
            return nil, err
        }

        instance = new(map[string]any)

        if err := json.Unmarshal(configData, instance); err != nil {
            return nil, err
        }
    }

    return instance, nil
}

func Get[T JSON_SUPPORTED_TYPES](key string, fallback T) T {
    cfg, err := loadConfig()

    if err != nil {
        panic(err)
    }

    if val, exists := (*cfg)[key]; exists {
        return val.(T)
    }
    
    return fallback
}
