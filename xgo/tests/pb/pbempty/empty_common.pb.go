// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: tests/proto/cases/empty/empty_common.proto

package pbempty

import (
	_ "github.com/yu31/protoc-plugin-defaults/xgo/pb/pbdefaults"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EnumCommon1 used tests cases in this file.
type EnumCommon1 int32

const (
	EnumCommon1_Zero  EnumCommon1 = 0
	EnumCommon1_One   EnumCommon1 = 2
	EnumCommon1_Two   EnumCommon1 = 3
	EnumCommon1_Three EnumCommon1 = 5
	EnumCommon1_Four  EnumCommon1 = 6
	EnumCommon1_Five  EnumCommon1 = 7
	EnumCommon1_Six   EnumCommon1 = 9
	EnumCommon1_Seven EnumCommon1 = 11
	EnumCommon1_Eight EnumCommon1 = 12
	EnumCommon1_Nine  EnumCommon1 = 15
	EnumCommon1_Ten   EnumCommon1 = 17
)

// Enum value maps for EnumCommon1.
var (
	EnumCommon1_name = map[int32]string{
		0:  "Zero",
		2:  "One",
		3:  "Two",
		5:  "Three",
		6:  "Four",
		7:  "Five",
		9:  "Six",
		11: "Seven",
		12: "Eight",
		15: "Nine",
		17: "Ten",
	}
	EnumCommon1_value = map[string]int32{
		"Zero":  0,
		"One":   2,
		"Two":   3,
		"Three": 5,
		"Four":  6,
		"Five":  7,
		"Six":   9,
		"Seven": 11,
		"Eight": 12,
		"Nine":  15,
		"Ten":   17,
	}
)

func (x EnumCommon1) Enum() *EnumCommon1 {
	p := new(EnumCommon1)
	*p = x
	return p
}

func (x EnumCommon1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumCommon1) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_proto_cases_empty_empty_common_proto_enumTypes[0].Descriptor()
}

func (EnumCommon1) Type() protoreflect.EnumType {
	return &file_tests_proto_cases_empty_empty_common_proto_enumTypes[0]
}

func (x EnumCommon1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumCommon1.Descriptor instead.
func (EnumCommon1) EnumDescriptor() ([]byte, []int) {
	return file_tests_proto_cases_empty_empty_common_proto_rawDescGZIP(), []int{0}
}

// MessageCommon1 used tests cases in this file.
type MessageCommon1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1 string `protobuf:"bytes,1,opt,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty"`
	FString2 string `protobuf:"bytes,2,opt,name=f_string2,json=fString2,proto3" json:"f_string2,omitempty"`
	FString3 string `protobuf:"bytes,3,opt,name=f_string3,json=fString3,proto3" json:"f_string3,omitempty"`
	FString4 string `protobuf:"bytes,4,opt,name=f_string4,json=fString4,proto3" json:"f_string4,omitempty"`
}

func (x *MessageCommon1) Reset() {
	*x = MessageCommon1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_empty_empty_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageCommon1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCommon1) ProtoMessage() {}

func (x *MessageCommon1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_empty_empty_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageCommon1.ProtoReflect.Descriptor instead.
func (*MessageCommon1) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_empty_empty_common_proto_rawDescGZIP(), []int{0}
}

func (x *MessageCommon1) GetFString1() string {
	if x != nil {
		return x.FString1
	}
	return ""
}

func (x *MessageCommon1) GetFString2() string {
	if x != nil {
		return x.FString2
	}
	return ""
}

func (x *MessageCommon1) GetFString3() string {
	if x != nil {
		return x.FString3
	}
	return ""
}

func (x *MessageCommon1) GetFString4() string {
	if x != nil {
		return x.FString4
	}
	return ""
}

type MessageCommon1_Embed1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1 string `protobuf:"bytes,1,opt,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty"`
	FString2 string `protobuf:"bytes,2,opt,name=f_string2,json=fString2,proto3" json:"f_string2,omitempty"`
	FString3 string `protobuf:"bytes,3,opt,name=f_string3,json=fString3,proto3" json:"f_string3,omitempty"`
	FString4 string `protobuf:"bytes,4,opt,name=f_string4,json=fString4,proto3" json:"f_string4,omitempty"`
}

func (x *MessageCommon1_Embed1) Reset() {
	*x = MessageCommon1_Embed1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_empty_empty_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageCommon1_Embed1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCommon1_Embed1) ProtoMessage() {}

func (x *MessageCommon1_Embed1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_empty_empty_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageCommon1_Embed1.ProtoReflect.Descriptor instead.
func (*MessageCommon1_Embed1) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_empty_empty_common_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MessageCommon1_Embed1) GetFString1() string {
	if x != nil {
		return x.FString1
	}
	return ""
}

func (x *MessageCommon1_Embed1) GetFString2() string {
	if x != nil {
		return x.FString2
	}
	return ""
}

func (x *MessageCommon1_Embed1) GetFString3() string {
	if x != nil {
		return x.FString3
	}
	return ""
}

func (x *MessageCommon1_Embed1) GetFString4() string {
	if x != nil {
		return x.FString4
	}
	return ""
}

type MessageCommon1_Embed1_Embed2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1 string `protobuf:"bytes,1,opt,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty"`
	FString2 string `protobuf:"bytes,2,opt,name=f_string2,json=fString2,proto3" json:"f_string2,omitempty"`
	FString3 string `protobuf:"bytes,3,opt,name=f_string3,json=fString3,proto3" json:"f_string3,omitempty"`
	FString4 string `protobuf:"bytes,4,opt,name=f_string4,json=fString4,proto3" json:"f_string4,omitempty"`
}

