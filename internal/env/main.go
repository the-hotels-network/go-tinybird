package env

import (
	"os"
	"strconv"
)

// Retrieve environment variable.
func Get(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Retrieve environment variable in boolean type.
func GetBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return b
	}
	return fallback
}

// Retrieve environment variable in int type.
func GetInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return b
	}
	return fallback
}
