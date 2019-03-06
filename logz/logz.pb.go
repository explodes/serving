// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logz/logz.proto

package logz

import (
	fmt "fmt"
	proto1 "github.com/explodes/serving/proto"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Level int32

const (
	Level_UNKNOWN Level = 0
	Level_DEBUG   Level = 1
	Level_INFO    Level = 2
	Level_WARN    Level = 3
	Level_ERROR   Level = 4
)

var Level_name = map[int32]string{
	0: "UNKNOWN",
	1: "DEBUG",
	2: "INFO",
	3: "WARN",
	4: "ERROR",
}

var Level_value = map[string]int32{
	"UNKNOWN": 0,
	"DEBUG":   1,
	"INFO":    2,
	"WARN":    3,
	"ERROR":   4,
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}

func (Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_623f03371c6d65cf, []int{0}
}

type Entry struct {
	// When the entry occurred.
	Timestamp *proto1.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Optional time when the entry ended.
	EndTimestamp *proto1.Timestamp `protobuf:"bytes,2,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
	// Level of the log.
	Level Level `protobuf:"varint,3,opt,name=level,proto3,enum=logz.Level" json:"level,omitempty"`
	// Log message.
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// Any additional information to record.
	Ext                  *any.Any `protobuf:"bytes,5,opt,name=ext,proto3" json:"ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_623f03371c6d65cf, []int{0}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetTimestamp() *proto1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Entry) GetEndTimestamp() *proto1.Timestamp {
	if m != nil {
		return m.EndTimestamp
	}
	return nil
}

func (m *Entry) GetLevel() Level {
	if m != nil {
		return m.Level
	}
	return Level_UNKNOWN
}

func (m *Entry) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Entry) GetExt() *any.Any {
	if m != nil {
		return m.Ext
	}
	return nil
}

