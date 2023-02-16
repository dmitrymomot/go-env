package env_test

import (
	"os"
	"testing"
	"time"

	env "github.com/dmitrymomot/go-env"
	"github.com/stretchr/testify/assert"
)

func TestGetString(t *testing.T) {
	assert.Equal(t, "default value", env.GetString("TEST_STRING", "default value"))
	os.Setenv("TEST_STRING", "test")
	assert.Equal(t, "test", env.GetString("TEST_STRING", "default value"))
}

func TestGetBool(t *testing.T) {
	assert.Equal(t, true, env.GetBool("TEST_BOOL", true))
	os.Setenv("TEST_BOOL", "false")
	assert.Equal(t, false, env.GetBool("TEST_BOOL", true))
	os.Setenv("TEST_BOOL", "wrong value")
	assert.Equal(t, true, env.GetBool("TEST_BOOL", true))
	os.Setenv("TEST_BOOL", "1")
	assert.Equal(t, true, env.GetBool("TEST_BOOL", false))
}

func TestGetInt(t *testing.T) {
	assert.Equal(t, int32(1), env.GetInt("TEST_INT", int32(1)))
	os.Setenv("TEST_INT", "123")
	assert.Equal(t, int(123), env.GetInt("TEST_INT", 0))
	assert.Equal(t, int16(123), env.GetInt[int16]("TEST_INT", 0))
	assert.Equal(t, int32(123), env.GetInt[int32]("TEST_INT", 0))
	assert.Equal(t, int64(123), env.GetInt[int64]("TEST_INT", 0))
	os.Setenv("TEST_INT", "wrong value")
	assert.Equal(t, int32(2), env.GetInt("TEST_INT", int32(2)))
}

func TestGetFloat(t *testing.T) {
	assert.Equal(t, float32(1.23), env.GetFloat[float32]("TEST_FLOAT", 1.23))
	os.Setenv("TEST_FLOAT", "123.45")
	assert.Equal(t, float32(123.45), env.GetFloat[float32]("TEST_FLOAT", 0.0))
	assert.Equal(t, float64(123.45), env.GetFloat("TEST_FLOAT", 0.0))
	os.Setenv("TEST_FLOAT", "wrong value")
	assert.Equal(t, float32(2.34), env.GetFloat[float32]("TEST_FLOAT", 2.34))
}

func TestGetDuration(t *testing.T) {
	assert.Equal(t, 1*time.Second, env.GetDuration("TEST_DURATION", 1*time.Second))
	os.Setenv("TEST_DURATION", "1m")
	assert.Equal(t, 1*time.Minute, env.GetDuration("TEST_DURATION", 0))
	os.Setenv("TEST_DURATION", "wrong value")
	assert.Equal(t, 2*time.Second, env.GetDuration("TEST_DURATION", 2*time.Second))
}

func TestGetTime(t *testing.T) {
	assert.Equal(t, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), env.GetTime("TEST_TIME", time.DateOnly, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)))

	os.Setenv("TEST_TIME", "2020-01-02T00:00:00Z")
	assert.Equal(t, time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), env.GetTime("TEST_TIME", "2006-01-02T15:04:05Z", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)))

	os.Setenv("TEST_TIME", "wrong value")
	assert.Equal(t, time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), env.GetTime("TEST_TIME", time.DateOnly, time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC)))

	os.Setenv("TEST_TIME", "2020-01-04")
	assert.Equal(t, time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), env.GetTime("TEST_TIME", time.DateOnly, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)))

	os.Setenv("TEST_TIME", "2020-01-05T00:00:00Z")
	assert.Equal(t, time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), env.GetTime("TEST_TIME", "", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)))
}

func TestGetBytes(t *testing.T) {
	assert.Equal(t, []byte("test"), env.GetBytes("TEST_BYTES", []byte("test")))
	os.Setenv("TEST_BYTES", "test")
	assert.Equal(t, []byte("test"), env.GetBytes("TEST_BYTES", []byte("default value")))
}

