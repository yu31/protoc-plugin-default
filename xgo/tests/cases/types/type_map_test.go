package types

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-default/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-default/xgo/tests/pb/pbtypes"
)

// Test cases the field value is nil pointer.
func Test_TypeMap1_NIL(t *testing.T) {
	data := pbtypes.TypeMap1{}
	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, 10, len(data.FString1))
		require.Equal(t, "ts1", data.FString1["ts1"])
		require.Equal(t, "", data.FString1[""])
		require.Equal(t, `""`, data.FString1[`""`])
		require.Equal(t, `"ts4"`, data.FString1[`"ts4"`])
		require.Equal(t, `"ts5"`, data.FString1[`"ts5"`])
		require.Equal(t, `"ts"6"`, data.FString1[`"ts"6"`])
		require.Equal(t, `"ts"7"`, data.FString1[`"ts"7"`])
		require.Equal(t, `[ts8]`, data.FString1[`[ts8]`])
		require.Equal(t, `{ts9}`, data.FString1[`{ts9}`])
		require.Equal(t, "", data.FString1["s10"])
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, 4, len(data.FInt32))
		require.Equal(t, int32(11), data.FInt32["11"])
		require.Equal(t, int32(12), data.FInt32["12"])
		require.Equal(t, int32(0), data.FInt32["13"])
		require.Equal(t, int32(0), data.FInt32[""])

		require.Equal(t, 4, len(data.FInt64))
		require.Equal(t, int64(21), data.FInt64["21"])
		require.Equal(t, int64(22), data.FInt64["22"])
		require.Equal(t, int64(0), data.FInt64["23"])
		require.Equal(t, int64(0), data.FInt64[""])

		require.Equal(t, 4, len(data.FUint32))
		require.Equal(t, uint32(31), data.FUint32["31"])
		require.Equal(t, uint32(32), data.FUint32["32"])
		require.Equal(t, uint32(0), data.FUint32["33"])
		require.Equal(t, uint32(0), data.FUint32[""])

		require.Equal(t, 4, len(data.FUint64))
		require.Equal(t, uint64(41), data.FUint64["41"])
		require.Equal(t, uint64(42), data.FUint64["42"])
		require.Equal(t, uint64(0), data.FUint64["43"])
		require.Equal(t, uint64(0), data.FUint64[""])

		require.Equal(t, 4, len(data.FSint32))
		require.Equal(t, int32(51), data.FSint32["51"])
		require.Equal(t, int32(52), data.FSint32["52"])
		require.Equal(t, int32(0), data.FSint32["53"])
		require.Equal(t, int32(0), data.FSint32[""])

		require.Equal(t, 4, len(data.FSint64))
		require.Equal(t, int64(61), data.FSint64["61"])
		require.Equal(t, int64(62), data.FSint64["62"])
		require.Equal(t, int64(0), data.FSint64["63"])
		require.Equal(t, int64(0), data.FSint64[""])

		require.Equal(t, 4, len(data.FSfixed32))
		require.Equal(t, int32(71), data.FSfixed32["71"])
		require.Equal(t, int32(72), data.FSfixed32["72"])
		require.Equal(t, int32(0), data.FSfixed32["73"])
		require.Equal(t, int32(0), data.FSfixed32[""])

		require.Equal(t, 4, len(data.FSfixed64))
		require.Equal(t, int64(81), data.FSfixed64["81"])
		require.Equal(t, int64(82), data.FSfixed64["82"])
		require.Equal(t, int64(0), data.FSfixed64["83"])
		require.Equal(t, int64(0), data.FSfixed64[""])

		require.Equal(t, 4, len(data.FFixed32))
		require.Equal(t, uint32(91), data.FFixed32["91"])
		require.Equal(t, uint32(92), data.FFixed32["92"])
		require.Equal(t, uint32(0), data.FFixed32["93"])
		require.Equal(t, uint32(0), data.FFixed32[""])

		require.Equal(t, 4, len(data.FFixed64))
		require.Equal(t, uint64(101), data.FFixed64["101"])
		require.Equal(t, uint64(102), data.FFixed64["102"])
		require.Equal(t, uint64(0), data.FFixed64["103"])
		require.Equal(t, uint64(0), data.FFixed64[""])

		require.Equal(t, 4, len(data.FFloat))
		require.Equal(t, float32(111.111), data.FFloat["111.111"])
		require.Equal(t, float32(112.112), data.FFloat["112.112"])
		require.Equal(t, float32(0), data.FFloat["113.113"])
		require.Equal(t, float32(0), data.FFloat[""])

		require.Equal(t, 4, len(data.FDouble))
		require.Equal(t, float64(121.121), data.FDouble["121.121"])
		require.Equal(t, float64(122.122), data.FDouble["122.122"])
		require.Equal(t, float64(0), data.FDouble["123.123"])
		require.Equal(t, float64(0), data.FDouble[""])
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, 4, len(data.FBool1))
		require.Equal(t, true, data.FBool1["t1"])
		require.Equal(t, false, data.FBool1["t2"])
		require.Equal(t, false, data.FBool1["t3"])
		require.Equal(t, false, data.FBool1[""])
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, 4, len(data.FBytes1))
		require.Equal(t, []byte("abc"), data.FBytes1["b1"])
		require.Equal(t, []byte("123"), data.FBytes1["b2"])
		require.Equal(t, []byte(""), data.FBytes1["b3"])
		require.Equal(t, []byte(""), data.FBytes1[""])
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, 4, len(data.FEnum1))
		require.Equal(t, pbtypes.EnumMap1(1), data.FEnum1["e1"])
		require.Equal(t, pbtypes.EnumMap1(3), data.FEnum1["e2"])
		require.Equal(t, pbtypes.EnumMap1(0), data.FEnum1["e3"])
		require.Equal(t, pbtypes.EnumMap1(0), data.FEnum1[""])

		require.Equal(t, 4, len(data.FEnum2))
		require.Equal(t, pbtypes.EnumMap1(2), data.FEnum2["e1"])
		require.Equal(t, pbtypes.EnumMap1(4), data.FEnum2["e2"])
		require.Equal(t, pbtypes.EnumMap1(0), data.FEnum2["e3"])
		require.Equal(t, pbtypes.EnumMap1(0), data.FEnum2[""])

		require.Equal(t, 4, len(data.FEnum3))
		require.Equal(t, pbexternal.Month1(1), data.FEnum3["e1"])
		require.Equal(t, pbexternal.Month1(3), data.FEnum3["e2"])
		require.Equal(t, pbexternal.Month1(0), data.FEnum3["e3"])
		require.Equal(t, pbexternal.Month1(0), data.FEnum3[""])

		require.Equal(t, 4, len(data.FEnum4))
		require.Equal(t, pbexternal.Month2(2), data.FEnum4["e1"])
		require.Equal(t, pbexternal.Month2(4), data.FEnum4["e2"])
		require.Equal(t, pbexternal.Month2(0), data.FEnum4["e3"])
		require.Equal(t, pbexternal.Month2(0), data.FEnum4[""])

		require.Equal(t, 4, len(data.FEnum5))
		require.Equal(t, pbexternal.EnumWeek1_Week(1), data.FEnum5["e1"])
		require.Equal(t, pbexternal.EnumWeek1_Week(3), data.FEnum5["e2"])
		require.Equal(t, pbexternal.EnumWeek1_Week(0), data.FEnum5["e3"])
		require.Equal(t, pbexternal.EnumWeek1_Week(0), data.FEnum5[""])

		require.Equal(t, 4, len(data.FEnum6))
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(2), data.FEnum6["e1"])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(4), data.FEnum6["e2"])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(0), data.FEnum6["e3"])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(0), data.FEnum6[""])

		require.Equal(t, 4, len(data.FEnum7))
		require.Equal(t, pbtypes.EnumCommon1(1), data.FEnum7["e1"])
		require.Equal(t, pbtypes.EnumCommon1(3), data.FEnum7["e2"])
		require.Equal(t, pbtypes.EnumCommon1(0), data.FEnum7["e3"])
		require.Equal(t, pbtypes.EnumCommon1(0), data.FEnum7[""])
	})

	t.Run("WKT", func(t *testing.T) {
		{
			require.Equal(t, 5, len(data.FDuration1))
			require.Equal(t, int64(100), data.FDuration1["d1"].Seconds)
			require.Equal(t, int32(200), data.FDuration1["d1"].Nanos)

			require.Equal(t, int64(101), data.FDuration1["d2"].Seconds)
			require.Equal(t, int32(201), data.FDuration1["d2"].Nanos)

			require.Equal(t, int64(0), data.FDuration1["d3"].Seconds)
			require.Equal(t, int32(0), data.FDuration1["d3"].Nanos)

			require.Equal(t, (*durationpb.Duration)(nil), data.FDuration1["d4"])
			require.Equal(t, (*durationpb.Duration)(nil), data.FDuration1[""])
		}
		{
			require.Equal(t, 5, len(data.FTimestamp1))
			require.Equal(t, int64(100), data.FTimestamp1["t1"].Seconds)
			require.Equal(t, int32(200), data.FTimestamp1["t1"].Nanos)

			require.Equal(t, int64(101), data.FTimestamp1["t2"].Seconds)
			require.Equal(t, int32(201), data.FTimestamp1["t2"].Nanos)

			require.Equal(t, int64(0), data.FTimestamp1["t3"].Seconds)
			require.Equal(t, int32(0), data.FTimestamp1["t3"].Nanos)

			require.Equal(t, (*timestamppb.Timestamp)(nil), data.FTimestamp1["t4"])
			require.Equal(t, (*timestamppb.Timestamp)(nil), data.FTimestamp1[""])
		}
	})

	t.Run("MESSAGE", func(t *testing.T) {
		{
			require.Equal(t, 5, len(data.FMessage1))
			require.NotNil(t, data.FMessage1["m1"])
			require.Equal(t, "a", data.FMessage1["m1"].FString1)
			require.Equal(t, "b", data.FMessage1["m1"].FString2)
			require.Equal(t, "c", data.FMessage1["m1"].FString3)
			require.Equal(t, (*pbtypes.MessageMap1)(nil), data.FMessage1["m2"])
			require.Equal(t, (*pbtypes.MessageMap1)(nil), data.FMessage1["m3"])
			require.Equal(t, (*pbtypes.MessageMap1)(nil), data.FMessage1["m4"])
			require.Equal(t, (*pbtypes.MessageMap1)(nil), data.FMessage1[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage2))
			require.NotNil(t, data.FMessage2["m1"])
			require.Equal(t, "", data.FMessage2["m1"].FString1)
			require.Equal(t, "", data.FMessage2["m1"].FString2)
			require.Equal(t, "", data.FMessage2["m1"].FString3)
			require.Equal(t, (*pbtypes.MessageMap1_Embed1)(nil), data.FMessage2["m2"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1)(nil), data.FMessage2["m3"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1)(nil), data.FMessage2["m4"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1)(nil), data.FMessage2[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage3))
			require.NotNil(t, data.FMessage3["m1"])
			require.Equal(t, "1", data.FMessage3["m1"].FString1)
			require.Equal(t, "2", data.FMessage3["m1"].FString2)
			require.Equal(t, "3", data.FMessage3["m1"].FString3)
			require.Equal(t, (*pbtypes.MessageMap1_Embed1_Embed2)(nil), data.FMessage3["m2"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1_Embed2)(nil), data.FMessage3["m3"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1_Embed2)(nil), data.FMessage3["m4"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1_Embed2)(nil), data.FMessage3[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage4))
			require.NotNil(t, data.FMessage4["m1"])
			require.Equal(t, "", data.FMessage4["m1"].TString)
			require.Equal(t, (*string)(nil), data.FMessage4["m1"].PString)
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4["m2"])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4["m3"])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4["m4"])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage5))
			require.NotNil(t, data.FMessage5["m1"])
			require.Equal(t, "", data.FMessage5["m1"].TString1)
			require.Equal(t, "", data.FMessage5["m1"].TString2)
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5["m2"])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5["m3"])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5["m4"])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage6))
			require.NotNil(t, data.FMessage6["m1"])
			require.Equal(t, "a", data.FMessage6["m1"].FString1)
			require.Equal(t, "b", data.FMessage6["m1"].FString2)
			require.Equal(t, "c", data.FMessage6["m1"].FString3)
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6["m2"])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6["m3"])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6["m4"])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[""])
		}
	})

	t.Run("ANY", func(t *testing.T) {
		require.Equal(t, 9, len(data.FAny1))
		{
			require.NotNil(t, data.FAny1["a1"])
			anyVal := &pbtypes.MessageMap1{}
			err := data.FAny1["a1"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1["a2"])
			anyVal := &pbtypes.MessageMap1_Embed1{}
			err := data.FAny1["a2"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "x", anyVal.FString1)
			require.Equal(t, "y", anyVal.FString2)
			require.Equal(t, "z", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1["a3"])
			anyVal := &pbtypes.MessageMap1_Embed1_Embed2{}
			err := data.FAny1["a3"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "1", anyVal.FString1)
			require.Equal(t, "2", anyVal.FString2)
			require.Equal(t, "3", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1["a4"])
			anyVal := &pbexternal.External1{}
			err := data.FAny1["a4"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString)
			require.Equal(t, (*string)(nil), anyVal.PString)
		}
		{
			require.NotNil(t, data.FAny1["a5"])
			anyVal := &pbexternal.External2_Embed1{}
			err := data.FAny1["a5"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString1)
			require.Equal(t, "", anyVal.TString2)
		}
		{
			require.NotNil(t, data.FAny1["a6"])
			anyVal := &timestamppb.Timestamp{}
			err := data.FAny1["a6"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, int64(0), anyVal.Seconds)
			require.Equal(t, int32(0), anyVal.Nanos)
		}
		{
			require.NotNil(t, data.FAny1["a7"])
			anyVal := &anypb.Any{}
			err := data.FAny1["a7"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TypeUrl)
			require.Equal(t, ([]byte)(nil), anyVal.Value)
		}
		{
			require.NotNil(t, data.FAny1["a8"])
			anyVal := &pbtypes.MessageCommon1{}
			err := data.FAny1["a8"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.Equal(t, (*anypb.Any)(nil), data.FAny1[""])
		}
	})
}

