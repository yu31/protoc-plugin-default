// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: default.proto

package io.github.yu31.protoc.pb.pbdefault;

/**
 * Protobuf type {@code default.TypeAny}
 */
public final class TypeAny extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:default.TypeAny)
    TypeAnyOrBuilder {
private static final long serialVersionUID = 0L;
  // Use TypeAny.newBuilder() to construct.
  private TypeAny(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private TypeAny() {
    typeUrl_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new TypeAny();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private TypeAny(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 90: {
            java.lang.String s = input.readStringRequireUtf8();
            bitField0_ |= 0x00000001;
            typeUrl_ = s;
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_TypeAny_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_TypeAny_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.github.yu31.protoc.pb.pbdefault.TypeAny.class, io.github.yu31.protoc.pb.pbdefault.TypeAny.Builder.class);
  }

  private int bitField0_;
  public static final int TYPE_URL_FIELD_NUMBER = 11;
  private volatile java.lang.Object typeUrl_;
  /**
   * <pre>
   * The type_url are combines with the import path(excludes suffix .proto) of proto file and the message name.
   * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions`.
   * You should use the point number connection to the parent message for embedded message.
   * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions.KV`.
   * You can use `.` represents the import path if the message is location in the current proto file.
   * And the entry format like `.MapOptions.KV`.
   * </pre>
   *
   * <code>optional string type_url = 11;</code>
   * @return Whether the typeUrl field is set.
   */
  @java.lang.Override
  public boolean hasTypeUrl() {
    return ((bitField0_ & 0x00000001) != 0);
  }
  /**
   * <pre>
   * The type_url are combines with the import path(excludes suffix .proto) of proto file and the message name.
   * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions`.
   * You should use the point number connection to the parent message for embedded message.
   * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions.KV`.
   * You can use `.` represents the import path if the message is location in the current proto file.
   * And the entry format like `.MapOptions.KV`.
   * </pre>
   *
   * <code>optional string type_url = 11;</code>
   * @return The typeUrl.
   */
  @java.lang.Override
  public java.lang.String getTypeUrl() {
    java.lang.Object ref = typeUrl_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      typeUrl_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * The type_url are combines with the import path(excludes suffix .proto) of proto file and the message name.
   * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions`.
   * You should use the point number connection to the parent message for embedded message.
   * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions.KV`.
   * You can use `.` represents the import path if the message is location in the current proto file.
   * And the entry format like `.MapOptions.KV`.
   * </pre>
   *
   * <code>optional string type_url = 11;</code>
   * @return The bytes for typeUrl.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getTypeUrlBytes() {
    java.lang.Object ref = typeUrl_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      typeUrl_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
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
    if (((bitField0_ & 0x00000001) != 0)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 11, typeUrl_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (((bitField0_ & 0x00000001) != 0)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(11, typeUrl_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof io.github.yu31.protoc.pb.pbdefault.TypeAny)) {
      return super.equals(obj);
    }
    io.github.yu31.protoc.pb.pbdefault.TypeAny other = (io.github.yu31.protoc.pb.pbdefault.TypeAny) obj;

    if (hasTypeUrl() != other.hasTypeUrl()) return false;
    if (hasTypeUrl()) {
      if (!getTypeUrl()
          .equals(other.getTypeUrl())) return false;
    }
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    if (hasTypeUrl()) {
      hash = (37 * hash) + TYPE_URL_FIELD_NUMBER;
      hash = (53 * hash) + getTypeUrl().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.github.yu31.protoc.pb.pbdefault.TypeAny parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeAny parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeAny parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeAny parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeAny parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeAny parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeAny parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeAny parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeAny parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeAny parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeAny parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeAny parseFrom(
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
  public static Builder newBuilder(io.github.yu31.protoc.pb.pbdefault.TypeAny prototype) {
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
   * Protobuf type {@code default.TypeAny}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:default.TypeAny)
      io.github.yu31.protoc.pb.pbdefault.TypeAnyOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_TypeAny_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_TypeAny_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.github.yu31.protoc.pb.pbdefault.TypeAny.class, io.github.yu31.protoc.pb.pbdefault.TypeAny.Builder.class);
    }

    // Construct using io.github.yu31.protoc.pb.pbdefault.TypeAny.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      typeUrl_ = "";
      bitField0_ = (bitField0_ & ~0x00000001);
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_TypeAny_descriptor;
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbdefault.TypeAny getDefaultInstanceForType() {
      return io.github.yu31.protoc.pb.pbdefault.TypeAny.getDefaultInstance();
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbdefault.TypeAny build() {
      io.github.yu31.protoc.pb.pbdefault.TypeAny result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbdefault.TypeAny buildPartial() {
      io.github.yu31.protoc.pb.pbdefault.TypeAny result = new io.github.yu31.protoc.pb.pbdefault.TypeAny(this);
      int from_bitField0_ = bitField0_;
      int to_bitField0_ = 0;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        to_bitField0_ |= 0x00000001;
      }
      result.typeUrl_ = typeUrl_;
      result.bitField0_ = to_bitField0_;
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
      if (other instanceof io.github.yu31.protoc.pb.pbdefault.TypeAny) {
        return mergeFrom((io.github.yu31.protoc.pb.pbdefault.TypeAny)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.github.yu31.protoc.pb.pbdefault.TypeAny other) {
      if (other == io.github.yu31.protoc.pb.pbdefault.TypeAny.getDefaultInstance()) return this;
      if (other.hasTypeUrl()) {
        bitField0_ |= 0x00000001;
        typeUrl_ = other.typeUrl_;
        onChanged();
      }
      this.mergeUnknownFields(other.unknownFields);
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
      io.github.yu31.protoc.pb.pbdefault.TypeAny parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (io.github.yu31.protoc.pb.pbdefault.TypeAny) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private java.lang.Object typeUrl_ = "";
    /**
     * <pre>
     * The type_url are combines with the import path(excludes suffix .proto) of proto file and the message name.
     * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions`.
     * You should use the point number connection to the parent message for embedded message.
     * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions.KV`.
     * You can use `.` represents the import path if the message is location in the current proto file.
     * And the entry format like `.MapOptions.KV`.
     * </pre>
     *
     * <code>optional string type_url = 11;</code>
     * @return Whether the typeUrl field is set.
     */
    public boolean hasTypeUrl() {
      return ((bitField0_ & 0x00000001) != 0);
    }
    /**
     * <pre>
     * The type_url are combines with the import path(excludes suffix .proto) of proto file and the message name.
     * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions`.
     * You should use the point number connection to the parent message for embedded message.
     * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions.KV`.
     * You can use `.` represents the import path if the message is location in the current proto file.
     * And the entry format like `.MapOptions.KV`.
     * </pre>
     *
     * <code>optional string type_url = 11;</code>
     * @return The typeUrl.
     */
    public java.lang.String getTypeUrl() {
      java.lang.Object ref = typeUrl_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        typeUrl_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     * The type_url are combines with the import path(excludes suffix .proto) of proto file and the message name.
     * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions`.
     * You should use the point number connection to the parent message for embedded message.
     * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions.KV`.
     * You can use `.` represents the import path if the message is location in the current proto file.
     * And the entry format like `.MapOptions.KV`.
     * </pre>
     *
     * <code>optional string type_url = 11;</code>
     * @return The bytes for typeUrl.
     */
    public com.google.protobuf.ByteString
        getTypeUrlBytes() {
      java.lang.Object ref = typeUrl_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        typeUrl_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * The type_url are combines with the import path(excludes suffix .proto) of proto file and the message name.
     * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions`.
     * You should use the point number connection to the parent message for embedded message.
     * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions.KV`.
     * You can use `.` represents the import path if the message is location in the current proto file.
     * And the entry format like `.MapOptions.KV`.
     * </pre>
     *
     * <code>optional string type_url = 11;</code>
     * @param value The typeUrl to set.
     * @return This builder for chaining.
     */
    public Builder setTypeUrl(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  bitField0_ |= 0x00000001;
      typeUrl_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * The type_url are combines with the import path(excludes suffix .proto) of proto file and the message name.
     * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions`.
     * You should use the point number connection to the parent message for embedded message.
     * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions.KV`.
     * You can use `.` represents the import path if the message is location in the current proto file.
     * And the entry format like `.MapOptions.KV`.
     * </pre>
     *
     * <code>optional string type_url = 11;</code>
     * @return This builder for chaining.
     */
    public Builder clearTypeUrl() {
      bitField0_ = (bitField0_ & ~0x00000001);
      typeUrl_ = getDefaultInstance().getTypeUrl();
      onChanged();
      return this;
    }
    /**
     * <pre>
     * The type_url are combines with the import path(excludes suffix .proto) of proto file and the message name.
     * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions`.
     * You should use the point number connection to the parent message for embedded message.
     * And the entry format like `github.com/yu31/protoc-plugin-default/proto/default.MapOptions.KV`.
     * You can use `.` represents the import path if the message is location in the current proto file.
     * And the entry format like `.MapOptions.KV`.
     * </pre>
     *
     * <code>optional string type_url = 11;</code>
     * @param value The bytes for typeUrl to set.
     * @return This builder for chaining.
     */
    public Builder setTypeUrlBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      bitField0_ |= 0x00000001;
      typeUrl_ = value;
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


    // @@protoc_insertion_point(builder_scope:default.TypeAny)
  }

  // @@protoc_insertion_point(class_scope:default.TypeAny)
  private static final io.github.yu31.protoc.pb.pbdefault.TypeAny DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.github.yu31.protoc.pb.pbdefault.TypeAny();
  }

  public static io.github.yu31.protoc.pb.pbdefault.TypeAny getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<TypeAny>
      PARSER = new com.google.protobuf.AbstractParser<TypeAny>() {
    @java.lang.Override
    public TypeAny parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new TypeAny(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<TypeAny> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<TypeAny> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.github.yu31.protoc.pb.pbdefault.TypeAny getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

