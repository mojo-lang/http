// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/http/method.proto

package org.mojo-lang.mojo.http;

/**
 * Protobuf enum {@code mojo.http.Method}
 */
public enum Method
    implements com.google.protobuf.ProtocolMessageEnum {
  /**
   * <code>METHOD_UNSEPECIFIED = 0;</code>
   */
  METHOD_UNSEPECIFIED(0),
  /**
   * <code>METHOD_GET = 1;</code>
   */
  METHOD_GET(1),
  /**
   * <code>METHOD_POST = 2;</code>
   */
  METHOD_POST(2),
  /**
   * <code>METHOD_PUT = 3;</code>
   */
  METHOD_PUT(3),
  /**
   * <code>METHOD_PATCH = 4;</code>
   */
  METHOD_PATCH(4),
  /**
   * <code>METHOD_DELETE = 5;</code>
   */
  METHOD_DELETE(5),
  /**
   * <code>METHOD_OPTIONS = 6;</code>
   */
  METHOD_OPTIONS(6),
  /**
   * <code>METHOD_HEAD = 7;</code>
   */
  METHOD_HEAD(7),
  /**
   * <code>METHOD_TRACE = 8;</code>
   */
  METHOD_TRACE(8),
  /**
   * <code>METHOD_CONNECT = 9;</code>
   */
  METHOD_CONNECT(9),
  UNRECOGNIZED(-1),
  ;

  /**
   * <code>METHOD_UNSEPECIFIED = 0;</code>
   */
  public static final int METHOD_UNSEPECIFIED_VALUE = 0;
  /**
   * <code>METHOD_GET = 1;</code>
   */
  public static final int METHOD_GET_VALUE = 1;
  /**
   * <code>METHOD_POST = 2;</code>
   */
  public static final int METHOD_POST_VALUE = 2;
  /**
   * <code>METHOD_PUT = 3;</code>
   */
  public static final int METHOD_PUT_VALUE = 3;
  /**
   * <code>METHOD_PATCH = 4;</code>
   */
  public static final int METHOD_PATCH_VALUE = 4;
  /**
   * <code>METHOD_DELETE = 5;</code>
   */
  public static final int METHOD_DELETE_VALUE = 5;
  /**
   * <code>METHOD_OPTIONS = 6;</code>
   */
  public static final int METHOD_OPTIONS_VALUE = 6;
  /**
   * <code>METHOD_HEAD = 7;</code>
   */
  public static final int METHOD_HEAD_VALUE = 7;
  /**
   * <code>METHOD_TRACE = 8;</code>
   */
  public static final int METHOD_TRACE_VALUE = 8;
  /**
   * <code>METHOD_CONNECT = 9;</code>
   */
  public static final int METHOD_CONNECT_VALUE = 9;


  public final int getNumber() {
    if (this == UNRECOGNIZED) {
      throw new java.lang.IllegalArgumentException(
          "Can't get the number of an unknown enum value.");
    }
    return value;
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   * @deprecated Use {@link #forNumber(int)} instead.
   */
  @java.lang.Deprecated
  public static Method valueOf(int value) {
    return forNumber(value);
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   */
  public static Method forNumber(int value) {
    switch (value) {
      case 0: return METHOD_UNSEPECIFIED;
      case 1: return METHOD_GET;
      case 2: return METHOD_POST;
      case 3: return METHOD_PUT;
      case 4: return METHOD_PATCH;
      case 5: return METHOD_DELETE;
      case 6: return METHOD_OPTIONS;
      case 7: return METHOD_HEAD;
      case 8: return METHOD_TRACE;
      case 9: return METHOD_CONNECT;
      default: return null;
    }
  }

  public static com.google.protobuf.Internal.EnumLiteMap<Method>
      internalGetValueMap() {
    return internalValueMap;
  }
  private static final com.google.protobuf.Internal.EnumLiteMap<
      Method> internalValueMap =
        new com.google.protobuf.Internal.EnumLiteMap<Method>() {
          public Method findValueByNumber(int number) {
            return Method.forNumber(number);
          }
        };

  public final com.google.protobuf.Descriptors.EnumValueDescriptor
      getValueDescriptor() {
    if (this == UNRECOGNIZED) {
      throw new java.lang.IllegalStateException(
          "Can't get the descriptor of an unrecognized enum value.");
    }
    return getDescriptor().getValues().get(ordinal());
  }
  public final com.google.protobuf.Descriptors.EnumDescriptor
      getDescriptorForType() {
    return getDescriptor();
  }
  public static final com.google.protobuf.Descriptors.EnumDescriptor
      getDescriptor() {
    return org.mojo-lang.mojo.http.MethodProto.getDescriptor().getEnumTypes().get(0);
  }

  private static final Method[] VALUES = values();

  public static Method valueOf(
      com.google.protobuf.Descriptors.EnumValueDescriptor desc) {
    if (desc.getType() != getDescriptor()) {
      throw new java.lang.IllegalArgumentException(
        "EnumValueDescriptor is not for this type.");
    }
    if (desc.getIndex() == -1) {
      return UNRECOGNIZED;
    }
    return VALUES[desc.getIndex()];
  }

  private final int value;

  private Method(int value) {
    this.value = value;
  }

  // @@protoc_insertion_point(enum_scope:mojo.http.Method)
}

