package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCookieHash(t *testing.T) {
	hash := "abc"

	u64, err := CookieHash(hash)
	assert.NoError(t, err)
	assert.True(t, u64 > 0)
}