func TestGetStrings(t *testing.T) {
	assert.Equal(t, []string{"test"}, env.GetStrings("TEST_STRINGS", ",", []string{"test"}))
	os.Setenv("TEST_STRINGS", "test1,test2")
	assert.Equal(t, []string{"test1", "test2"}, env.GetStrings("TEST_STRINGS", ",", []string{"default value"}))
	os.Setenv("TEST_STRINGS", ",")
	assert.Equal(t, []string{"default value"}, env.GetStrings("TEST_STRINGS", ",", []string{"default value"}))
}

func TestGetInts(t *testing.T) {
	assert.Equal(t, []int{1, 2}, env.GetInts("TEST_INTS", ",", []int{1, 2}))
	os.Setenv("TEST_INTS", "1,2")
	assert.Equal(t, []int{1, 2}, env.GetInts("TEST_INTS", ",", []int{3, 4}))
	os.Setenv("TEST_INTS", ",")
	assert.Equal(t, []int{3, 4}, env.GetInts("TEST_INTS", ",", []int{3, 4}))
	os.Setenv("TEST_INTS", "1,2,three")
	assert.Equal(t, []int{1, 2, 3}, env.GetInts("TEST_INTS", ",", []int{1, 2, 3}))
}

func TestGetFloats(t *testing.T) {
	assert.Equal(t, []float64{1.23, 4.56}, env.GetFloats("TEST_FLOATS", ",", []float64{1.23, 4.56}))
	os.Setenv("TEST_FLOATS", "1.23,4.56")
	assert.Equal(t, []float64{1.23, 4.56}, env.GetFloats("TEST_FLOATS", ",", []float64{7.89, 10.11}))
	os.Setenv("TEST_FLOATS", ",")
	assert.Equal(t, []float64{7.89, 10.11}, env.GetFloats("TEST_FLOATS", ",", []float64{7.89, 10.11}))
	os.Setenv("TEST_FLOATS", "1.23,4.56,seven")
	assert.Equal(t, []float64{1.23, 4.56, 7.0}, env.GetFloats("TEST_FLOATS", ",", []float64{1.23, 4.56, 7.0}))
}

func TestGetStringsMap(t *testing.T) {
	assert.Equal(t,
		map[string]string{"key1": "value1", "key2": "value2"},
		env.GetStringsMap("TEST_STRING_MAP", "", "", map[string]string{"key1": "value1", "key2": "value2"}),
	)

	os.Setenv("TEST_STRING_MAP", "key1=value1,key2=value2")
	assert.Equal(t,
		map[string]string{"key1": "value1", "key2": "value2"},
		env.GetStringsMap("TEST_STRING_MAP", "", "", map[string]string{"key3": "value3", "key4": "value4"}),
	)

	os.Setenv("TEST_STRING_MAP", ",")
	assert.Equal(t,
		map[string]string{"key3": "value3", "key4": "value4"},
		env.GetStringsMap("TEST_STRING_MAP", "", "", map[string]string{"key3": "value3", "key4": "value4"}),
	)

	os.Setenv("TEST_STRING_MAP", "=,=,=")
	assert.Equal(t,
		map[string]string{"key3": "value3", "key4": "value4"},
		env.GetStringsMap("TEST_STRING_MAP", "", "", map[string]string{"key3": "value3", "key4": "value4"}),
	)

	os.Setenv("TEST_STRING_MAP", "key1=value1,key2=value2,key3")
	assert.Equal(t,
		map[string]string{"key3": "value3", "key4": "value4"},
		env.GetStringsMap("TEST_STRING_MAP", "", "", map[string]string{"key3": "value3", "key4": "value4"}),
	)

	os.Setenv("TEST_STRING_MAP", "key1=value1,key2=value2,key3=")
	assert.Equal(t,
		map[string]string{"key1": "value1", "key2": "value2", "key3": ""},
		env.GetStringsMap("TEST_STRING_MAP", "", "", map[string]string{"key3": "value3", "key4": "value4"}),
	)
}

