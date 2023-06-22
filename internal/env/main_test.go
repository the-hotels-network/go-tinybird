package env_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/the-hotels-network/go-tinybird/internal/env"
)

func TestGet(t *testing.T) {
	e := env.Get("TB_TEST_GET", "a")
	assert.Equal(t, e, "a")

	t.Setenv("TB_TEST_GET", "b")
	e = env.Get("TB_TEST_GET", "a")
	assert.Equal(t, e, "b")
}

func TestGetBool(t *testing.T) {
	e := env.GetBool("TB_TEST_GET", false)
	assert.Equal(t, e, false)

	t.Setenv("TB_TEST_GET", "true")
	e = env.GetBool("TB_TEST_GET", false)
	assert.Equal(t, e, true)

	t.Setenv("TB_TEST_GET", "asdf")
	e = env.GetBool("TB_TEST_GET", false)
	assert.Equal(t, e, false)
}

func TestGetInt(t *testing.T) {
	e := env.GetInt("TB_TEST_GET", 1)
	assert.Equal(t, e, 1)

	t.Setenv("TB_TEST_GET", "3")
	e = env.GetInt("TB_TEST_GET", 1)
	assert.Equal(t, e, 3)

	t.Setenv("TB_TEST_GET", "a")
	e = env.GetInt("TB_TEST_GET", 1)
	assert.Equal(t, e, 1)
}
