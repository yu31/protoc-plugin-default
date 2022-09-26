package types

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-defaults/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-defaults/xgo/tests/pb/pbtypes"
)

// Test case for field in `oneof` parts.
func Test_TypeOneOf1(t *testing.T) {
	data := &pbtypes.TypeOneOf1{}

	t.Run("NIL", func(t *testing.T) {
		data.SetDefaults()
		require.Nil(t, data.OneType1)
	})

	t.Run("STRING-Empty", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FString1{FString1: ""}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "ts1", ov.FString1)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString2{FString2: ""}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "", ov.FString2)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString3{FString3: ""}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, `""`, ov.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString4{FString4: ""}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, `"ts4"`, ov.FString4)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString5{FString5: ""}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, `"ts5"`, ov.FString5)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString6{FString6: ""}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, `"ts"6"`, ov.FString6)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString7{FString7: ""}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, `"ts"7"`, ov.FString7)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString8{FString8: ""}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "[ts8]", ov.FString8)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString9{FString9: ""}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "{ts9}", ov.FString9)
		}
	})
	t.Run("STRING-Ignore", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FString1{FString1: "101"}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "101", ov.FString1)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString2{FString2: "102"}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "102", ov.FString2)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString3{FString3: "103"}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "103", ov.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString4{FString4: "104"}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "104", ov.FString4)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString5{FString5: "105"}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "105", ov.FString5)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString6{FString6: "106"}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "106", ov.FString6)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString7{FString7: "107"}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "107", ov.FString7)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString8{FString8: "108"}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "108", ov.FString8)
		}
		{
			ov := &pbtypes.TypeOneOf1_FString9{FString9: "109"}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, "109", ov.FString9)
		}
	})

	t.Run("NUMBER-Empty", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FInt32{FInt32: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, int32(1), ov.FInt32)
		}
		{
			ov := &pbtypes.TypeOneOf1_FInt64{FInt64: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, int64(2), ov.FInt64)
		}
		{
			ov := &pbtypes.TypeOneOf1_FUint32{FUint32: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, uint32(3), ov.FUint32)
		}
		{
			ov := &pbtypes.TypeOneOf1_FUint64{FUint64: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, uint64(4), ov.FUint64)
		}
		{
			ov := &pbtypes.TypeOneOf1_FSint32{FSint32: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, int32(5), ov.FSint32)
		}
		{
			ov := &pbtypes.TypeOneOf1_FSint64{FSint64: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, int64(6), ov.FSint64)
		}
		{
			ov := &pbtypes.TypeOneOf1_FSfixed32{FSfixed32: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, int32(7), ov.FSfixed32)
		}
		{
			ov := &pbtypes.TypeOneOf1_FSfixed64{FSfixed64: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, int64(8), ov.FSfixed64)
		}
		{
			ov := &pbtypes.TypeOneOf1_FFixed32{FFixed32: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, uint32(9), ov.FFixed32)
		}
		{
			ov := &pbtypes.TypeOneOf1_FFixed64{FFixed64: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, uint64(10), ov.FFixed64)
		}
		{
			ov := &pbtypes.TypeOneOf1_FFloat{FFloat: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, float32(11.11), ov.FFloat)
		}
		{
			ov := &pbtypes.TypeOneOf1_FDouble{FDouble: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, float64(12.12), ov.FDouble)
		}
	})
	t.Run("NUMBER-Ignore", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FInt32{FInt32: 101}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, int32(101), ov.FInt32)
		}
		{
			ov := &pbtypes.TypeOneOf1_FInt64{FInt64: 102}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, int64(102), ov.FInt64)
		}
		{
			ov := &pbtypes.TypeOneOf1_FUint32{FUint32: 103}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, uint32(103), ov.FUint32)
		}
		{
			ov := &pbtypes.TypeOneOf1_FUint64{FUint64: 104}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, uint64(104), ov.FUint64)
		}
		{
			ov := &pbtypes.TypeOneOf1_FSint32{FSint32: 105}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, int32(105), ov.FSint32)
		}
		{
			ov := &pbtypes.TypeOneOf1_FSint64{FSint64: 106}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, int64(106), ov.FSint64)
		}
		{
			ov := &pbtypes.TypeOneOf1_FSfixed32{FSfixed32: 107}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, int32(107), ov.FSfixed32)
		}
		{
			ov := &pbtypes.TypeOneOf1_FSfixed64{FSfixed64: 108}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, int64(108), ov.FSfixed64)
		}
		{
			ov := &pbtypes.TypeOneOf1_FFixed32{FFixed32: 109}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, uint32(109), ov.FFixed32)
		}
		{
			ov := &pbtypes.TypeOneOf1_FFixed64{FFixed64: 110}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, uint64(110), ov.FFixed64)
		}
		{
			ov := &pbtypes.TypeOneOf1_FFloat{FFloat: 111.111}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, float32(111.111), ov.FFloat)
		}
		{
			ov := &pbtypes.TypeOneOf1_FDouble{FDouble: 112.112}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, float64(112.112), ov.FDouble)
		}
	})

	t.Run("BOOL-Empty", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FBool1{FBool1: false}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, true, ov.FBool1)
		}
		{
			ov := &pbtypes.TypeOneOf1_FBool2{FBool2: false}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, false, ov.FBool2)
		}
	})
	t.Run("BOOL-Ignore", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FBool1{FBool1: true}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, true, ov.FBool1)
		}
		{
			ov := &pbtypes.TypeOneOf1_FBool2{FBool2: true}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, true, ov.FBool2)
		}
	})

	t.Run("BYTES-NIL", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FBytes1{FBytes1: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, []byte(""), ov.FBytes1)
		}
		{
			ov := &pbtypes.TypeOneOf1_FBytes2{FBytes2: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, []byte("abc"), ov.FBytes2)
		}
	})
	t.Run("BYTES-Empty", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FBytes1{FBytes1: []byte("")}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, []byte(""), ov.FBytes1)
		}
		{
			ov := &pbtypes.TypeOneOf1_FBytes2{FBytes2: []byte("")}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, []byte("abc"), ov.FBytes2)
		}
	})
	t.Run("BYTES-Ignore", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FBytes1{FBytes1: []byte("bytes")}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, []byte("bytes"), ov.FBytes1)
		}
		{
			ov := &pbtypes.TypeOneOf1_FBytes2{FBytes2: []byte("bytes")}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, []byte("bytes"), ov.FBytes2)
		}
	})

	t.Run("ENUM-Empty", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FEnum1{FEnum1: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbtypes.EnumOneOf1(2), ov.FEnum1)
		}
		{
			ov := &pbtypes.TypeOneOf1_FEnum2{FEnum2: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbtypes.EnumOneOf1(10), ov.FEnum2)
		}
		{
			ov := &pbtypes.TypeOneOf1_FEnum3{FEnum3: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbexternal.Month1(1), ov.FEnum3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FEnum4{FEnum4: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbexternal.Month2(7), ov.FEnum4)
		}
		{
			ov := &pbtypes.TypeOneOf1_FEnum5{FEnum5: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbexternal.EnumWeek1_Week(3), ov.FEnum5)
		}
		{
			ov := &pbtypes.TypeOneOf1_FEnum6{FEnum6: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(13), ov.FEnum6)
		}
		{
			ov := &pbtypes.TypeOneOf1_FEnum7{FEnum7: 0}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbtypes.EnumCommon1(9), ov.FEnum7)
		}
	})
	t.Run("ENUM-Ignore", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FEnum1{FEnum1: 201}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbtypes.EnumOneOf1(201), ov.FEnum1)
		}
		{
			ov := &pbtypes.TypeOneOf1_FEnum2{FEnum2: 202}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbtypes.EnumOneOf1(202), ov.FEnum2)
		}
		{
			ov := &pbtypes.TypeOneOf1_FEnum3{FEnum3: 203}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbexternal.Month1(203), ov.FEnum3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FEnum4{FEnum4: 204}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbexternal.Month2(204), ov.FEnum4)
		}
		{
			ov := &pbtypes.TypeOneOf1_FEnum5{FEnum5: 205}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbexternal.EnumWeek1_Week(205), ov.FEnum5)
		}
		{
			ov := &pbtypes.TypeOneOf1_FEnum6{FEnum6: 206}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(206), ov.FEnum6)
		}
		{
			ov := &pbtypes.TypeOneOf1_FEnum7{FEnum7: 207}
			data.OneType1 = ov
			data.SetDefaults()
			require.Equal(t, pbtypes.EnumCommon1(207), ov.FEnum7)
		}
	})

	t.Run("WKT-NIL", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FDuration1{FDuration1: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FDuration1)
			require.Equal(t, int64(100), ov.FDuration1.Seconds)
			require.Equal(t, int32(200), ov.FDuration1.Nanos)
		}
		{
			ov := &pbtypes.TypeOneOf1_FDuration2{FDuration2: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FDuration2)
			require.Equal(t, int64(0), ov.FDuration2.Seconds)
			require.Equal(t, int32(0), ov.FDuration2.Nanos)
		}
		{
			ov := &pbtypes.TypeOneOf1_FTimestamp1{FTimestamp1: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FTimestamp1)
			require.Equal(t, int64(200), ov.FTimestamp1.Seconds)
			require.Equal(t, int32(300), ov.FTimestamp1.Nanos)
		}
		{
			ov := &pbtypes.TypeOneOf1_FTimestamp2{FTimestamp2: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FTimestamp2)
			require.Equal(t, int64(0), ov.FTimestamp2.Seconds)
			require.Equal(t, int32(0), ov.FTimestamp2.Nanos)
		}
	})
	t.Run("WKT-Empty", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FDuration1{FDuration1: &durationpb.Duration{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FDuration1)
			require.Equal(t, int64(100), ov.FDuration1.Seconds)
			require.Equal(t, int32(200), ov.FDuration1.Nanos)
		}
		{
			ov := &pbtypes.TypeOneOf1_FDuration2{FDuration2: &durationpb.Duration{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FDuration2)
			require.Equal(t, int64(0), ov.FDuration2.Seconds)
			require.Equal(t, int32(0), ov.FDuration2.Nanos)
		}
		{
			ov := &pbtypes.TypeOneOf1_FTimestamp1{FTimestamp1: &timestamppb.Timestamp{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FTimestamp1)
			require.Equal(t, int64(200), ov.FTimestamp1.Seconds)
			require.Equal(t, int32(300), ov.FTimestamp1.Nanos)
		}
		{
			ov := &pbtypes.TypeOneOf1_FTimestamp2{FTimestamp2: &timestamppb.Timestamp{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FTimestamp2)
			require.Equal(t, int64(0), ov.FTimestamp2.Seconds)
			require.Equal(t, int32(0), ov.FTimestamp2.Nanos)
		}
	})
	t.Run("WKT-Ignore", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FDuration1{FDuration1: &durationpb.Duration{Seconds: 300, Nanos: 301}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FDuration1)
			require.Equal(t, int64(300), ov.FDuration1.Seconds)
			require.Equal(t, int32(301), ov.FDuration1.Nanos)
		}
		{
			ov := &pbtypes.TypeOneOf1_FDuration2{FDuration2: &durationpb.Duration{Seconds: 400, Nanos: 401}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FDuration2)
			require.Equal(t, int64(400), ov.FDuration2.Seconds)
			require.Equal(t, int32(401), ov.FDuration2.Nanos)
		}
		{
			ov := &pbtypes.TypeOneOf1_FTimestamp1{FTimestamp1: &timestamppb.Timestamp{Seconds: 500, Nanos: 501}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FTimestamp1)
			require.Equal(t, int64(500), ov.FTimestamp1.Seconds)
			require.Equal(t, int32(501), ov.FTimestamp1.Nanos)
		}
		{
			ov := &pbtypes.TypeOneOf1_FTimestamp2{FTimestamp2: &timestamppb.Timestamp{Seconds: 600, Nanos: 601}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FTimestamp2)
			require.Equal(t, int64(600), ov.FTimestamp2.Seconds)
			require.Equal(t, int32(601), ov.FTimestamp2.Nanos)
		}
	})

	t.Run("MESSAGE1-NIL", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FMessage11{FMessage11: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage11)
			require.Equal(t, "a", ov.FMessage11.FString1)
			require.Equal(t, "b", ov.FMessage11.FString2)
			require.Equal(t, "c", ov.FMessage11.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage12{FMessage12: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage12)
			require.Equal(t, "", ov.FMessage12.FString1)
			require.Equal(t, "", ov.FMessage12.FString2)
			require.Equal(t, "", ov.FMessage12.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage13{FMessage13: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage13)
			require.Equal(t, "1", ov.FMessage13.FString1)
			require.Equal(t, "2", ov.FMessage13.FString2)
			require.Equal(t, "3", ov.FMessage13.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage14{FMessage14: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage14)
			require.Equal(t, "", ov.FMessage14.TString)
			require.Equal(t, (*string)(nil), ov.FMessage14.PString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage15{FMessage15: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage15)
			require.Equal(t, "", ov.FMessage15.TString1)
			require.Equal(t, "", ov.FMessage15.TString2)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage16{FMessage16: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage16)
			require.Equal(t, "", ov.FMessage16.FString1)
			require.Equal(t, "", ov.FMessage16.FString2)
			require.Equal(t, "", ov.FMessage16.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage17{FMessage17: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage17)
			require.Equal(t, "x", ov.FMessage17.FString1)
			require.Equal(t, "y", ov.FMessage17.FString2)
			require.Equal(t, "z", ov.FMessage17.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage18{FMessage18: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage18)
			require.Equal(t, "", ov.FMessage18.FString1)
			require.Equal(t, "", ov.FMessage18.FString2)
			require.Equal(t, "", ov.FMessage18.FString3)
		}
	})
	t.Run("MESSAGE1-Empty", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FMessage11{FMessage11: &pbtypes.MessageOneOf1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage11)
			require.Equal(t, "a", ov.FMessage11.FString1)
			require.Equal(t, "b", ov.FMessage11.FString2)
			require.Equal(t, "c", ov.FMessage11.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage12{FMessage12: &pbtypes.MessageOneOf1_Embed1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage12)
			require.Equal(t, "", ov.FMessage12.FString1)
			require.Equal(t, "", ov.FMessage12.FString2)
			require.Equal(t, "", ov.FMessage12.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage13{FMessage13: &pbtypes.MessageOneOf1_Embed1_Embed2{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage13)
			require.Equal(t, "1", ov.FMessage13.FString1)
			require.Equal(t, "2", ov.FMessage13.FString2)
			require.Equal(t, "3", ov.FMessage13.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage14{FMessage14: &pbexternal.External1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage14)
			require.Equal(t, "", ov.FMessage14.TString)
			require.Equal(t, (*string)(nil), ov.FMessage14.PString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage15{FMessage15: &pbexternal.External2_Embed1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage15)
			require.Equal(t, "", ov.FMessage15.TString1)
			require.Equal(t, "", ov.FMessage15.TString2)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage16{FMessage16: &pbtypes.MessageCommon1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage16)
			require.Equal(t, "", ov.FMessage16.FString1)
			require.Equal(t, "", ov.FMessage16.FString2)
			require.Equal(t, "", ov.FMessage16.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage17{FMessage17: &pbtypes.MessageCommon1_Embed1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage17)
			require.Equal(t, "x", ov.FMessage17.FString1)
			require.Equal(t, "y", ov.FMessage17.FString2)
			require.Equal(t, "z", ov.FMessage17.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage18{FMessage18: &pbtypes.MessageCommon1_Embed1_Embed2{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage18)
			require.Equal(t, "", ov.FMessage18.FString1)
			require.Equal(t, "", ov.FMessage18.FString2)
			require.Equal(t, "", ov.FMessage18.FString3)
		}
	})
	t.Run("MESSAGE1-Ignore", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FMessage11{FMessage11: &pbtypes.MessageOneOf1{FString1: "1101", FString2: "1102", FString3: "1103"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage11)
			require.Equal(t, "1101", ov.FMessage11.FString1)
			require.Equal(t, "1102", ov.FMessage11.FString2)
			require.Equal(t, "1103", ov.FMessage11.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage12{FMessage12: &pbtypes.MessageOneOf1_Embed1{FString1: "1201", FString2: "1202", FString3: "1203"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage12)
			require.Equal(t, "1201", ov.FMessage12.FString1)
			require.Equal(t, "1202", ov.FMessage12.FString2)
			require.Equal(t, "1203", ov.FMessage12.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage13{FMessage13: &pbtypes.MessageOneOf1_Embed1_Embed2{FString1: "1301", FString2: "1302", FString3: "1303"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage13)
			require.Equal(t, "1301", ov.FMessage13.FString1)
			require.Equal(t, "1302", ov.FMessage13.FString2)
			require.Equal(t, "1303", ov.FMessage13.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage14{FMessage14: &pbexternal.External1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage14)
			require.Equal(t, "", ov.FMessage14.TString)
			require.Equal(t, (*string)(nil), ov.FMessage14.PString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage15{FMessage15: &pbexternal.External2_Embed1{TString1: "1501", TString2: "1502"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage15)
			require.Equal(t, "1501", ov.FMessage15.TString1)
			require.Equal(t, "1502", ov.FMessage15.TString2)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage16{FMessage16: &pbtypes.MessageCommon1{FString1: "1601", FString2: "1602", FString3: "1603"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage16)
			require.Equal(t, "1601", ov.FMessage16.FString1)
			require.Equal(t, "1602", ov.FMessage16.FString2)
			require.Equal(t, "1603", ov.FMessage16.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage17{FMessage17: &pbtypes.MessageCommon1_Embed1{FString1: "1701", FString2: "1702", FString3: "1703"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage17)
			require.Equal(t, "1701", ov.FMessage17.FString1)
			require.Equal(t, "1702", ov.FMessage17.FString2)
			require.Equal(t, "1703", ov.FMessage17.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage18{FMessage18: &pbtypes.MessageCommon1_Embed1_Embed2{FString1: "1801", FString2: "1802", FString3: "1803"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage18)
			require.Equal(t, "1801", ov.FMessage18.FString1)
			require.Equal(t, "1802", ov.FMessage18.FString2)
			require.Equal(t, "1803", ov.FMessage18.FString3)
		}
	})

	t.Run("MESSAGE2-NIL", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FMessage21{FMessage21: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.Nil(t, ov.FMessage21)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage22{FMessage22: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.Nil(t, ov.FMessage22)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage23{FMessage23: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.Nil(t, ov.FMessage23)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage24{FMessage24: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.Nil(t, ov.FMessage24)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage25{FMessage25: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.Nil(t, ov.FMessage25)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage26{FMessage26: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.Nil(t, ov.FMessage26)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage27{FMessage27: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.Nil(t, ov.FMessage27)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage28{FMessage28: nil}
			data.OneType1 = ov
			data.SetDefaults()
			require.Nil(t, ov.FMessage28)
		}
	})
	t.Run("MESSAGE2-Empty", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FMessage21{FMessage21: &pbtypes.MessageOneOf1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage21)
			require.Equal(t, "a", ov.FMessage21.FString1)
			require.Equal(t, "b", ov.FMessage21.FString2)
			require.Equal(t, "c", ov.FMessage21.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage22{FMessage22: &pbtypes.MessageOneOf1_Embed1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage22)
			require.Equal(t, "", ov.FMessage22.FString1)
			require.Equal(t, "", ov.FMessage22.FString2)
			require.Equal(t, "", ov.FMessage22.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage23{FMessage23: &pbtypes.MessageOneOf1_Embed1_Embed2{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage23)
			require.Equal(t, "1", ov.FMessage23.FString1)
			require.Equal(t, "2", ov.FMessage23.FString2)
			require.Equal(t, "3", ov.FMessage23.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage24{FMessage24: &pbexternal.External1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage24)
			require.Equal(t, "", ov.FMessage24.TString)
			require.Equal(t, (*string)(nil), ov.FMessage24.PString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage25{FMessage25: &pbexternal.External2_Embed1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage25)
			require.Equal(t, "", ov.FMessage25.TString1)
			require.Equal(t, "", ov.FMessage25.TString2)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage26{FMessage26: &pbtypes.MessageCommon1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage26)
			require.Equal(t, "", ov.FMessage26.FString1)
			require.Equal(t, "", ov.FMessage26.FString2)
			require.Equal(t, "", ov.FMessage26.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage27{FMessage27: &pbtypes.MessageCommon1_Embed1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage27)
			require.Equal(t, "x", ov.FMessage27.FString1)
			require.Equal(t, "y", ov.FMessage27.FString2)
			require.Equal(t, "z", ov.FMessage27.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage28{FMessage28: &pbtypes.MessageCommon1_Embed1_Embed2{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage28)
			require.Equal(t, "", ov.FMessage28.FString1)
			require.Equal(t, "", ov.FMessage28.FString2)
			require.Equal(t, "", ov.FMessage28.FString3)
		}
	})
	t.Run("MESSAGE2-Ignore", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FMessage21{FMessage21: &pbtypes.MessageOneOf1{FString1: "2101", FString2: "2102", FString3: "2103"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage21)
			require.Equal(t, "2101", ov.FMessage21.FString1)
			require.Equal(t, "2102", ov.FMessage21.FString2)
			require.Equal(t, "2103", ov.FMessage21.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage22{FMessage22: &pbtypes.MessageOneOf1_Embed1{FString1: "2201", FString2: "2202", FString3: "2203"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage22)
			require.Equal(t, "2201", ov.FMessage22.FString1)
			require.Equal(t, "2202", ov.FMessage22.FString2)
			require.Equal(t, "2203", ov.FMessage22.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage23{FMessage23: &pbtypes.MessageOneOf1_Embed1_Embed2{FString1: "2301", FString2: "2302", FString3: "2303"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage23)
			require.Equal(t, "2301", ov.FMessage23.FString1)
			require.Equal(t, "2302", ov.FMessage23.FString2)
			require.Equal(t, "2303", ov.FMessage23.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage24{FMessage24: &pbexternal.External1{}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage24)
			require.Equal(t, "", ov.FMessage24.TString)
			require.Equal(t, (*string)(nil), ov.FMessage24.PString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage25{FMessage25: &pbexternal.External2_Embed1{TString1: "2501", TString2: "2502"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage25)
			require.Equal(t, "2501", ov.FMessage25.TString1)
			require.Equal(t, "2502", ov.FMessage25.TString2)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage26{FMessage26: &pbtypes.MessageCommon1{FString1: "2601", FString2: "2602", FString3: "2603"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage26)
			require.Equal(t, "2601", ov.FMessage26.FString1)
			require.Equal(t, "2602", ov.FMessage26.FString2)
			require.Equal(t, "2603", ov.FMessage26.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage27{FMessage27: &pbtypes.MessageCommon1_Embed1{FString1: "2701", FString2: "2702", FString3: "2703"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage27)
			require.Equal(t, "2701", ov.FMessage27.FString1)
			require.Equal(t, "2702", ov.FMessage27.FString2)
			require.Equal(t, "2703", ov.FMessage27.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FMessage28{FMessage28: &pbtypes.MessageCommon1_Embed1_Embed2{FString1: "2801", FString2: "2802", FString3: "2803"}}
			data.OneType1 = ov
			data.SetDefaults()
			require.NotNil(t, ov.FMessage28)
			require.Equal(t, "2801", ov.FMessage28.FString1)
			require.Equal(t, "2802", ov.FMessage28.FString2)
			require.Equal(t, "2803", ov.FMessage28.FString3)
		}
	})

	t.Run("ANY-NIL", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FAny1{FAny1: nil}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny1)
			anyVal := &pbtypes.MessageOneOf1{}
			err := ov.FAny1.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny2{FAny2: nil}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny2)
			anyVal := &pbtypes.MessageOneOf1{}
			err := ov.FAny2.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.FString1)
			require.Equal(t, "", anyVal.FString2)
			require.Equal(t, "", anyVal.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny3{FAny3: nil}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny3)
			anyVal := &pbtypes.MessageOneOf1_Embed1{}
			err := ov.FAny3.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.FString1)
			require.Equal(t, "", anyVal.FString2)
			require.Equal(t, "", anyVal.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny4{FAny4: nil}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny4)
			anyVal := &pbtypes.MessageOneOf1_Embed1_Embed2{}
			err := ov.FAny4.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "1", anyVal.FString1)
			require.Equal(t, "2", anyVal.FString2)
			require.Equal(t, "3", anyVal.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny5{FAny5: nil}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny5)
			anyVal := &pbexternal.External1{}
			err := ov.FAny5.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString)
			require.Equal(t, (*string)(nil), anyVal.PString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny6{FAny6: nil}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny6)
			anyVal := &pbexternal.External2_Embed1{}
			err := ov.FAny6.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString1)
			require.Equal(t, "", anyVal.TString2)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny7{FAny7: nil}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny7)
			anyVal := &timestamppb.Timestamp{Seconds: 0, Nanos: 0}
			err := ov.FAny7.UnmarshalTo(anyVal)
			require.Nil(t, err)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny8{FAny8: nil}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny8)
			anyVal := &anypb.Any{}
			err := ov.FAny8.UnmarshalTo(anyVal)
			require.Nil(t, err)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny9{FAny9: nil}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny9)
			anyVal := &pbtypes.MessageCommon1{}
			err := ov.FAny9.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
	})
	t.Run("ANY-Empty", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FAny1{FAny1: &anypb.Any{}}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny1)
			anyVal := &pbtypes.MessageOneOf1{}
			err := ov.FAny1.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny2{FAny2: &anypb.Any{}}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny2)
			anyVal := &pbtypes.MessageOneOf1{}
			err := ov.FAny2.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.FString1)
			require.Equal(t, "", anyVal.FString2)
			require.Equal(t, "", anyVal.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny3{FAny3: &anypb.Any{}}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny3)
			anyVal := &pbtypes.MessageOneOf1_Embed1{}
			err := ov.FAny3.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.FString1)
			require.Equal(t, "", anyVal.FString2)
			require.Equal(t, "", anyVal.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny4{FAny4: &anypb.Any{}}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny4)
			anyVal := &pbtypes.MessageOneOf1_Embed1_Embed2{}
			err := ov.FAny4.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "1", anyVal.FString1)
			require.Equal(t, "2", anyVal.FString2)
			require.Equal(t, "3", anyVal.FString3)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny5{FAny5: &anypb.Any{}}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny5)
			anyVal := &pbexternal.External1{}
			err := ov.FAny5.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString)
			require.Equal(t, (*string)(nil), anyVal.PString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny6{FAny6: &anypb.Any{}}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny6)
			anyVal := &pbexternal.External2_Embed1{}
			err := ov.FAny6.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString1)
			require.Equal(t, "", anyVal.TString2)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny7{FAny7: &anypb.Any{}}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny7)
			anyVal := &timestamppb.Timestamp{Seconds: 0, Nanos: 0}
			err := ov.FAny7.UnmarshalTo(anyVal)
			require.Nil(t, err)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny8{FAny8: &anypb.Any{}}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny8)
			anyVal := &anypb.Any{}
			err := ov.FAny8.UnmarshalTo(anyVal)
			require.Nil(t, err)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny9{FAny9: &anypb.Any{}}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny9)
			anyVal := &pbtypes.MessageCommon1{}
			err := ov.FAny9.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
	})
	t.Run("ANY-Ignore", func(t *testing.T) {
		{
			ov := &pbtypes.TypeOneOf1_FAny1{FAny1: mustNewAny(&pbexternal.External3{TString: "any1"})}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny1)
			anyVal := &pbexternal.External3{}
			err := ov.FAny1.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any1", anyVal.TString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny2{FAny2: mustNewAny(&pbexternal.External3{TString: "any2"})}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny2)
			anyVal := &pbexternal.External3{}
			err := ov.FAny2.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any2", anyVal.TString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny3{FAny3: mustNewAny(&pbexternal.External3{TString: "any3"})}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny3)
			anyVal := &pbexternal.External3{}
			err := ov.FAny3.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any3", anyVal.TString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny4{FAny4: mustNewAny(&pbexternal.External3{TString: "any4"})}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny4)
			anyVal := &pbexternal.External3{}
			err := ov.FAny4.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any4", anyVal.TString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny5{FAny5: mustNewAny(&pbexternal.External3{TString: "any5"})}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny5)
			anyVal := &pbexternal.External3{}
			err := ov.FAny5.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any5", anyVal.TString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny6{FAny6: mustNewAny(&pbexternal.External3{TString: "any6"})}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny6)
			anyVal := &pbexternal.External3{}
			err := ov.FAny6.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any6", anyVal.TString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny7{FAny7: mustNewAny(&pbexternal.External3{TString: "any7"})}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny7)
			anyVal := &pbexternal.External3{}
			err := ov.FAny7.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any7", anyVal.TString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny8{FAny8: mustNewAny(&pbexternal.External3{TString: "any8"})}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny8)
			anyVal := &pbexternal.External3{}
			err := ov.FAny8.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any8", anyVal.TString)
		}
		{
			ov := &pbtypes.TypeOneOf1_FAny9{FAny9: mustNewAny(&pbexternal.External3{TString: "any9"})}
			data.OneType1 = ov
			data.SetDefaults()

			require.NotNil(t, ov.FAny9)
			anyVal := &pbexternal.External3{}
			err := ov.FAny9.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any9", anyVal.TString)
		}
	})
}

// Test case for oneof field.
func Test_TypeOneOf2(t *testing.T) {
	data := pbtypes.TypeOneOf2{}
	data.SetDefaults()

	t.Run("OneStr1", func(t *testing.T) {
		require.NotNil(t, data.OneStr1)
		ov, ok := data.OneStr1.(*pbtypes.TypeOneOf2_FString1)
		require.True(t, ok)
		require.NotNil(t, ov)
		require.Equal(t, "a", ov.FString1)
	})

	t.Run("OneStr2", func(t *testing.T) {
		require.NotNil(t, data.OneStr2)
		ov, ok := data.OneStr2.(*pbtypes.TypeOneOf2_FString4)
		require.True(t, ok)
		require.NotNil(t, ov)
		require.Equal(t, "d", ov.FString4)
	})

	t.Run("OneMsg1", func(t *testing.T) {
		require.NotNil(t, data.OneMsg1)
		ov, ok := data.OneMsg1.(*pbtypes.TypeOneOf2_FMessage11)
		require.True(t, ok)
		require.NotNil(t, ov)
		require.NotNil(t, ov.FMessage11)
		require.Equal(t, "a", ov.FMessage11.FString1)
		require.Equal(t, "b", ov.FMessage11.FString2)
		require.Equal(t, "c", ov.FMessage11.FString3)
	})

	t.Run("OneMsg2", func(t *testing.T) {
		require.NotNil(t, data.OneMsg2)
		ov, ok := data.OneMsg2.(*pbtypes.TypeOneOf2_FMessage22)
		require.True(t, ok)
		require.NotNil(t, ov)
		require.NotNil(t, ov.FMessage22)
		require.Equal(t, "", ov.FMessage22.FString1)
		require.Equal(t, "", ov.FMessage22.FString2)
		require.Equal(t, "", ov.FMessage22.FString3)
	})

	t.Run("OneMsg3", func(t *testing.T) {
		require.NotNil(t, data.OneMsg3)
		ov, ok := data.OneMsg3.(*pbtypes.TypeOneOf2_FMessage33)
		require.True(t, ok)
		require.NotNil(t, ov)
		require.NotNil(t, ov.FMessage33)
		require.Equal(t, "1", ov.FMessage33.FString1)
		require.Equal(t, "2", ov.FMessage33.FString2)
		require.Equal(t, "3", ov.FMessage33.FString3)
	})

	t.Run("OneMsg4", func(t *testing.T) {
		require.NotNil(t, data.OneMsg4)
		ov, ok := data.OneMsg4.(*pbtypes.TypeOneOf2_FMessage44)
		require.True(t, ok)
		require.NotNil(t, ov)
		require.NotNil(t, ov.FMessage44)
		require.Equal(t, "", ov.FMessage44.TString)
		require.Equal(t, (*string)(nil), ov.FMessage44.PString)
	})

	t.Run("OneMsg5", func(t *testing.T) {
		require.NotNil(t, data.OneMsg5)
		ov, ok := data.OneMsg5.(*pbtypes.TypeOneOf2_FMessage55)
		require.True(t, ok)
		require.NotNil(t, ov)
		require.NotNil(t, ov.FMessage55)
		require.Equal(t, "", ov.FMessage55.TString1)
		require.Equal(t, "", ov.FMessage55.TString1)
	})
}