// Test cases the field value is empty.
func Test_TypeMap1_Empty(t *testing.T) {
	data := pbtypes.TypeMap1{
		FString1:    map[string]string{},
		FInt32:      map[string]int32{},
		FInt64:      map[string]int64{},
		FUint32:     map[string]uint32{},
		FUint64:     map[string]uint64{},
		FSint32:     map[string]int32{},
		FSint64:     map[string]int64{},
		FSfixed32:   map[string]int32{},
		FSfixed64:   map[string]int64{},
		FFixed32:    map[string]uint32{},
		FFixed64:    map[string]uint64{},
		FFloat:      map[string]float32{},
		FDouble:     map[string]float64{},
		FBool1:      map[string]bool{},
		FBytes1:     map[string][]byte{},
		FEnum1:      map[string]pbtypes.EnumMap1{},
		FEnum2:      map[string]pbtypes.EnumMap1{},
		FEnum3:      map[string]pbexternal.Month1{},
		FEnum4:      map[string]pbexternal.Month2{},
		FEnum5:      map[string]pbexternal.EnumWeek1_Week{},
		FEnum6:      map[string]pbexternal.EnumWeek2_Embed1_Week{},
		FEnum7:      map[string]pbtypes.EnumCommon1{},
		FDuration1:  map[string]*durationpb.Duration{},
		FTimestamp1: map[string]*timestamppb.Timestamp{},
		FMessage1:   map[string]*pbtypes.MessageMap1{},
		FMessage2:   map[string]*pbtypes.MessageMap1_Embed1{},
		FMessage3:   map[string]*pbtypes.MessageMap1_Embed1_Embed2{},
		FMessage4:   map[string]*pbexternal.External1{},
		FMessage5:   map[string]*pbexternal.External2_Embed1{},
		FMessage6:   map[string]*pbtypes.MessageCommon1{},
		FAny1:       map[string]*anypb.Any{},
	}
	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, 10, len(data.FString1))
		require.Equal(t, "ts1", data.FString1["ts1"])
		require.Equal(t, "", data.FString1[""])
		require.Equal(t, `""`, data.FString1[`""`])
		require.Equal(t, `"ts4"`, data.FString1[`"ts4"`])
		require.Equal(t, `"ts5"`, data.FString1[`"ts5"`])
		require.Equal(t, `"ts"6"`, data.FString1[`"ts"6"`])
		require.Equal(t, `"ts"7"`, data.FString1[`"ts"7"`])
		require.Equal(t, `[ts8]`, data.FString1[`[ts8]`])
		require.Equal(t, `{ts9}`, data.FString1[`{ts9}`])
		require.Equal(t, "", data.FString1["s10"])
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, 4, len(data.FInt32))
		require.Equal(t, int32(11), data.FInt32["11"])
		require.Equal(t, int32(12), data.FInt32["12"])
		require.Equal(t, int32(0), data.FInt32["13"])
		require.Equal(t, int32(0), data.FInt32[""])

		require.Equal(t, 4, len(data.FInt64))
		require.Equal(t, int64(21), data.FInt64["21"])
		require.Equal(t, int64(22), data.FInt64["22"])
		require.Equal(t, int64(0), data.FInt64["23"])
		require.Equal(t, int64(0), data.FInt64[""])

		require.Equal(t, 4, len(data.FUint32))
		require.Equal(t, uint32(31), data.FUint32["31"])
		require.Equal(t, uint32(32), data.FUint32["32"])
		require.Equal(t, uint32(0), data.FUint32["33"])
		require.Equal(t, uint32(0), data.FUint32[""])

		require.Equal(t, 4, len(data.FUint64))
		require.Equal(t, uint64(41), data.FUint64["41"])
		require.Equal(t, uint64(42), data.FUint64["42"])
		require.Equal(t, uint64(0), data.FUint64["43"])
		require.Equal(t, uint64(0), data.FUint64[""])

		require.Equal(t, 4, len(data.FSint32))
		require.Equal(t, int32(51), data.FSint32["51"])
		require.Equal(t, int32(52), data.FSint32["52"])
		require.Equal(t, int32(0), data.FSint32["53"])
		require.Equal(t, int32(0), data.FSint32[""])

		require.Equal(t, 4, len(data.FSint64))
		require.Equal(t, int64(61), data.FSint64["61"])
		require.Equal(t, int64(62), data.FSint64["62"])
		require.Equal(t, int64(0), data.FSint64["63"])
		require.Equal(t, int64(0), data.FSint64[""])

		require.Equal(t, 4, len(data.FSfixed32))
		require.Equal(t, int32(71), data.FSfixed32["71"])
		require.Equal(t, int32(72), data.FSfixed32["72"])
		require.Equal(t, int32(0), data.FSfixed32["73"])
		require.Equal(t, int32(0), data.FSfixed32[""])

		require.Equal(t, 4, len(data.FSfixed64))
		require.Equal(t, int64(81), data.FSfixed64["81"])
		require.Equal(t, int64(82), data.FSfixed64["82"])
		require.Equal(t, int64(0), data.FSfixed64["83"])
		require.Equal(t, int64(0), data.FSfixed64[""])

		require.Equal(t, 4, len(data.FFixed32))
		require.Equal(t, uint32(91), data.FFixed32["91"])
		require.Equal(t, uint32(92), data.FFixed32["92"])
		require.Equal(t, uint32(0), data.FFixed32["93"])
		require.Equal(t, uint32(0), data.FFixed32[""])

		require.Equal(t, 4, len(data.FFixed64))
		require.Equal(t, uint64(101), data.FFixed64["101"])
		require.Equal(t, uint64(102), data.FFixed64["102"])
		require.Equal(t, uint64(0), data.FFixed64["103"])
		require.Equal(t, uint64(0), data.FFixed64[""])

		require.Equal(t, 4, len(data.FFloat))
		require.Equal(t, float32(111.111), data.FFloat["111.111"])
		require.Equal(t, float32(112.112), data.FFloat["112.112"])
		require.Equal(t, float32(0), data.FFloat["113.113"])
		require.Equal(t, float32(0), data.FFloat[""])

		require.Equal(t, 4, len(data.FDouble))
		require.Equal(t, float64(121.121), data.FDouble["121.121"])
		require.Equal(t, float64(122.122), data.FDouble["122.122"])
		require.Equal(t, float64(0), data.FDouble["123.123"])
		require.Equal(t, float64(0), data.FDouble[""])
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, 4, len(data.FBool1))
		require.Equal(t, true, data.FBool1["t1"])
		require.Equal(t, false, data.FBool1["t2"])
		require.Equal(t, false, data.FBool1["t3"])
		require.Equal(t, false, data.FBool1[""])
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, 4, len(data.FBytes1))
		require.Equal(t, []byte("abc"), data.FBytes1["b1"])
		require.Equal(t, []byte("123"), data.FBytes1["b2"])
		require.Equal(t, []byte(""), data.FBytes1["b3"])
		require.Equal(t, []byte(""), data.FBytes1[""])
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, 4, len(data.FEnum1))
		require.Equal(t, pbtypes.EnumMap1(1), data.FEnum1["e1"])
		require.Equal(t, pbtypes.EnumMap1(3), data.FEnum1["e2"])
		require.Equal(t, pbtypes.EnumMap1(0), data.FEnum1["e3"])
		require.Equal(t, pbtypes.EnumMap1(0), data.FEnum1[""])

		require.Equal(t, 4, len(data.FEnum2))
		require.Equal(t, pbtypes.EnumMap1(2), data.FEnum2["e1"])
		require.Equal(t, pbtypes.EnumMap1(4), data.FEnum2["e2"])
		require.Equal(t, pbtypes.EnumMap1(0), data.FEnum2["e3"])
		require.Equal(t, pbtypes.EnumMap1(0), data.FEnum2[""])

		require.Equal(t, 4, len(data.FEnum3))
		require.Equal(t, pbexternal.Month1(1), data.FEnum3["e1"])
		require.Equal(t, pbexternal.Month1(3), data.FEnum3["e2"])
		require.Equal(t, pbexternal.Month1(0), data.FEnum3["e3"])
		require.Equal(t, pbexternal.Month1(0), data.FEnum3[""])

		require.Equal(t, 4, len(data.FEnum4))
		require.Equal(t, pbexternal.Month2(2), data.FEnum4["e1"])
		require.Equal(t, pbexternal.Month2(4), data.FEnum4["e2"])
		require.Equal(t, pbexternal.Month2(0), data.FEnum4["e3"])
		require.Equal(t, pbexternal.Month2(0), data.FEnum4[""])

		require.Equal(t, 4, len(data.FEnum5))
		require.Equal(t, pbexternal.EnumWeek1_Week(1), data.FEnum5["e1"])
		require.Equal(t, pbexternal.EnumWeek1_Week(3), data.FEnum5["e2"])
		require.Equal(t, pbexternal.EnumWeek1_Week(0), data.FEnum5["e3"])
		require.Equal(t, pbexternal.EnumWeek1_Week(0), data.FEnum5[""])

		require.Equal(t, 4, len(data.FEnum6))
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(2), data.FEnum6["e1"])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(4), data.FEnum6["e2"])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(0), data.FEnum6["e3"])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(0), data.FEnum6[""])

		require.Equal(t, 4, len(data.FEnum7))
		require.Equal(t, pbtypes.EnumCommon1(1), data.FEnum7["e1"])
		require.Equal(t, pbtypes.EnumCommon1(3), data.FEnum7["e2"])
		require.Equal(t, pbtypes.EnumCommon1(0), data.FEnum7["e3"])
		require.Equal(t, pbtypes.EnumCommon1(0), data.FEnum7[""])
	})

	t.Run("WKT", func(t *testing.T) {
		{
			require.Equal(t, 5, len(data.FDuration1))
			require.Equal(t, int64(100), data.FDuration1["d1"].Seconds)
			require.Equal(t, int32(200), data.FDuration1["d1"].Nanos)

			require.Equal(t, int64(101), data.FDuration1["d2"].Seconds)
			require.Equal(t, int32(201), data.FDuration1["d2"].Nanos)

			require.Equal(t, int64(0), data.FDuration1["d3"].Seconds)
			require.Equal(t, int32(0), data.FDuration1["d3"].Nanos)

			require.Equal(t, (*durationpb.Duration)(nil), data.FDuration1["d4"])
			require.Equal(t, (*durationpb.Duration)(nil), data.FDuration1[""])
		}
		{
			require.Equal(t, 5, len(data.FTimestamp1))
			require.Equal(t, int64(100), data.FTimestamp1["t1"].Seconds)
			require.Equal(t, int32(200), data.FTimestamp1["t1"].Nanos)

			require.Equal(t, int64(101), data.FTimestamp1["t2"].Seconds)
			require.Equal(t, int32(201), data.FTimestamp1["t2"].Nanos)

			require.Equal(t, int64(0), data.FTimestamp1["t3"].Seconds)
			require.Equal(t, int32(0), data.FTimestamp1["t3"].Nanos)

			require.Equal(t, (*timestamppb.Timestamp)(nil), data.FTimestamp1["t4"])
			require.Equal(t, (*timestamppb.Timestamp)(nil), data.FTimestamp1[""])
		}
	})

	t.Run("MESSAGE", func(t *testing.T) {
		{
			require.Equal(t, 5, len(data.FMessage1))
			require.NotNil(t, data.FMessage1["m1"])
			require.Equal(t, "a", data.FMessage1["m1"].FString1)
			require.Equal(t, "b", data.FMessage1["m1"].FString2)
			require.Equal(t, "c", data.FMessage1["m1"].FString3)
			require.Equal(t, (*pbtypes.MessageMap1)(nil), data.FMessage1["m2"])
			require.Equal(t, (*pbtypes.MessageMap1)(nil), data.FMessage1["m3"])
			require.Equal(t, (*pbtypes.MessageMap1)(nil), data.FMessage1["m4"])
			require.Equal(t, (*pbtypes.MessageMap1)(nil), data.FMessage1[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage2))
			require.NotNil(t, data.FMessage2["m1"])
			require.Equal(t, "", data.FMessage2["m1"].FString1)
			require.Equal(t, "", data.FMessage2["m1"].FString2)
			require.Equal(t, "", data.FMessage2["m1"].FString3)
			require.Equal(t, (*pbtypes.MessageMap1_Embed1)(nil), data.FMessage2["m2"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1)(nil), data.FMessage2["m3"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1)(nil), data.FMessage2["m4"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1)(nil), data.FMessage2[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage3))
			require.NotNil(t, data.FMessage3["m1"])
			require.Equal(t, "1", data.FMessage3["m1"].FString1)
			require.Equal(t, "2", data.FMessage3["m1"].FString2)
			require.Equal(t, "3", data.FMessage3["m1"].FString3)
			require.Equal(t, (*pbtypes.MessageMap1_Embed1_Embed2)(nil), data.FMessage3["m2"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1_Embed2)(nil), data.FMessage3["m3"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1_Embed2)(nil), data.FMessage3["m4"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1_Embed2)(nil), data.FMessage3[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage4))
			require.NotNil(t, data.FMessage4["m1"])
			require.Equal(t, "", data.FMessage4["m1"].TString)
			require.Equal(t, (*string)(nil), data.FMessage4["m1"].PString)
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4["m2"])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4["m3"])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4["m4"])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage5))
			require.NotNil(t, data.FMessage5["m1"])
			require.Equal(t, "", data.FMessage5["m1"].TString1)
			require.Equal(t, "", data.FMessage5["m1"].TString2)
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5["m2"])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5["m3"])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5["m4"])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage6))
			require.NotNil(t, data.FMessage6["m1"])
			require.Equal(t, "a", data.FMessage6["m1"].FString1)
			require.Equal(t, "b", data.FMessage6["m1"].FString2)
			require.Equal(t, "c", data.FMessage6["m1"].FString3)
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6["m2"])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6["m3"])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6["m4"])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[""])
		}
	})

	t.Run("ANY", func(t *testing.T) {
		require.Equal(t, 9, len(data.FAny1))
		{
			require.NotNil(t, data.FAny1["a1"])
			anyVal := &pbtypes.MessageMap1{}
			err := data.FAny1["a1"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1["a2"])
			anyVal := &pbtypes.MessageMap1_Embed1{}
			err := data.FAny1["a2"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "x", anyVal.FString1)
			require.Equal(t, "y", anyVal.FString2)
			require.Equal(t, "z", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1["a3"])
			anyVal := &pbtypes.MessageMap1_Embed1_Embed2{}
			err := data.FAny1["a3"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "1", anyVal.FString1)
			require.Equal(t, "2", anyVal.FString2)
			require.Equal(t, "3", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1["a4"])
			anyVal := &pbexternal.External1{}
			err := data.FAny1["a4"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString)
			require.Equal(t, (*string)(nil), anyVal.PString)
		}
		{
			require.NotNil(t, data.FAny1["a5"])
			anyVal := &pbexternal.External2_Embed1{}
			err := data.FAny1["a5"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString1)
			require.Equal(t, "", anyVal.TString2)
		}
		{
			require.NotNil(t, data.FAny1["a6"])
			anyVal := &timestamppb.Timestamp{}
			err := data.FAny1["a6"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, int64(0), anyVal.Seconds)
			require.Equal(t, int32(0), anyVal.Nanos)
		}
		{
			require.NotNil(t, data.FAny1["a7"])
			anyVal := &anypb.Any{}
			err := data.FAny1["a7"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TypeUrl)
			require.Equal(t, ([]byte)(nil), anyVal.Value)
		}
		{
			require.NotNil(t, data.FAny1["a8"])
			anyVal := &pbtypes.MessageCommon1{}
			err := data.FAny1["a8"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.Equal(t, (*anypb.Any)(nil), data.FAny1[""])
		}
	})
}