func TestGetIntsMap(t *testing.T) {
	assert.Equal(t,
		map[string]int{"key1": 1, "key2": 2},
		env.GetIntsMap("TEST_INT_MAP", "", "", map[string]int{"key1": 1, "key2": 2}),
	)

	os.Setenv("TEST_INT_MAP", "key1=1,key2=2")
	assert.Equal(t,
		map[string]int{"key1": 1, "key2": 2},
		env.GetIntsMap("TEST_INT_MAP", "", "", map[string]int{"key3": 3, "key4": 4}),
	)

	os.Setenv("TEST_INT_MAP", ",")
	assert.Equal(t,
		map[string]int{"key3": 3, "key4": 4},
		env.GetIntsMap("TEST_INT_MAP", "", "", map[string]int{"key3": 3, "key4": 4}),
	)

	os.Setenv("TEST_INT_MAP", "=,=,=")
	assert.Equal(t,
		map[string]int{"key3": 3, "key4": 4},
		env.GetIntsMap("TEST_INT_MAP", "", "", map[string]int{"key3": 3, "key4": 4}),
	)

	os.Setenv("TEST_INT_MAP", "key1=1,key2=2,key3")
	assert.Equal(t,
		map[string]int{"key3": 3, "key4": 4},
		env.GetIntsMap("TEST_INT_MAP", "", "", map[string]int{"key3": 3, "key4": 4}),
	)

	os.Setenv("TEST_INT_MAP", "key1=1,key2=2,key3=")
	assert.Equal(t,
		map[string]int32{"key1": 1, "key2": 2, "key3": 0},
		env.GetIntsMap("TEST_INT_MAP", "", "", map[string]int32{"key3": 3, "key4": 4}),
	)

	os.Setenv("TEST_INT_MAP", "key1=1,key2=2,key3=three")
	assert.Equal(t,
		map[string]int{"key3": 3, "key4": 4},
		env.GetIntsMap("TEST_INT_MAP", "", "", map[string]int{"key3": 3, "key4": 4}),
	)
}

func TestGetFloatsMap(t *testing.T) {
	assert.Equal(t,
		map[string]float64{"key1": 1.23, "key2": 4.56},
		env.GetFloatsMap("TEST_FLOAT_MAP", "", "", map[string]float64{"key1": 1.23, "key2": 4.56}),
	)

	os.Setenv("TEST_FLOAT_MAP", "key1=1.23,key2=4.56")
	assert.Equal(t,
		map[string]float64{"key1": 1.23, "key2": 4.56},
		env.GetFloatsMap("TEST_FLOAT_MAP", "", "", map[string]float64{"key3": 7.89, "key4": 10.11}),
	)

	os.Setenv("TEST_FLOAT_MAP", ",")
	assert.Equal(t,
		map[string]float64{"key3": 7.89, "key4": 10.11},
		env.GetFloatsMap("TEST_FLOAT_MAP", "", "", map[string]float64{"key3": 7.89, "key4": 10.11}),
	)

	os.Setenv("TEST_FLOAT_MAP", "=,=,=")
	assert.Equal(t,
		map[string]float64{"key3": 7.89, "key4": 10.11},
		env.GetFloatsMap("TEST_FLOAT_MAP", "", "", map[string]float64{"key3": 7.89, "key4": 10.11}),
	)

	os.Setenv("TEST_FLOAT_MAP", "key1=1.23,key2=4.56,key3")
	assert.Equal(t,
		map[string]float64{"key3": 7.89, "key4": 10.11},
		env.GetFloatsMap("TEST_FLOAT_MAP", "", "", map[string]float64{"key3": 7.89, "key4": 10.11}),
	)

	os.Setenv("TEST_FLOAT_MAP", "key1=1.23,key2=4.56,key3=")
	assert.Equal(t,
		map[string]float64{"key1": 1.23, "key2": 4.56, "key3": 0.0},
		env.GetFloatsMap("TEST_FLOAT_MAP", "", "", map[string]float64{"key3": 7.89, "key4": 10.11}),
	)

	os.Setenv("TEST_FLOAT_MAP", "key1=1.23,key2=4.56,key3=seven")
	assert.Equal(t,
		map[string]float64{"key3": 7.89, "key4": 10.11},
		env.GetFloatsMap("TEST_FLOAT_MAP", "", "", map[string]float64{"key3": 7.89, "key4": 10.11}),
	)
}
