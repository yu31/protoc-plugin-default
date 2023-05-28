// Code generated by protoc-gen-godefault. DO NOT EDIT.
// versions:
// 		protoc-gen-godefault 0.0.1
// source: tests/proto/cases/types/type_map.proto

package pbtypes

import (
	_ "github.com/yu31/protoc-plugin-default/xgo/pb/pbdefault"
	pbexternal "github.com/yu31/protoc-plugin-default/xgo/tests/pb/pbexternal"
	proto "google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// Set default value for message MessageMap1 that in file tests/proto/cases/types/type_map.proto
func (x *MessageMap1) SetDefault() {
	if x == nil {
		return
	}
	if x.FString1 == "" {
		x.FString1 = "a"
	}
	if x.FString2 == "" {
		x.FString2 = "b"
	}
	if x.FString3 == "" {
		x.FString3 = "c"
	}
}

// Set default value for message Embed1 that in file tests/proto/cases/types/type_map.proto
func (x *MessageMap1_Embed1) SetDefault() {
	if x == nil {
		return
	}
	if x.FString1 == "" {
		x.FString1 = "x"
	}
	if x.FString2 == "" {
		x.FString2 = "y"
	}
	if x.FString3 == "" {
		x.FString3 = "z"
	}
}

// Set default value for message Embed2 that in file tests/proto/cases/types/type_map.proto
func (x *MessageMap1_Embed1_Embed2) SetDefault() {
	if x == nil {
		return
	}
	if x.FString1 == "" {
		x.FString1 = "1"
	}
	if x.FString2 == "" {
		x.FString2 = "2"
	}
	if x.FString3 == "" {
		x.FString3 = "3"
	}
}

// Set default value for message TypeMap1 that in file tests/proto/cases/types/type_map.proto
func (x *TypeMap1) SetDefault() {
	if x == nil {
		return
	}
	if len(x.FString1) == 0 {
		x.FString1 = map[string]string{
			"ts1":       "ts1",
			"":          "",
			"\"\"":      "\"\"",
			"\"ts4\"":   "\"ts4\"",
			"\"ts5\"":   "\"ts5\"",
			"\"ts\"6\"": "\"ts\"6\"",
			"\"ts\"7\"": "\"ts\"7\"",
			"[ts8]":     "[ts8]",
			"{ts9}":     "{ts9}",
			"s10":       "",
		}
	}
	if len(x.FInt32) == 0 {
		x.FInt32 = map[string]int32{
			"11": 11,
			"12": 12,
			"13": 0,
			"":   0,
		}
	}
	if len(x.FInt64) == 0 {
		x.FInt64 = map[string]int64{
			"21": 21,
			"22": 22,
			"23": 0,
			"":   0,
		}
	}
	if len(x.FUint32) == 0 {
		x.FUint32 = map[string]uint32{
			"31": 31,
			"32": 32,
			"33": 0,
			"":   0,
		}
	}
	if len(x.FUint64) == 0 {
		x.FUint64 = map[string]uint64{
			"41": 41,
			"42": 42,
			"43": 0,
			"":   0,
		}
	}
	if len(x.FSint32) == 0 {
		x.FSint32 = map[string]int32{
			"51": 51,
			"52": 52,
			"53": 0,
			"":   0,
		}
	}
	if len(x.FSint64) == 0 {
		x.FSint64 = map[string]int64{
			"61": 61,
			"62": 62,
			"63": 0,
			"":   0,
		}
	}
	if len(x.FSfixed32) == 0 {
		x.FSfixed32 = map[string]int32{
			"71": 71,
			"72": 72,
			"73": 0,
			"":   0,
		}
	}
	if len(x.FSfixed64) == 0 {
		x.FSfixed64 = map[string]int64{
			"81": 81,
			"82": 82,
			"83": 0,
			"":   0,
		}
	}
	if len(x.FFixed32) == 0 {
		x.FFixed32 = map[string]uint32{
			"91": 91,
			"92": 92,
			"93": 0,
			"":   0,
		}
	}
	if len(x.FFixed64) == 0 {
		x.FFixed64 = map[string]uint64{
			"101": 101,
			"102": 102,
			"103": 0,
			"":    0,
		}
	}
	if len(x.FFloat) == 0 {
		x.FFloat = map[string]float32{
			"111.111": 111.11100006103516,
			"112.112": 112.11199951171875,
			"113.113": 0,
			"":        0,
		}
	}
	if len(x.FDouble) == 0 {
		x.FDouble = map[string]float64{
			"121.121": 121.121,
			"122.122": 122.122,
			"123.123": 0,
			"":        0,
		}
	}
	if len(x.FBool1) == 0 {
		x.FBool1 = map[string]bool{
			"t1": true,
			"t2": false,
			"t3": false,
			"":   false,
		}
	}
	if len(x.FBytes1) == 0 {
		x.FBytes1 = map[string][]byte{
			"b1": []byte{97, 98, 99},
			"b2": []byte{49, 50, 51},
			"b3": []byte{},
			"":   []byte{},
		}
	}
	if len(x.FEnum1) == 0 {
		x.FEnum1 = map[string]EnumMap1{
			"e1": 1,
			"e2": 3,
			"e3": 0,
			"":   0,
		}
	}
	if len(x.FEnum2) == 0 {
		x.FEnum2 = map[string]EnumMap1{
			"e1": 2,
			"e2": 4,
			"e3": 0,
			"":   0,
		}
	}
	if len(x.FEnum3) == 0 {
		x.FEnum3 = map[string]pbexternal.Month1{
			"e1": 1,
			"e2": 3,
			"e3": 0,
			"":   0,
		}
	}
	if len(x.FEnum4) == 0 {
		x.FEnum4 = map[string]pbexternal.Month2{
			"e1": 2,
			"e2": 4,
			"e3": 0,
			"":   0,
		}
	}
	if len(x.FEnum5) == 0 {
		x.FEnum5 = map[string]pbexternal.EnumWeek1_Week{
			"e1": 1,
			"e2": 3,
			"e3": 0,
			"":   0,
		}
	}
	if len(x.FEnum6) == 0 {
		x.FEnum6 = map[string]pbexternal.EnumWeek2_Embed1_Week{
			"e1": 2,
			"e2": 4,
			"e3": 0,
			"":   0,
		}
	}
	if len(x.FEnum7) == 0 {
		x.FEnum7 = map[string]EnumCommon1{
			"e1": 1,
			"e2": 3,
			"e3": 0,
			"":   0,
		}
	}
	if len(x.FDuration1) == 0 {
		x.FDuration1 = map[string]*durationpb.Duration{
			"d1": &durationpb.Duration{Seconds: 100, Nanos: 200},
			"d2": &durationpb.Duration{Seconds: 101, Nanos: 201},
			"d3": &durationpb.Duration{Seconds: 0, Nanos: 0},
			"d4": nil,
			"":   nil,
		}
	}
	if len(x.FTimestamp1) == 0 {
		x.FTimestamp1 = map[string]*timestamppb.Timestamp{
			"t1": &timestamppb.Timestamp{Seconds: 100, Nanos: 200},
			"t2": &timestamppb.Timestamp{Seconds: 101, Nanos: 201},
			"t3": &timestamppb.Timestamp{Seconds: 0, Nanos: 0},
			"t4": nil,
			"":   nil,
		}
	}
	if len(x.FMessage1) == 0 {
		x.FMessage1 = map[string]*MessageMap1{
			"m1": &MessageMap1{},
			"m2": nil,
			"m3": nil,
			"m4": nil,
			"":   nil,
		}
	}
	for _, mv := range x.FMessage1 {
		if mv != nil {
			mv.SetDefault()
		}
	}
	if len(x.FMessage2) == 0 {
		x.FMessage2 = map[string]*MessageMap1_Embed1{
			"m1": &MessageMap1_Embed1{},
			"m2": nil,
			"m3": nil,
			"m4": nil,
			"":   nil,
		}
	}
	if len(x.FMessage3) == 0 {
		x.FMessage3 = map[string]*MessageMap1_Embed1_Embed2{
			"m1": &MessageMap1_Embed1_Embed2{},
			"m2": nil,
			"m3": nil,
			"m4": nil,
			"":   nil,
		}
	}
	for _, mv := range x.FMessage3 {
		if mv != nil {
			mv.SetDefault()
		}
	}
	if len(x.FMessage4) == 0 {
		x.FMessage4 = map[string]*pbexternal.External1{
			"m1": &pbexternal.External1{},
			"m2": nil,
			"m3": nil,
			"m4": nil,
			"":   nil,
		}
	}
	if len(x.FMessage5) == 0 {
		x.FMessage5 = map[string]*pbexternal.External2_Embed1{
			"m1": &pbexternal.External2_Embed1{},
			"m2": nil,
			"m3": nil,
			"m4": nil,
			"":   nil,
		}
	}
	for _, mv := range x.FMessage5 {
		if mv != nil {
			mv.SetDefault()
		}
	}
	if len(x.FMessage6) == 0 {
		x.FMessage6 = map[string]*MessageCommon1{
			"m1": &MessageCommon1{},
			"m2": nil,
			"m3": nil,
			"m4": nil,
			"":   nil,
		}
	}
	for _, mv := range x.FMessage6 {
		if mv != nil {
			mv.SetDefault()
		}
	}
	if len(x.FAny1) == 0 {
		x.FAny1 = map[string]*anypb.Any{
			"a1": x.xxx_protoc_plugin_default_mustNewAny(&MessageMap1{}, false),
			"a2": x.xxx_protoc_plugin_default_mustNewAny(&MessageMap1_Embed1{}, false),
			"a3": x.xxx_protoc_plugin_default_mustNewAny(&MessageMap1_Embed1_Embed2{}, false),
			"a4": x.xxx_protoc_plugin_default_mustNewAny(&pbexternal.External1{}, false),
			"a5": x.xxx_protoc_plugin_default_mustNewAny(&pbexternal.External2_Embed1{}, false),
			"a6": x.xxx_protoc_plugin_default_mustNewAny(&timestamppb.Timestamp{}, false),
			"a7": x.xxx_protoc_plugin_default_mustNewAny(&anypb.Any{}, false),
			"a8": x.xxx_protoc_plugin_default_mustNewAny(&MessageCommon1{}, false),
			"":   nil,
		}
	}
}

// xxx_protoc_plugin_default_mustNewAny used to create instance of type any
func (*TypeMap1) xxx_protoc_plugin_default_mustNewAny(src proto.Message, skipEval bool) *anypb.Any {
	if !skipEval {
		if dt, ok := interface{}(src).(interface{ SetDefault() }); ok {
			dt.SetDefault()
		}
	}
	v, err := anypb.New(src)
	if err != nil {
		panic(err)
	}
	return v
}

// Set default value for message TypeMap2 that in file tests/proto/cases/types/type_map.proto
func (x *TypeMap2) SetDefault() {
	if x == nil {
		return
	}
	if x.FString1 == nil {
		x.FString1 = map[string]string{
			"ts1":       "ts1",
			"":          "",
			"\"\"":      "\"\"",
			"\"ts4\"":   "\"ts4\"",
			"\"ts5\"":   "\"ts5\"",
			"\"ts\"6\"": "\"ts\"6\"",
			"\"ts\"7\"": "\"ts\"7\"",
			"[ts8]":     "[ts8]",
			"{ts9}":     "{ts9}",
			"s10":       "",
		}
	}
	if x.FInt32 == nil {
		x.FInt32 = map[string]int32{
			"11": 11,
			"12": 12,
			"13": 0,
			"":   0,
		}
	}
	if x.FInt64 == nil {
		x.FInt64 = map[string]int64{
			"21": 21,
			"22": 22,
			"23": 0,
			"":   0,
		}
	}
	if x.FUint32 == nil {
		x.FUint32 = map[string]uint32{
			"31": 31,
			"32": 32,
			"33": 0,
			"":   0,
		}
	}
	if x.FUint64 == nil {
		x.FUint64 = map[string]uint64{
			"41": 41,
			"42": 42,
			"43": 0,
			"":   0,
		}
	}
	if x.FSint32 == nil {
		x.FSint32 = map[string]int32{
			"51": 51,
			"52": 52,
			"53": 0,
			"":   0,
		}
	}
	if x.FSint64 == nil {
		x.FSint64 = map[string]int64{
			"61": 61,
			"62": 62,
			"63": 0,
			"":   0,
		}
	}
	if x.FSfixed32 == nil {
		x.FSfixed32 = map[string]int32{
			"71": 71,
			"72": 72,
			"73": 0,
			"":   0,
		}
	}
	if x.FSfixed64 == nil {
		x.FSfixed64 = map[string]int64{
			"81": 81,
			"82": 82,
			"83": 0,
			"":   0,
		}
	}
	if x.FFixed32 == nil {
		x.FFixed32 = map[string]uint32{
			"91": 91,
			"92": 92,
			"93": 0,
			"":   0,
		}
	}
	if x.FFixed64 == nil {
		x.FFixed64 = map[string]uint64{
			"101": 101,
			"102": 102,
			"103": 0,
			"":    0,
		}
	}
	if x.FFloat == nil {
		x.FFloat = map[string]float32{
			"111.111": 111.11100006103516,
			"112.112": 112.11199951171875,
			"113.113": 0,
			"":        0,
		}
	}
	if x.FDouble == nil {
		x.FDouble = map[string]float64{
			"121.121": 121.121,
			"122.122": 122.122,
			"123.123": 0,
			"":        0,
		}
	}
	if x.FBool1 == nil {
		x.FBool1 = map[string]bool{
			"t1": true,
			"t2": false,
			"t3": false,
			"":   false,
		}
	}
	if x.FBytes1 == nil {
		x.FBytes1 = map[string][]byte{
			"b1": []byte{97, 98, 99},
			"b2": []byte{49, 50, 51},
			"b3": []byte{},
			"":   []byte{},
		}
	}
	if x.FEnum1 == nil {
		x.FEnum1 = map[string]EnumMap1{
			"e1": 1,
			"e2": 3,
			"e3": 0,
			"":   0,
		}
	}
	if x.FEnum2 == nil {
		x.FEnum2 = map[string]EnumMap1{
			"e1": 2,
			"e2": 4,
			"e3": 0,
			"":   0,
		}
	}
	if x.FEnum3 == nil {
		x.FEnum3 = map[string]pbexternal.Month1{
			"e1": 1,
			"e2": 3,
			"e3": 0,
			"":   0,
		}
	}
	if x.FEnum4 == nil {
		x.FEnum4 = map[string]pbexternal.Month2{
			"e1": 2,
			"e2": 4,
			"e3": 0,
			"":   0,
		}
	}
	if x.FEnum5 == nil {
		x.FEnum5 = map[string]pbexternal.EnumWeek1_Week{
			"e1": 1,
			"e2": 3,
			"e3": 0,
			"":   0,
		}
	}
	if x.FEnum6 == nil {
		x.FEnum6 = map[string]pbexternal.EnumWeek2_Embed1_Week{
			"e1": 2,
			"e2": 4,
			"e3": 0,
			"":   0,
		}
	}
	if x.FEnum7 == nil {
		x.FEnum7 = map[string]EnumCommon1{
			"e1": 1,
			"e2": 3,
			"e3": 0,
			"":   0,
		}
	}
	if x.FDuration1 == nil {
		x.FDuration1 = map[string]*durationpb.Duration{
			"d1": &durationpb.Duration{Seconds: 100, Nanos: 200},
			"d2": &durationpb.Duration{Seconds: 101, Nanos: 201},
			"d3": &durationpb.Duration{Seconds: 0, Nanos: 0},
			"d4": nil,
			"":   nil,
		}
	}
	if x.FTimestamp1 == nil {
		x.FTimestamp1 = map[string]*timestamppb.Timestamp{
			"t1": &timestamppb.Timestamp{Seconds: 100, Nanos: 200},
			"t2": &timestamppb.Timestamp{Seconds: 101, Nanos: 201},
			"t3": &timestamppb.Timestamp{Seconds: 0, Nanos: 0},
			"t4": nil,
			"":   nil,
		}
	}
	if x.FMessage1 == nil {
		x.FMessage1 = map[string]*MessageMap1{
			"m1": &MessageMap1{},
			"m2": nil,
			"m3": nil,
			"m4": nil,
			"":   nil,
		}
	}
	for _, mv := range x.FMessage1 {
		if mv != nil {
			mv.SetDefault()
		}
	}
	if x.FMessage2 == nil {
		x.FMessage2 = map[string]*MessageMap1_Embed1{
			"m1": &MessageMap1_Embed1{},
			"m2": nil,
			"m3": nil,
			"m4": nil,
			"":   nil,
		}
	}
	if x.FMessage3 == nil {
		x.FMessage3 = map[string]*MessageMap1_Embed1_Embed2{
			"m1": &MessageMap1_Embed1_Embed2{},
			"m2": nil,
			"m3": nil,
			"m4": nil,
			"":   nil,
		}
	}
	for _, mv := range x.FMessage3 {
		if mv != nil {
			mv.SetDefault()
		}
	}
	if x.FMessage4 == nil {
		x.FMessage4 = map[string]*pbexternal.External1{
			"m1": &pbexternal.External1{},
			"m2": nil,
			"m3": nil,
			"m4": nil,
			"":   nil,
		}
	}
	if x.FMessage5 == nil {
		x.FMessage5 = map[string]*pbexternal.External2_Embed1{
			"m1": &pbexternal.External2_Embed1{},
			"m2": nil,
			"m3": nil,
			"m4": nil,
			"":   nil,
		}
	}
	for _, mv := range x.FMessage5 {
		if mv != nil {
			mv.SetDefault()
		}
	}
	if x.FMessage6 == nil {
		x.FMessage6 = map[string]*MessageCommon1{
			"m1": &MessageCommon1{},
			"m2": nil,
			"m3": nil,
			"m4": nil,
			"":   nil,
		}
	}
	for _, mv := range x.FMessage6 {
		if mv != nil {
			mv.SetDefault()
		}
	}
	if x.FAny1 == nil {
		x.FAny1 = map[string]*anypb.Any{
			"a1": x.xxx_protoc_plugin_default_mustNewAny(&MessageMap1{}, false),
			"a2": x.xxx_protoc_plugin_default_mustNewAny(&MessageMap1_Embed1{}, false),
			"a3": x.xxx_protoc_plugin_default_mustNewAny(&MessageMap1_Embed1_Embed2{}, false),
			"a4": x.xxx_protoc_plugin_default_mustNewAny(&pbexternal.External1{}, false),
			"a5": x.xxx_protoc_plugin_default_mustNewAny(&pbexternal.External2_Embed1{}, false),
			"a6": x.xxx_protoc_plugin_default_mustNewAny(&timestamppb.Timestamp{}, false),
			"a7": x.xxx_protoc_plugin_default_mustNewAny(&anypb.Any{}, false),
			"a8": x.xxx_protoc_plugin_default_mustNewAny(&MessageCommon1{}, false),
			"":   nil,
		}
	}
}

// xxx_protoc_plugin_default_mustNewAny used to create instance of type any
func (*TypeMap2) xxx_protoc_plugin_default_mustNewAny(src proto.Message, skipEval bool) *anypb.Any {
	if !skipEval {
		if dt, ok := interface{}(src).(interface{ SetDefault() }); ok {
			dt.SetDefault()
		}
	}
	v, err := anypb.New(src)
	if err != nil {
		panic(err)
	}
	return v
}

// Set default value for message TypeMap3 that in file tests/proto/cases/types/type_map.proto
func (x *TypeMap3) SetDefault() {
	if x == nil {
		return
	}
	if x.FString1 == nil {
		x.FString1 = map[string]string{
			"s1": "v1",
			"s2": "v2",
			"s3": "",
			"":   "",
		}
	}
	if len(x.FInt32) == 0 {
		x.FInt32 = map[int32]int32{
			11: 11,
			12: 12,
			13: 0,
			0:  0,
		}
	}
	if x.FInt64 == nil {
		x.FInt64 = map[int64]int64{
			21: 21,
			22: 22,
			23: 0,
			0:  0,
		}
	}
	if len(x.FUint32) == 0 {
		x.FUint32 = map[uint32]uint32{
			31: 31,
			32: 32,
			33: 0,
			0:  0,
		}
	}
	if x.FUint64 == nil {
		x.FUint64 = map[uint64]uint64{
			41: 41,
			42: 42,
			43: 0,
			0:  0,
		}
	}
	if len(x.FSint32) == 0 {
		x.FSint32 = map[int32]int32{
			51: 51,
			52: 52,
			53: 0,
			0:  0,
		}
	}
	if x.FSint64 == nil {
		x.FSint64 = map[int64]int64{
			61: 61,
			62: 62,
			63: 0,
			0:  0,
		}
	}
	if len(x.FSfixed32) == 0 {
		x.FSfixed32 = map[int32]int32{
			71: 71,
			72: 72,
			73: 0,
			0:  0,
		}
	}
	if x.FSfixed64 == nil {
		x.FSfixed64 = map[int64]int64{
			81: 81,
			82: 82,
			83: 0,
			0:  0,
		}
	}
	if len(x.FFixed32) == 0 {
		x.FFixed32 = map[uint32]uint32{
			91: 91,
			92: 92,
			93: 0,
			0:  0,
		}
	}
	if x.FFixed64 == nil {
		x.FFixed64 = map[uint64]uint64{
			101: 101,
			102: 102,
			103: 0,
			0:   0,
		}
	}
	if len(x.FBool1) == 0 {
		x.FBool1 = map[bool]bool{
			true:  true,
			false: false,
		}
	}
}