// Test cases the field value are have valid value, expected ignore default value.
func Test_TypeMap1_Ignore(t *testing.T) {
	data := pbtypes.TypeMap1{
		FString1:    map[string]string{"sk1": "101"},
		FInt32:      map[string]int32{"101": 101},
		FInt64:      map[string]int64{"102": 102},
		FUint32:     map[string]uint32{"103": 103},
		FUint64:     map[string]uint64{"104": 104},
		FSint32:     map[string]int32{"105": 105},
		FSint64:     map[string]int64{"106": 106},
		FSfixed32:   map[string]int32{"107": 107},
		FSfixed64:   map[string]int64{"108": 108},
		FFixed32:    map[string]uint32{"109": 109},
		FFixed64:    map[string]uint64{"110": 110},
		FFloat:      map[string]float32{"111.111": 111.111},
		FDouble:     map[string]float64{"112.112": 112.112},
		FBool1:      map[string]bool{"bk1": true},
		FBytes1:     map[string][]byte{"bk1": []byte("bytes1")},
		FEnum1:      map[string]pbtypes.EnumMap1{"201": 201},
		FEnum2:      map[string]pbtypes.EnumMap1{"202": 202},
		FEnum3:      map[string]pbexternal.Month1{"203": 203},
		FEnum4:      map[string]pbexternal.Month2{"204": 204},
		FEnum5:      map[string]pbexternal.EnumWeek1_Week{"205": 205},
		FEnum6:      map[string]pbexternal.EnumWeek2_Embed1_Week{"206": 206},
		FEnum7:      map[string]pbtypes.EnumCommon1{"207": 207},
		FDuration1:  map[string]*durationpb.Duration{"dk1": {Seconds: 300, Nanos: 301}},
		FTimestamp1: map[string]*timestamppb.Timestamp{"tk1": {Seconds: 500, Nanos: 501}},
		FMessage1:   map[string]*pbtypes.MessageMap1{"mk1": {FString1: "1101", FString2: "1102", FString3: "1103"}},
		FMessage2:   map[string]*pbtypes.MessageMap1_Embed1{"mk2": {FString1: "1201", FString2: "1202", FString3: "1203"}},
		FMessage3:   map[string]*pbtypes.MessageMap1_Embed1_Embed2{"mk3": {FString1: "1301", FString2: "1302", FString3: "1303"}},
		FMessage4:   map[string]*pbexternal.External1{"mk4": {TString: "1401"}},
		FMessage5:   map[string]*pbexternal.External2_Embed1{"mk5": {TString1: "1501", TString2: "1502"}},
		FMessage6:   map[string]*pbtypes.MessageCommon1{"mk6": {FString1: "1601", FString2: "1602", FString3: "1603"}},
		FAny1:       map[string]*anypb.Any{"ak1": mustNewAny(&pbexternal.External3{TString: "any1"})},
	}
	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, 1, len(data.FString1))
		require.Equal(t, "101", data.FString1["sk1"])
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, 1, len(data.FInt32))
		require.Equal(t, int32(101), data.FInt32["101"])

		require.Equal(t, 1, len(data.FInt64))
		require.Equal(t, int64(102), data.FInt64["102"])

		require.Equal(t, 1, len(data.FUint32))
		require.Equal(t, uint32(103), data.FUint32["103"])

		require.Equal(t, 1, len(data.FUint64))
		require.Equal(t, uint64(104), data.FUint64["104"])

		require.Equal(t, 1, len(data.FSint32))
		require.Equal(t, int32(105), data.FSint32["105"])

		require.Equal(t, 1, len(data.FSint64))
		require.Equal(t, int64(106), data.FSint64["106"])

		require.Equal(t, 1, len(data.FSfixed32))
		require.Equal(t, int32(107), data.FSfixed32["107"])

		require.Equal(t, 1, len(data.FSfixed64))
		require.Equal(t, int64(108), data.FSfixed64["108"])

		require.Equal(t, 1, len(data.FFixed32))
		require.Equal(t, uint32(109), data.FFixed32["109"])

		require.Equal(t, 1, len(data.FFixed64))
		require.Equal(t, uint64(110), data.FFixed64["110"])

		require.Equal(t, 1, len(data.FFloat))
		require.Equal(t, float32(111.111), data.FFloat["111.111"])

		require.Equal(t, 1, len(data.FDouble))
		require.Equal(t, float64(112.112), data.FDouble["112.112"])
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, 1, len(data.FBool1))
		require.Equal(t, true, data.FBool1["bk1"])
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, 1, len(data.FBytes1))
		require.Equal(t, []byte("bytes1"), data.FBytes1["bk1"])
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, 1, len(data.FEnum1))
		require.Equal(t, pbtypes.EnumMap1(201), data.FEnum1["201"])

		require.Equal(t, 1, len(data.FEnum2))
		require.Equal(t, pbtypes.EnumMap1(202), data.FEnum2["202"])

		require.Equal(t, 1, len(data.FEnum3))
		require.Equal(t, pbexternal.Month1(203), data.FEnum3["203"])

		require.Equal(t, 1, len(data.FEnum4))
		require.Equal(t, pbexternal.Month2(204), data.FEnum4["204"])

		require.Equal(t, 1, len(data.FEnum5))
		require.Equal(t, pbexternal.EnumWeek1_Week(205), data.FEnum5["205"])

		require.Equal(t, 1, len(data.FEnum6))
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(206), data.FEnum6["206"])

		require.Equal(t, 1, len(data.FEnum7))
		require.Equal(t, pbtypes.EnumCommon1(207), data.FEnum7["207"])
	})

	t.Run("WKT", func(t *testing.T) {
		{
			require.Equal(t, 1, len(data.FDuration1))
			require.Equal(t, int64(300), data.FDuration1["dk1"].Seconds)
			require.Equal(t, int32(301), data.FDuration1["dk1"].Nanos)
		}
		{
			require.Equal(t, 1, len(data.FTimestamp1))
			require.Equal(t, int64(500), data.FTimestamp1["tk1"].Seconds)
			require.Equal(t, int32(501), data.FTimestamp1["tk1"].Nanos)
		}
	})

	t.Run("MESSAGE", func(t *testing.T) {
		{
			require.Equal(t, 1, len(data.FMessage1))
			require.NotNil(t, data.FMessage1["mk1"])
			require.Equal(t, "1101", data.FMessage1["mk1"].FString1)
			require.Equal(t, "1102", data.FMessage1["mk1"].FString2)
			require.Equal(t, "1103", data.FMessage1["mk1"].FString3)
		}
		{
			require.Equal(t, 1, len(data.FMessage2))
			require.NotNil(t, data.FMessage2["mk2"])
			require.Equal(t, "1201", data.FMessage2["mk2"].FString1)
			require.Equal(t, "1202", data.FMessage2["mk2"].FString2)
			require.Equal(t, "1203", data.FMessage2["mk2"].FString3)
		}
		{
			require.Equal(t, 1, len(data.FMessage3))
			require.NotNil(t, data.FMessage3["mk3"])
			require.Equal(t, "1301", data.FMessage3["mk3"].FString1)
			require.Equal(t, "1302", data.FMessage3["mk3"].FString2)
			require.Equal(t, "1303", data.FMessage3["mk3"].FString3)
		}
		{
			require.Equal(t, 1, len(data.FMessage4))
			require.NotNil(t, data.FMessage4["mk4"])
			require.Equal(t, "1401", data.FMessage4["mk4"].TString)
			require.Equal(t, (*string)(nil), data.FMessage4["mk4"].PString)
		}
		{
			require.Equal(t, 1, len(data.FMessage5))
			require.NotNil(t, data.FMessage5["mk5"])
			require.Equal(t, "1501", data.FMessage5["mk5"].TString1)
			require.Equal(t, "1502", data.FMessage5["mk5"].TString2)
		}
		{
			require.Equal(t, 1, len(data.FMessage6))
			require.NotNil(t, data.FMessage6["mk6"])
			require.Equal(t, "1601", data.FMessage6["mk6"].FString1)
			require.Equal(t, "1602", data.FMessage6["mk6"].FString2)
			require.Equal(t, "1603", data.FMessage6["mk6"].FString3)
		}
	})

	t.Run("ANY", func(t *testing.T) {
		require.Equal(t, 1, len(data.FAny1))
		{
			require.NotNil(t, data.FAny1["ak1"])
			anyVal := &pbexternal.External3{}
			err := data.FAny1["ak1"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any1", anyVal.TString)
		}
	})
}

