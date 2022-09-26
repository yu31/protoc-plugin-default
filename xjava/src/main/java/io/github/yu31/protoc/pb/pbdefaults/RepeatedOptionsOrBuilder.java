// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: defaults.proto

package io.github.yu31.protoc.pb.pbdefaults;

public interface RepeatedOptionsOrBuilder extends
    // @@protoc_insertion_point(interface_extends:defaults.RepeatedOptions)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * items declares the default value for field.
   * </pre>
   *
   * <code>repeated .defaults.FieldType items = 1;</code>
   */
  java.util.List<io.github.yu31.protoc.pb.pbdefaults.FieldType> 
      getItemsList();
  /**
   * <pre>
   * items declares the default value for field.
   * </pre>
   *
   * <code>repeated .defaults.FieldType items = 1;</code>
   */
  io.github.yu31.protoc.pb.pbdefaults.FieldType getItems(int index);
  /**
   * <pre>
   * items declares the default value for field.
   * </pre>
   *
   * <code>repeated .defaults.FieldType items = 1;</code>
   */
  int getItemsCount();
  /**
   * <pre>
   * items declares the default value for field.
   * </pre>
   *
   * <code>repeated .defaults.FieldType items = 1;</code>
   */
  java.util.List<? extends io.github.yu31.protoc.pb.pbdefaults.FieldTypeOrBuilder> 
      getItemsOrBuilderList();
  /**
   * <pre>
   * items declares the default value for field.
   * </pre>
   *
   * <code>repeated .defaults.FieldType items = 1;</code>
   */
  io.github.yu31.protoc.pb.pbdefaults.FieldTypeOrBuilder getItemsOrBuilder(
      int index);

  /**
   * <pre>
   * By default, The default value will be applied if the length of field is zero.
   * If `ignore_empty` is true, Only apply the default value if the field is a nil pointer.
   * Only effective for language Go.
   * </pre>
   *
   * <code>optional bool ignore_empty = 2;</code>
   * @return Whether the ignoreEmpty field is set.
   */
  boolean hasIgnoreEmpty();
  /**
   * <pre>
   * By default, The default value will be applied if the length of field is zero.
   * If `ignore_empty` is true, Only apply the default value if the field is a nil pointer.
   * Only effective for language Go.
   * </pre>
   *
   * <code>optional bool ignore_empty = 2;</code>
   * @return The ignoreEmpty.
   */
  boolean getIgnoreEmpty();

  /**
   * <pre>
   * By default, We will try to eval call the method of defaults for type of repeated items are message and any.
   * If `skip_eval` is true, the method of defaults will not be called.
   * Only effective for language Go.
   * </pre>
   *
   * <code>optional bool skip_eval = 11;</code>
   * @return Whether the skipEval field is set.
   */
  boolean hasSkipEval();
  /**
   * <pre>
   * By default, We will try to eval call the method of defaults for type of repeated items are message and any.
   * If `skip_eval` is true, the method of defaults will not be called.
   * Only effective for language Go.
   * </pre>
   *
   * <code>optional bool skip_eval = 11;</code>
   * @return The skipEval.
   */
  boolean getSkipEval();
}