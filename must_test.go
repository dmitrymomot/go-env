package env_test

import (
	"os"
	"testing"
	"time"

	env "github.com/dmitrymomot/go-env"
	"github.com/stretchr/testify/assert"
)

func TestMustString(t *testing.T) {
	os.Setenv("TEST_ENV", "test")
	assert.Equal(t, "test", env.MustString("TEST_ENV"))
	assert.Panics(t, func() { env.MustString("TEST_ENV_2") })
}

func TestMustBool(t *testing.T) {
	os.Setenv("TEST_ENV", "true")
	assert.Equal(t, true, env.MustBool("TEST_ENV"))
	assert.Panics(t, func() { env.MustBool("TEST_ENV_2") })
	os.Setenv("TEST_ENV", "wrong value")
	assert.Panics(t, func() { env.MustBool("TEST_ENV") })
	os.Setenv("TEST_ENV", "1")
	assert.Equal(t, true, env.MustBool("TEST_ENV"))
	os.Setenv("TEST_ENV", "0")
	assert.Equal(t, false, env.MustBool("TEST_ENV"))
	os.Setenv("TEST_ENV", "false")
	assert.Equal(t, false, env.MustBool("TEST_ENV"))
}

func TestMustInt(t *testing.T) {
	assert.Panics(t, func() { env.MustInt[int]("TEST_ENV_2") })
	os.Setenv("TEST_ENV", "wrong value")
	assert.Panics(t, func() { env.MustInt[int]("TEST_ENV") })

	os.Setenv("TEST_ENV", "1")
	assert.Equal(t, int(1), env.MustInt[int]("TEST_ENV"))
	assert.Equal(t, int16(1), env.MustInt[int16]("TEST_ENV"))
	assert.Equal(t, int32(1), env.MustInt[int32]("TEST_ENV"))
	assert.Equal(t, int64(1), env.MustInt[int64]("TEST_ENV"))
}

func TestMustFloat(t *testing.T) {
	assert.Panics(t, func() { env.MustFloat[float32]("TEST_ENV_2") })
	os.Setenv("TEST_ENV", "wrong value")
	assert.Panics(t, func() { env.MustFloat[float32]("TEST_ENV") })

	os.Setenv("TEST_ENV", "1.1")
	assert.Equal(t, float32(1.1), env.MustFloat[float32]("TEST_ENV"))
	assert.Equal(t, float64(1.1), env.MustFloat[float64]("TEST_ENV"))
}

func TestMustDuration(t *testing.T) {
	assert.Panics(t, func() { env.MustDuration("TEST_ENV_2") })
	os.Setenv("TEST_ENV", "wrong value")
	assert.Panics(t, func() { env.MustDuration("TEST_ENV") })

	os.Setenv("TEST_ENV", "1s")
	assert.Equal(t, 1*time.Second, env.MustDuration("TEST_ENV"))
}

func TestMustTime(t *testing.T) {
	assert.Panics(t, func() { env.MustTime("TEST_ENV_2", time.DateOnly) })
	os.Setenv("TEST_ENV", "wrong value")
	assert.Panics(t, func() { env.MustTime("TEST_ENV", time.DateOnly) })
	os.Setenv("TEST_ENV", "1th Jan 2019")
	assert.Panics(t, func() { env.MustTime("TEST_ENV", time.DateOnly) })

	os.Setenv("TEST_ENV", "2019-01-01")
	assert.Equal(t, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), env.MustTime("TEST_ENV", time.DateOnly))

	os.Setenv("TEST_ENV", "2019-01-01 00:00:00")
	assert.Equal(t, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), env.MustTime("TEST_ENV", time.DateTime))

	os.Setenv("TEST_ENV", "2019-01-01T00:00:00Z")
	assert.Equal(t, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), env.MustTime("TEST_ENV", ""))
}

func TestMustBytes(t *testing.T) {
	assert.Panics(t, func() { env.MustBytes("TEST_ENV_2") })
	os.Setenv("TEST_ENV", "test")
	assert.Equal(t, []byte("test"), env.MustBytes("TEST_ENV"))
}
