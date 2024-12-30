package config

import (
    "encoding/json"
    "os"
)

type Config struct {
    Port         string `json:"port"`
    CacheType    string `json:"cache_type"`
    RedisAddress string `json:"redis_address"`
}

func LoadConfig(filePath string) (*Config, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    config := &Config{}
    decoder := json.NewDecoder(file)
    err = decoder.Decode(config)
    if err != nil {
        return nil, err
    }

    return config, nil
}

func LoadEnvConfig() *Config {
    return &Config{
        Port:         getEnv("CACHE_PORT", "8080"),
        CacheType:    getEnv("CACHE_TYPE", "memory"),
        RedisAddress: getEnv("REDIS_ADDRESS", ""),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}