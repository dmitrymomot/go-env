package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// MustString func returns environment variable value as a string value,
// If variable doesn't exist or is not set, exits from the runtime
func MustString(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
	}

	return value
}

// MustBool func returns environment variable value as a boolean value,
// If variable doesn't exist or is not set, exits from the runtime
func MustBool(key string) bool {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
	}

	if value == "true" || value == "1" {
		return true
	} else if value == "false" || value == "0" {
		return false
	}

	panic(fmt.Errorf("required ENV %q must be a boolean but it's %q", key, value))
}

// MustInt func returns environment variable value as an integer value,
// If variable doesn't exist or is not set, exits from the runtime
func MustInt[T int | int16 | int32 | int64](key string) T {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
	}

	res, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic(fmt.Errorf("required ENV %q must be an integer but it's %q", key, value))
	}

	return T(res)
}

// MustFloat func returns environment variable value as a float value,
// If variable doesn't exist or is not set, exits from the runtime
func MustFloat[T float32 | float64](key string) T {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
	}

	res, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(fmt.Errorf("required ENV %q must be a float but it's %q", key, value))
	}

	return T(res)
}

// MustDuration func returns environment variable value as a parsed duration value,
// If variable doesn't exist, is not set or unparsable, then panics
func MustDuration(key string) time.Duration {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
	}

	res, err := time.ParseDuration(value)
	if err != nil {
		panic(fmt.Errorf("required ENV %q must be a parsable duration but it's %q: %v", key, value, err))
	}

	return res
}

// MustTime func returns environment variable value as a parsed time value,
// If variable doesn't exist, is not set or unparsable, then panics.
// If format is empty, then time.RFC3339 is used.
// See default time formats: https://golang.org/pkg/time/#pkg-constants
func MustTime(key string, format string) time.Time {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
	}

	if format == "" {
		format = time.RFC3339
	}

	res, err := time.Parse(format, value)
	if err != nil {
		panic(fmt.Errorf("required ENV %q must be a parsable time but it's %q: %v", key, value, err))
	}

	return res
}

// MustBytes func returns environment variable value as a bytes slice.
// If variable doesn't exist or is not set, exits from the runtime.
func MustBytes(key string) []byte {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
	}

	return []byte(value)
}
