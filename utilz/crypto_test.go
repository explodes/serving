package utilz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	const passphrase = "test"
	original := []byte("some text")
	b, err := Encrypt(original, passphrase)
	assert.NoError(t, err)
	assert.NotEqual(t, original, b)

	decrypted, err := Decrypt(b, passphrase)
	assert.NoError(t, err)

	assert.Equal(t, original, decrypted)
}

func TestEncryptBase64(t *testing.T) {
	const passphrase = "test"
	original := []byte("some text")
	b64, err := EncryptToBase64String(original, passphrase)
	assert.NoError(t, err)
	assert.NotEqual(t, "some text", b64)

	decrypted, err := DecryptFromBase64String(b64, passphrase)
	assert.NoError(t, err)

	assert.Equal(t, original, decrypted)
}