// Test cases the field value is nil pointer.
func Test_TypeMap2_NIL(t *testing.T) {
	data := pbtypes.TypeMap2{}
	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, 10, len(data.FString1))
		require.Equal(t, "ts1", data.FString1["ts1"])
		require.Equal(t, "", data.FString1[""])
		require.Equal(t, `""`, data.FString1[`""`])
		require.Equal(t, `"ts4"`, data.FString1[`"ts4"`])
		require.Equal(t, `"ts5"`, data.FString1[`"ts5"`])
		require.Equal(t, `"ts"6"`, data.FString1[`"ts"6"`])
		require.Equal(t, `"ts"7"`, data.FString1[`"ts"7"`])
		require.Equal(t, `[ts8]`, data.FString1[`[ts8]`])
		require.Equal(t, `{ts9}`, data.FString1[`{ts9}`])
		require.Equal(t, "", data.FString1["s10"])
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, 4, len(data.FInt32))
		require.Equal(t, int32(11), data.FInt32["11"])
		require.Equal(t, int32(12), data.FInt32["12"])
		require.Equal(t, int32(0), data.FInt32["13"])
		require.Equal(t, int32(0), data.FInt32[""])

		require.Equal(t, 4, len(data.FInt64))
		require.Equal(t, int64(21), data.FInt64["21"])
		require.Equal(t, int64(22), data.FInt64["22"])
		require.Equal(t, int64(0), data.FInt64["23"])
		require.Equal(t, int64(0), data.FInt64[""])

		require.Equal(t, 4, len(data.FUint32))
		require.Equal(t, uint32(31), data.FUint32["31"])
		require.Equal(t, uint32(32), data.FUint32["32"])
		require.Equal(t, uint32(0), data.FUint32["33"])
		require.Equal(t, uint32(0), data.FUint32[""])

		require.Equal(t, 4, len(data.FUint64))
		require.Equal(t, uint64(41), data.FUint64["41"])
		require.Equal(t, uint64(42), data.FUint64["42"])
		require.Equal(t, uint64(0), data.FUint64["43"])
		require.Equal(t, uint64(0), data.FUint64[""])

		require.Equal(t, 4, len(data.FSint32))
		require.Equal(t, int32(51), data.FSint32["51"])
		require.Equal(t, int32(52), data.FSint32["52"])
		require.Equal(t, int32(0), data.FSint32["53"])
		require.Equal(t, int32(0), data.FSint32[""])

		require.Equal(t, 4, len(data.FSint64))
		require.Equal(t, int64(61), data.FSint64["61"])
		require.Equal(t, int64(62), data.FSint64["62"])
		require.Equal(t, int64(0), data.FSint64["63"])
		require.Equal(t, int64(0), data.FSint64[""])

		require.Equal(t, 4, len(data.FSfixed32))
		require.Equal(t, int32(71), data.FSfixed32["71"])
		require.Equal(t, int32(72), data.FSfixed32["72"])
		require.Equal(t, int32(0), data.FSfixed32["73"])
		require.Equal(t, int32(0), data.FSfixed32[""])

		require.Equal(t, 4, len(data.FSfixed64))
		require.Equal(t, int64(81), data.FSfixed64["81"])
		require.Equal(t, int64(82), data.FSfixed64["82"])
		require.Equal(t, int64(0), data.FSfixed64["83"])
		require.Equal(t, int64(0), data.FSfixed64[""])

		require.Equal(t, 4, len(data.FFixed32))
		require.Equal(t, uint32(91), data.FFixed32["91"])
		require.Equal(t, uint32(92), data.FFixed32["92"])
		require.Equal(t, uint32(0), data.FFixed32["93"])
		require.Equal(t, uint32(0), data.FFixed32[""])

		require.Equal(t, 4, len(data.FFixed64))
		require.Equal(t, uint64(101), data.FFixed64["101"])
		require.Equal(t, uint64(102), data.FFixed64["102"])
		require.Equal(t, uint64(0), data.FFixed64["103"])
		require.Equal(t, uint64(0), data.FFixed64[""])

		require.Equal(t, 4, len(data.FFloat))
		require.Equal(t, float32(111.111), data.FFloat["111.111"])
		require.Equal(t, float32(112.112), data.FFloat["112.112"])
		require.Equal(t, float32(0), data.FFloat["113.113"])
		require.Equal(t, float32(0), data.FFloat[""])

		require.Equal(t, 4, len(data.FDouble))
		require.Equal(t, float64(121.121), data.FDouble["121.121"])
		require.Equal(t, float64(122.122), data.FDouble["122.122"])
		require.Equal(t, float64(0), data.FDouble["123.123"])
		require.Equal(t, float64(0), data.FDouble[""])
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, 4, len(data.FBool1))
		require.Equal(t, true, data.FBool1["t1"])
		require.Equal(t, false, data.FBool1["t2"])
		require.Equal(t, false, data.FBool1["t3"])
		require.Equal(t, false, data.FBool1[""])
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, 4, len(data.FBytes1))
		require.Equal(t, []byte("abc"), data.FBytes1["b1"])
		require.Equal(t, []byte("123"), data.FBytes1["b2"])
		require.Equal(t, []byte(""), data.FBytes1["b3"])
		require.Equal(t, []byte(""), data.FBytes1[""])
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, 4, len(data.FEnum1))
		require.Equal(t, pbtypes.EnumMap1(1), data.FEnum1["e1"])
		require.Equal(t, pbtypes.EnumMap1(3), data.FEnum1["e2"])
		require.Equal(t, pbtypes.EnumMap1(0), data.FEnum1["e3"])
		require.Equal(t, pbtypes.EnumMap1(0), data.FEnum1[""])

		require.Equal(t, 4, len(data.FEnum2))
		require.Equal(t, pbtypes.EnumMap1(2), data.FEnum2["e1"])
		require.Equal(t, pbtypes.EnumMap1(4), data.FEnum2["e2"])
		require.Equal(t, pbtypes.EnumMap1(0), data.FEnum2["e3"])
		require.Equal(t, pbtypes.EnumMap1(0), data.FEnum2[""])

		require.Equal(t, 4, len(data.FEnum3))
		require.Equal(t, pbexternal.Month1(1), data.FEnum3["e1"])
		require.Equal(t, pbexternal.Month1(3), data.FEnum3["e2"])
		require.Equal(t, pbexternal.Month1(0), data.FEnum3["e3"])
		require.Equal(t, pbexternal.Month1(0), data.FEnum3[""])

		require.Equal(t, 4, len(data.FEnum4))
		require.Equal(t, pbexternal.Month2(2), data.FEnum4["e1"])
		require.Equal(t, pbexternal.Month2(4), data.FEnum4["e2"])
		require.Equal(t, pbexternal.Month2(0), data.FEnum4["e3"])
		require.Equal(t, pbexternal.Month2(0), data.FEnum4[""])

		require.Equal(t, 4, len(data.FEnum5))
		require.Equal(t, pbexternal.EnumWeek1_Week(1), data.FEnum5["e1"])
		require.Equal(t, pbexternal.EnumWeek1_Week(3), data.FEnum5["e2"])
		require.Equal(t, pbexternal.EnumWeek1_Week(0), data.FEnum5["e3"])
		require.Equal(t, pbexternal.EnumWeek1_Week(0), data.FEnum5[""])

		require.Equal(t, 4, len(data.FEnum6))
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(2), data.FEnum6["e1"])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(4), data.FEnum6["e2"])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(0), data.FEnum6["e3"])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(0), data.FEnum6[""])

		require.Equal(t, 4, len(data.FEnum7))
		require.Equal(t, pbtypes.EnumCommon1(1), data.FEnum7["e1"])
		require.Equal(t, pbtypes.EnumCommon1(3), data.FEnum7["e2"])
		require.Equal(t, pbtypes.EnumCommon1(0), data.FEnum7["e3"])
		require.Equal(t, pbtypes.EnumCommon1(0), data.FEnum7[""])
	})

	t.Run("WKT", func(t *testing.T) {
		{
			require.Equal(t, 5, len(data.FDuration1))
			require.Equal(t, int64(100), data.FDuration1["d1"].Seconds)
			require.Equal(t, int32(200), data.FDuration1["d1"].Nanos)

			require.Equal(t, int64(101), data.FDuration1["d2"].Seconds)
			require.Equal(t, int32(201), data.FDuration1["d2"].Nanos)

			require.Equal(t, int64(0), data.FDuration1["d3"].Seconds)
			require.Equal(t, int32(0), data.FDuration1["d3"].Nanos)

			require.Equal(t, (*durationpb.Duration)(nil), data.FDuration1["d4"])
			require.Equal(t, (*durationpb.Duration)(nil), data.FDuration1[""])
		}
		{
			require.Equal(t, 5, len(data.FTimestamp1))
			require.Equal(t, int64(100), data.FTimestamp1["t1"].Seconds)
			require.Equal(t, int32(200), data.FTimestamp1["t1"].Nanos)

			require.Equal(t, int64(101), data.FTimestamp1["t2"].Seconds)
			require.Equal(t, int32(201), data.FTimestamp1["t2"].Nanos)

			require.Equal(t, int64(0), data.FTimestamp1["t3"].Seconds)
			require.Equal(t, int32(0), data.FTimestamp1["t3"].Nanos)

			require.Equal(t, (*timestamppb.Timestamp)(nil), data.FTimestamp1["t4"])
			require.Equal(t, (*timestamppb.Timestamp)(nil), data.FTimestamp1[""])
		}
	})

	t.Run("MESSAGE", func(t *testing.T) {
		{
			require.Equal(t, 5, len(data.FMessage1))
			require.NotNil(t, data.FMessage1["m1"])
			require.Equal(t, "a", data.FMessage1["m1"].FString1)
			require.Equal(t, "b", data.FMessage1["m1"].FString2)
			require.Equal(t, "c", data.FMessage1["m1"].FString3)
			require.Equal(t, (*pbtypes.MessageMap1)(nil), data.FMessage1["m2"])
			require.Equal(t, (*pbtypes.MessageMap1)(nil), data.FMessage1["m3"])
			require.Equal(t, (*pbtypes.MessageMap1)(nil), data.FMessage1["m4"])
			require.Equal(t, (*pbtypes.MessageMap1)(nil), data.FMessage1[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage2))
			require.NotNil(t, data.FMessage2["m1"])
			require.Equal(t, "", data.FMessage2["m1"].FString1)
			require.Equal(t, "", data.FMessage2["m1"].FString2)
			require.Equal(t, "", data.FMessage2["m1"].FString3)
			require.Equal(t, (*pbtypes.MessageMap1_Embed1)(nil), data.FMessage2["m2"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1)(nil), data.FMessage2["m3"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1)(nil), data.FMessage2["m4"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1)(nil), data.FMessage2[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage3))
			require.NotNil(t, data.FMessage3["m1"])
			require.Equal(t, "1", data.FMessage3["m1"].FString1)
			require.Equal(t, "2", data.FMessage3["m1"].FString2)
			require.Equal(t, "3", data.FMessage3["m1"].FString3)
			require.Equal(t, (*pbtypes.MessageMap1_Embed1_Embed2)(nil), data.FMessage3["m2"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1_Embed2)(nil), data.FMessage3["m3"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1_Embed2)(nil), data.FMessage3["m4"])
			require.Equal(t, (*pbtypes.MessageMap1_Embed1_Embed2)(nil), data.FMessage3[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage4))
			require.NotNil(t, data.FMessage4["m1"])
			require.Equal(t, "", data.FMessage4["m1"].TString)
			require.Equal(t, (*string)(nil), data.FMessage4["m1"].PString)
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4["m2"])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4["m3"])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4["m4"])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage5))
			require.NotNil(t, data.FMessage5["m1"])
			require.Equal(t, "", data.FMessage5["m1"].TString1)
			require.Equal(t, "", data.FMessage5["m1"].TString2)
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5["m2"])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5["m3"])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5["m4"])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[""])
		}
		{
			require.Equal(t, 5, len(data.FMessage6))
			require.NotNil(t, data.FMessage6["m1"])
			require.Equal(t, "a", data.FMessage6["m1"].FString1)
			require.Equal(t, "b", data.FMessage6["m1"].FString2)
			require.Equal(t, "c", data.FMessage6["m1"].FString3)
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6["m2"])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6["m3"])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6["m4"])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[""])
		}
	})

	t.Run("ANY", func(t *testing.T) {
		require.Equal(t, 9, len(data.FAny1))
		{
			require.NotNil(t, data.FAny1["a1"])
			anyVal := &pbtypes.MessageMap1{}
			err := data.FAny1["a1"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1["a2"])
			anyVal := &pbtypes.MessageMap1_Embed1{}
			err := data.FAny1["a2"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "x", anyVal.FString1)
			require.Equal(t, "y", anyVal.FString2)
			require.Equal(t, "z", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1["a3"])
			anyVal := &pbtypes.MessageMap1_Embed1_Embed2{}
			err := data.FAny1["a3"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "1", anyVal.FString1)
			require.Equal(t, "2", anyVal.FString2)
			require.Equal(t, "3", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1["a4"])
			anyVal := &pbexternal.External1{}
			err := data.FAny1["a4"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString)
			require.Equal(t, (*string)(nil), anyVal.PString)
		}
		{
			require.NotNil(t, data.FAny1["a5"])
			anyVal := &pbexternal.External2_Embed1{}
			err := data.FAny1["a5"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString1)
			require.Equal(t, "", anyVal.TString2)
		}
		{
			require.NotNil(t, data.FAny1["a6"])
			anyVal := &timestamppb.Timestamp{}
			err := data.FAny1["a6"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, int64(0), anyVal.Seconds)
			require.Equal(t, int32(0), anyVal.Nanos)
		}
		{
			require.NotNil(t, data.FAny1["a7"])
			anyVal := &anypb.Any{}
			err := data.FAny1["a7"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TypeUrl)
			require.Equal(t, ([]byte)(nil), anyVal.Value)
		}
		{
			require.NotNil(t, data.FAny1["a8"])
			anyVal := &pbtypes.MessageCommon1{}
			err := data.FAny1["a8"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.Equal(t, (*anypb.Any)(nil), data.FAny1[""])
		}
	})
}

