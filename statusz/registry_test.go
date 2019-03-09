package statusz_test

import (
	"github.com/explodes/serving/statusz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister_duplicate(t *testing.T) {
	v := noopVar{}
	statusz.Register("foo", v)

	defer func() {
		assert.NotNil(t, recover())
	}()
	statusz.Register("foo", v)
}