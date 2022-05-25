// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: mojo/http/cookie.proto

package http

import (
	core "github.com/mojo-lang/core/go/pkg/mojo/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SameSite int32

const (
	SameSite_SAME_SITE_UNSEPECIFIED SameSite = 0
	SameSite_SAME_SITE_LAX          SameSite = 1
	SameSite_SAME_SITE_STRICT       SameSite = 2
	SameSite_SAME_SITE_NONE         SameSite = 3
)

// Enum value maps for SameSite.
var (
	SameSite_name = map[int32]string{
		0: "SAME_SITE_UNSEPECIFIED",
		1: "SAME_SITE_LAX",
		2: "SAME_SITE_STRICT",
		3: "SAME_SITE_NONE",
	}
	SameSite_value = map[string]int32{
		"SAME_SITE_UNSEPECIFIED": 0,
		"SAME_SITE_LAX":          1,
		"SAME_SITE_STRICT":       2,
		"SAME_SITE_NONE":         3,
	}
)

func (x SameSite) Enum() *SameSite {
	p := new(SameSite)
	*p = x
	return p
}

func (x SameSite) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SameSite) Descriptor() protoreflect.EnumDescriptor {
	return file_mojo_http_cookie_proto_enumTypes[0].Descriptor()
}

func (SameSite) Type() protoreflect.EnumType {
	return &file_mojo_http_cookie_proto_enumTypes[0]
}

func (x SameSite) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SameSite.Descriptor instead.
func (SameSite) EnumDescriptor() ([]byte, []int) {
	return file_mojo_http_cookie_proto_rawDescGZIP(), []int{0}
}

type Cookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value    string          `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Path     string          `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Domain   string          `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	Expires  *core.Timestamp `protobuf:"bytes,5,opt,name=expires,proto3" json:"expires,omitempty"`
	MaxAge   int32           `protobuf:"varint,6,opt,name=max_age,json=maxAge,proto3" json:"maxAge,omitempty"`
	Secure   bool            `protobuf:"varint,7,opt,name=secure,proto3" json:"secure,omitempty"`
	HttpOnly bool            `protobuf:"varint,8,opt,name=http_only,json=httpOnly,proto3" json:"httpOnly,omitempty"`
	SameSite SameSite        `protobuf:"varint,9,opt,name=same_site,json=sameSite,proto3,enum=mojo.http.SameSite" json:"sameSite,omitempty"`
}

func (x *Cookie) Reset() {
	*x = Cookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_http_cookie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cookie) ProtoMessage() {}

func (x *Cookie) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_http_cookie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cookie.ProtoReflect.Descriptor instead.
func (*Cookie) Descriptor() ([]byte, []int) {
	return file_mojo_http_cookie_proto_rawDescGZIP(), []int{0}
}

func (x *Cookie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cookie) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Cookie) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Cookie) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Cookie) GetExpires() *core.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

func (x *Cookie) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *Cookie) GetSecure() bool {
	if x != nil {
		return x.Secure
	}
	return false
}

func (x *Cookie) GetHttpOnly() bool {
	if x != nil {
		return x.HttpOnly
	}
	return false
}

func (x *Cookie) GetSameSite() SameSite {
	if x != nil {
		return x.SameSite
	}
	return SameSite_SAME_SITE_UNSEPECIFIED
}

var File_mojo_http_cookie_proto protoreflect.FileDescriptor

var file_mojo_http_cookie_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x1a, 0x14, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x06, 0x43, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x2e, 0x0a, 0x07, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f,
	0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61,
	0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x78,
	0x41, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x68, 0x74, 0x74, 0x70, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x61, 0x6d, 0x65,
	0x5f, 0x73, 0x69, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x6f,
	0x6a, 0x6f, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x53, 0x61, 0x6d, 0x65, 0x53, 0x69, 0x74, 0x65,
	0x52, 0x08, 0x73, 0x61, 0x6d, 0x65, 0x53, 0x69, 0x74, 0x65, 0x2a, 0x63, 0x0a, 0x08, 0x53, 0x61,
	0x6d, 0x65, 0x53, 0x69, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x41, 0x4d, 0x45, 0x5f, 0x53,
	0x49, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x49, 0x54, 0x45, 0x5f,
	0x4c, 0x41, 0x58, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x49,
	0x54, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x41, 0x4d, 0x45, 0x5f, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x42,
	0x59, 0x0a, 0x17, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x42, 0x0b, 0x43, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x3b, 0x68, 0x74, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mojo_http_cookie_proto_rawDescOnce sync.Once
	file_mojo_http_cookie_proto_rawDescData = file_mojo_http_cookie_proto_rawDesc
)

func file_mojo_http_cookie_proto_rawDescGZIP() []byte {
	file_mojo_http_cookie_proto_rawDescOnce.Do(func() {
		file_mojo_http_cookie_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_http_cookie_proto_rawDescData)
	})
	return file_mojo_http_cookie_proto_rawDescData
}

var file_mojo_http_cookie_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mojo_http_cookie_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mojo_http_cookie_proto_goTypes = []interface{}{
	(SameSite)(0),          // 0: mojo.http.SameSite
	(*Cookie)(nil),         // 1: mojo.http.Cookie
	(*core.Timestamp)(nil), // 2: mojo.core.Timestamp
}
var file_mojo_http_cookie_proto_depIdxs = []int32{
	2, // 0: mojo.http.Cookie.expires:type_name -> mojo.core.Timestamp
	0, // 1: mojo.http.Cookie.same_site:type_name -> mojo.http.SameSite
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mojo_http_cookie_proto_init() }
func file_mojo_http_cookie_proto_init() {
	if File_mojo_http_cookie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mojo_http_cookie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cookie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mojo_http_cookie_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_http_cookie_proto_goTypes,
		DependencyIndexes: file_mojo_http_cookie_proto_depIdxs,
		EnumInfos:         file_mojo_http_cookie_proto_enumTypes,
		MessageInfos:      file_mojo_http_cookie_proto_msgTypes,
	}.Build()
	File_mojo_http_cookie_proto = out.File
	file_mojo_http_cookie_proto_rawDesc = nil
	file_mojo_http_cookie_proto_goTypes = nil
	file_mojo_http_cookie_proto_depIdxs = nil
}