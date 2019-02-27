package proto

import (
	"encoding/base64"
	"github.com/golang/protobuf/proto"
)

// DeserializeProtoBase64 deserializes a base-64 encoded string into a target Message.
func DeserializeProtoBase64(s string, pb proto.Message)error {
	b, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		return  err
	}
	return  proto.Unmarshal(b, pb)
}

// SerializeProtoBase64 serializes Message to a base-64 encoded string.
func SerializeProtoBase64(pb proto.Message) (string, error) {
	b, err := proto.Marshal(pb)
	if err != nil {
		return "", err
	}
	s := base64.RawStdEncoding.EncodeToString(b)
	return s, nil
}