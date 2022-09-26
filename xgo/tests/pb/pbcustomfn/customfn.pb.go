// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: tests/proto/customfn/customfn.proto

package pbcustomfn

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

type CustomMethodFuncName1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	F1 string `protobuf:"bytes,1,opt,name=f1,proto3" json:"f1,omitempty"`
}

func (x *CustomMethodFuncName1) Reset() {
	*x = CustomMethodFuncName1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_customfn_customfn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomMethodFuncName1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomMethodFuncName1) ProtoMessage() {}

func (x *CustomMethodFuncName1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_customfn_customfn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomMethodFuncName1.ProtoReflect.Descriptor instead.
func (*CustomMethodFuncName1) Descriptor() ([]byte, []int) {
	return file_tests_proto_customfn_customfn_proto_rawDescGZIP(), []int{0}
}

func (x *CustomMethodFuncName1) GetF1() string {
	if x != nil {
		return x.F1
	}
	return ""
}

type CustomMethodFuncName2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	F1 string                 `protobuf:"bytes,1,opt,name=f1,proto3" json:"f1,omitempty"`
	M1 *CustomMethodFuncName1 `protobuf:"bytes,11,opt,name=m1,proto3" json:"m1,omitempty"`
}

func (x *CustomMethodFuncName2) Reset() {
	*x = CustomMethodFuncName2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_customfn_customfn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomMethodFuncName2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomMethodFuncName2) ProtoMessage() {}

func (x *CustomMethodFuncName2) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_customfn_customfn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomMethodFuncName2.ProtoReflect.Descriptor instead.
func (*CustomMethodFuncName2) Descriptor() ([]byte, []int) {
	return file_tests_proto_customfn_customfn_proto_rawDescGZIP(), []int{1}
}

func (x *CustomMethodFuncName2) GetF1() string {
	if x != nil {
		return x.F1
	}
	return ""
}

func (x *CustomMethodFuncName2) GetM1() *CustomMethodFuncName1 {
	if x != nil {
		return x.M1
	}
	return nil
}

var File_tests_proto_customfn_customfn_proto protoreflect.FileDescriptor

var file_tests_proto_customfn_customfn_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x66, 0x6e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x66, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x66, 0x6e, 0x1a,
	0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x15, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x31, 0x12, 0x1d,
	0x0a, 0x02, 0x66, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xa2, 0xa1, 0x1f, 0x09,
	0x0a, 0x07, 0x0a, 0x05, 0x7a, 0x03, 0x31, 0x32, 0x33, 0x52, 0x02, 0x66, 0x31, 0x22, 0x76, 0x0a,
	0x15, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x46, 0x75, 0x6e,
	0x63, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x12, 0x1d, 0x0a, 0x02, 0x66, 0x31, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0d, 0xa2, 0xa1, 0x1f, 0x09, 0x0a, 0x07, 0x0a, 0x05, 0x7a, 0x03, 0x31, 0x32,
	0x33, 0x52, 0x02, 0x66, 0x31, 0x12, 0x3e, 0x0a, 0x02, 0x6d, 0x31, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x66, 0x6e, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x31, 0x42, 0x0d, 0xa2, 0xa1, 0x1f, 0x09, 0x0a, 0x07, 0x0a, 0x05, 0x8a, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x02, 0x6d, 0x31, 0x42, 0x19, 0x5a, 0x17, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x66, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_proto_customfn_customfn_proto_rawDescOnce sync.Once
	file_tests_proto_customfn_customfn_proto_rawDescData = file_tests_proto_customfn_customfn_proto_rawDesc
)

func file_tests_proto_customfn_customfn_proto_rawDescGZIP() []byte {
	file_tests_proto_customfn_customfn_proto_rawDescOnce.Do(func() {
		file_tests_proto_customfn_customfn_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_customfn_customfn_proto_rawDescData)
	})
	return file_tests_proto_customfn_customfn_proto_rawDescData
}

var file_tests_proto_customfn_customfn_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tests_proto_customfn_customfn_proto_goTypes = []interface{}{
	(*CustomMethodFuncName1)(nil), // 0: customfn.CustomMethodFuncName1
	(*CustomMethodFuncName2)(nil), // 1: customfn.CustomMethodFuncName2
}
var file_tests_proto_customfn_customfn_proto_depIdxs = []int32{
	0, // 0: customfn.CustomMethodFuncName2.m1:type_name -> customfn.CustomMethodFuncName1
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tests_proto_customfn_customfn_proto_init() }
func file_tests_proto_customfn_customfn_proto_init() {
	if File_tests_proto_customfn_customfn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_customfn_customfn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomMethodFuncName1); i {
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
		file_tests_proto_customfn_customfn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomMethodFuncName2); i {
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
			RawDescriptor: file_tests_proto_customfn_customfn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_customfn_customfn_proto_goTypes,
		DependencyIndexes: file_tests_proto_customfn_customfn_proto_depIdxs,
		MessageInfos:      file_tests_proto_customfn_customfn_proto_msgTypes,
	}.Build()
	File_tests_proto_customfn_customfn_proto = out.File
	file_tests_proto_customfn_customfn_proto_rawDesc = nil
	file_tests_proto_customfn_customfn_proto_goTypes = nil
	file_tests_proto_customfn_customfn_proto_depIdxs = nil
}