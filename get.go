package env

import (
	"os"
	"strconv"
	"strings"
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

// GetStrings func returns environment variable value as a string slice
// If variable doesn't exist or is not set, returns fallback value
func GetStrings(key string, sep string, fallback []string) []string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	strings := strings.Split(value, sep)

	// filter empty strings
	for i := 0; i < len(strings); i++ {
		if strings[i] == "" {
			strings = append(strings[:i], strings[i+1:]...)
			i--
		}
	}

	if len(strings) == 0 {
		return fallback
	}

	return strings
}

// GetInts func returns environment variable value as a integer slice
// If variable doesn't exist or is not set, returns fallback value
func GetInts[T int | int16 | int32 | int64](key string, sep string, fallback []T) []T {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	strings := strings.Split(value, sep)

	// filter empty strings
	for i := 0; i < len(strings); i++ {
		if strings[i] == "" {
			strings = append(strings[:i], strings[i+1:]...)
			i--
		}
	}

	if len(strings) == 0 {
		return fallback
	}

	var ints []T
	for _, s := range strings {
		res, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return fallback
		}

		ints = append(ints, T(res))
	}

	return ints
}

// GetFloats func returns environment variable value as a float slice
// If variable doesn't exist or is not set, returns fallback value
func GetFloats[T float32 | float64](key string, sep string, fallback []T) []T {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	strings := strings.Split(value, sep)

	// filter empty strings
	for i := 0; i < len(strings); i++ {
		if strings[i] == "" {
			strings = append(strings[:i], strings[i+1:]...)
			i--
		}
	}

	if len(strings) == 0 {
		return fallback
	}

	var floats []T
	for _, s := range strings {
		res, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return fallback
		}

		floats = append(floats, T(res))
	}

	return floats
}

// GetStringsMap func returns environment variable value as a map[string]string
// If variable doesn't exist or is not set, returns fallback value
// Example: key=value1,key2=value2
// sep - key value separator, default is ","
// kvSep - key value separator, default is "="
func GetStringsMap(key string, sep string, kvSep string, fallback map[string]string) map[string]string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	if sep == "" {
		sep = ","
	}
	if kvSep == "" {
		kvSep = "="
	}

	stringsSlice := strings.Split(value, sep)

	// filter empty strings
	for i := 0; i < len(stringsSlice); i++ {
		if stringsSlice[i] == "" {
			stringsSlice = append(stringsSlice[:i], stringsSlice[i+1:]...)
			i--
		}
	}

	if len(stringsSlice) == 0 {
		return fallback
	}

	m := make(map[string]string)
	for _, s := range stringsSlice {
		kv := strings.Split(s, kvSep)
		if len(kv) != 2 {
			return fallback
		}

		// filter empty map keys
		if kv[0] == "" {
			continue
		}

		m[kv[0]] = kv[1]
	}

	if len(m) == 0 {
		return fallback
	}

	return m
}

// GetIntsMap func returns environment variable value as a map[string]int[16|32|64]
// If variable doesn't exist or is not set, returns fallback value
// Example: key=2,key2=32
// sep - key value separator, default is ","
// kvSep - key value separator, default is "="
func GetIntsMap[T int | int16 | int32 | int64](key string, sep string, kvSep string, fallback map[string]T) map[string]T {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	if sep == "" {
		sep = ","
	}
	if kvSep == "" {
		kvSep = "="
	}

	stringsSlice := strings.Split(value, sep)

	// filter empty strings
	for i := 0; i < len(stringsSlice); i++ {
		if stringsSlice[i] == "" {
			stringsSlice = append(stringsSlice[:i], stringsSlice[i+1:]...)
			i--
		}
	}

	if len(stringsSlice) == 0 {
		return fallback
	}

	m := make(map[string]T)
	for _, s := range stringsSlice {
		kv := strings.Split(s, kvSep)
		if len(kv) != 2 {
			return fallback
		}

		// filter empty map keys
		if kv[0] == "" {
			continue
		}

		if kv[1] == "" {
			m[kv[0]] = T(0)
			continue
		}

		res, err := strconv.ParseInt(kv[1], 10, 32)
		if err != nil {
			return fallback
		}

		m[kv[0]] = T(res)
	}

	if len(m) == 0 {
		return fallback
	}

	return m
}

// GetFloatsMap func returns environment variable value as a map[string]float[32|64]
// If variable doesn't exist or is not set, returns fallback value
// Example: key=1.2,key2=0.3
// sep - key value separator, default is ","
// kvSep - key value separator, default is "="
func GetFloatsMap[T float32 | float64](key string, sep string, kvSep string, fallback map[string]T) map[string]T {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	if sep == "" {
		sep = ","
	}
	if kvSep == "" {
		kvSep = "="
	}

	stringsSlice := strings.Split(value, sep)

	// filter empty strings
	for i := 0; i < len(stringsSlice); i++ {
		if stringsSlice[i] == "" {
			stringsSlice = append(stringsSlice[:i], stringsSlice[i+1:]...)
			i--
		}
	}

	if len(stringsSlice) == 0 {
		return fallback
	}

	m := make(map[string]T)
	for _, s := range stringsSlice {
		kv := strings.Split(s, kvSep)
		if len(kv) != 2 {
			return fallback
		}

		// filter empty map keys
		if kv[0] == "" {
			continue
		}

		if kv[1] == "" {
			m[kv[0]] = T(0)
			continue
		}

		res, err := strconv.ParseFloat(kv[1], 64)
		if err != nil {
			return fallback
		}

		m[kv[0]] = T(res)
	}

	if len(m) == 0 {
		return fallback
	}

	return m
}
