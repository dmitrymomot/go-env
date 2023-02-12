package env

import (
	"os"
	"strconv"
	"time"
)

// GetString func returns environment variable value as a string value,
// If variable doesn't exist or is not set, returns fallback value
func GetString(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

// GetBool func returns environment variable value as a boolean value,
// If variable doesn't exist or is not set, returns fallback value
func GetBool(key string, fallback bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	if value == "true" || value == "1" {
		return true
	} else if value == "false" || value == "0" {
		return false
	}

	return fallback
}

// GetInt func returns environment variable value as a integer value,
// If variable doesn't exist or is not set, returns fallback value
func GetInt[T int | int16 | int32 | int64](key string, fallback T) T {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	res, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return fallback
	}

	return T(res)
}

// GetFloat func returns environment variable value as a float value,
// If variable doesn't exist or is not set, returns fallback value
func GetFloat[T float32 | float64](key string, fallback T) T {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	res, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fallback
	}

	return T(res)
}

// GetDuration func returns environment variable value as a parsed duration value,
// If variable doesn't exist, is not set or unparsable, returns fallback value
func GetDuration(key string, fallback time.Duration) time.Duration {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	res, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return res
}

// GetTime func returns environment variable value as a parsed time value,
// If variable doesn't exist, is not set or unparsable, returns fallback value.
// If format is empty, then time.RFC3339 is used.
func GetTime(key, format string, fallback time.Time) time.Time {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	if format == "" {
		format = time.RFC3339
	}

	res, err := time.Parse(format, value)
	if err != nil {
		return fallback
	}

	return res
}

// GetBytes func returns environment variable value as a bytes slice
// If variable doesn't exist or is not set, returns fallback value
func GetBytes(key string, fallback []byte) []byte {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	return []byte(value)
}
