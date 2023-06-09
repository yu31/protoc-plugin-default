// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: default.proto

package io.github.yu31.protoc.pb.pbdefault;

/**
 * Protobuf type {@code default.TypeMessage}
 */
public final class TypeMessage extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:default.TypeMessage)
    TypeMessageOrBuilder {
private static final long serialVersionUID = 0L;
  // Use TypeMessage.newBuilder() to construct.
  private TypeMessage(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private TypeMessage() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new TypeMessage();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private TypeMessage(
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
          case 8: {
            bitField0_ |= 0x00000001;
            init_ = input.readBool();
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
    return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_TypeMessage_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_TypeMessage_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.github.yu31.protoc.pb.pbdefault.TypeMessage.class, io.github.yu31.protoc.pb.pbdefault.TypeMessage.Builder.class);
  }

  private int bitField0_;
  public static final int INIT_FIELD_NUMBER = 1;
  private boolean init_;
  /**
   * <pre>
   * Init indicates the `message` to empty value if it is nil pointer.
   * Only effective for language Go.
   * </pre>
   *
   * <code>optional bool init = 1;</code>
   * @return Whether the init field is set.
   */
  @java.lang.Override
  public boolean hasInit() {
    return ((bitField0_ & 0x00000001) != 0);
  }
  /**
   * <pre>
   * Init indicates the `message` to empty value if it is nil pointer.
   * Only effective for language Go.
   * </pre>
   *
   * <code>optional bool init = 1;</code>
   * @return The init.
   */
  @java.lang.Override
  public boolean getInit() {
    return init_;
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
      output.writeBool(1, init_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (((bitField0_ & 0x00000001) != 0)) {
      size += com.google.protobuf.CodedOutputStream
        .computeBoolSize(1, init_);
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
    if (!(obj instanceof io.github.yu31.protoc.pb.pbdefault.TypeMessage)) {
      return super.equals(obj);
    }
    io.github.yu31.protoc.pb.pbdefault.TypeMessage other = (io.github.yu31.protoc.pb.pbdefault.TypeMessage) obj;

    if (hasInit() != other.hasInit()) return false;
    if (hasInit()) {
      if (getInit()
          != other.getInit()) return false;
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
    if (hasInit()) {
      hash = (37 * hash) + INIT_FIELD_NUMBER;
      hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
          getInit());
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage parseFrom(
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
  public static Builder newBuilder(io.github.yu31.protoc.pb.pbdefault.TypeMessage prototype) {
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
   * Protobuf type {@code default.TypeMessage}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:default.TypeMessage)
      io.github.yu31.protoc.pb.pbdefault.TypeMessageOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_TypeMessage_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_TypeMessage_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.github.yu31.protoc.pb.pbdefault.TypeMessage.class, io.github.yu31.protoc.pb.pbdefault.TypeMessage.Builder.class);
    }

    // Construct using io.github.yu31.protoc.pb.pbdefault.TypeMessage.newBuilder()
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
      init_ = false;
      bitField0_ = (bitField0_ & ~0x00000001);
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.github.yu31.protoc.pb.pbdefault.PBDefault.internal_static_default_TypeMessage_descriptor;
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbdefault.TypeMessage getDefaultInstanceForType() {
      return io.github.yu31.protoc.pb.pbdefault.TypeMessage.getDefaultInstance();
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbdefault.TypeMessage build() {
      io.github.yu31.protoc.pb.pbdefault.TypeMessage result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.github.yu31.protoc.pb.pbdefault.TypeMessage buildPartial() {
      io.github.yu31.protoc.pb.pbdefault.TypeMessage result = new io.github.yu31.protoc.pb.pbdefault.TypeMessage(this);
      int from_bitField0_ = bitField0_;
      int to_bitField0_ = 0;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.init_ = init_;
        to_bitField0_ |= 0x00000001;
      }
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
      if (other instanceof io.github.yu31.protoc.pb.pbdefault.TypeMessage) {
        return mergeFrom((io.github.yu31.protoc.pb.pbdefault.TypeMessage)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.github.yu31.protoc.pb.pbdefault.TypeMessage other) {
      if (other == io.github.yu31.protoc.pb.pbdefault.TypeMessage.getDefaultInstance()) return this;
      if (other.hasInit()) {
        setInit(other.getInit());
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
      io.github.yu31.protoc.pb.pbdefault.TypeMessage parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (io.github.yu31.protoc.pb.pbdefault.TypeMessage) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private boolean init_ ;
    /**
     * <pre>
     * Init indicates the `message` to empty value if it is nil pointer.
     * Only effective for language Go.
     * </pre>
     *
     * <code>optional bool init = 1;</code>
     * @return Whether the init field is set.
     */
    @java.lang.Override
    public boolean hasInit() {
      return ((bitField0_ & 0x00000001) != 0);
    }
    /**
     * <pre>
     * Init indicates the `message` to empty value if it is nil pointer.
     * Only effective for language Go.
     * </pre>
     *
     * <code>optional bool init = 1;</code>
     * @return The init.
     */
    @java.lang.Override
    public boolean getInit() {
      return init_;
    }
    /**
     * <pre>
     * Init indicates the `message` to empty value if it is nil pointer.
     * Only effective for language Go.
     * </pre>
     *
     * <code>optional bool init = 1;</code>
     * @param value The init to set.
     * @return This builder for chaining.
     */
    public Builder setInit(boolean value) {
      bitField0_ |= 0x00000001;
      init_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * Init indicates the `message` to empty value if it is nil pointer.
     * Only effective for language Go.
     * </pre>
     *
     * <code>optional bool init = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearInit() {
      bitField0_ = (bitField0_ & ~0x00000001);
      init_ = false;
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


    // @@protoc_insertion_point(builder_scope:default.TypeMessage)
  }

  // @@protoc_insertion_point(class_scope:default.TypeMessage)
  private static final io.github.yu31.protoc.pb.pbdefault.TypeMessage DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.github.yu31.protoc.pb.pbdefault.TypeMessage();
  }

  public static io.github.yu31.protoc.pb.pbdefault.TypeMessage getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<TypeMessage>
      PARSER = new com.google.protobuf.AbstractParser<TypeMessage>() {
    @java.lang.Override
    public TypeMessage parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new TypeMessage(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<TypeMessage> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<TypeMessage> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.github.yu31.protoc.pb.pbdefault.TypeMessage getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