// Test cases the field value is empty.
func Test_TypeMap2_Empty(t *testing.T) {
	data := pbtypes.TypeMap2{
		FString1:    map[string]string{},
		FInt32:      map[string]int32{},
		FInt64:      map[string]int64{},
		FUint32:     map[string]uint32{},
		FUint64:     map[string]uint64{},
		FSint32:     map[string]int32{},
		FSint64:     map[string]int64{},
		FSfixed32:   map[string]int32{},
		FSfixed64:   map[string]int64{},
		FFixed32:    map[string]uint32{},
		FFixed64:    map[string]uint64{},
		FFloat:      map[string]float32{},
		FDouble:     map[string]float64{},
		FBool1:      map[string]bool{},
		FBytes1:     map[string][]byte{},
		FEnum1:      map[string]pbtypes.EnumMap1{},
		FEnum2:      map[string]pbtypes.EnumMap1{},
		FEnum3:      map[string]pbexternal.Month1{},
		FEnum4:      map[string]pbexternal.Month2{},
		FEnum5:      map[string]pbexternal.EnumWeek1_Week{},
		FEnum6:      map[string]pbexternal.EnumWeek2_Embed1_Week{},
		FEnum7:      map[string]pbtypes.EnumCommon1{},
		FDuration1:  map[string]*durationpb.Duration{},
		FTimestamp1: map[string]*timestamppb.Timestamp{},
		FMessage1:   map[string]*pbtypes.MessageMap1{},
		FMessage2:   map[string]*pbtypes.MessageMap1_Embed1{},
		FMessage3:   map[string]*pbtypes.MessageMap1_Embed1_Embed2{},
		FMessage4:   map[string]*pbexternal.External1{},
		FMessage5:   map[string]*pbexternal.External2_Embed1{},
		FMessage6:   map[string]*pbtypes.MessageCommon1{},
		FAny1:       map[string]*anypb.Any{},
	}
	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, 0, len(data.FString1))
		require.NotNil(t, data.FString1)
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, 0, len(data.FInt32))
		require.NotNil(t, data.FInt32)

		require.Equal(t, 0, len(data.FInt64))
		require.NotNil(t, data.FInt64)

		require.Equal(t, 0, len(data.FUint32))
		require.NotNil(t, data.FUint32)

		require.Equal(t, 0, len(data.FUint64))
		require.NotNil(t, data.FUint64)

		require.Equal(t, 0, len(data.FSint32))
		require.NotNil(t, data.FSint32)

		require.Equal(t, 0, len(data.FSint64))
		require.NotNil(t, data.FSint64)

		require.Equal(t, 0, len(data.FSfixed32))
		require.NotNil(t, data.FSfixed32)

		require.Equal(t, 0, len(data.FSfixed64))
		require.NotNil(t, data.FSfixed64)

		require.Equal(t, 0, len(data.FFixed32))
		require.NotNil(t, data.FFixed32)

		require.Equal(t, 0, len(data.FFixed64))
		require.NotNil(t, data.FFixed64)

		require.Equal(t, 0, len(data.FFloat))
		require.NotNil(t, data.FFloat)

		require.Equal(t, 0, len(data.FDouble))
		require.NotNil(t, data.FDouble)
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, 0, len(data.FBool1))
		require.NotNil(t, data.FBool1)
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, 0, len(data.FBytes1))
		require.NotNil(t, data.FBytes1)
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, 0, len(data.FEnum1))
		require.NotNil(t, data.FEnum1)

		require.Equal(t, 0, len(data.FEnum2))
		require.NotNil(t, data.FEnum2)

		require.Equal(t, 0, len(data.FEnum3))
		require.NotNil(t, data.FEnum3)

		require.Equal(t, 0, len(data.FEnum4))
		require.NotNil(t, data.FEnum4)

		require.Equal(t, 0, len(data.FEnum5))
		require.NotNil(t, data.FEnum5)

		require.Equal(t, 0, len(data.FEnum6))
		require.NotNil(t, data.FEnum6)

		require.Equal(t, 0, len(data.FEnum7))
		require.NotNil(t, data.FEnum7)
	})

	t.Run("WKT", func(t *testing.T) {
		require.Equal(t, 0, len(data.FDuration1))
		require.NotNil(t, data.FDuration1)

		require.Equal(t, 0, len(data.FTimestamp1))
		require.NotNil(t, data.FTimestamp1)
	})

	t.Run("MESSAGE", func(t *testing.T) {
		{
			require.Equal(t, 0, len(data.FMessage1))
			require.NotNil(t, data.FMessage1)
		}
		{
			require.Equal(t, 0, len(data.FMessage2))
			require.NotNil(t, data.FMessage2)
		}
		{
			require.Equal(t, 0, len(data.FMessage3))
			require.NotNil(t, data.FMessage3)
		}
		{
			require.Equal(t, 0, len(data.FMessage4))
			require.NotNil(t, data.FMessage4)
		}
		{
			require.Equal(t, 0, len(data.FMessage5))
			require.NotNil(t, data.FMessage5)
		}
		{
			require.Equal(t, 0, len(data.FMessage6))
			require.NotNil(t, data.FMessage6)
		}
	})

	t.Run("ANY", func(t *testing.T) {
		require.Equal(t, 0, len(data.FAny1))
		require.NotNil(t, data.FAny1)
	})
}

