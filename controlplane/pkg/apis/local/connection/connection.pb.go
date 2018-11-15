// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connection.proto

package connection

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

type MechanismType int32

const (
	MechanismType_DEFAULT_INTERFACE MechanismType = 0
	MechanismType_KERNEL_INTERFACE  MechanismType = 1
	MechanismType_VHOST_INTERFACE   MechanismType = 2
	MechanismType_MEM_INTERFACE     MechanismType = 3
	MechanismType_SRIOV_INTERFACE   MechanismType = 4
	MechanismType_HW_INTERFACE      MechanismType = 5
)

var MechanismType_name = map[int32]string{
	0: "DEFAULT_INTERFACE",
	1: "KERNEL_INTERFACE",
	2: "VHOST_INTERFACE",
	3: "MEM_INTERFACE",
	4: "SRIOV_INTERFACE",
	5: "HW_INTERFACE",
}

var MechanismType_value = map[string]int32{
	"DEFAULT_INTERFACE": 0,
	"KERNEL_INTERFACE":  1,
	"VHOST_INTERFACE":   2,
	"MEM_INTERFACE":     3,
	"SRIOV_INTERFACE":   4,
	"HW_INTERFACE":      5,
}

func (x MechanismType) String() string {
	return proto.EnumName(MechanismType_name, int32(x))
}

func (MechanismType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{0}
}

type Mechanism struct {
	Type                 MechanismType     `protobuf:"varint,1,opt,name=type,proto3,enum=local.connection.MechanismType" json:"type,omitempty"`
	Parameters           map[string]string `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Mechanism) Reset()         { *m = Mechanism{} }
func (m *Mechanism) String() string { return proto.CompactTextString(m) }
func (*Mechanism) ProtoMessage()    {}
func (*Mechanism) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{0}
}

func (m *Mechanism) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mechanism.Unmarshal(m, b)
}
func (m *Mechanism) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mechanism.Marshal(b, m, deterministic)
}
func (m *Mechanism) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mechanism.Merge(m, src)
}
func (m *Mechanism) XXX_Size() int {
	return xxx_messageInfo_Mechanism.Size(m)
}
func (m *Mechanism) XXX_DiscardUnknown() {
	xxx_messageInfo_Mechanism.DiscardUnknown(m)
}

var xxx_messageInfo_Mechanism proto.InternalMessageInfo

func (m *Mechanism) GetType() MechanismType {
	if m != nil {
		return m.Type
	}
	return MechanismType_DEFAULT_INTERFACE
}

func (m *Mechanism) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type Connection struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NetworkService       string            `protobuf:"bytes,2,opt,name=network_service,json=networkService,proto3" json:"network_service,omitempty"`
	Mechanism            *Mechanism        `protobuf:"bytes,3,opt,name=mechanism,proto3" json:"mechanism,omitempty"`
	Context              map[string]string `protobuf:"bytes,4,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{1}
}

func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (m *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(m, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Connection) GetNetworkService() string {
	if m != nil {
		return m.NetworkService
	}
	return ""
}

func (m *Connection) GetMechanism() *Mechanism {
	if m != nil {
		return m.Mechanism
	}
	return nil
}

func (m *Connection) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Connection) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterEnum("local.connection.MechanismType", MechanismType_name, MechanismType_value)
	proto.RegisterType((*Mechanism)(nil), "local.connection.Mechanism")
	proto.RegisterMapType((map[string]string)(nil), "local.connection.Mechanism.ParametersEntry")
	proto.RegisterType((*Connection)(nil), "local.connection.Connection")
	proto.RegisterMapType((map[string]string)(nil), "local.connection.Connection.ContextEntry")
	proto.RegisterMapType((map[string]string)(nil), "local.connection.Connection.LabelsEntry")
}

func init() { proto.RegisterFile("connection.proto", fileDescriptor_51baa40a1cc6b48b) }