type Frame struct {
	// ID of and end-to-end series of RPCs.
	StackId string `protobuf:"bytes,1,opt,name=stack_id,json=stackId,proto3" json:"stack_id,omitempty"`
	// ID of an individual RPC.
	FrameId string `protobuf:"bytes,2,opt,name=frame_id,json=frameId,proto3" json:"frame_id,omitempty"`
	// Name of an individual RPC.
	FrameName string `protobuf:"bytes,3,opt,name=frame_name,json=frameName,proto3" json:"frame_name,omitempty"`
	// Parent operation ID of an individual RPC.
	ParentFrameId        string   `protobuf:"bytes,4,opt,name=parent_frame_id,json=parentFrameId,proto3" json:"parent_frame_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_623f03371c6d65cf, []int{1}
}

func (m *Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Frame.Unmarshal(m, b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
}
func (m *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(m, src)
}
func (m *Frame) XXX_Size() int {
	return xxx_messageInfo_Frame.Size(m)
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

func (m *Frame) GetStackId() string {
	if m != nil {
		return m.StackId
	}
	return ""
}

func (m *Frame) GetFrameId() string {
	if m != nil {
		return m.FrameId
	}
	return ""
}

func (m *Frame) GetFrameName() string {
	if m != nil {
		return m.FrameName
	}
	return ""
}

func (m *Frame) GetParentFrameId() string {
	if m != nil {
		return m.ParentFrameId
	}
	return ""
}

func init() {
	proto.RegisterEnum("logz.Level", Level_name, Level_value)
	proto.RegisterType((*Entry)(nil), "logz.Entry")
	proto.RegisterType((*Frame)(nil), "logz.Frame")
}

func init() { proto.RegisterFile("logz/logz.proto", fileDescriptor_623f03371c6d65cf) }

var fileDescriptor_623f03371c6d65cf = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x8b, 0xda, 0x40,
	0x18, 0xc5, 0x1b, 0x4d, 0xaa, 0xf9, 0xac, 0x35, 0x0c, 0x3d, 0xa4, 0x42, 0x21, 0x15, 0x2a, 0xd2,
	0xc3, 0xa4, 0xd8, 0x43, 0x6f, 0x05, 0x65, 0x75, 0x91, 0x5d, 0x22, 0x0c, 0x2b, 0xc2, 0x5e, 0xc2,
	0x68, 0xc6, 0x18, 0x36, 0x99, 0x84, 0x64, 0x14, 0xb3, 0xe7, 0xfd, 0x1f, 0xf7, 0xdf, 0x59, 0x66,
	0x66, 0xd5, 0xd3, 0x5e, 0xc2, 0xbc, 0xf7, 0x7b, 0xdf, 0x97, 0x79, 0x21, 0xd0, 0x4b, 0xf3, 0xf8,
	0xd9, 0x97, 0x0f, 0x5c, 0x94, 0xb9, 0xc8, 0x91, 0x29, 0xcf, 0x7d, 0x47, 0x09, 0x5f, 0x24, 0x19,
	0xd3, 0x7e, 0xff, 0x97, 0xd8, 0x27, 0x65, 0x14, 0x16, 0xb4, 0x14, 0xb5, 0x1f, 0xe7, 0x79, 0x9c,
	0x32, 0x5f, 0x91, 0xcd, 0x61, 0xe7, 0x53, 0x5e, 0xeb, 0xd8, 0xe0, 0xd5, 0x00, 0x6b, 0xc6, 0x45,
	0x59, 0xa3, 0x3f, 0x60, 0xcb, 0xf1, 0x4a, 0xd0, 0xac, 0x70, 0x0d, 0xcf, 0x18, 0x75, 0xc6, 0x08,
	0x57, 0xac, 0x3c, 0x26, 0x3c, 0xc6, 0x0f, 0x67, 0x42, 0xae, 0x21, 0xf4, 0x0f, 0xba, 0x8c, 0x47,
	0xe1, 0x75, 0xaa, 0xf1, 0xe1, 0xd4, 0x17, 0xc6, 0xa3, 0x8b, 0x42, 0x3f, 0xc1, 0x4a, 0xd9, 0x91,
	0xa5, 0x6e, 0xd3, 0x33, 0x46, 0x5f, 0xc7, 0x1d, 0xac, 0xfa, 0xdc, 0x4b, 0x8b, 0x68, 0x82, 0x5c,
	0x68, 0x65, 0xac, 0xaa, 0x68, 0xcc, 0x5c, 0xd3, 0x33, 0x46, 0x36, 0x39, 0x4b, 0x34, 0x84, 0x26,
	0x3b, 0x09, 0xd7, 0x52, 0xef, 0xfa, 0x86, 0x75, 0x35, 0x7c, 0xae, 0x86, 0x27, 0xbc, 0x26, 0x32,
	0x30, 0x78, 0x31, 0xc0, 0x9a, 0x97, 0x34, 0x63, 0xe8, 0x3b, 0xb4, 0x2b, 0x41, 0xb7, 0x4f, 0x61,
	0x12, 0xa9, 0x62, 0x36, 0x69, 0x29, 0xbd, 0x88, 0x24, 0xda, 0xc9, 0x8c, 0x44, 0x0d, 0x8d, 0x94,
	0x5e, 0x44, 0xe8, 0x07, 0x80, 0x46, 0x9c, 0x66, 0x4c, 0xdd, 0xd4, 0x26, 0xb6, 0x72, 0x02, 0xb9,
	0x74, 0x08, 0xbd, 0x82, 0x96, 0x8c, 0x8b, 0xf0, 0xb2, 0x40, 0x5f, 0xb4, 0xab, 0xed, 0xb9, 0x5e,
	0xf3, 0xfb, 0x3f, 0x58, 0xaa, 0x18, 0xea, 0x40, 0x6b, 0x15, 0xdc, 0x05, 0xcb, 0x75, 0xe0, 0x7c,
	0x42, 0x36, 0x58, 0x37, 0xb3, 0xe9, 0xea, 0xd6, 0x31, 0x50, 0x1b, 0xcc, 0x45, 0x30, 0x5f, 0x3a,
	0x0d, 0x79, 0x5a, 0x4f, 0x48, 0xe0, 0x34, 0x25, 0x9e, 0x11, 0xb2, 0x24, 0x8e, 0x39, 0x1d, 0x3c,
	0x7a, 0x71, 0x22, 0xf6, 0x87, 0x0d, 0xde, 0xe6, 0x99, 0xcf, 0x4e, 0x45, 0x9a, 0x47, 0xac, 0xf2,
	0xdf, 0x3f, 0xb1, 0xfa, 0x13, 0x36, 0x9f, 0x55, 0xfb, 0xbf, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x8e, 0x1b, 0x1d, 0xcb, 0x1d, 0x02, 0x00, 0x00,
}
