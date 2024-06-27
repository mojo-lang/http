// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/http/request.proto

package org.mojolang.mojo.http;

public interface RequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.http.Request)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.mojo.http.Method method = 1;</code>
   * @return The enum numeric value on the wire for method.
   */
  int getMethodValue();
  /**
   * <code>.mojo.http.Method method = 1;</code>
   * @return The method.
   */
  org.mojolang.mojo.http.Method getMethod();

  /**
   * <code>.mojo.core.Url url = 2;</code>
   * @return Whether the url field is set.
   */
  boolean hasUrl();
  /**
   * <code>.mojo.core.Url url = 2;</code>
   * @return The url.
   */
  org.mojolang.mojo.core.Url getUrl();
  /**
   * <code>.mojo.core.Url url = 2;</code>
   */
  org.mojolang.mojo.core.UrlOrBuilder getUrlOrBuilder();

  /**
   * <code>.mojo.http.Version version = 3;</code>
   * @return Whether the version field is set.
   */
  boolean hasVersion();
  /**
   * <code>.mojo.http.Version version = 3;</code>
   * @return The version.
   */
  org.mojolang.mojo.http.Version getVersion();
  /**
   * <code>.mojo.http.Version version = 3;</code>
   */
  org.mojolang.mojo.http.VersionOrBuilder getVersionOrBuilder();

  /**
   * <code>.mojo.http.Headers headers = 4;</code>
   * @return Whether the headers field is set.
   */
  boolean hasHeaders();
  /**
   * <code>.mojo.http.Headers headers = 4;</code>
   * @return The headers.
   */
  org.mojolang.mojo.http.Headers getHeaders();
  /**
   * <code>.mojo.http.Headers headers = 4;</code>
   */
  org.mojolang.mojo.http.HeadersOrBuilder getHeadersOrBuilder();

  /**
   * <code>.mojo.core.Value body = 5;</code>
   * @return Whether the body field is set.
   */
  boolean hasBody();
  /**
   * <code>.mojo.core.Value body = 5;</code>
   * @return The body.
   */
  org.mojolang.mojo.core.Value getBody();
  /**
   * <code>.mojo.core.Value body = 5;</code>
   */
  org.mojolang.mojo.core.ValueOrBuilder getBodyOrBuilder();
}