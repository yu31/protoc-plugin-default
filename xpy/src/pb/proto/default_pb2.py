# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/default.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import descriptor_pb2 as google_dot_protobuf_dot_descriptor__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13proto/default.proto\x12\x07\x64\x65\x66\x61ult\x1a google/protobuf/descriptor.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\",\n\x0cOneOfOptions\x12\x12\n\x05\x66ield\x18\x01 \x01(\tH\x00\x88\x01\x01\x42\x08\n\x06_field\"\x90\x01\n\x0c\x46ieldOptions\x12&\n\x05plain\x18\x01 \x01(\x0b\x32\x15.default.PlainOptionsH\x00\x12,\n\x08repeated\x18\x02 \x01(\x0b\x32\x18.default.RepeatedOptionsH\x00\x12\"\n\x03map\x18\x03 \x01(\x0b\x32\x13.default.MapOptionsH\x00\x42\x06\n\x04Kind\"\x83\x01\n\x0cPlainOptions\x12!\n\x05value\x18\x01 \x01(\x0b\x32\x12.default.FieldType\x12\x19\n\x0cignore_empty\x18\x02 \x01(\x08H\x00\x88\x01\x01\x12\x16\n\tskip_eval\x18\x0b \x01(\x08H\x01\x88\x01\x01\x42\x0f\n\r_ignore_emptyB\x0c\n\n_skip_eval\"\x86\x01\n\x0fRepeatedOptions\x12!\n\x05items\x18\x01 \x03(\x0b\x32\x12.default.FieldType\x12\x19\n\x0cignore_empty\x18\x02 \x01(\x08H\x00\x88\x01\x01\x12\x16\n\tskip_eval\x18\x0b \x01(\x08H\x01\x88\x01\x01\x42\x0f\n\r_ignore_emptyB\x0c\n\n_skip_eval\"\xcd\x01\n\nMapOptions\x12#\n\x03kvs\x18\x01 \x03(\x0b\x32\x16.default.MapOptions.KV\x12\x19\n\x0cignore_empty\x18\x02 \x01(\x08H\x00\x88\x01\x01\x12\x16\n\tskip_eval\x18\x0b \x01(\x08H\x01\x88\x01\x01\x1aH\n\x02KV\x12\x1f\n\x03key\x18\x01 \x01(\x0b\x32\x12.default.FieldType\x12!\n\x05value\x18\x02 \x01(\x0b\x32\x12.default.FieldTypeB\x0f\n\r_ignore_emptyB\x0c\n\n_skip_eval\"\xdb\x03\n\tFieldType\x12\x0f\n\x05int32\x18\x01 \x01(\x05H\x00\x12\x0f\n\x05int64\x18\x02 \x01(\x03H\x00\x12\x10\n\x06sint32\x18\x03 \x01(\x11H\x00\x12\x10\n\x06sint64\x18\x04 \x01(\x12H\x00\x12\x12\n\x08sfixed32\x18\x05 \x01(\x0fH\x00\x12\x12\n\x08sfixed64\x18\x06 \x01(\x10H\x00\x12\x10\n\x06uint32\x18\x07 \x01(\rH\x00\x12\x10\n\x06uint64\x18\x08 \x01(\x04H\x00\x12\x11\n\x07\x66ixed32\x18\t \x01(\x07H\x00\x12\x11\n\x07\x66ixed64\x18\n \x01(\x06H\x00\x12\x0f\n\x05\x66loat\x18\x0b \x01(\x02H\x00\x12\x10\n\x06\x64ouble\x18\x0c \x01(\x01H\x00\x12\x0e\n\x04\x62ool\x18\r \x01(\x08H\x00\x12\x0e\n\x04\x65num\x18\x0e \x01(\x05H\x00\x12\x10\n\x06string\x18\x0f \x01(\tH\x00\x12\x0f\n\x05\x62ytes\x18\x10 \x01(\x0cH\x00\x12\'\n\x07message\x18\x11 \x01(\x0b\x32\x14.default.TypeMessageH\x00\x12\x1f\n\x03\x61ny\x18\x15 \x01(\x0b\x32\x10.default.TypeAnyH\x00\x12-\n\x08\x64uration\x18\x16 \x01(\x0b\x32\x19.google.protobuf.DurationH\x00\x12/\n\ttimestamp\x18\x17 \x01(\x0b\x32\x1a.google.protobuf.TimestampH\x00\x42\x06\n\x04Kind\")\n\x0bTypeMessage\x12\x11\n\x04init\x18\x01 \x01(\x08H\x00\x88\x01\x01\x42\x07\n\x05_init\"-\n\x07TypeAny\x12\x15\n\x08type_url\x18\x0b \x01(\tH\x00\x88\x01\x01\x42\x0b\n\t_type_url:E\n\x05\x66ield\x12\x1d.google.protobuf.FieldOptions\x18\xaf\xec\x03 \x01(\x0b\x32\x15.default.FieldOptions:E\n\x05oneof\x12\x1d.google.protobuf.OneofOptions\x18\xb0\xec\x03 \x01(\x0b\x32\x15.default.OneOfOptionsBi\n\"io.github.yu31.protoc.pb.pbdefaultB\tPBDefaultP\x01Z6github.com/yu31/protoc-plugin-default/xgo/pb/pbdefaultb\x06proto3')


FIELD_FIELD_NUMBER = 63023
field = DESCRIPTOR.extensions_by_name['field']
ONEOF_FIELD_NUMBER = 63024
oneof = DESCRIPTOR.extensions_by_name['oneof']

