// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/http/template_header.proto

package org.mojolang.mojo.http;

public interface TemplateHeaderOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.http.TemplateHeader)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string name = 1;</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <code>string name = 1;</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <code>.mojo.core.TemplateString value = 2;</code>
   * @return Whether the value field is set.
   */
  boolean hasValue();
  /**
   * <code>.mojo.core.TemplateString value = 2;</code>
   * @return The value.
   */
  org.mojolang.mojo.core.TemplateString getValue();
  /**
   * <code>.mojo.core.TemplateString value = 2;</code>
   */
  org.mojolang.mojo.core.TemplateStringOrBuilder getValueOrBuilder();
}