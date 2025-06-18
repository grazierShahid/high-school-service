package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return intVal
}

func GetEnvBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return boolVal
}

func GetEnvDuration(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return duration
}

func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		panic(fmt.Sprintf("environment variable %s not set", key))
	}
	return value
}
