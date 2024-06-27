// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: mojo/http/method.proto

package http

import (
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

type Method int32

const (
	Method_METHOD_UNSEPECIFIED Method = 0
	Method_METHOD_GET          Method = 1
	Method_METHOD_POST         Method = 2
	Method_METHOD_PUT          Method = 3
	Method_METHOD_PATCH        Method = 4
	Method_METHOD_DELETE       Method = 5
	Method_METHOD_OPTIONS      Method = 6
	Method_METHOD_HEAD         Method = 7
	Method_METHOD_TRACE        Method = 8
	Method_METHOD_CONNECT      Method = 9
)

// Enum value maps for Method.
var (
	Method_name = map[int32]string{
		0: "METHOD_UNSEPECIFIED",
		1: "METHOD_GET",
		2: "METHOD_POST",
		3: "METHOD_PUT",
		4: "METHOD_PATCH",
		5: "METHOD_DELETE",
		6: "METHOD_OPTIONS",
		7: "METHOD_HEAD",
		8: "METHOD_TRACE",
		9: "METHOD_CONNECT",
	}
	Method_value = map[string]int32{
		"METHOD_UNSEPECIFIED": 0,
		"METHOD_GET":          1,
		"METHOD_POST":         2,
		"METHOD_PUT":          3,
		"METHOD_PATCH":        4,
		"METHOD_DELETE":       5,
		"METHOD_OPTIONS":      6,
		"METHOD_HEAD":         7,
		"METHOD_TRACE":        8,
		"METHOD_CONNECT":      9,
	}
)

func (x Method) Enum() *Method {
	p := new(Method)
	*p = x
	return p
}

func (x Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Method) Descriptor() protoreflect.EnumDescriptor {
	return file_mojo_http_method_proto_enumTypes[0].Descriptor()
}

func (Method) Type() protoreflect.EnumType {
	return &file_mojo_http_method_proto_enumTypes[0]
}

func (x Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Method.Descriptor instead.
func (Method) EnumDescriptor() ([]byte, []int) {
	return file_mojo_http_method_proto_rawDescGZIP(), []int{0}
}

var File_mojo_http_method_proto protoreflect.FileDescriptor

var file_mojo_http_method_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2a, 0xc2, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x17,
	0x0a, 0x13, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x45, 0x54, 0x48, 0x4f,
	0x44, 0x5f, 0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x45, 0x54, 0x48, 0x4f,
	0x44, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x45, 0x54, 0x48,
	0x4f, 0x44, 0x5f, 0x50, 0x55, 0x54, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x54, 0x48,
	0x4f, 0x44, 0x5f, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x05, 0x12, 0x12, 0x0a,
	0x0e, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10,
	0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x48, 0x45, 0x41, 0x44,
	0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x54, 0x52, 0x41,
	0x43, 0x45, 0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x09, 0x42, 0x58, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e,
	0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x42, 0x0b, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f,
	0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x3b, 0x68, 0x74,
	0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_http_method_proto_rawDescOnce sync.Once
	file_mojo_http_method_proto_rawDescData = file_mojo_http_method_proto_rawDesc
)

func file_mojo_http_method_proto_rawDescGZIP() []byte {
	file_mojo_http_method_proto_rawDescOnce.Do(func() {
		file_mojo_http_method_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_http_method_proto_rawDescData)
	})
	return file_mojo_http_method_proto_rawDescData
}

var file_mojo_http_method_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mojo_http_method_proto_goTypes = []interface{}{
	(Method)(0), // 0: mojo.http.Method
}
var file_mojo_http_method_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mojo_http_method_proto_init() }
func file_mojo_http_method_proto_init() {
	if File_mojo_http_method_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mojo_http_method_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_http_method_proto_goTypes,
		DependencyIndexes: file_mojo_http_method_proto_depIdxs,
		EnumInfos:         file_mojo_http_method_proto_enumTypes,
	}.Build()
	File_mojo_http_method_proto = out.File
	file_mojo_http_method_proto_rawDesc = nil
	file_mojo_http_method_proto_goTypes = nil
	file_mojo_http_method_proto_depIdxs = nil
}
