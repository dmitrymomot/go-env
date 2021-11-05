package env

import (
	"log"
	"os"
	"strconv"
	"time"
)

// MustString func returns environment variable value as a string value,
// If variable doesn't exist or is not set, exits from the runtime
func MustString(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("required ENV %q is not set", key)
	}
	if value == "" {
		log.Fatalf("required ENV %q is empty", key)
	}

	return value
}

// MustBool func returns environment variable value as a boolean value,
// If variable doesn't exist or is not set, exits from the runtime
func MustBool(key string) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("required ENV %q is not set", key)
	}

	if value == "true" || value == "1" {
		return true
	}

	return false
}

// MustInt func returns environment variable value as an integer value,
// If variable doesn't exist or is not set, exits from the runtime
func MustInt(key string) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("required ENV %q is not set", key)
	}
	if value == "" {
		log.Fatalf("required ENV %q is empty", key)
	}

	res, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		log.Fatalf("required ENV %q must be a number but it's %q", key, value)
	}

	return int(res)
}

// MustFloat func returns environment variable value as a float value,
// If variable doesn't exist or is not set, exits from the runtime
func MustFloat(key string) float64 {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("required ENV %q is not set", key)
	}
	if value == "" {
		log.Fatalf("required ENV %q is empty", key)
	}

	res, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatalf("required ENV %q must be a float but it's %q", key, value)
	}

	return res
}

// MustDuration func returns environment variable value as a parsed duration value,
// If variable doesn't exist, is not set or unparsable, then panics
func MustDuration(key string) time.Duration {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("required ENV %q is not set", key)
	}

	res, err := time.ParseDuration(value)
	if err != nil {
		log.Fatalf("required ENV %q must be a parsable duration but it's %q: %v", key, value, err)
	}

	return res
}

// MustBytes func returns environment variable value as a bytes slice.
// If variable doesn't exist or is not set, exits from the runtime.
func MustBytes(key string) []byte {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("required ENV %q is not set", key)
	}
	if value == "" {
		log.Fatalf("required ENV %q is empty", key)
	}

	return []byte(value)
}