// Test cases the field value are have valid value, expected ignore default value.
func Test_TypeMap2_Ignore(t *testing.T) {
	data := pbtypes.TypeMap2{
		FString1:    map[string]string{"sk1": "101"},
		FInt32:      map[string]int32{"101": 101},
		FInt64:      map[string]int64{"102": 102},
		FUint32:     map[string]uint32{"103": 103},
		FUint64:     map[string]uint64{"104": 104},
		FSint32:     map[string]int32{"105": 105},
		FSint64:     map[string]int64{"106": 106},
		FSfixed32:   map[string]int32{"107": 107},
		FSfixed64:   map[string]int64{"108": 108},
		FFixed32:    map[string]uint32{"109": 109},
		FFixed64:    map[string]uint64{"110": 110},
		FFloat:      map[string]float32{"111.111": 111.111},
		FDouble:     map[string]float64{"112.112": 112.112},
		FBool1:      map[string]bool{"bk1": true},
		FBytes1:     map[string][]byte{"bk1": []byte("bytes1")},
		FEnum1:      map[string]pbtypes.EnumMap1{"201": 201},
		FEnum2:      map[string]pbtypes.EnumMap1{"202": 202},
		FEnum3:      map[string]pbexternal.Month1{"203": 203},
		FEnum4:      map[string]pbexternal.Month2{"204": 204},
		FEnum5:      map[string]pbexternal.EnumWeek1_Week{"205": 205},
		FEnum6:      map[string]pbexternal.EnumWeek2_Embed1_Week{"206": 206},
		FEnum7:      map[string]pbtypes.EnumCommon1{"207": 207},
		FDuration1:  map[string]*durationpb.Duration{"dk1": {Seconds: 300, Nanos: 301}},
		FTimestamp1: map[string]*timestamppb.Timestamp{"tk1": {Seconds: 500, Nanos: 501}},
		FMessage1:   map[string]*pbtypes.MessageMap1{"mk1": {FString1: "1101", FString2: "1102", FString3: "1103"}},
		FMessage2:   map[string]*pbtypes.MessageMap1_Embed1{"mk2": {FString1: "1201", FString2: "1202", FString3: "1203"}},
		FMessage3:   map[string]*pbtypes.MessageMap1_Embed1_Embed2{"mk3": {FString1: "1301", FString2: "1302", FString3: "1303"}},
		FMessage4:   map[string]*pbexternal.External1{"mk4": {TString: "1401"}},
		FMessage5:   map[string]*pbexternal.External2_Embed1{"mk5": {TString1: "1501", TString2: "1502"}},
		FMessage6:   map[string]*pbtypes.MessageCommon1{"mk6": {FString1: "1601", FString2: "1602", FString3: "1603"}},
		FAny1:       map[string]*anypb.Any{"ak1": mustNewAny(&pbexternal.External3{TString: "any1"})},
	}
	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, 1, len(data.FString1))
		require.Equal(t, "101", data.FString1["sk1"])
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, 1, len(data.FInt32))
		require.Equal(t, int32(101), data.FInt32["101"])

		require.Equal(t, 1, len(data.FInt64))
		require.Equal(t, int64(102), data.FInt64["102"])

		require.Equal(t, 1, len(data.FUint32))
		require.Equal(t, uint32(103), data.FUint32["103"])

		require.Equal(t, 1, len(data.FUint64))
		require.Equal(t, uint64(104), data.FUint64["104"])

		require.Equal(t, 1, len(data.FSint32))
		require.Equal(t, int32(105), data.FSint32["105"])

		require.Equal(t, 1, len(data.FSint64))
		require.Equal(t, int64(106), data.FSint64["106"])

		require.Equal(t, 1, len(data.FSfixed32))
		require.Equal(t, int32(107), data.FSfixed32["107"])

		require.Equal(t, 1, len(data.FSfixed64))
		require.Equal(t, int64(108), data.FSfixed64["108"])

		require.Equal(t, 1, len(data.FFixed32))
		require.Equal(t, uint32(109), data.FFixed32["109"])

		require.Equal(t, 1, len(data.FFixed64))
		require.Equal(t, uint64(110), data.FFixed64["110"])

		require.Equal(t, 1, len(data.FFloat))
		require.Equal(t, float32(111.111), data.FFloat["111.111"])

		require.Equal(t, 1, len(data.FDouble))
		require.Equal(t, float64(112.112), data.FDouble["112.112"])
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, 1, len(data.FBool1))
		require.Equal(t, true, data.FBool1["bk1"])
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, 1, len(data.FBytes1))
		require.Equal(t, []byte("bytes1"), data.FBytes1["bk1"])
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, 1, len(data.FEnum1))
		require.Equal(t, pbtypes.EnumMap1(201), data.FEnum1["201"])

		require.Equal(t, 1, len(data.FEnum2))
		require.Equal(t, pbtypes.EnumMap1(202), data.FEnum2["202"])

		require.Equal(t, 1, len(data.FEnum3))
		require.Equal(t, pbexternal.Month1(203), data.FEnum3["203"])

		require.Equal(t, 1, len(data.FEnum4))
		require.Equal(t, pbexternal.Month2(204), data.FEnum4["204"])

		require.Equal(t, 1, len(data.FEnum5))
		require.Equal(t, pbexternal.EnumWeek1_Week(205), data.FEnum5["205"])

		require.Equal(t, 1, len(data.FEnum6))
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(206), data.FEnum6["206"])

		require.Equal(t, 1, len(data.FEnum7))
		require.Equal(t, pbtypes.EnumCommon1(207), data.FEnum7["207"])
	})

	t.Run("WKT", func(t *testing.T) {
		{
			require.Equal(t, 1, len(data.FDuration1))
			require.Equal(t, int64(300), data.FDuration1["dk1"].Seconds)
			require.Equal(t, int32(301), data.FDuration1["dk1"].Nanos)
		}
		{
			require.Equal(t, 1, len(data.FTimestamp1))
			require.Equal(t, int64(500), data.FTimestamp1["tk1"].Seconds)
			require.Equal(t, int32(501), data.FTimestamp1["tk1"].Nanos)
		}
	})

	t.Run("MESSAGE", func(t *testing.T) {
		{
			require.Equal(t, 1, len(data.FMessage1))
			require.NotNil(t, data.FMessage1["mk1"])
			require.Equal(t, "1101", data.FMessage1["mk1"].FString1)
			require.Equal(t, "1102", data.FMessage1["mk1"].FString2)
			require.Equal(t, "1103", data.FMessage1["mk1"].FString3)
		}
		{
			require.Equal(t, 1, len(data.FMessage2))
			require.NotNil(t, data.FMessage2["mk2"])
			require.Equal(t, "1201", data.FMessage2["mk2"].FString1)
			require.Equal(t, "1202", data.FMessage2["mk2"].FString2)
			require.Equal(t, "1203", data.FMessage2["mk2"].FString3)
		}
		{
			require.Equal(t, 1, len(data.FMessage3))
			require.NotNil(t, data.FMessage3["mk3"])
			require.Equal(t, "1301", data.FMessage3["mk3"].FString1)
			require.Equal(t, "1302", data.FMessage3["mk3"].FString2)
			require.Equal(t, "1303", data.FMessage3["mk3"].FString3)
		}
		{
			require.Equal(t, 1, len(data.FMessage4))
			require.NotNil(t, data.FMessage4["mk4"])
			require.Equal(t, "1401", data.FMessage4["mk4"].TString)
			require.Equal(t, (*string)(nil), data.FMessage4["mk4"].PString)
		}
		{
			require.Equal(t, 1, len(data.FMessage5))
			require.NotNil(t, data.FMessage5["mk5"])
			require.Equal(t, "1501", data.FMessage5["mk5"].TString1)
			require.Equal(t, "1502", data.FMessage5["mk5"].TString2)
		}
		{
			require.Equal(t, 1, len(data.FMessage6))
			require.NotNil(t, data.FMessage6["mk6"])
			require.Equal(t, "1601", data.FMessage6["mk6"].FString1)
			require.Equal(t, "1602", data.FMessage6["mk6"].FString2)
			require.Equal(t, "1603", data.FMessage6["mk6"].FString3)
		}
	})

	t.Run("ANY", func(t *testing.T) {
		require.Equal(t, 1, len(data.FAny1))
		{
			require.NotNil(t, data.FAny1["ak1"])
			anyVal := &pbexternal.External3{}
			err := data.FAny1["ak1"].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any1", anyVal.TString)
		}
	})
}