func (x *MessageCommon1_Embed1_Embed2) Reset() {
	*x = MessageCommon1_Embed1_Embed2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_empty_empty_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageCommon1_Embed1_Embed2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCommon1_Embed1_Embed2) ProtoMessage() {}

func (x *MessageCommon1_Embed1_Embed2) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_empty_empty_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageCommon1_Embed1_Embed2.ProtoReflect.Descriptor instead.
func (*MessageCommon1_Embed1_Embed2) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_empty_empty_common_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *MessageCommon1_Embed1_Embed2) GetFString1() string {
	if x != nil {
		return x.FString1
	}
	return ""
}

func (x *MessageCommon1_Embed1_Embed2) GetFString2() string {
	if x != nil {
		return x.FString2
	}
	return ""
}

func (x *MessageCommon1_Embed1_Embed2) GetFString3() string {
	if x != nil {
		return x.FString3
	}
	return ""
}

func (x *MessageCommon1_Embed1_Embed2) GetFString4() string {
	if x != nil {
		return x.FString4
	}
	return ""
}

var File_tests_proto_cases_empty_empty_common_proto protoreflect.FileDescriptor

var file_tests_proto_cases_empty_empty_common_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xca, 0x03, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x31, 0x12, 0x25, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xa2, 0xa1, 0x1f, 0x04, 0x0a, 0x02, 0x0a, 0x00, 0x52,
	0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x12, 0x23, 0x0a, 0x09, 0x66, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xa2, 0xa1,
	0x1f, 0x02, 0x0a, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x12, 0x21,
	0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0xa2, 0xa1, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x33, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x34, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x34, 0x1a, 0xab,
	0x02, 0x0a, 0x06, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x31, 0x12, 0x25, 0x0a, 0x09, 0x66, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xa2, 0xa1,
	0x1f, 0x04, 0x0a, 0x02, 0x0a, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31,
	0x12, 0x23, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xa2, 0xa1, 0x1f, 0x02, 0x0a, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x32, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xa2, 0xa1, 0x1f, 0x00, 0x52, 0x08,
	0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x33, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x34, 0x1a, 0x94, 0x01, 0x0a, 0x06, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x32,
	0x12, 0x25, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xa2, 0xa1, 0x1f, 0x04, 0x0a, 0x02, 0x0a, 0x00, 0x52, 0x08, 0x66,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x12, 0x23, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xa2, 0xa1, 0x1f, 0x02,
	0x0a, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x12, 0x21, 0x0a, 0x09,
	0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0xa2, 0xa1, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x33, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x34, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x34, 0x2a, 0x7a, 0x0a, 0x0b,
	0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x31, 0x12, 0x08, 0x0a, 0x04, 0x5a,
	0x65, 0x72, 0x6f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x6e, 0x65, 0x10, 0x02, 0x12, 0x07,
	0x0a, 0x03, 0x54, 0x77, 0x6f, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x68, 0x72, 0x65, 0x65,
	0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x6f, 0x75, 0x72, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x69, 0x76, 0x65, 0x10, 0x07, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x69, 0x78, 0x10, 0x09, 0x12,
	0x09, 0x0a, 0x05, 0x53, 0x65, 0x76, 0x65, 0x6e, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x69,
	0x67, 0x68, 0x74, 0x10, 0x0c, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x69, 0x6e, 0x65, 0x10, 0x0f, 0x12,
	0x07, 0x0a, 0x03, 0x54, 0x65, 0x6e, 0x10, 0x11, 0x42, 0x16, 0x5a, 0x14, 0x78, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_proto_cases_empty_empty_common_proto_rawDescOnce sync.Once
	file_tests_proto_cases_empty_empty_common_proto_rawDescData = file_tests_proto_cases_empty_empty_common_proto_rawDesc
)

func file_tests_proto_cases_empty_empty_common_proto_rawDescGZIP() []byte {
	file_tests_proto_cases_empty_empty_common_proto_rawDescOnce.Do(func() {
		file_tests_proto_cases_empty_empty_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_cases_empty_empty_common_proto_rawDescData)
	})
	return file_tests_proto_cases_empty_empty_common_proto_rawDescData
}

var file_tests_proto_cases_empty_empty_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tests_proto_cases_empty_empty_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tests_proto_cases_empty_empty_common_proto_goTypes = []interface{}{
	(EnumCommon1)(0),                     // 0: type_common.EnumCommon1
	(*MessageCommon1)(nil),               // 1: type_common.MessageCommon1
	(*MessageCommon1_Embed1)(nil),        // 2: type_common.MessageCommon1.Embed1
	(*MessageCommon1_Embed1_Embed2)(nil), // 3: type_common.MessageCommon1.Embed1.Embed2
}
var file_tests_proto_cases_empty_empty_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_proto_cases_empty_empty_common_proto_init() }
func file_tests_proto_cases_empty_empty_common_proto_init() {
	if File_tests_proto_cases_empty_empty_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_cases_empty_empty_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageCommon1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_proto_cases_empty_empty_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageCommon1_Embed1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_proto_cases_empty_empty_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageCommon1_Embed1_Embed2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_proto_cases_empty_empty_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_cases_empty_empty_common_proto_goTypes,
		DependencyIndexes: file_tests_proto_cases_empty_empty_common_proto_depIdxs,
		EnumInfos:         file_tests_proto_cases_empty_empty_common_proto_enumTypes,
		MessageInfos:      file_tests_proto_cases_empty_empty_common_proto_msgTypes,
	}.Build()
	File_tests_proto_cases_empty_empty_common_proto = out.File
	file_tests_proto_cases_empty_empty_common_proto_rawDesc = nil
	file_tests_proto_cases_empty_empty_common_proto_goTypes = nil
	file_tests_proto_cases_empty_empty_common_proto_depIdxs = nil
}
