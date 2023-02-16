package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

// MustStrings func returns environment variable value as a string slice.
// If variable doesn't exist or is not set, exits from the runtime.
func MustStrings(key string, sep string) []string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
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
		panic(fmt.Errorf("required ENV %q must be a string slice but it's %q", key, value))
	}

	return strings
}

// MustInts func returns environment variable value as an integer slice.
// If variable doesn't exist or is not set, exits from the runtime.
func MustInts[T int | int16 | int32 | int64](key string, sep string) []T {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
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
		panic(fmt.Errorf("required ENV %q must be an integer slice but it's %q", key, value))
	}

	ints := make([]T, len(strings))
	for i, str := range strings {
		res, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			panic(fmt.Errorf("required ENV %q must be an integer slice but it's %q", key, value))
		}

		ints[i] = T(res)
	}

	return ints
}

// MustFloats func returns environment variable value as a float slice.
// If variable doesn't exist or is not set, exits from the runtime.
func MustFloats[T float32 | float64](key string, sep string) []T {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
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
		panic(fmt.Errorf("required ENV %q must be a float slice but it's %q", key, value))
	}

	floats := make([]T, len(strings))
	for i, str := range strings {
		res, err := strconv.ParseFloat(str, 64)
		if err != nil {
			panic(fmt.Errorf("required ENV %q must be a float slice but it's %q", key, value))
		}

		floats[i] = T(res)
	}

	return floats
}

// MustStringsMap func returns environment variable value as a string map.
// If variable doesn't exist or is not set, exits from the runtime.
func MustStringsMap(key string, sep string, kvSep string) map[string]string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
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
		panic(fmt.Errorf("required ENV %q must be a string map but it's %q", key, value))
	}

	m := make(map[string]string)
	for _, s := range stringsSlice {
		kv := strings.Split(s, kvSep)
		if len(kv) != 2 {
			panic(fmt.Errorf("required ENV %q must be a string map but it's %q", key, value))
		}

		// filter empty map keys
		if kv[0] == "" {
			continue
		}

		m[kv[0]] = kv[1]
	}

	if len(m) == 0 {
		panic(fmt.Errorf("required ENV %q must be a string map but it's %q", key, value))
	}

	return m
}

// MustIntsMap func returns environment variable value as a map[string]int[16|32|64]
// If variable doesn't exist or is not set, exits from the runtime.
// Example: key=2,key2=32
// sep - key value separator, default is ","
// kvSep - key value separator, default is "="
func MustIntsMap[T int | int16 | int32 | int64](key string, sep string, kvSep string) map[string]T {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
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
		panic(fmt.Errorf("required ENV %q must be an integer map but it's %q", key, value))
	}

	m := make(map[string]T)
	for _, s := range stringsSlice {
		kv := strings.Split(s, kvSep)
		if len(kv) != 2 {
			panic(fmt.Errorf("required ENV %q must be an integer map but it's %q", key, value))
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
			panic(fmt.Errorf("required ENV %q must be an integer map but it's %q", key, value))
		}

		m[kv[0]] = T(res)
	}

	if len(m) == 0 {
		panic(fmt.Errorf("required ENV %q must be an integer map but it's %q", key, value))
	}

	return m
}

// MustFloatsMap func returns environment variable value as a map[string]float[32|64]
// If variable doesn't exist or is not set, exits from the runtime.
// Example: key=1.2,key2=0.3
// sep - key value separator, default is ","
// kvSep - key value separator, default is "="
func MustFloatsMap[T float32 | float64](key string, sep string, kvSep string) map[string]T {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic(fmt.Errorf("required ENV %q is not set", key))
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
		panic(fmt.Errorf("required ENV %q must be a float map but it's %q", key, value))
	}

	m := make(map[string]T)
	for _, s := range stringsSlice {
		kv := strings.Split(s, kvSep)
		if len(kv) != 2 {
			panic(fmt.Errorf("required ENV %q must be a float map but it's %q", key, value))
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
			panic(fmt.Errorf("required ENV %q must be a float map but it's %q", key, value))
		}

		m[kv[0]] = T(res)
	}

	if len(m) == 0 {
		panic(fmt.Errorf("required ENV %q must be a float map but it's %q", key, value))
	}

	return m
}