_ONEOFOPTIONS = DESCRIPTOR.message_types_by_name['OneOfOptions']
_FIELDOPTIONS = DESCRIPTOR.message_types_by_name['FieldOptions']
_PLAINOPTIONS = DESCRIPTOR.message_types_by_name['PlainOptions']
_REPEATEDOPTIONS = DESCRIPTOR.message_types_by_name['RepeatedOptions']
_MAPOPTIONS = DESCRIPTOR.message_types_by_name['MapOptions']
_MAPOPTIONS_KV = _MAPOPTIONS.nested_types_by_name['KV']
_FIELDTYPE = DESCRIPTOR.message_types_by_name['FieldType']
_TYPEMESSAGE = DESCRIPTOR.message_types_by_name['TypeMessage']
_TYPEANY = DESCRIPTOR.message_types_by_name['TypeAny']
OneOfOptions = _reflection.GeneratedProtocolMessageType('OneOfOptions', (_message.Message,), {
  'DESCRIPTOR' : _ONEOFOPTIONS,
  '__module__' : 'proto.default_pb2'
  # @@protoc_insertion_point(class_scope:default.OneOfOptions)
  })
_sym_db.RegisterMessage(OneOfOptions)

FieldOptions = _reflection.GeneratedProtocolMessageType('FieldOptions', (_message.Message,), {
  'DESCRIPTOR' : _FIELDOPTIONS,
  '__module__' : 'proto.default_pb2'
  # @@protoc_insertion_point(class_scope:default.FieldOptions)
  })
_sym_db.RegisterMessage(FieldOptions)

PlainOptions = _reflection.GeneratedProtocolMessageType('PlainOptions', (_message.Message,), {
  'DESCRIPTOR' : _PLAINOPTIONS,
  '__module__' : 'proto.default_pb2'
  # @@protoc_insertion_point(class_scope:default.PlainOptions)
  })
_sym_db.RegisterMessage(PlainOptions)

RepeatedOptions = _reflection.GeneratedProtocolMessageType('RepeatedOptions', (_message.Message,), {
  'DESCRIPTOR' : _REPEATEDOPTIONS,
  '__module__' : 'proto.default_pb2'
  # @@protoc_insertion_point(class_scope:default.RepeatedOptions)
  })
_sym_db.RegisterMessage(RepeatedOptions)

MapOptions = _reflection.GeneratedProtocolMessageType('MapOptions', (_message.Message,), {

  'KV' : _reflection.GeneratedProtocolMessageType('KV', (_message.Message,), {
    'DESCRIPTOR' : _MAPOPTIONS_KV,
    '__module__' : 'proto.default_pb2'
    # @@protoc_insertion_point(class_scope:default.MapOptions.KV)
    })
  ,
  'DESCRIPTOR' : _MAPOPTIONS,
  '__module__' : 'proto.default_pb2'
  # @@protoc_insertion_point(class_scope:default.MapOptions)
  })
_sym_db.RegisterMessage(MapOptions)
_sym_db.RegisterMessage(MapOptions.KV)

FieldType = _reflection.GeneratedProtocolMessageType('FieldType', (_message.Message,), {
  'DESCRIPTOR' : _FIELDTYPE,
  '__module__' : 'proto.default_pb2'
  # @@protoc_insertion_point(class_scope:default.FieldType)
  })
_sym_db.RegisterMessage(FieldType)

TypeMessage = _reflection.GeneratedProtocolMessageType('TypeMessage', (_message.Message,), {
  'DESCRIPTOR' : _TYPEMESSAGE,
  '__module__' : 'proto.default_pb2'
  # @@protoc_insertion_point(class_scope:default.TypeMessage)
  })
_sym_db.RegisterMessage(TypeMessage)

TypeAny = _reflection.GeneratedProtocolMessageType('TypeAny', (_message.Message,), {
  'DESCRIPTOR' : _TYPEANY,
  '__module__' : 'proto.default_pb2'
  # @@protoc_insertion_point(class_scope:default.TypeAny)
  })
_sym_db.RegisterMessage(TypeAny)

if _descriptor._USE_C_DESCRIPTORS == False:
  google_dot_protobuf_dot_descriptor__pb2.FieldOptions.RegisterExtension(field)
  google_dot_protobuf_dot_descriptor__pb2.OneofOptions.RegisterExtension(oneof)

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\"io.github.yu31.protoc.pb.pbdefaultB\tPBDefaultP\001Z6github.com/yu31/protoc-plugin-default/xgo/pb/pbdefault'
  _ONEOFOPTIONS._serialized_start=131
  _ONEOFOPTIONS._serialized_end=175
  _FIELDOPTIONS._serialized_start=178
  _FIELDOPTIONS._serialized_end=322
  _PLAINOPTIONS._serialized_start=325
  _PLAINOPTIONS._serialized_end=456
  _REPEATEDOPTIONS._serialized_start=459
  _REPEATEDOPTIONS._serialized_end=593
  _MAPOPTIONS._serialized_start=596
  _MAPOPTIONS._serialized_end=801
  _MAPOPTIONS_KV._serialized_start=698
  _MAPOPTIONS_KV._serialized_end=770
  _FIELDTYPE._serialized_start=804
  _FIELDTYPE._serialized_end=1279
  _TYPEMESSAGE._serialized_start=1281
  _TYPEMESSAGE._serialized_end=1322
  _TYPEANY._serialized_start=1324
  _TYPEANY._serialized_end=1369
# @@protoc_insertion_point(module_scope)
