// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/http/form_data_content.proto

package org.mojo-lang.mojo.http;

public interface FormDataContentOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.http.FormDataContent)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.mojo.http.FormDataContent.Disposition disposition = 1;</code>
   * @return Whether the disposition field is set.
   */
  boolean hasDisposition();
  /**
   * <code>.mojo.http.FormDataContent.Disposition disposition = 1;</code>
   * @return The disposition.
   */
  org.mojo-lang.mojo.http.FormDataContent.Disposition getDisposition();
  /**
   * <code>.mojo.http.FormDataContent.Disposition disposition = 1;</code>
   */
  org.mojo-lang.mojo.http.FormDataContent.DispositionOrBuilder getDispositionOrBuilder();

  /**
   * <code>.mojo.core.MediaType type = 2;</code>
   * @return Whether the type field is set.
   */
  boolean hasType();
  /**
   * <code>.mojo.core.MediaType type = 2;</code>
   * @return The type.
   */
  org.mojolang.mojo.core.MediaType getType();
  /**
   * <code>.mojo.core.MediaType type = 2;</code>
   */
  org.mojolang.mojo.core.MediaTypeOrBuilder getTypeOrBuilder();

  /**
   * <code>string transfer_encoding = 3;</code>
   * @return The transferEncoding.
   */
  java.lang.String getTransferEncoding();
  /**
   * <code>string transfer_encoding = 3;</code>
   * @return The bytes for transferEncoding.
   */
  com.google.protobuf.ByteString
      getTransferEncodingBytes();

  /**
   * <code>.mojo.http.Headers headers = 4;</code>
   * @return Whether the headers field is set.
   */
  boolean hasHeaders();
  /**
   * <code>.mojo.http.Headers headers = 4;</code>
   * @return The headers.
   */
  org.mojo-lang.mojo.http.Headers getHeaders();
  /**
   * <code>.mojo.http.Headers headers = 4;</code>
   */
  org.mojo-lang.mojo.http.HeadersOrBuilder getHeadersOrBuilder();

  /**
   * <code>.mojo.core.Value value = 5;</code>
   * @return Whether the value field is set.
   */
  boolean hasValue();
  /**
   * <code>.mojo.core.Value value = 5;</code>
   * @return The value.
   */
  org.mojolang.mojo.core.Value getValue();
  /**
   * <code>.mojo.core.Value value = 5;</code>
   */
  org.mojolang.mojo.core.ValueOrBuilder getValueOrBuilder();
}
