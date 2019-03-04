package userz

import (
	"encoding/base64"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"hash/fnv"
)

func SerializeCookie(cookie *Cookie) (string, error) {
	b, err := proto.Marshal(cookie)
	if err != nil {
		return "", errors.Wrap(err, "error serializing cookie")
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func DeserializeCookie(s string) (*Cookie, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding cookie")
	}
	cookie := &Cookie{}
	if err := proto.Unmarshal(b, cookie); err != nil {
		return nil, errors.Wrap(err, "error marshalling cookie")
	}
	return cookie, nil
}

func CookieHash(s string) (uint64, error) {
	h := fnv.New64()
	_, err := h.Write([]byte(s))
	if err != nil {
		return 0, errors.Wrap(err, "unable to update cookie hash")
	}
	return h.Sum64(), nil
}

func (m *Cookie) Hash() (uint64, error) {
	s, err := SerializeCookie(m)
	if err != nil {
		return 0, errors.Wrap(err, "unable to hash cookie")
	}
	return CookieHash(s)
}
