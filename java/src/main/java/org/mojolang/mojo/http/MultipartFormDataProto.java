// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/http/multipart_form_data.proto

package org.mojolang.mojo.http;

public final class MultipartFormDataProto {
  private MultipartFormDataProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_http_MultipartFormData_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_http_MultipartFormData_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n#mojo/http/multipart_form_data.proto\022\tm" +
      "ojo.http\032!mojo/http/form_data_content.pr" +
      "oto\"P\n\021MultipartFormData\022\020\n\010boundary\030\001 \001" +
      "(\t\022)\n\005parts\030\002 \003(\0132\032.mojo.http.FormDataCo" +
      "ntentBc\n\026org.mojolang.mojo.httpB\026Multipa" +
      "rtFormDataProtoP\001Z/github.com/mojo-lang/" +
      "http/go/pkg/mojo/http;httpb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.http.FormDataContentProto.getDescriptor(),
        });
    internal_static_mojo_http_MultipartFormData_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_http_MultipartFormData_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_http_MultipartFormData_descriptor,
        new java.lang.String[] { "Boundary", "Parts", });
    org.mojolang.mojo.http.FormDataContentProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
