// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/http/request.proto

package org.mojolang.mojo.http;

/**
 * Protobuf type {@code mojo.http.RequestOptions}
 */
public final class RequestOptions extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.http.RequestOptions)
    RequestOptionsOrBuilder {
private static final long serialVersionUID = 0L;
  // Use RequestOptions.newBuilder() to construct.
  private RequestOptions(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private RequestOptions() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new RequestOptions();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.http.RequestProto.internal_static_mojo_http_RequestOptions_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.http.RequestProto.internal_static_mojo_http_RequestOptions_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.http.RequestOptions.class, org.mojolang.mojo.http.RequestOptions.Builder.class);
  }

  public static final int TIMEOUT_FIELD_NUMBER = 1;
  private org.mojolang.mojo.core.Duration timeout_;
  /**
   * <code>.mojo.core.Duration timeout = 1;</code>
   * @return Whether the timeout field is set.
   */
  @java.lang.Override
  public boolean hasTimeout() {
    return timeout_ != null;
  }
  /**
   * <code>.mojo.core.Duration timeout = 1;</code>
   * @return The timeout.
   */
  @java.lang.Override
  public org.mojolang.mojo.core.Duration getTimeout() {
    return timeout_ == null ? org.mojolang.mojo.core.Duration.getDefaultInstance() : timeout_;
  }
  /**
   * <code>.mojo.core.Duration timeout = 1;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.DurationOrBuilder getTimeoutOrBuilder() {
    return getTimeout();
  }

  public static final int MAX_RETRIES_FIELD_NUMBER = 2;
  private int maxRetries_;
  /**
   * <code>int32 max_retries = 2;</code>
   * @return The maxRetries.
   */
  @java.lang.Override
  public int getMaxRetries() {
    return maxRetries_;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (timeout_ != null) {
      output.writeMessage(1, getTimeout());
    }
    if (maxRetries_ != 0) {
      output.writeInt32(2, maxRetries_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (timeout_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, getTimeout());
    }
    if (maxRetries_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(2, maxRetries_);
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.mojolang.mojo.http.RequestOptions)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.http.RequestOptions other = (org.mojolang.mojo.http.RequestOptions) obj;

    if (hasTimeout() != other.hasTimeout()) return false;
    if (hasTimeout()) {
      if (!getTimeout()
          .equals(other.getTimeout())) return false;
    }
    if (getMaxRetries()
        != other.getMaxRetries()) return false;
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    if (hasTimeout()) {
      hash = (37 * hash) + TIMEOUT_FIELD_NUMBER;
      hash = (53 * hash) + getTimeout().hashCode();
    }
    hash = (37 * hash) + MAX_RETRIES_FIELD_NUMBER;
    hash = (53 * hash) + getMaxRetries();
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.http.RequestOptions parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.http.RequestOptions parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.http.RequestOptions parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.http.RequestOptions parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.http.RequestOptions parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.http.RequestOptions parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.http.RequestOptions parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.http.RequestOptions parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.http.RequestOptions parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.http.RequestOptions parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.http.RequestOptions parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.http.RequestOptions parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(org.mojolang.mojo.http.RequestOptions prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code mojo.http.RequestOptions}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.http.RequestOptions)
      org.mojolang.mojo.http.RequestOptionsOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.http.RequestProto.internal_static_mojo_http_RequestOptions_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.http.RequestProto.internal_static_mojo_http_RequestOptions_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.http.RequestOptions.class, org.mojolang.mojo.http.RequestOptions.Builder.class);
    }

    // Construct using org.mojolang.mojo.http.RequestOptions.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      if (timeoutBuilder_ == null) {
        timeout_ = null;
      } else {
        timeout_ = null;
        timeoutBuilder_ = null;
      }
      maxRetries_ = 0;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.http.RequestProto.internal_static_mojo_http_RequestOptions_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.http.RequestOptions getDefaultInstanceForType() {
      return org.mojolang.mojo.http.RequestOptions.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.http.RequestOptions build() {
      org.mojolang.mojo.http.RequestOptions result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.http.RequestOptions buildPartial() {
      org.mojolang.mojo.http.RequestOptions result = new org.mojolang.mojo.http.RequestOptions(this);
      if (timeoutBuilder_ == null) {
        result.timeout_ = timeout_;
      } else {
        result.timeout_ = timeoutBuilder_.build();
      }
      result.maxRetries_ = maxRetries_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof org.mojolang.mojo.http.RequestOptions) {
        return mergeFrom((org.mojolang.mojo.http.RequestOptions)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.http.RequestOptions other) {
      if (other == org.mojolang.mojo.http.RequestOptions.getDefaultInstance()) return this;
      if (other.hasTimeout()) {
        mergeTimeout(other.getTimeout());
      }
      if (other.getMaxRetries() != 0) {
        setMaxRetries(other.getMaxRetries());
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              input.readMessage(
                  getTimeoutFieldBuilder().getBuilder(),
                  extensionRegistry);

              break;
            } // case 10
            case 16: {
              maxRetries_ = input.readInt32();

              break;
            } // case 16
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }

    private org.mojolang.mojo.core.Duration timeout_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Duration, org.mojolang.mojo.core.Duration.Builder, org.mojolang.mojo.core.DurationOrBuilder> timeoutBuilder_;
    /**
     * <code>.mojo.core.Duration timeout = 1;</code>
     * @return Whether the timeout field is set.
     */
    public boolean hasTimeout() {
      return timeoutBuilder_ != null || timeout_ != null;
    }
    /**
     * <code>.mojo.core.Duration timeout = 1;</code>
     * @return The timeout.
     */
    public org.mojolang.mojo.core.Duration getTimeout() {
      if (timeoutBuilder_ == null) {
        return timeout_ == null ? org.mojolang.mojo.core.Duration.getDefaultInstance() : timeout_;
      } else {
        return timeoutBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.core.Duration timeout = 1;</code>
     */
    public Builder setTimeout(org.mojolang.mojo.core.Duration value) {
      if (timeoutBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        timeout_ = value;
        onChanged();
      } else {
        timeoutBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Duration timeout = 1;</code>
     */
    public Builder setTimeout(
        org.mojolang.mojo.core.Duration.Builder builderForValue) {
      if (timeoutBuilder_ == null) {
        timeout_ = builderForValue.build();
        onChanged();
      } else {
        timeoutBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.core.Duration timeout = 1;</code>
     */
    public Builder mergeTimeout(org.mojolang.mojo.core.Duration value) {
      if (timeoutBuilder_ == null) {
        if (timeout_ != null) {
          timeout_ =
            org.mojolang.mojo.core.Duration.newBuilder(timeout_).mergeFrom(value).buildPartial();
        } else {
          timeout_ = value;
        }
        onChanged();
      } else {
        timeoutBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Duration timeout = 1;</code>
     */
    public Builder clearTimeout() {
      if (timeoutBuilder_ == null) {
        timeout_ = null;
        onChanged();
      } else {
        timeout_ = null;
        timeoutBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.core.Duration timeout = 1;</code>
     */
    public org.mojolang.mojo.core.Duration.Builder getTimeoutBuilder() {
      
      onChanged();
      return getTimeoutFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.core.Duration timeout = 1;</code>
     */
    public org.mojolang.mojo.core.DurationOrBuilder getTimeoutOrBuilder() {
      if (timeoutBuilder_ != null) {
        return timeoutBuilder_.getMessageOrBuilder();
      } else {
        return timeout_ == null ?
            org.mojolang.mojo.core.Duration.getDefaultInstance() : timeout_;
      }
    }
    /**
     * <code>.mojo.core.Duration timeout = 1;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Duration, org.mojolang.mojo.core.Duration.Builder, org.mojolang.mojo.core.DurationOrBuilder> 
        getTimeoutFieldBuilder() {
      if (timeoutBuilder_ == null) {
        timeoutBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.core.Duration, org.mojolang.mojo.core.Duration.Builder, org.mojolang.mojo.core.DurationOrBuilder>(
                getTimeout(),
                getParentForChildren(),
                isClean());
        timeout_ = null;
      }
      return timeoutBuilder_;
    }

    private int maxRetries_ ;
    /**
     * <code>int32 max_retries = 2;</code>
     * @return The maxRetries.
     */
    @java.lang.Override
    public int getMaxRetries() {
      return maxRetries_;
    }
    /**
     * <code>int32 max_retries = 2;</code>
     * @param value The maxRetries to set.
     * @return This builder for chaining.
     */
    public Builder setMaxRetries(int value) {
      
      maxRetries_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 max_retries = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearMaxRetries() {
      
      maxRetries_ = 0;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:mojo.http.RequestOptions)
  }

  // @@protoc_insertion_point(class_scope:mojo.http.RequestOptions)
  private static final org.mojolang.mojo.http.RequestOptions DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.http.RequestOptions();
  }

  public static org.mojolang.mojo.http.RequestOptions getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<RequestOptions>
      PARSER = new com.google.protobuf.AbstractParser<RequestOptions>() {
    @java.lang.Override
    public RequestOptions parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<RequestOptions> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<RequestOptions> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.http.RequestOptions getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

