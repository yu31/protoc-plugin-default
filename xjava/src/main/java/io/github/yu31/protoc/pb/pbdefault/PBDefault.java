// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: default.proto

package io.github.yu31.protoc.pb.pbdefault;

public final class PBDefault {
  private PBDefault() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
    registry.add(io.github.yu31.protoc.pb.pbdefault.PBDefault.field);
    registry.add(io.github.yu31.protoc.pb.pbdefault.PBDefault.oneof);
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public static final int FIELD_FIELD_NUMBER = 64020;
  /**
   * <code>extend .google.protobuf.FieldOptions { ... }</code>
   */
  public static final
    com.google.protobuf.GeneratedMessage.GeneratedExtension<
      com.google.protobuf.DescriptorProtos.FieldOptions,
      io.github.yu31.protoc.pb.pbdefault.FieldOptions> field = com.google.protobuf.GeneratedMessage
          .newFileScopedGeneratedExtension(
        io.github.yu31.protoc.pb.pbdefault.FieldOptions.class,
        io.github.yu31.protoc.pb.pbdefault.FieldOptions.getDefaultInstance());
  public static final int ONEOF_FIELD_NUMBER = 64021;
  /**
   * <code>extend .google.protobuf.OneofOptions { ... }</code>
   */
  public static final
    com.google.protobuf.GeneratedMessage.GeneratedExtension<
      com.google.protobuf.DescriptorProtos.OneofOptions,
      io.github.yu31.protoc.pb.pbdefault.OneOfOptions> oneof = com.google.protobuf.GeneratedMessage
          .newFileScopedGeneratedExtension(
        io.github.yu31.protoc.pb.pbdefault.OneOfOptions.class,
        io.github.yu31.protoc.pb.pbdefault.OneOfOptions.getDefaultInstance());
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_default_OneOfOptions_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_default_OneOfOptions_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_default_FieldOptions_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_default_FieldOptions_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_default_PlainOptions_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_default_PlainOptions_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_default_RepeatedOptions_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_default_RepeatedOptions_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_default_MapOptions_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_default_MapOptions_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_default_MapOptions_KV_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_default_MapOptions_KV_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_default_FieldType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_default_FieldType_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_default_TypeMessage_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_default_TypeMessage_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_default_TypeAny_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_default_TypeAny_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\rdefault.proto\022\007default\032 google/protobu" +
      "f/descriptor.proto\032\036google/protobuf/dura" +
      "tion.proto\032\037google/protobuf/timestamp.pr" +
      "oto\",\n\014OneOfOptions\022\022\n\005field\030\001 \001(\tH\000\210\001\001B" +
      "\010\n\006_field\"\220\001\n\014FieldOptions\022&\n\005plain\030\001 \001(" +
      "\0132\025.default.PlainOptionsH\000\022,\n\010repeated\030\002" +
      " \001(\0132\030.default.RepeatedOptionsH\000\022\"\n\003map\030" +
      "\003 \001(\0132\023.default.MapOptionsH\000B\006\n\004Kind\"\203\001\n" +
      "\014PlainOptions\022!\n\005value\030\001 \001(\0132\022.default.F" +
      "ieldType\022\031\n\014ignore_empty\030\002 \001(\010H\000\210\001\001\022\026\n\ts" +
      "kip_eval\030\013 \001(\010H\001\210\001\001B\017\n\r_ignore_emptyB\014\n\n" +
      "_skip_eval\"\206\001\n\017RepeatedOptions\022!\n\005items\030" +
      "\001 \003(\0132\022.default.FieldType\022\031\n\014ignore_empt" +
      "y\030\002 \001(\010H\000\210\001\001\022\026\n\tskip_eval\030\013 \001(\010H\001\210\001\001B\017\n\r" +
      "_ignore_emptyB\014\n\n_skip_eval\"\315\001\n\nMapOptio" +
      "ns\022#\n\003kvs\030\001 \003(\0132\026.default.MapOptions.KV\022" +
      "\031\n\014ignore_empty\030\002 \001(\010H\000\210\001\001\022\026\n\tskip_eval\030" +
      "\013 \001(\010H\001\210\001\001\032H\n\002KV\022\037\n\003key\030\001 \001(\0132\022.default." +
      "FieldType\022!\n\005value\030\002 \001(\0132\022.default.Field" +
      "TypeB\017\n\r_ignore_emptyB\014\n\n_skip_eval\"\333\003\n\t" +
      "FieldType\022\017\n\005int32\030\001 \001(\005H\000\022\017\n\005int64\030\002 \001(" +
      "\003H\000\022\020\n\006sint32\030\003 \001(\021H\000\022\020\n\006sint64\030\004 \001(\022H\000\022" +
      "\022\n\010sfixed32\030\005 \001(\017H\000\022\022\n\010sfixed64\030\006 \001(\020H\000\022" +
      "\020\n\006uint32\030\007 \001(\rH\000\022\020\n\006uint64\030\010 \001(\004H\000\022\021\n\007f" +
      "ixed32\030\t \001(\007H\000\022\021\n\007fixed64\030\n \001(\006H\000\022\017\n\005flo" +
      "at\030\013 \001(\002H\000\022\020\n\006double\030\014 \001(\001H\000\022\016\n\004bool\030\r \001" +
      "(\010H\000\022\016\n\004enum\030\016 \001(\005H\000\022\020\n\006string\030\017 \001(\tH\000\022\017" +
      "\n\005bytes\030\020 \001(\014H\000\022\'\n\007message\030\021 \001(\0132\024.defau" +
      "lt.TypeMessageH\000\022\037\n\003any\030\025 \001(\0132\020.default." +
      "TypeAnyH\000\022-\n\010duration\030\026 \001(\0132\031.google.pro" +
      "tobuf.DurationH\000\022/\n\ttimestamp\030\027 \001(\0132\032.go" +
      "ogle.protobuf.TimestampH\000B\006\n\004Kind\")\n\013Typ" +
      "eMessage\022\021\n\004init\030\001 \001(\010H\000\210\001\001B\007\n\005_init\"-\n\007" +
      "TypeAny\022\025\n\010type_url\030\013 \001(\tH\000\210\001\001B\013\n\t_type_" +
      "url:E\n\005field\022\035.google.protobuf.FieldOpti" +
      "ons\030\224\364\003 \001(\0132\025.default.FieldOptions:E\n\005on" +
      "eof\022\035.google.protobuf.OneofOptions\030\225\364\003 \001" +
      "(\0132\025.default.OneOfOptionsBi\n\"io.github.y" +
      "u31.protoc.pb.pbdefaultB\tPBDefaultP\001Z6gi" +
      "thub.com/yu31/protoc-plugin-default/xgo/" +
      "pb/pbdefaultb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.DescriptorProtos.getDescriptor(),
          com.google.protobuf.DurationProto.getDescriptor(),
          com.google.protobuf.TimestampProto.getDescriptor(),
        });
    internal_static_default_OneOfOptions_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_default_OneOfOptions_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_default_OneOfOptions_descriptor,
        new java.lang.String[] { "Field", "Field", });
    internal_static_default_FieldOptions_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_default_FieldOptions_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_default_FieldOptions_descriptor,
        new java.lang.String[] { "Plain", "Repeated", "Map", "Kind", });
    internal_static_default_PlainOptions_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_default_PlainOptions_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_default_PlainOptions_descriptor,
        new java.lang.String[] { "Value", "IgnoreEmpty", "SkipEval", "IgnoreEmpty", "SkipEval", });
    internal_static_default_RepeatedOptions_descriptor =
      getDescriptor().getMessageTypes().get(3);
    internal_static_default_RepeatedOptions_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_default_RepeatedOptions_descriptor,
        new java.lang.String[] { "Items", "IgnoreEmpty", "SkipEval", "IgnoreEmpty", "SkipEval", });
    internal_static_default_MapOptions_descriptor =
      getDescriptor().getMessageTypes().get(4);
    internal_static_default_MapOptions_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_default_MapOptions_descriptor,
        new java.lang.String[] { "Kvs", "IgnoreEmpty", "SkipEval", "IgnoreEmpty", "SkipEval", });
    internal_static_default_MapOptions_KV_descriptor =
      internal_static_default_MapOptions_descriptor.getNestedTypes().get(0);
    internal_static_default_MapOptions_KV_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_default_MapOptions_KV_descriptor,
        new java.lang.String[] { "Key", "Value", });
    internal_static_default_FieldType_descriptor =
      getDescriptor().getMessageTypes().get(5);
    internal_static_default_FieldType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_default_FieldType_descriptor,
        new java.lang.String[] { "Int32", "Int64", "Sint32", "Sint64", "Sfixed32", "Sfixed64", "Uint32", "Uint64", "Fixed32", "Fixed64", "Float", "Double", "Bool", "Enum", "String", "Bytes", "Message", "Any", "Duration", "Timestamp", "Kind", });
    internal_static_default_TypeMessage_descriptor =
      getDescriptor().getMessageTypes().get(6);
    internal_static_default_TypeMessage_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_default_TypeMessage_descriptor,
        new java.lang.String[] { "Init", "Init", });
    internal_static_default_TypeAny_descriptor =
      getDescriptor().getMessageTypes().get(7);
    internal_static_default_TypeAny_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_default_TypeAny_descriptor,
        new java.lang.String[] { "TypeUrl", "TypeUrl", });
    field.internalInit(descriptor.getExtensions().get(0));
    oneof.internalInit(descriptor.getExtensions().get(1));
    com.google.protobuf.DescriptorProtos.getDescriptor();
    com.google.protobuf.DurationProto.getDescriptor();
    com.google.protobuf.TimestampProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}