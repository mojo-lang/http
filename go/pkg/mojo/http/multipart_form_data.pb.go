// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: mojo/http/multipart_form_data.proto

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

type MultipartFormData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boundary string             `protobuf:"bytes,1,opt,name=boundary,proto3" json:"boundary,omitempty"`
	Parts    []*FormDataContent `protobuf:"bytes,2,rep,name=parts,proto3" json:"parts,omitempty"`
}

func (x *MultipartFormData) Reset() {
	*x = MultipartFormData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_http_multipart_form_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipartFormData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipartFormData) ProtoMessage() {}

func (x *MultipartFormData) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_http_multipart_form_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipartFormData.ProtoReflect.Descriptor instead.
func (*MultipartFormData) Descriptor() ([]byte, []int) {
	return file_mojo_http_multipart_form_data_proto_rawDescGZIP(), []int{0}
}

func (x *MultipartFormData) GetBoundary() string {
	if x != nil {
		return x.Boundary
	}
	return ""
}

func (x *MultipartFormData) GetParts() []*FormDataContent {
	if x != nil {
		return x.Parts
	}
	return nil
}

var File_mojo_http_multipart_form_data_proto protoreflect.FileDescriptor

var file_mojo_http_multipart_form_data_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x1a, 0x21, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x11, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74,
	0x46, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x46, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x42, 0x63, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f,
	0x6a, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x42, 0x16, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a,
	0x6f, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x3b, 0x68, 0x74, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mojo_http_multipart_form_data_proto_rawDescOnce sync.Once
	file_mojo_http_multipart_form_data_proto_rawDescData = file_mojo_http_multipart_form_data_proto_rawDesc
)

func file_mojo_http_multipart_form_data_proto_rawDescGZIP() []byte {
	file_mojo_http_multipart_form_data_proto_rawDescOnce.Do(func() {
		file_mojo_http_multipart_form_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_http_multipart_form_data_proto_rawDescData)
	})
	return file_mojo_http_multipart_form_data_proto_rawDescData
}

var file_mojo_http_multipart_form_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mojo_http_multipart_form_data_proto_goTypes = []interface{}{
	(*MultipartFormData)(nil), // 0: mojo.http.MultipartFormData
	(*FormDataContent)(nil),   // 1: mojo.http.FormDataContent
}
var file_mojo_http_multipart_form_data_proto_depIdxs = []int32{
	1, // 0: mojo.http.MultipartFormData.parts:type_name -> mojo.http.FormDataContent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mojo_http_multipart_form_data_proto_init() }
func file_mojo_http_multipart_form_data_proto_init() {
	if File_mojo_http_multipart_form_data_proto != nil {
		return
	}
	file_mojo_http_form_data_content_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mojo_http_multipart_form_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipartFormData); i {
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
			RawDescriptor: file_mojo_http_multipart_form_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_http_multipart_form_data_proto_goTypes,
		DependencyIndexes: file_mojo_http_multipart_form_data_proto_depIdxs,
		MessageInfos:      file_mojo_http_multipart_form_data_proto_msgTypes,
	}.Build()
	File_mojo_http_multipart_form_data_proto = out.File
	file_mojo_http_multipart_form_data_proto_rawDesc = nil
	file_mojo_http_multipart_form_data_proto_goTypes = nil
	file_mojo_http_multipart_form_data_proto_depIdxs = nil
}
