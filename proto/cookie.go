package proto

import (
	"github.com/pkg/errors"
	"hash/fnv"
)

func CookieHash(s string) (uint64, error) {
	h := fnv.New64()
	_, err := h.Write([]byte(s))
	if err != nil {
		return 0, errors.Wrap(err, "unable to update cookie hash")
	}
	return h.Sum64(), nil
}