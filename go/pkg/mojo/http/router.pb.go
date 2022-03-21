// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/http/router.proto

package http

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	core "github.com/mojo-lang/core/go/pkg/mojo/core"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Router struct {
	Path                 *core.TemplateString `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Headers              []*TemplateHeader    `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	Method               Method               `protobuf:"varint,10,opt,name=method,proto3,enum=mojo.http.Method" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte               `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32                `json:"-" gorm:"-" xml:"-" bson:"-"`
}

func (m *Router) Reset()         { *m = Router{} }
func (m *Router) String() string { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()    {}
func (*Router) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc516f2d963cf1e3, []int{0}
}
func (m *Router) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Router.Unmarshal(m, b)
}
func (m *Router) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Router.Marshal(b, m, deterministic)
}
func (m *Router) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Router.Merge(m, src)
}
func (m *Router) XXX_Size() int {
	return xxx_messageInfo_Router.Size(m)
}
func (m *Router) XXX_DiscardUnknown() {
	xxx_messageInfo_Router.DiscardUnknown(m)
}

var xxx_messageInfo_Router proto.InternalMessageInfo

func (m *Router) GetPath() *core.TemplateString {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Router) GetHeaders() []*TemplateHeader {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Router) GetMethod() Method {
	if m != nil {
		return m.Method
	}
	return Method_METHOD_UNSEPECIFIED
}

func init() {
	proto.RegisterType((*Router)(nil), "mojo.http.Router")
}

func init() { proto.RegisterFile("mojo/http/router.proto", fileDescriptor_bc516f2d963cf1e3) }

var fileDescriptor_bc516f2d963cf1e3 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x18, 0xc4, 0x09, 0xa0, 0x20, 0x5c, 0xa9, 0xa2, 0x11, 0x7f, 0x42, 0x07, 0x12, 0x31, 0x75, 0x80,
	0x58, 0x2a, 0x23, 0x03, 0xc2, 0x13, 0x42, 0x42, 0xaa, 0x0a, 0x13, 0x12, 0x42, 0x69, 0x6b, 0xd9,
	0x81, 0xba, 0x9f, 0xe5, 0x7e, 0x1d, 0x78, 0x13, 0x1e, 0x8d, 0x29, 0x0b, 0x5b, 0x9e, 0xa2, 0xb2,
	0x9d, 0x26, 0x59, 0x32, 0x5c, 0xee, 0x7e, 0xe7, 0xef, 0xc8, 0xb9, 0x82, 0x2f, 0xa0, 0x12, 0x51,
	0x53, 0x03, 0x1b, 0xe4, 0x26, 0xd3, 0x06, 0x10, 0xa2, 0x63, 0xab, 0x67, 0x56, 0x1f, 0x26, 0xce,
	0x32, 0x07, 0xc3, 0x29, 0x72, 0xa5, 0x97, 0x39, 0xf2, 0xcf, 0x35, 0x9a, 0x62, 0x25, 0xbc, 0x77,
	0xd8, 0x61, 0x28, 0x8e, 0x12, 0x16, 0xb5, 0x9e, 0xb4, 0x7a, 0x13, 0x94, 0x3c, 0x5f, 0xec, 0x4a,
	0xae, 0xff, 0x02, 0x12, 0x4e, 0x5d, 0x6b, 0xf4, 0x48, 0x0e, 0x75, 0x8e, 0x32, 0x0e, 0xd2, 0x60,
	0xd4, 0x1b, 0x5f, 0x66, 0xae, 0xde, 0x76, 0x66, 0x6f, 0x75, 0xf4, 0xd5, 0x55, 0xb2, 0xa8, 0x2a,
	0x93, 0xbe, 0xb5, 0xde, 0x80, 0x2a, 0x2c, 0x16, 0x7f, 0xa6, 0x2e, 0x1a, 0x3d, 0x93, 0x23, 0x4f,
	0x5f, 0xc7, 0xfb, 0xe9, 0x41, 0x4b, 0xb1, 0x0f, 0x68, 0x28, 0x4f, 0xce, 0xc1, 0xce, 0xaa, 0x32,
	0x19, 0xd4, 0xee, 0x0e, 0x68, 0x07, 0x88, 0x1e, 0x48, 0xe8, 0x4f, 0x89, 0x49, 0x1a, 0x8c, 0xfa,
	0xe3, 0x41, 0x07, 0xf5, 0xe2, 0x7e, 0xb0, 0xd3, 0xaa, 0x4c, 0x4e, 0xbc, 0xa9, 0x43, 0xa8, 0x63,
	0xec, 0xe3, 0xf7, 0xff, 0x6a, 0x8f, 0x5c, 0x80, 0x11, 0x2e, 0x79, 0xbb, 0xcc, 0x57, 0xa2, 0x65,
	0xb0, 0x9e, 0x3f, 0x7b, 0x62, 0x67, 0x98, 0x04, 0xef, 0x54, 0x14, 0x28, 0x37, 0xb3, 0x6c, 0x0e,
	0x8a, 0x36, 0x76, 0xbf, 0x9c, 0x00, 0xaa, 0xbf, 0x05, 0x6d, 0xa6, 0xbc, 0xb7, 0x9f, 0x59, 0xe8,
	0x06, 0xbc, 0xdb, 0x06, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xbf, 0x0c, 0x49, 0xbf, 0x01, 0x00, 0x00,
}