// Test cases the field value is nil pointer.
func Test_TypeMap3_NIL(t *testing.T) {
	data := pbtypes.TypeMap3{}
	data.SetDefault()

	t.Run("f_string1", func(t *testing.T) {
		require.Equal(t, 4, len(data.FString1))
		require.Equal(t, "v1", data.FString1["s1"])
		require.Equal(t, "v2", data.FString1["s2"])
		require.Equal(t, "", data.FString1["s3"])
		require.Equal(t, "", data.FString1[""])
	})

	t.Run("f_int32", func(t *testing.T) {
		require.Equal(t, 4, len(data.FInt32))
		require.Equal(t, int32(11), data.FInt32[11])
		require.Equal(t, int32(12), data.FInt32[12])
		require.Equal(t, int32(0), data.FInt32[13])
		require.Equal(t, int32(0), data.FInt32[0])
	})

	t.Run("f_int64", func(t *testing.T) {
		require.Equal(t, 4, len(data.FInt64))
		require.Equal(t, int64(21), data.FInt64[21])
		require.Equal(t, int64(22), data.FInt64[22])
		require.Equal(t, int64(0), data.FInt64[23])
		require.Equal(t, int64(0), data.FInt64[0])
	})

	t.Run("f_uint32", func(t *testing.T) {
		require.Equal(t, 4, len(data.FUint32))
		require.Equal(t, uint32(31), data.FUint32[31])
		require.Equal(t, uint32(32), data.FUint32[32])
		require.Equal(t, uint32(0), data.FUint32[33])
		require.Equal(t, uint32(0), data.FUint32[0])
	})

	t.Run("f_uint64", func(t *testing.T) {
		require.Equal(t, 4, len(data.FUint64))
		require.Equal(t, uint64(41), data.FUint64[41])
		require.Equal(t, uint64(42), data.FUint64[42])
		require.Equal(t, uint64(0), data.FUint64[43])
		require.Equal(t, uint64(0), data.FUint64[0])
	})

	t.Run("f_sint32", func(t *testing.T) {
		require.Equal(t, 4, len(data.FSint32))
		require.Equal(t, int32(51), data.FSint32[51])
		require.Equal(t, int32(52), data.FSint32[52])
		require.Equal(t, int32(0), data.FSint32[53])
		require.Equal(t, int32(0), data.FSint32[0])
	})

	t.Run("f_sint64", func(t *testing.T) {
		require.Equal(t, 4, len(data.FSint64))
		require.Equal(t, int64(61), data.FSint64[61])
		require.Equal(t, int64(62), data.FSint64[62])
		require.Equal(t, int64(0), data.FSint64[63])
		require.Equal(t, int64(0), data.FSint64[0])
	})

	t.Run("f_sfixed32", func(t *testing.T) {
		require.Equal(t, 4, len(data.FSfixed32))
		require.Equal(t, int32(71), data.FSfixed32[71])
		require.Equal(t, int32(72), data.FSfixed32[72])
		require.Equal(t, int32(0), data.FSfixed32[73])
		require.Equal(t, int32(0), data.FSfixed32[0])
	})

	t.Run("f_sfixed64", func(t *testing.T) {
		require.Equal(t, 4, len(data.FSfixed64))
		require.Equal(t, int64(81), data.FSfixed64[81])
		require.Equal(t, int64(82), data.FSfixed64[82])
		require.Equal(t, int64(0), data.FSfixed64[83])
		require.Equal(t, int64(0), data.FSfixed64[0])
	})

	t.Run("f_fixed32", func(t *testing.T) {
		require.Equal(t, 4, len(data.FFixed32))
		require.Equal(t, uint32(91), data.FFixed32[91])
		require.Equal(t, uint32(92), data.FFixed32[92])
		require.Equal(t, uint32(0), data.FFixed32[93])
		require.Equal(t, uint32(0), data.FFixed32[0])
	})

	t.Run("f_fixed64", func(t *testing.T) {
		require.Equal(t, 4, len(data.FFixed64))
		require.Equal(t, uint64(101), data.FFixed64[101])
		require.Equal(t, uint64(102), data.FFixed64[102])
		require.Equal(t, uint64(0), data.FFixed64[103])
		require.Equal(t, uint64(0), data.FFixed64[0])
	})

	t.Run("f_bool1", func(t *testing.T) {
		require.Equal(t, 2, len(data.FBool1))
		require.Equal(t, true, data.FBool1[true])
		require.Equal(t, false, data.FBool1[false])
	})
}

