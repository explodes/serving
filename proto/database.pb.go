// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/database.proto

package proto

import (
	fmt "fmt"
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

type DatabaseAddress struct {
	// Types that are valid to be assigned to DatabaseAddress:
	//	*DatabaseAddress_Postgres
	//	*DatabaseAddress_Sqlite3
	DatabaseAddress      isDatabaseAddress_DatabaseAddress `protobuf_oneof:"database_address"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *DatabaseAddress) Reset()         { *m = DatabaseAddress{} }
func (m *DatabaseAddress) String() string { return proto.CompactTextString(m) }
func (*DatabaseAddress) ProtoMessage()    {}
func (*DatabaseAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bac758e969973a9, []int{0}
}

func (m *DatabaseAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseAddress.Unmarshal(m, b)
}
func (m *DatabaseAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseAddress.Marshal(b, m, deterministic)
}
func (m *DatabaseAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseAddress.Merge(m, src)
}
func (m *DatabaseAddress) XXX_Size() int {
	return xxx_messageInfo_DatabaseAddress.Size(m)
}
func (m *DatabaseAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseAddress.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseAddress proto.InternalMessageInfo

type isDatabaseAddress_DatabaseAddress interface {
	isDatabaseAddress_DatabaseAddress()
}

type DatabaseAddress_Postgres struct {
	Postgres *PostgresAddress `protobuf:"bytes,1,opt,name=postgres,proto3,oneof"`
}

type DatabaseAddress_Sqlite3 struct {
	Sqlite3 *Sqlite3Address `protobuf:"bytes,2,opt,name=sqlite3,proto3,oneof"`
}

func (*DatabaseAddress_Postgres) isDatabaseAddress_DatabaseAddress() {}

func (*DatabaseAddress_Sqlite3) isDatabaseAddress_DatabaseAddress() {}

func (m *DatabaseAddress) GetDatabaseAddress() isDatabaseAddress_DatabaseAddress {
	if m != nil {
		return m.DatabaseAddress
	}
	return nil
}

func (m *DatabaseAddress) GetPostgres() *PostgresAddress {
	if x, ok := m.GetDatabaseAddress().(*DatabaseAddress_Postgres); ok {
		return x.Postgres
	}
	return nil
}

func (m *DatabaseAddress) GetSqlite3() *Sqlite3Address {
	if x, ok := m.GetDatabaseAddress().(*DatabaseAddress_Sqlite3); ok {
		return x.Sqlite3
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DatabaseAddress) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DatabaseAddress_OneofMarshaler, _DatabaseAddress_OneofUnmarshaler, _DatabaseAddress_OneofSizer, []interface{}{
		(*DatabaseAddress_Postgres)(nil),
		(*DatabaseAddress_Sqlite3)(nil),
	}
}

func _DatabaseAddress_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DatabaseAddress)
	// database_address
	switch x := m.DatabaseAddress.(type) {
	case *DatabaseAddress_Postgres:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Postgres); err != nil {
			return err
		}
	case *DatabaseAddress_Sqlite3:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Sqlite3); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DatabaseAddress.DatabaseAddress has unexpected type %T", x)
	}
	return nil
}

func _DatabaseAddress_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DatabaseAddress)
	switch tag {
	case 1: // database_address.postgres
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PostgresAddress)
		err := b.DecodeMessage(msg)
		m.DatabaseAddress = &DatabaseAddress_Postgres{msg}
		return true, err
	case 2: // database_address.sqlite3
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Sqlite3Address)
		err := b.DecodeMessage(msg)
		m.DatabaseAddress = &DatabaseAddress_Sqlite3{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DatabaseAddress_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DatabaseAddress)
	// database_address
	switch x := m.DatabaseAddress.(type) {
	case *DatabaseAddress_Postgres:
		s := proto.Size(x.Postgres)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DatabaseAddress_Sqlite3:
		s := proto.Size(x.Sqlite3)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PostgresAddress struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgresAddress) Reset()         { *m = PostgresAddress{} }
func (m *PostgresAddress) String() string { return proto.CompactTextString(m) }
func (*PostgresAddress) ProtoMessage()    {}
func (*PostgresAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bac758e969973a9, []int{1}
}

func (m *PostgresAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgresAddress.Unmarshal(m, b)
}
func (m *PostgresAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgresAddress.Marshal(b, m, deterministic)
}
func (m *PostgresAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgresAddress.Merge(m, src)
}
func (m *PostgresAddress) XXX_Size() int {
	return xxx_messageInfo_PostgresAddress.Size(m)
}
func (m *PostgresAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgresAddress.DiscardUnknown(m)
}

var xxx_messageInfo_PostgresAddress proto.InternalMessageInfo

func (m *PostgresAddress) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Sqlite3Address struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sqlite3Address) Reset()         { *m = Sqlite3Address{} }
func (m *Sqlite3Address) String() string { return proto.CompactTextString(m) }
func (*Sqlite3Address) ProtoMessage()    {}
func (*Sqlite3Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bac758e969973a9, []int{2}
}

func (m *Sqlite3Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sqlite3Address.Unmarshal(m, b)
}
func (m *Sqlite3Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sqlite3Address.Marshal(b, m, deterministic)
}
func (m *Sqlite3Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sqlite3Address.Merge(m, src)
}
func (m *Sqlite3Address) XXX_Size() int {
	return xxx_messageInfo_Sqlite3Address.Size(m)
}
func (m *Sqlite3Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Sqlite3Address.DiscardUnknown(m)
}

var xxx_messageInfo_Sqlite3Address proto.InternalMessageInfo

func (m *Sqlite3Address) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*DatabaseAddress)(nil), "serving.DatabaseAddress")
	proto.RegisterType((*PostgresAddress)(nil), "serving.PostgresAddress")
	proto.RegisterType((*Sqlite3Address)(nil), "serving.Sqlite3Address")
}

func init() { proto.RegisterFile("proto/database.proto", fileDescriptor_8bac758e969973a9) }

var fileDescriptor_8bac758e969973a9 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x49, 0x2c, 0x49, 0x4c, 0x4a, 0x2c, 0x4e, 0xd5, 0x03, 0x73, 0x85, 0xd8, 0x8b,
	0x53, 0x8b, 0xca, 0x32, 0xf3, 0xd2, 0x95, 0x26, 0x31, 0x72, 0xf1, 0xbb, 0x40, 0xe5, 0x1c, 0x53,
	0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x85, 0xcc, 0xb8, 0x38, 0x0a, 0xf2, 0x8b, 0x4b, 0xd2, 0x8b, 0x52,
	0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x24, 0xf4, 0xa0, 0xea, 0xf5, 0x02, 0xa0, 0x12,
	0x50, 0xb5, 0x1e, 0x0c, 0x41, 0x70, 0xb5, 0x42, 0xc6, 0x5c, 0xec, 0xc5, 0x85, 0x39, 0x99, 0x25,
	0xa9, 0xc6, 0x12, 0x4c, 0x60, 0x6d, 0xe2, 0x70, 0x6d, 0xc1, 0x10, 0x71, 0x84, 0x2e, 0x98, 0x4a,
	0x27, 0x21, 0x2e, 0x01, 0x98, 0xdb, 0xe2, 0x13, 0x21, 0xd2, 0x4a, 0xca, 0x5c, 0xfc, 0x68, 0xf6,
	0x08, 0x09, 0x70, 0x31, 0x97, 0x16, 0xe5, 0x80, 0x9d, 0xc3, 0x19, 0x04, 0x62, 0x2a, 0x29, 0x71,
	0xf1, 0xa1, 0x9a, 0x8a, 0xa9, 0xc6, 0x49, 0x39, 0x4a, 0x31, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x57, 0x3f, 0xb5, 0xa2, 0x20, 0x27, 0x3f, 0x25, 0xb5, 0x58, 0x1f, 0xea, 0x2a,
	0x7d, 0x70, 0x58, 0x24, 0xb1, 0x81, 0x29, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0xf4,
	0x2a, 0xff, 0x2a, 0x01, 0x00, 0x00,
}