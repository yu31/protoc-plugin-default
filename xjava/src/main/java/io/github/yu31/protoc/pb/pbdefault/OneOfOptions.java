// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: default.proto

package io.github.yu31.protoc.pb.pbdefault;

/**
 * <pre>
 * OneOfOptions is the options for the oneof field.
 * </pre>
 *
 * Protobuf type {@code default.OneOfOptions}
 */
public final class OneOfOptions extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:default.OneOfOptions)
    OneOfOptionsOrBuilder {
private static final long serialVersionUID = 0L;
  // Use OneOfOptions.newBuilder() to construct.
  private OneOfOptions(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private OneOfOptions() {
    field_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new OneOfOptions();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private OneOfOptions(
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
          case 10: {
            java.lang.String s = input.readStringRequireUtf8();
            bitField0_ |= 0x00000001;
            field_ = s;
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
    return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_OneOfOptions_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_OneOfOptions_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.github.yu31.protoc.pb.pbdefault.OneOfOptions.class, io.github.yu31.protoc.pb.pbdefault.OneOfOptions.Builder.class);
  }

  private int bitField0_;
  public static final int FIELD_FIELD_NUMBER = 1;
  private volatile java.lang.Object field_;
  /**
   * <pre>
   * The name of filed in `oneof` part.
   * </pre>
   *
   * <code>optional string field = 1;</code>
   * @return Whether the field field is set.
   */
  @java.lang.Override
  public boolean hasField() {
    return ((bitField0_ & 0x00000001) != 0);
  }
  /**
   * <pre>
   * The name of filed in `oneof` part.
   * </pre>
   *
   * <code>optional string field = 1;</code>
   * @return The field.
   */
  @java.lang.Override
  public java.lang.String getField() {
    java.lang.Object ref = field_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      field_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * The name of filed in `oneof` part.
   * </pre>
   *
   * <code>optional string field = 1;</code>
   * @return The bytes for field.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getFieldBytes() {
    java.lang.Object ref = field_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      field_ = b;
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
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, field_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (((bitField0_ & 0x00000001) != 0)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, field_);
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
    if (!(obj instanceof io.github.yu31.protoc.pb.pbdefault.OneOfOptions)) {
      return super.equals(obj);
    }
    io.github.yu31.protoc.pb.pbdefault.OneOfOptions other = (io.github.yu31.protoc.pb.pbdefault.OneOfOptions) obj;

    if (hasField() != other.hasField()) return false;
    if (hasField()) {
      if (!getField()
          .equals(other.getField())) return false;
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
    if (hasField()) {
      hash = (37 * hash) + FIELD_FIELD_NUMBER;
      hash = (53 * hash) + getField().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions parseFrom(
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
  public static Builder newBuilder(io.github.yu31.protoc.pb.pbdefault.OneOfOptions prototype) {
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
   * <pre>
   * OneOfOptions is the options for the oneof field.
   * </pre>
   *
   * Protobuf type {@code default.OneOfOptions}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:default.OneOfOptions)
      io.github.yu31.protoc.pb.pbdefault.OneOfOptionsOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_OneOfOptions_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_OneOfOptions_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.github.yu31.protoc.pb.pbdefault.OneOfOptions.class, io.github.yu31.protoc.pb.pbdefault.OneOfOptions.Builder.class);
    }

    // Construct using io.github.yu31.protoc.pb.pbdefault.OneOfOptions.newBuilder()
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
      field_ = "";
      bitField0_ = (bitField0_ & ~0x00000001);
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_OneOfOptions_descriptor;
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbdefault.OneOfOptions getDefaultInstanceForType() {
      return io.github.yu31.protoc.pb.pbdefault.OneOfOptions.getDefaultInstance();
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbdefault.OneOfOptions build() {
      io.github.yu31.protoc.pb.pbdefault.OneOfOptions result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbdefault.OneOfOptions buildPartial() {
      io.github.yu31.protoc.pb.pbdefault.OneOfOptions result = new io.github.yu31.protoc.pb.pbdefault.OneOfOptions(this);
      int from_bitField0_ = bitField0_;
      int to_bitField0_ = 0;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        to_bitField0_ |= 0x00000001;
      }
      result.field_ = field_;
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
      if (other instanceof io.github.yu31.protoc.pb.pbdefault.OneOfOptions) {
        return mergeFrom((io.github.yu31.protoc.pb.pbdefault.OneOfOptions)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.github.yu31.protoc.pb.pbdefault.OneOfOptions other) {
      if (other == io.github.yu31.protoc.pb.pbdefault.OneOfOptions.getDefaultInstance()) return this;
      if (other.hasField()) {
        bitField0_ |= 0x00000001;
        field_ = other.field_;
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
      io.github.yu31.protoc.pb.pbdefault.OneOfOptions parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (io.github.yu31.protoc.pb.pbdefault.OneOfOptions) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private java.lang.Object field_ = "";
    /**
     * <pre>
     * The name of filed in `oneof` part.
     * </pre>
     *
     * <code>optional string field = 1;</code>
     * @return Whether the field field is set.
     */
    public boolean hasField() {
      return ((bitField0_ & 0x00000001) != 0);
    }
    /**
     * <pre>
     * The name of filed in `oneof` part.
     * </pre>
     *
     * <code>optional string field = 1;</code>
     * @return The field.
     */
    public java.lang.String getField() {
      java.lang.Object ref = field_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        field_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     * The name of filed in `oneof` part.
     * </pre>
     *
     * <code>optional string field = 1;</code>
     * @return The bytes for field.
     */
    public com.google.protobuf.ByteString
        getFieldBytes() {
      java.lang.Object ref = field_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        field_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * The name of filed in `oneof` part.
     * </pre>
     *
     * <code>optional string field = 1;</code>
     * @param value The field to set.
     * @return This builder for chaining.
     */
    public Builder setField(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  bitField0_ |= 0x00000001;
      field_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * The name of filed in `oneof` part.
     * </pre>
     *
     * <code>optional string field = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearField() {
      bitField0_ = (bitField0_ & ~0x00000001);
      field_ = getDefaultInstance().getField();
      onChanged();
      return this;
    }
    /**
     * <pre>
     * The name of filed in `oneof` part.
     * </pre>
     *
     * <code>optional string field = 1;</code>
     * @param value The bytes for field to set.
     * @return This builder for chaining.
     */
    public Builder setFieldBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      bitField0_ |= 0x00000001;
      field_ = value;
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


    // @@protoc_insertion_point(builder_scope:default.OneOfOptions)
  }

  // @@protoc_insertion_point(class_scope:default.OneOfOptions)
  private static final io.github.yu31.protoc.pb.pbdefault.OneOfOptions DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.github.yu31.protoc.pb.pbdefault.OneOfOptions();
  }

  public static io.github.yu31.protoc.pb.pbdefault.OneOfOptions getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<OneOfOptions>
      PARSER = new com.google.protobuf.AbstractParser<OneOfOptions>() {
    @java.lang.Override
    public OneOfOptions parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new OneOfOptions(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<OneOfOptions> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<OneOfOptions> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.github.yu31.protoc.pb.pbdefault.OneOfOptions getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}
