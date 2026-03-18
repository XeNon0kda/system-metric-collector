package config

import (
    "os"
    "strconv"
    "time"
)

type Config struct {
    ServerPort    string
    ReadTimeout   time.Duration
    WriteTimeout  time.Duration
    MetricsUpdateInterval time.Duration
}

func Load() *Config {
    port := getEnv("SERVER_PORT", "8080")
    readTimeout := getEnvAsInt("READ_TIMEOUT_SEC", 5)
    writeTimeout := getEnvAsInt("WRITE_TIMEOUT_SEC", 10)
    updateInterval := getEnvAsInt("METRICS_UPDATE_INTERVAL_SEC", 2)

    return &Config{
        ServerPort:    ":" + port,
        ReadTimeout:   time.Duration(readTimeout) * time.Second,
        WriteTimeout:  time.Duration(writeTimeout) * time.Second,
        MetricsUpdateInterval: time.Duration(updateInterval) * time.Second,
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
    if value, exists := os.LookupEnv(key); exists {
        if intVal, err := strconv.Atoi(value); err == nil {
            return intVal
        }
    }
    return defaultValue
}