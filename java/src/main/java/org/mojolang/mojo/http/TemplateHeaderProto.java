// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/http/template_header.proto

package org.mojolang.mojo.http;

public final class TemplateHeaderProto {
  private TemplateHeaderProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_http_TemplateHeader_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_http_TemplateHeader_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\037mojo/http/template_header.proto\022\tmojo." +
      "http\032\037mojo/core/template_string.proto\"H\n" +
      "\016TemplateHeader\022\014\n\004name\030\001 \001(\t\022(\n\005value\030\002" +
      " \001(\0132\031.mojo.core.TemplateStringB`\n\026org.m" +
      "ojolang.mojo.httpB\023TemplateHeaderProtoP\001" +
      "Z/github.com/mojo-lang/http/go/pkg/mojo/" +
      "http;httpb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.core.TemplateStringProto.getDescriptor(),
        });
    internal_static_mojo_http_TemplateHeader_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_http_TemplateHeader_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_http_TemplateHeader_descriptor,
        new java.lang.String[] { "Name", "Value", });
    org.mojolang.mojo.core.TemplateStringProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