// Test cases the field value is empty.
func Test_TypeMap3_Empty(t *testing.T) {
	data := pbtypes.TypeMap3{
		FString1:  map[string]string{},
		FInt32:    map[int32]int32{},
		FInt64:    map[int64]int64{},
		FUint32:   map[uint32]uint32{},
		FUint64:   map[uint64]uint64{},
		FSint32:   map[int32]int32{},
		FSint64:   map[int64]int64{},
		FSfixed32: map[int32]int32{},
		FSfixed64: map[int64]int64{},
		FFixed32:  map[uint32]uint32{},
		FFixed64:  map[uint64]uint64{},
		FBool1:    map[bool]bool{},
	}
	data.SetDefault()

	t.Run("f_string1", func(t *testing.T) {
		require.Equal(t, 0, len(data.FString1))
		require.NotNil(t, data.FString1)
	})

	t.Run("f_int32", func(t *testing.T) {
		require.Equal(t, 4, len(data.FInt32))
		require.Equal(t, int32(11), data.FInt32[11])
		require.Equal(t, int32(12), data.FInt32[12])
		require.Equal(t, int32(0), data.FInt32[13])
		require.Equal(t, int32(0), data.FInt32[0])
	})

	t.Run("f_int64", func(t *testing.T) {
		require.Equal(t, 0, len(data.FInt64))
		require.NotNil(t, data.FInt64)
	})

	t.Run("f_uint32", func(t *testing.T) {
		require.Equal(t, 4, len(data.FUint32))
		require.Equal(t, uint32(31), data.FUint32[31])
		require.Equal(t, uint32(32), data.FUint32[32])
		require.Equal(t, uint32(0), data.FUint32[33])
		require.Equal(t, uint32(0), data.FUint32[0])
	})

	t.Run("f_uint64", func(t *testing.T) {
		require.Equal(t, 0, len(data.FUint64))
		require.NotNil(t, data.FUint64)
	})

	t.Run("f_sint32", func(t *testing.T) {
		require.Equal(t, 4, len(data.FSint32))
		require.Equal(t, int32(51), data.FSint32[51])
		require.Equal(t, int32(52), data.FSint32[52])
		require.Equal(t, int32(0), data.FSint32[53])
		require.Equal(t, int32(0), data.FSint32[0])
	})

	t.Run("f_sint64", func(t *testing.T) {
		require.Equal(t, 0, len(data.FSint64))
		require.NotNil(t, data.FSint64)
	})

	t.Run("f_sfixed32", func(t *testing.T) {
		require.Equal(t, 4, len(data.FSfixed32))
		require.Equal(t, int32(71), data.FSfixed32[71])
		require.Equal(t, int32(72), data.FSfixed32[72])
		require.Equal(t, int32(0), data.FSfixed32[73])
		require.Equal(t, int32(0), data.FSfixed32[0])
	})

	t.Run("f_sfixed64", func(t *testing.T) {
		require.Equal(t, 0, len(data.FSfixed64))
		require.NotNil(t, data.FSfixed64)
	})

	t.Run("f_fixed32", func(t *testing.T) {
		require.Equal(t, 4, len(data.FFixed32))
		require.Equal(t, uint32(91), data.FFixed32[91])
		require.Equal(t, uint32(92), data.FFixed32[92])
		require.Equal(t, uint32(0), data.FFixed32[93])
		require.Equal(t, uint32(0), data.FFixed32[0])
	})

	t.Run("f_fixed64", func(t *testing.T) {
		require.Equal(t, 0, len(data.FFixed64))
		require.NotNil(t, data.FFixed64)
	})

	t.Run("f_bool1", func(t *testing.T) {
		require.Equal(t, 2, len(data.FBool1))
		require.Equal(t, true, data.FBool1[true])
		require.Equal(t, false, data.FBool1[false])
	})
}

// Test cases the field value are have valid value, expected ignore default value.
func Test_TypeMap3_Ignore(t *testing.T) {
	data := pbtypes.TypeMap3{
		FString1:  map[string]string{"sk1": "101"},
		FInt32:    map[int32]int32{101: 101},
		FInt64:    map[int64]int64{102: 102},
		FUint32:   map[uint32]uint32{103: 103},
		FUint64:   map[uint64]uint64{104: 104},
		FSint32:   map[int32]int32{105: 105},
		FSint64:   map[int64]int64{106: 106},
		FSfixed32: map[int32]int32{107: 107},
		FSfixed64: map[int64]int64{108: 108},
		FFixed32:  map[uint32]uint32{109: 109},
		FFixed64:  map[uint64]uint64{110: 110},
		FBool1:    map[bool]bool{true: true},
	}
	data.SetDefault()

	t.Run("f_string1", func(t *testing.T) {
		require.Equal(t, 1, len(data.FString1))
		require.Equal(t, "101", data.FString1["sk1"])
	})

	t.Run("f_int32", func(t *testing.T) {
		require.Equal(t, 1, len(data.FInt32))
		require.Equal(t, int32(101), data.FInt32[101])
	})

	t.Run("f_int64", func(t *testing.T) {
		require.Equal(t, 1, len(data.FInt64))
		require.Equal(t, int64(102), data.FInt64[102])
	})

	t.Run("f_uint32", func(t *testing.T) {
		require.Equal(t, 1, len(data.FUint32))
		require.Equal(t, uint32(103), data.FUint32[103])
	})

	t.Run("f_uint64", func(t *testing.T) {
		require.Equal(t, 1, len(data.FUint64))
		require.Equal(t, uint64(104), data.FUint64[104])
	})

	t.Run("f_sint32", func(t *testing.T) {
		require.Equal(t, 1, len(data.FSint32))
		require.Equal(t, int32(105), data.FSint32[105])
	})

	t.Run("f_sint64", func(t *testing.T) {
		require.Equal(t, 1, len(data.FSint64))
		require.Equal(t, int64(106), data.FSint64[106])
	})

	t.Run("f_sfixed32", func(t *testing.T) {
		require.Equal(t, 1, len(data.FSfixed32))
		require.Equal(t, int32(107), data.FSfixed32[107])
	})

	t.Run("f_sfixed64", func(t *testing.T) {
		require.Equal(t, 1, len(data.FSfixed64))
		require.Equal(t, int64(108), data.FSfixed64[108])
	})

	t.Run("f_fixed32", func(t *testing.T) {
		require.Equal(t, 1, len(data.FFixed32))
		require.Equal(t, uint32(109), data.FFixed32[109])
	})

	t.Run("f_fixed64", func(t *testing.T) {
		require.Equal(t, 1, len(data.FFixed64))
		require.Equal(t, uint64(110), data.FFixed64[110])
	})

	t.Run("f_bool1", func(t *testing.T) {
		require.Equal(t, 1, len(data.FBool1))
		require.Equal(t, true, data.FBool1[true])
	})
}
