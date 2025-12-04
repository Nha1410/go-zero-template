package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/logx"
)

func LoadEnv() error {
	if os.Getenv("DOCKER_CONTAINER") == "true" {
		logx.Infof("Running in Docker, using environment variables from container")
		return nil
	}

	if err := godotenv.Load(); err != nil {
		if err := godotenv.Load("../.env"); err != nil {
			if err := godotenv.Load("../../.env"); err != nil {
				logx.Infof("No .env file found, using environment variables only")
				return nil
			}
		}
	}
	return nil
}

func GetString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		logx.Errorf("Failed to parse %s as int: %v, using default: %d", key, err, defaultValue)
		return defaultValue
	}
	return intValue
}

func GetBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		logx.Errorf("Failed to parse %s as bool: %v, using default: %v", key, err, defaultValue)
		return defaultValue
	}
	return boolValue
}

func GetStringSlice(key string, defaultValue []string) []string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return strings.Split(value, ",")
}