var fileDescriptor_51baa40a1cc6b48b = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4d, 0x6f, 0xa2, 0x50,
	0x14, 0x86, 0x07, 0x50, 0x27, 0x1e, 0xbf, 0xf0, 0x8c, 0x93, 0x10, 0x67, 0x31, 0xc6, 0xcd, 0x30,
	0x33, 0x09, 0x0b, 0xdd, 0x8c, 0x26, 0x93, 0xd4, 0x52, 0x8c, 0xc6, 0xaf, 0x06, 0xad, 0x4d, 0xba,
	0x31, 0x88, 0x37, 0x29, 0x11, 0x81, 0xc0, 0xad, 0x2d, 0xbf, 0xa1, 0x7f, 0xad, 0xbf, 0xa9, 0x69,
	0x04, 0x3f, 0xae, 0x36, 0x31, 0x71, 0x77, 0x79, 0x78, 0xde, 0x97, 0xc3, 0x09, 0x80, 0x68, 0xba,
	0x8e, 0x43, 0x4c, 0x6a, 0xb9, 0x8e, 0xe2, 0xf9, 0x2e, 0x75, 0x51, 0xb4, 0x5d, 0xd3, 0xb0, 0x95,
	0x03, 0xaf, 0xbe, 0x71, 0x90, 0x1e, 0x10, 0xf3, 0xd1, 0x70, 0xac, 0x60, 0x85, 0x75, 0x48, 0xd0,
	0xd0, 0x23, 0x12, 0x57, 0xe1, 0xe4, 0x7c, 0xed, 0xa7, 0x72, 0xaa, 0x2b, 0x7b, 0x75, 0x12, 0x7a,
	0x44, 0x8f, 0x64, 0xec, 0x01, 0x78, 0x86, 0x6f, 0xac, 0x08, 0x25, 0x7e, 0x20, 0xf1, 0x15, 0x41,
	0xce, 0xd4, 0xfe, 0x9e, 0x89, 0x2a, 0xb7, 0x7b, 0x5b, 0x73, 0xa8, 0x1f, 0xea, 0x4c, 0xbc, 0xfc,
	0x1f, 0x0a, 0x27, 0xb7, 0x51, 0x04, 0x61, 0x49, 0xc2, 0x68, 0xa6, 0xb4, 0xbe, 0x39, 0x62, 0x09,
	0x92, 0x6b, 0xc3, 0x7e, 0x22, 0x12, 0x1f, 0xb1, 0xf8, 0xa2, 0xc9, 0xff, 0xe3, 0xaa, 0xef, 0x3c,
	0x80, 0xba, 0x7f, 0x26, 0xe6, 0x81, 0xb7, 0x16, 0xdb, 0x24, 0x6f, 0x2d, 0xf0, 0x17, 0x14, 0x1c,
	0x42, 0x9f, 0x5d, 0x7f, 0x39, 0x0b, 0x88, 0xbf, 0xb6, 0xcc, 0x5d, 0x45, 0x7e, 0x8b, 0xc7, 0x31,
	0xc5, 0x06, 0xa4, 0x57, 0xbb, 0x79, 0x25, 0xa1, 0xc2, 0xc9, 0x99, 0xda, 0x8f, 0x33, 0xaf, 0xa4,
	0x1f, 0x6c, 0x54, 0xe1, 0xab, 0xe9, 0x3a, 0x94, 0xbc, 0x50, 0x29, 0x11, 0xed, 0xe2, 0xf7, 0xe7,
	0xa0, 0x7a, 0x74, 0xdc, 0xb8, 0xf1, 0x26, 0x76, 0x49, 0xbc, 0x82, 0x94, 0x6d, 0xcc, 0x89, 0x1d,
	0x48, 0xc9, 0xa8, 0x43, 0x3e, 0xdb, 0xd1, 0x8f, 0xd4, 0xb8, 0x62, 0x9b, 0x2b, 0x37, 0x21, 0xcb,
	0x56, 0x5f, 0xb2, 0xc5, 0x72, 0x03, 0x32, 0x4c, 0xe5, 0x25, 0xd1, 0x3f, 0xaf, 0x1c, 0xe4, 0x8e,
	0x3e, 0x12, 0xfc, 0x0e, 0xc5, 0x1b, 0xad, 0xdd, 0xba, 0xeb, 0x4f, 0x66, 0xdd, 0xe1, 0x44, 0xd3,
	0xdb, 0x2d, 0x55, 0x13, 0xbf, 0x60, 0x09, 0xc4, 0x9e, 0xa6, 0x0f, 0xb5, 0x3e, 0x43, 0x39, 0xfc,
	0x06, 0x85, 0x69, 0x67, 0x34, 0x66, 0x55, 0x1e, 0x8b, 0x90, 0x1b, 0x68, 0x03, 0x06, 0x09, 0x1b,
	0x6f, 0xac, 0x77, 0x47, 0x53, 0x06, 0x26, 0x50, 0x84, 0x6c, 0xe7, 0x9e, 0x21, 0xc9, 0xeb, 0xec,
	0x03, 0x1c, 0x36, 0x36, 0x4f, 0x45, 0x3f, 0x41, 0xfd, 0x23, 0x00, 0x00, 0xff, 0xff, 0x17, 0x7e,
	0xaa, 0xe2, 0x18, 0x03, 0x00, 0x00,
}
