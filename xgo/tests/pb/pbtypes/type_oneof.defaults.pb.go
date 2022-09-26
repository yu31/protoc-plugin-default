// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.1
// source: tests/proto/cases/types/type_oneof.proto

package pbtypes

import (
	_ "github.com/yu31/protoc-plugin-defaults/xgo/pb/pbdefaults"
	pbexternal "github.com/yu31/protoc-plugin-defaults/xgo/tests/pb/pbexternal"
	proto "google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// Set default value for message MessageOneOf1 that in file tests/proto/cases/types/type_oneof.proto
func (x *MessageOneOf1) SetDefaults() {
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

// Set default value for message Embed1 that in file tests/proto/cases/types/type_oneof.proto
func (x *MessageOneOf1_Embed1) SetDefaults() {
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

// Set default value for message Embed2 that in file tests/proto/cases/types/type_oneof.proto
func (x *MessageOneOf1_Embed1_Embed2) SetDefaults() {
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

// Set default value for message TypeOneOf1 that in file tests/proto/cases/types/type_oneof.proto
func (x *TypeOneOf1) SetDefaults() {
	if x == nil {
		return
	}
	switch o := x.OneType1.(type) {
	case *TypeOneOf1_FString1:
		if o.FString1 == "" {
			o.FString1 = "ts1"
		}
	case *TypeOneOf1_FString2:
		if o.FString2 == "" {
			o.FString2 = ""
		}
	case *TypeOneOf1_FString3:
		if o.FString3 == "" {
			o.FString3 = "\"\""
		}
	case *TypeOneOf1_FString4:
		if o.FString4 == "" {
			o.FString4 = "\"ts4\""
		}
	case *TypeOneOf1_FString5:
		if o.FString5 == "" {
			o.FString5 = "\"ts5\""
		}
	case *TypeOneOf1_FString6:
		if o.FString6 == "" {
			o.FString6 = "\"ts\"6\""
		}
	case *TypeOneOf1_FString7:
		if o.FString7 == "" {
			o.FString7 = "\"ts\"7\""
		}
	case *TypeOneOf1_FString8:
		if o.FString8 == "" {
			o.FString8 = "[ts8]"
		}
	case *TypeOneOf1_FString9:
		if o.FString9 == "" {
			o.FString9 = "{ts9}"
		}
	case *TypeOneOf1_FInt32:
		if o.FInt32 == 0 {
			o.FInt32 = 1
		}
	case *TypeOneOf1_FInt64:
		if o.FInt64 == 0 {
			o.FInt64 = 2
		}
	case *TypeOneOf1_FUint32:
		if o.FUint32 == 0 {
			o.FUint32 = 3
		}
	case *TypeOneOf1_FUint64:
		if o.FUint64 == 0 {
			o.FUint64 = 4
		}
	case *TypeOneOf1_FSint32:
		if o.FSint32 == 0 {
			o.FSint32 = 5
		}
	case *TypeOneOf1_FSint64:
		if o.FSint64 == 0 {
			o.FSint64 = 6
		}
	case *TypeOneOf1_FSfixed32:
		if o.FSfixed32 == 0 {
			o.FSfixed32 = 7
		}
	case *TypeOneOf1_FSfixed64:
		if o.FSfixed64 == 0 {
			o.FSfixed64 = 8
		}
	case *TypeOneOf1_FFixed32:
		if o.FFixed32 == 0 {
			o.FFixed32 = 9
		}
	case *TypeOneOf1_FFixed64:
		if o.FFixed64 == 0 {
			o.FFixed64 = 10
		}
	case *TypeOneOf1_FFloat:
		if o.FFloat == 0 {
			o.FFloat = 11.109999656677246
		}
	case *TypeOneOf1_FDouble:
		if o.FDouble == 0 {
			o.FDouble = 12.12
		}
	case *TypeOneOf1_FBool1:
		if o.FBool1 == false {
			o.FBool1 = true
		}
	case *TypeOneOf1_FBool2:
		if o.FBool2 == false {
			o.FBool2 = false
		}
	case *TypeOneOf1_FBytes1:
		if len(o.FBytes1) == 0 {
			o.FBytes1 = []byte{}
		}
	case *TypeOneOf1_FBytes2:
		if len(o.FBytes2) == 0 {
			o.FBytes2 = []byte{97, 98, 99}
		}
	case *TypeOneOf1_FEnum1:
		if o.FEnum1 == 0 {
			o.FEnum1 = 2
		}
	case *TypeOneOf1_FEnum2:
		if o.FEnum2 == 0 {
			o.FEnum2 = 10
		}
	case *TypeOneOf1_FEnum3:
		if o.FEnum3 == 0 {
			o.FEnum3 = 1
		}
	case *TypeOneOf1_FEnum4:
		if o.FEnum4 == 0 {
			o.FEnum4 = 7
		}
	case *TypeOneOf1_FEnum5:
		if o.FEnum5 == 0 {
			o.FEnum5 = 3
		}
	case *TypeOneOf1_FEnum6:
		if o.FEnum6 == 0 {
			o.FEnum6 = 13
		}
	case *TypeOneOf1_FEnum7:
		if o.FEnum7 == 0 {
			o.FEnum7 = 9
		}
	case *TypeOneOf1_FDuration1:
		if o.FDuration1.AsDuration() == 0 {
			o.FDuration1 = &durationpb.Duration{Seconds: 100, Nanos: 200}
		}
	case *TypeOneOf1_FDuration2:
		if o.FDuration2.AsDuration() == 0 {
			o.FDuration2 = &durationpb.Duration{Seconds: 0, Nanos: 0}
		}
	case *TypeOneOf1_FTimestamp1:
		if o.FTimestamp1.AsTime().UnixNano() == 0 {
			o.FTimestamp1 = &timestamppb.Timestamp{Seconds: 200, Nanos: 300}
		}
	case *TypeOneOf1_FTimestamp2:
		if o.FTimestamp2.AsTime().UnixNano() == 0 {
			o.FTimestamp2 = &timestamppb.Timestamp{Seconds: 0, Nanos: 0}
		}
	case *TypeOneOf1_FMessage11:
		if o.FMessage11 == nil {
			o.FMessage11 = &MessageOneOf1{}
		}
		if o.FMessage11 != nil {
			o.FMessage11.SetDefaults()
		}
	case *TypeOneOf1_FMessage12:
		if o.FMessage12 == nil {
			o.FMessage12 = &MessageOneOf1_Embed1{}
		}
	case *TypeOneOf1_FMessage13:
		if o.FMessage13 == nil {
			o.FMessage13 = &MessageOneOf1_Embed1_Embed2{}
		}
		if o.FMessage13 != nil {
			o.FMessage13.SetDefaults()
		}
	case *TypeOneOf1_FMessage14:
		if o.FMessage14 == nil {
			o.FMessage14 = &pbexternal.External1{}
		}
	case *TypeOneOf1_FMessage15:
		if o.FMessage15 == nil {
			o.FMessage15 = &pbexternal.External2_Embed1{}
		}
		if o.FMessage15 != nil {
			o.FMessage15.SetDefaults()
		}
	case *TypeOneOf1_FMessage16:
		if o.FMessage16 == nil {
			o.FMessage16 = &MessageCommon1{}
		}
	case *TypeOneOf1_FMessage17:
		if o.FMessage17 == nil {
			o.FMessage17 = &MessageCommon1_Embed1{}
		}
		if o.FMessage17 != nil {
			o.FMessage17.SetDefaults()
		}
	case *TypeOneOf1_FMessage18:
		if o.FMessage18 == nil {
			o.FMessage18 = &MessageCommon1_Embed1_Embed2{}
		}
	case *TypeOneOf1_FMessage21:
		if o.FMessage21 != nil {
			o.FMessage21.SetDefaults()
		}
	case *TypeOneOf1_FMessage22:
	case *TypeOneOf1_FMessage23:
		if o.FMessage23 != nil {
			o.FMessage23.SetDefaults()
		}
	case *TypeOneOf1_FMessage24:
	case *TypeOneOf1_FMessage25:
		if o.FMessage25 != nil {
			o.FMessage25.SetDefaults()
		}
	case *TypeOneOf1_FMessage26:
	case *TypeOneOf1_FMessage27:
		if o.FMessage27 != nil {
			o.FMessage27.SetDefaults()
		}
	case *TypeOneOf1_FMessage28:
	case *TypeOneOf1_FAny1:
		if o.FAny1.GetTypeUrl() == "" {
			o.FAny1 = x.xxxPGD_mustNewAny(&MessageOneOf1{}, false)
		}
	case *TypeOneOf1_FAny2:
		if o.FAny2.GetTypeUrl() == "" {
			o.FAny2 = x.xxxPGD_mustNewAny(&MessageOneOf1{}, true)
		}
	case *TypeOneOf1_FAny3:
		if o.FAny3.GetTypeUrl() == "" {
			o.FAny3 = x.xxxPGD_mustNewAny(&MessageOneOf1_Embed1{}, true)
		}
	case *TypeOneOf1_FAny4:
		if o.FAny4.GetTypeUrl() == "" {
			o.FAny4 = x.xxxPGD_mustNewAny(&MessageOneOf1_Embed1_Embed2{}, false)
		}
	case *TypeOneOf1_FAny5:
		if o.FAny5.GetTypeUrl() == "" {
			o.FAny5 = x.xxxPGD_mustNewAny(&pbexternal.External1{}, true)
		}
	case *TypeOneOf1_FAny6:
		if o.FAny6.GetTypeUrl() == "" {
			o.FAny6 = x.xxxPGD_mustNewAny(&pbexternal.External2_Embed1{}, false)
		}
	case *TypeOneOf1_FAny7:
		if o.FAny7.GetTypeUrl() == "" {
			o.FAny7 = x.xxxPGD_mustNewAny(&timestamppb.Timestamp{}, false)
		}
	case *TypeOneOf1_FAny8:
		if o.FAny8.GetTypeUrl() == "" {
			o.FAny8 = x.xxxPGD_mustNewAny(&anypb.Any{}, false)
		}
	case *TypeOneOf1_FAny9:
		if o.FAny9.GetTypeUrl() == "" {
			o.FAny9 = x.xxxPGD_mustNewAny(&MessageCommon1{}, false)
		}
	default:
		_ = o // to avoid unused panic
	}
}

// xxxPGD_mustNewAny used to create instance of type any
func (*TypeOneOf1) xxxPGD_mustNewAny(src proto.Message, skipEval bool) *anypb.Any {
	if !skipEval {
		if dt, ok := interface{}(src).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	v, err := anypb.New(src)
	if err != nil {
		panic(err)
	}
	return v
}

// Set default value for message TypeOneOf2 that in file tests/proto/cases/types/type_oneof.proto
func (x *TypeOneOf2) SetDefaults() {
	if x == nil {
		return
	}
	if x.OneStr1 == nil {
		x.OneStr1 = new(TypeOneOf2_FString1)
	}
	switch o := x.OneStr1.(type) {
	case *TypeOneOf2_FString1:
		if o.FString1 == "" {
			o.FString1 = "a"
		}
	case *TypeOneOf2_FString2:
		if o.FString2 == "" {
			o.FString2 = "b"
		}
	default:
		_ = o // to avoid unused panic
	}
	if x.OneStr2 == nil {
		x.OneStr2 = new(TypeOneOf2_FString4)
	}
	switch o := x.OneStr2.(type) {
	case *TypeOneOf2_FString3:
		if o.FString3 == "" {
			o.FString3 = "c"
		}
	case *TypeOneOf2_FString4:
		if o.FString4 == "" {
			o.FString4 = "d"
		}
	default:
		_ = o // to avoid unused panic
	}
	if x.OneMsg1 == nil {
		x.OneMsg1 = new(TypeOneOf2_FMessage11)
	}
	switch o := x.OneMsg1.(type) {
	case *TypeOneOf2_FMessage11:
		if o.FMessage11 == nil {
			o.FMessage11 = &MessageOneOf1{}
		}
		if o.FMessage11 != nil {
			o.FMessage11.SetDefaults()
		}
	case *TypeOneOf2_FMessage12:
		if o.FMessage12 == nil {
			o.FMessage12 = &MessageOneOf1_Embed1{}
		}
	case *TypeOneOf2_FMessage13:
		if o.FMessage13 == nil {
			o.FMessage13 = &MessageOneOf1_Embed1_Embed2{}
		}
		if o.FMessage13 != nil {
			o.FMessage13.SetDefaults()
		}
	case *TypeOneOf2_FMessage14:
		if o.FMessage14 == nil {
			o.FMessage14 = &pbexternal.External1{}
		}
	case *TypeOneOf2_FMessage15:
		if o.FMessage15 == nil {
			o.FMessage15 = &pbexternal.External2_Embed1{}
		}
		if o.FMessage15 != nil {
			o.FMessage15.SetDefaults()
		}
	default:
		_ = o // to avoid unused panic
	}
	if x.OneMsg2 == nil {
		x.OneMsg2 = new(TypeOneOf2_FMessage22)
	}
	switch o := x.OneMsg2.(type) {
	case *TypeOneOf2_FMessage21:
		if o.FMessage21 == nil {
			o.FMessage21 = &MessageOneOf1{}
		}
		if o.FMessage21 != nil {
			o.FMessage21.SetDefaults()
		}
	case *TypeOneOf2_FMessage22:
		if o.FMessage22 == nil {
			o.FMessage22 = &MessageOneOf1_Embed1{}
		}
	case *TypeOneOf2_FMessage23:
		if o.FMessage23 == nil {
			o.FMessage23 = &MessageOneOf1_Embed1_Embed2{}
		}
		if o.FMessage23 != nil {
			o.FMessage23.SetDefaults()
		}
	case *TypeOneOf2_FMessage24:
		if o.FMessage24 == nil {
			o.FMessage24 = &pbexternal.External1{}
		}
	case *TypeOneOf2_FMessage25:
		if o.FMessage25 == nil {
			o.FMessage25 = &pbexternal.External2_Embed1{}
		}
		if o.FMessage25 != nil {
			o.FMessage25.SetDefaults()
		}
	default:
		_ = o // to avoid unused panic
	}
	if x.OneMsg3 == nil {
		x.OneMsg3 = new(TypeOneOf2_FMessage33)
	}
	switch o := x.OneMsg3.(type) {
	case *TypeOneOf2_FMessage31:
		if o.FMessage31 == nil {
			o.FMessage31 = &MessageOneOf1{}
		}
		if o.FMessage31 != nil {
			o.FMessage31.SetDefaults()
		}
	case *TypeOneOf2_FMessage32:
		if o.FMessage32 == nil {
			o.FMessage32 = &MessageOneOf1_Embed1{}
		}
	case *TypeOneOf2_FMessage33:
		if o.FMessage33 == nil {
			o.FMessage33 = &MessageOneOf1_Embed1_Embed2{}
		}
		if o.FMessage33 != nil {
			o.FMessage33.SetDefaults()
		}
	case *TypeOneOf2_FMessage34:
		if o.FMessage34 == nil {
			o.FMessage34 = &pbexternal.External1{}
		}
	case *TypeOneOf2_FMessage35:
		if o.FMessage35 == nil {
			o.FMessage35 = &pbexternal.External2_Embed1{}
		}
		if o.FMessage35 != nil {
			o.FMessage35.SetDefaults()
		}
	default:
		_ = o // to avoid unused panic
	}
	if x.OneMsg4 == nil {
		x.OneMsg4 = new(TypeOneOf2_FMessage44)
	}
	switch o := x.OneMsg4.(type) {
	case *TypeOneOf2_FMessage41:
		if o.FMessage41 == nil {
			o.FMessage41 = &MessageOneOf1{}
		}
		if o.FMessage41 != nil {
			o.FMessage41.SetDefaults()
		}
	case *TypeOneOf2_FMessage42:
		if o.FMessage42 == nil {
			o.FMessage42 = &MessageOneOf1_Embed1{}
		}
	case *TypeOneOf2_FMessage43:
		if o.FMessage43 == nil {
			o.FMessage43 = &MessageOneOf1_Embed1_Embed2{}
		}
		if o.FMessage43 != nil {
			o.FMessage43.SetDefaults()
		}
	case *TypeOneOf2_FMessage44:
		if o.FMessage44 == nil {
			o.FMessage44 = &pbexternal.External1{}
		}
	case *TypeOneOf2_FMessage45:
		if o.FMessage45 == nil {
			o.FMessage45 = &pbexternal.External2_Embed1{}
		}
		if o.FMessage45 != nil {
			o.FMessage45.SetDefaults()
		}
	default:
		_ = o // to avoid unused panic
	}
	if x.OneMsg5 == nil {
		x.OneMsg5 = new(TypeOneOf2_FMessage55)
	}
	switch o := x.OneMsg5.(type) {
	case *TypeOneOf2_FMessage51:
		if o.FMessage51 == nil {
			o.FMessage51 = &MessageOneOf1{}
		}
		if o.FMessage51 != nil {
			o.FMessage51.SetDefaults()
		}
	case *TypeOneOf2_FMessage52:
		if o.FMessage52 == nil {
			o.FMessage52 = &MessageOneOf1_Embed1{}
		}
	case *TypeOneOf2_FMessage53:
		if o.FMessage53 == nil {
			o.FMessage53 = &MessageOneOf1_Embed1_Embed2{}
		}
		if o.FMessage53 != nil {
			o.FMessage53.SetDefaults()
		}
	case *TypeOneOf2_FMessage54:
		if o.FMessage54 == nil {
			o.FMessage54 = &pbexternal.External1{}
		}
	case *TypeOneOf2_FMessage55:
		if o.FMessage55 == nil {
			o.FMessage55 = &pbexternal.External2_Embed1{}
		}
		if o.FMessage55 != nil {
			o.FMessage55.SetDefaults()
		}
	default:
		_ = o // to avoid unused panic
	}
}