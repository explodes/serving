package expz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModFlags_setDefault(t *testing.T) {
	someFlag := &Flag{}
	mods := NewModFlags()
	mods.setDefault("foo", someFlag)

	assert.Equal(t, someFlag, mods[0]["foo"])
	assert.Equal(t, someFlag, mods[MaxMods-1]["foo"])
}

func TestModFlags_setRange(t *testing.T) {
	someFlag := &Flag{}
	mods := NewModFlags()
	mods.setRange("foo", 10, 15, someFlag)

	assert.Nil(t, mods[9]["foo"])
	assert.Equal(t, someFlag, mods[10]["foo"])
	assert.Equal(t, someFlag, mods[15]["foo"])
	assert.Nil(t, mods[16]["foo"])

}

func TestModFlags_rangeContains(t *testing.T) {
	someFlag := &Flag{}
	mods := NewModFlags()
	mods.setRange("foo", 10, 15, someFlag)

	assert.False(t, mods.rangeContains("foo", 0, 9))
	assert.True(t, mods.rangeContains("foo", 0, 10))
	assert.True(t, mods.rangeContains("foo", 0, 11))
	assert.True(t, mods.rangeContains("foo", 11, 15))
	assert.True(t, mods.rangeContains("foo", 15, 16))
	assert.False(t, mods.rangeContains("foo", 16, 99))
}
