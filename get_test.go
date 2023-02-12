package env_test

import (
	"os"
	"testing"
	"time"

	"github.com/dmitrymomot/go-env"
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
