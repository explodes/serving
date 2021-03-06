// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/serialization.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_SerializationOf = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "serving.serialization_of",
	Tag:           "bytes,50000,opt,name=serialization_of",
	Filename:      "proto/serialization.proto",
}

func init() {
	proto.RegisterExtension(E_SerializationOf)
}

func init() { proto.RegisterFile("proto/serialization.proto", fileDescriptor_e6ceea3e8e713876) }

var fileDescriptor_e6ceea3e8e713876 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0xca, 0x4c, 0xcc, 0xc9, 0xac, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3,
	0x03, 0x8b, 0x09, 0xb1, 0x17, 0xa7, 0x16, 0x95, 0x65, 0xe6, 0xa5, 0x4b, 0xe9, 0x94, 0x64, 0x64,
	0x16, 0xa5, 0xc4, 0x17, 0x24, 0x16, 0x95, 0x54, 0xea, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea,
	0x83, 0x95, 0x24, 0x95, 0xa6, 0xe9, 0xa7, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x16, 0x94, 0xe4, 0x17,
	0x41, 0xb4, 0x59, 0x79, 0x71, 0x09, 0xa0, 0x98, 0x16, 0x9f, 0x9f, 0x26, 0x24, 0xab, 0x07, 0xd1,
	0xa6, 0x07, 0xd3, 0xa6, 0xe7, 0x96, 0x99, 0x9a, 0x93, 0xe2, 0x5f, 0x00, 0x52, 0x50, 0x2c, 0x71,
	0xa1, 0x8d, 0x59, 0x81, 0x51, 0x83, 0x33, 0x88, 0x1f, 0x45, 0xa3, 0x7f, 0x9a, 0x93, 0x72, 0x94,
	0x62, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x6a, 0x45, 0x41, 0x4e,
	0x7e, 0x4a, 0x6a, 0xb1, 0x3e, 0xd4, 0x61, 0x50, 0x47, 0xb0, 0x81, 0x29, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x71, 0x11, 0x9c, 0x95, 0xcb, 0x00, 0x00, 0x00,
}
