package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddress_Address(t *testing.T) {
	cases := []struct {
		name     string
		host     string
		port     uint32
		expected string
	}{
		{"empty_host", "", 0, ":0"},
		{"domain_name", "foo", 999, "foo:999"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			addr := &Address{Host: tc.host, Port: tc.port}
			out := addr.Address()
			assert.Equal(t, tc.expected, out)
		})
	}
}
