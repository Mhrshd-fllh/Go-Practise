package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valAsInt, error := strconv.Atoi(val)

	if error != nil {
		return fallback
	}

	return valAsInt
}
