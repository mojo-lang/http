// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/http/protocol.proto

package org.mojo-lang.mojo.http;

public final class ProtocolProto {
  private ProtocolProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_http_Protocol_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_http_Protocol_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\030mojo/http/protocol.proto\022\tmojo.http\"(\n" +
      "\010Protocol\022\r\n\005major\030\001 \001(\005\022\r\n\005minor\030\002 \001(\005B" +
      "[\n\027org.mojo-lang.mojo.httpB\rProtocolProt" +
      "oP\001Z/github.com/mojo-lang/http/go/pkg/mo" +
      "jo/http;httpb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_mojo_http_Protocol_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_http_Protocol_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_http_Protocol_descriptor,
        new java.lang.String[] { "Major", "Minor", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}