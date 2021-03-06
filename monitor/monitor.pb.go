// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor/monitor.proto

package monitor

import (
	fmt "fmt"
	proto1 "github.com/explodes/serving/proto"
	proto "github.com/golang/protobuf/proto"
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

type Config struct {
	Services             []*Config_Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d5950496a7550d, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetServices() []*Config_Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type Config_Service struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GrpcServer           *proto1.GrpcServer `protobuf:"bytes,2,opt,name=grpc_server,json=grpcServer,proto3" json:"grpc_server,omitempty"`
	UpdateFrequency      *proto1.Duration   `protobuf:"bytes,3,opt,name=update_frequency,json=updateFrequency,proto3" json:"update_frequency,omitempty"`
	Timeout              *proto1.Duration   `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Config_Service) Reset()         { *m = Config_Service{} }
func (m *Config_Service) String() string { return proto.CompactTextString(m) }
func (*Config_Service) ProtoMessage()    {}
func (*Config_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d5950496a7550d, []int{0, 0}
}

func (m *Config_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_Service.Unmarshal(m, b)
}
func (m *Config_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_Service.Marshal(b, m, deterministic)
}
func (m *Config_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_Service.Merge(m, src)
}
func (m *Config_Service) XXX_Size() int {
	return xxx_messageInfo_Config_Service.Size(m)
}
func (m *Config_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Config_Service proto.InternalMessageInfo

func (m *Config_Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Config_Service) GetGrpcServer() *proto1.GrpcServer {
	if m != nil {
		return m.GrpcServer
	}
	return nil
}

func (m *Config_Service) GetUpdateFrequency() *proto1.Duration {
	if m != nil {
		return m.UpdateFrequency
	}
	return nil
}

func (m *Config_Service) GetTimeout() *proto1.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "monitor.Config")
	proto.RegisterType((*Config_Service)(nil), "monitor.Config.Service")
}

func init() { proto.RegisterFile("monitor/monitor.proto", fileDescriptor_94d5950496a7550d) }

var fileDescriptor_94d5950496a7550d = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0xf4, 0x30,
	0x14, 0x86, 0xc9, 0x37, 0xc3, 0xf4, 0x33, 0x5d, 0x38, 0x66, 0x10, 0x43, 0x57, 0x45, 0x11, 0x0a,
	0x42, 0x0a, 0x33, 0x2e, 0x5d, 0xa9, 0xe8, 0xbe, 0xb3, 0x73, 0x33, 0x74, 0xd2, 0x33, 0x31, 0x60,
	0x93, 0x98, 0x1f, 0xd1, 0x9b, 0xf3, 0xd2, 0x44, 0x26, 0x69, 0xbb, 0x73, 0xd5, 0x73, 0xde, 0xf7,
	0x79, 0x0a, 0x39, 0xf8, 0xbc, 0xd7, 0x4a, 0x7a, 0x6d, 0xeb, 0xe1, 0xcb, 0x8c, 0xd5, 0x5e, 0x93,
	0x6c, 0x58, 0x8b, 0x55, 0xdc, 0xeb, 0xb6, 0xeb, 0x2c, 0x38, 0x97, 0xda, 0x62, 0x99, 0x42, 0x2f,
	0x7b, 0x48, 0xc9, 0xe5, 0x0f, 0xc2, 0x8b, 0x07, 0xad, 0x0e, 0x52, 0x90, 0x0d, 0xfe, 0xef, 0xc0,
	0x7e, 0x48, 0x0e, 0x8e, 0xa2, 0x72, 0x56, 0xe5, 0xeb, 0x0b, 0x36, 0xfe, 0x3c, 0x21, 0x6c, 0x9b,
	0xfa, 0x66, 0x02, 0x8b, 0x6f, 0x84, 0xb3, 0x21, 0x25, 0x04, 0xcf, 0x55, 0xdb, 0x03, 0x45, 0x25,
	0xaa, 0x4e, 0x9a, 0x38, 0x93, 0x5b, 0x9c, 0x0b, 0x6b, 0xf8, 0xee, 0x28, 0x80, 0xa5, 0xff, 0x4a,
	0x54, 0xe5, 0xeb, 0x15, 0x8b, 0xbe, 0x12, 0xec, 0xd9, 0x1a, 0xbe, 0x8d, 0x55, 0x83, 0xc5, 0x34,
	0x93, 0x3b, 0xbc, 0x0c, 0xa6, 0x6b, 0x3d, 0xec, 0x0e, 0x16, 0xde, 0x03, 0x28, 0xfe, 0x45, 0x67,
	0x51, 0x3d, 0x9b, 0xd4, 0xc7, 0x60, 0x5b, 0x2f, 0xb5, 0x6a, 0x4e, 0x13, 0xfa, 0x34, 0x92, 0xe4,
	0x06, 0x67, 0xc7, 0x17, 0xea, 0xe0, 0xe9, 0xfc, 0x2f, 0x69, 0x24, 0xee, 0xaf, 0x5f, 0xae, 0x84,
	0xf4, 0xaf, 0x61, 0xcf, 0xb8, 0xee, 0x6b, 0xf8, 0x34, 0x6f, 0xba, 0x03, 0x57, 0x0f, 0xc2, 0x78,
	0xdd, 0xfd, 0x22, 0x9e, 0x6b, 0xf3, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x95, 0xcf, 0xda, 0xa8, 0x77,
	0x01, 0x00, 0x00,
}
