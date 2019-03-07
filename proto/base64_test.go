package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerializeProtoBase64(t *testing.T) {
	addr := &Address{Host: "hello", Port: 123}

	s, err := SerializeProtoBase64(addr)
	assert.NoError(t, err)

	addr2 := &Address{}
	err = DeserializeProtoBase64(s, addr2)
	assert.NoError(t, err)

	assert.Equal(t, addr.Host, addr2.Host)
	assert.Equal(t, addr.Port, addr2.Port)
}

func TestDeserializeProtoBase64_invalid(t *testing.T) {
	assert.Error(t, DeserializeProtoBase64("!", &Address{}))
}
