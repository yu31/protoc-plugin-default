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
func Test_TypeRepeated1_NIL(t *testing.T) {
	data := pbtypes.TypeRepeated1{}
	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, 11, len(data.FString1))
		require.Equal(t, "ts1", data.FString1[0])
		require.Equal(t, "", data.FString1[1])
		require.Equal(t, `""`, data.FString1[2])
		require.Equal(t, `"ts4"`, data.FString1[3])
		require.Equal(t, `"ts5"`, data.FString1[4])
		require.Equal(t, `"ts"6"`, data.FString1[5])
		require.Equal(t, `"ts"7"`, data.FString1[6])
		require.Equal(t, `[ts8]`, data.FString1[7])
		require.Equal(t, `{ts9}`, data.FString1[8])
		require.Equal(t, "", data.FString1[9])
		require.Equal(t, "", data.FString1[10])
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, 5, len(data.FInt32))
		require.Equal(t, int32(11), data.FInt32[0])
		require.Equal(t, int32(12), data.FInt32[1])
		require.Equal(t, int32(13), data.FInt32[2])
		require.Equal(t, int32(0), data.FInt32[3])
		require.Equal(t, int32(0), data.FInt32[4])

		require.Equal(t, 5, len(data.FInt64))
		require.Equal(t, int64(21), data.FInt64[0])
		require.Equal(t, int64(22), data.FInt64[1])
		require.Equal(t, int64(23), data.FInt64[2])
		require.Equal(t, int64(0), data.FInt64[3])
		require.Equal(t, int64(0), data.FInt64[4])

		require.Equal(t, 5, len(data.FUint32))
		require.Equal(t, uint32(31), data.FUint32[0])
		require.Equal(t, uint32(32), data.FUint32[1])
		require.Equal(t, uint32(33), data.FUint32[2])
		require.Equal(t, uint32(0), data.FUint32[3])
		require.Equal(t, uint32(0), data.FUint32[4])

		require.Equal(t, 5, len(data.FUint64))
		require.Equal(t, uint64(41), data.FUint64[0])
		require.Equal(t, uint64(42), data.FUint64[1])
		require.Equal(t, uint64(43), data.FUint64[2])
		require.Equal(t, uint64(0), data.FUint64[3])
		require.Equal(t, uint64(0), data.FUint64[4])

		require.Equal(t, 5, len(data.FSint32))
		require.Equal(t, int32(51), data.FSint32[0])
		require.Equal(t, int32(52), data.FSint32[1])
		require.Equal(t, int32(53), data.FSint32[2])
		require.Equal(t, int32(0), data.FSint32[3])
		require.Equal(t, int32(0), data.FSint32[4])

		require.Equal(t, 5, len(data.FSint64))
		require.Equal(t, int64(61), data.FSint64[0])
		require.Equal(t, int64(62), data.FSint64[1])
		require.Equal(t, int64(63), data.FSint64[2])
		require.Equal(t, int64(0), data.FSint64[3])
		require.Equal(t, int64(0), data.FSint64[4])

		require.Equal(t, 5, len(data.FSfixed32))
		require.Equal(t, int32(71), data.FSfixed32[0])
		require.Equal(t, int32(72), data.FSfixed32[1])
		require.Equal(t, int32(73), data.FSfixed32[2])
		require.Equal(t, int32(0), data.FSfixed32[3])
		require.Equal(t, int32(0), data.FSfixed32[4])

		require.Equal(t, 5, len(data.FSfixed64))
		require.Equal(t, int64(81), data.FSfixed64[0])
		require.Equal(t, int64(82), data.FSfixed64[1])
		require.Equal(t, int64(83), data.FSfixed64[2])
		require.Equal(t, int64(0), data.FSfixed64[3])
		require.Equal(t, int64(0), data.FSfixed64[4])

		require.Equal(t, 5, len(data.FFixed32))
		require.Equal(t, uint32(91), data.FFixed32[0])
		require.Equal(t, uint32(92), data.FFixed32[1])
		require.Equal(t, uint32(93), data.FFixed32[2])
		require.Equal(t, uint32(0), data.FFixed32[3])
		require.Equal(t, uint32(0), data.FFixed32[4])

		require.Equal(t, 5, len(data.FFixed64))
		require.Equal(t, uint64(101), data.FFixed64[0])
		require.Equal(t, uint64(102), data.FFixed64[1])
		require.Equal(t, uint64(103), data.FFixed64[2])
		require.Equal(t, uint64(0), data.FFixed64[3])
		require.Equal(t, uint64(0), data.FFixed64[4])

		require.Equal(t, 5, len(data.FFloat))
		require.Equal(t, float32(111.111), data.FFloat[0])
		require.Equal(t, float32(112.112), data.FFloat[1])
		require.Equal(t, float32(113.113), data.FFloat[2])
		require.Equal(t, float32(0), data.FFloat[3])
		require.Equal(t, float32(0), data.FFloat[4])

		require.Equal(t, 5, len(data.FDouble))
		require.Equal(t, float64(121.121), data.FDouble[0])
		require.Equal(t, float64(122.122), data.FDouble[1])
		require.Equal(t, float64(123.123), data.FDouble[2])
		require.Equal(t, float64(0), data.FDouble[3])
		require.Equal(t, float64(0), data.FDouble[4])
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, 5, len(data.FBool1))
		require.Equal(t, true, data.FBool1[0])
		require.Equal(t, false, data.FBool1[1])
		require.Equal(t, true, data.FBool1[2])
		require.Equal(t, false, data.FBool1[3])
		require.Equal(t, false, data.FBool1[4])
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, 5, len(data.FBytes1))
		require.Equal(t, []byte("abc"), data.FBytes1[0])
		require.Equal(t, []byte(""), data.FBytes1[1])
		require.Equal(t, []byte("123"), data.FBytes1[2])
		require.Equal(t, []byte(""), data.FBytes1[3])
		require.Equal(t, []byte(""), data.FBytes1[4])
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, 5, len(data.FEnum1))
		require.Equal(t, pbtypes.EnumRepeated1(1), data.FEnum1[0])
		require.Equal(t, pbtypes.EnumRepeated1(3), data.FEnum1[1])
		require.Equal(t, pbtypes.EnumRepeated1(5), data.FEnum1[2])
		require.Equal(t, pbtypes.EnumRepeated1(0), data.FEnum1[3])
		require.Equal(t, pbtypes.EnumRepeated1(0), data.FEnum1[4])

		require.Equal(t, 5, len(data.FEnum2))
		require.Equal(t, pbtypes.EnumRepeated1(2), data.FEnum2[0])
		require.Equal(t, pbtypes.EnumRepeated1(4), data.FEnum2[1])
		require.Equal(t, pbtypes.EnumRepeated1(6), data.FEnum2[2])
		require.Equal(t, pbtypes.EnumRepeated1(0), data.FEnum2[3])
		require.Equal(t, pbtypes.EnumRepeated1(0), data.FEnum2[4])

		require.Equal(t, 5, len(data.FEnum3))
		require.Equal(t, pbexternal.Month1(1), data.FEnum3[0])
		require.Equal(t, pbexternal.Month1(3), data.FEnum3[1])
		require.Equal(t, pbexternal.Month1(5), data.FEnum3[2])
		require.Equal(t, pbexternal.Month1(0), data.FEnum3[3])
		require.Equal(t, pbexternal.Month1(0), data.FEnum3[4])

		require.Equal(t, 5, len(data.FEnum4))
		require.Equal(t, pbexternal.Month2(2), data.FEnum4[0])
		require.Equal(t, pbexternal.Month2(4), data.FEnum4[1])
		require.Equal(t, pbexternal.Month2(6), data.FEnum4[2])
		require.Equal(t, pbexternal.Month2(0), data.FEnum4[3])
		require.Equal(t, pbexternal.Month2(0), data.FEnum4[4])

		require.Equal(t, 5, len(data.FEnum5))
		require.Equal(t, pbexternal.EnumWeek1_Week(1), data.FEnum5[0])
		require.Equal(t, pbexternal.EnumWeek1_Week(3), data.FEnum5[1])
		require.Equal(t, pbexternal.EnumWeek1_Week(5), data.FEnum5[2])
		require.Equal(t, pbexternal.EnumWeek1_Week(0), data.FEnum5[3])
		require.Equal(t, pbexternal.EnumWeek1_Week(0), data.FEnum5[4])

		require.Equal(t, 5, len(data.FEnum6))
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(2), data.FEnum6[0])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(4), data.FEnum6[1])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(6), data.FEnum6[2])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(0), data.FEnum6[3])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(0), data.FEnum6[4])

		require.Equal(t, 5, len(data.FEnum7))
		require.Equal(t, pbtypes.EnumCommon1(2), data.FEnum7[0])
		require.Equal(t, pbtypes.EnumCommon1(4), data.FEnum7[1])
		require.Equal(t, pbtypes.EnumCommon1(6), data.FEnum7[2])
		require.Equal(t, pbtypes.EnumCommon1(0), data.FEnum7[3])
		require.Equal(t, pbtypes.EnumCommon1(0), data.FEnum7[4])
	})

	t.Run("WKT", func(t *testing.T) {
		{
			require.Equal(t, 6, len(data.FDuration1))
			require.Equal(t, int64(100), data.FDuration1[0].Seconds)
			require.Equal(t, int32(200), data.FDuration1[0].Nanos)

			require.Equal(t, int64(0), data.FDuration1[1].Seconds)
			require.Equal(t, int32(0), data.FDuration1[1].Nanos)

			require.Equal(t, int64(102), data.FDuration1[2].Seconds)
			require.Equal(t, int32(202), data.FDuration1[2].Nanos)

			require.Equal(t, int64(0), data.FDuration1[3].Seconds)
			require.Equal(t, int32(0), data.FDuration1[3].Nanos)

			require.Equal(t, (*durationpb.Duration)(nil), data.FDuration1[4])
			require.Equal(t, (*durationpb.Duration)(nil), data.FDuration1[5])
		}
		{
			require.Equal(t, 6, len(data.FTimestamp1))
			require.Equal(t, int64(200), data.FTimestamp1[0].Seconds)
			require.Equal(t, int32(300), data.FTimestamp1[0].Nanos)

			require.Equal(t, int64(0), data.FTimestamp1[1].Seconds)
			require.Equal(t, int32(0), data.FTimestamp1[1].Nanos)

			require.Equal(t, int64(202), data.FTimestamp1[2].Seconds)
			require.Equal(t, int32(302), data.FTimestamp1[2].Nanos)

			require.Equal(t, int64(0), data.FTimestamp1[3].Seconds)
			require.Equal(t, int32(0), data.FTimestamp1[3].Nanos)

			require.Equal(t, (*timestamppb.Timestamp)(nil), data.FTimestamp1[4])
			require.Equal(t, (*timestamppb.Timestamp)(nil), data.FTimestamp1[5])
		}
	})

	t.Run("MESSAGE", func(t *testing.T) {
		{
			require.Equal(t, 6, len(data.FMessage1))
			require.NotNil(t, data.FMessage1[0])
			require.Equal(t, "a", data.FMessage1[0].FString1)
			require.Equal(t, "b", data.FMessage1[0].FString2)
			require.Equal(t, "c", data.FMessage1[0].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1)(nil), data.FMessage1[1])
			require.NotNil(t, data.FMessage1[2])
			require.Equal(t, "a", data.FMessage1[2].FString1)
			require.Equal(t, "b", data.FMessage1[2].FString2)
			require.Equal(t, "c", data.FMessage1[2].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1)(nil), data.FMessage1[3])
			require.Equal(t, (*pbtypes.MessageRepeated1)(nil), data.FMessage1[4])
			require.Equal(t, (*pbtypes.MessageRepeated1)(nil), data.FMessage1[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage2))
			require.NotNil(t, data.FMessage2[0])
			require.Equal(t, "", data.FMessage2[0].FString1)
			require.Equal(t, "", data.FMessage2[0].FString2)
			require.Equal(t, "", data.FMessage2[0].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1)(nil), data.FMessage2[1])
			require.NotNil(t, data.FMessage2[2])
			require.Equal(t, "", data.FMessage2[2].FString1)
			require.Equal(t, "", data.FMessage2[2].FString2)
			require.Equal(t, "", data.FMessage2[2].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1)(nil), data.FMessage2[3])
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1)(nil), data.FMessage2[4])
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1)(nil), data.FMessage2[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage3))
			require.NotNil(t, data.FMessage3[0])
			require.Equal(t, "1", data.FMessage3[0].FString1)
			require.Equal(t, "2", data.FMessage3[0].FString2)
			require.Equal(t, "3", data.FMessage3[0].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1_Embed2)(nil), data.FMessage3[1])
			require.NotNil(t, data.FMessage3[2])
			require.Equal(t, "1", data.FMessage3[2].FString1)
			require.Equal(t, "2", data.FMessage3[2].FString2)
			require.Equal(t, "3", data.FMessage3[2].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1_Embed2)(nil), data.FMessage3[3])
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1_Embed2)(nil), data.FMessage3[4])
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1_Embed2)(nil), data.FMessage3[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage4))
			require.NotNil(t, data.FMessage4[0])
			require.Equal(t, "", data.FMessage4[0].TString)
			require.Equal(t, (*string)(nil), data.FMessage4[0].PString)
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[1])
			require.NotNil(t, data.FMessage4[2])
			require.Equal(t, "", data.FMessage4[2].TString)
			require.Equal(t, (*string)(nil), data.FMessage4[2].PString)
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[3])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[4])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage5))
			require.NotNil(t, data.FMessage5[0])
			require.Equal(t, "", data.FMessage5[0].TString1)
			require.Equal(t, "", data.FMessage5[0].TString2)
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[1])
			require.NotNil(t, data.FMessage5[2])
			require.Equal(t, "", data.FMessage5[2].TString1)
			require.Equal(t, "", data.FMessage5[2].TString2)
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[3])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[4])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage6))
			require.NotNil(t, data.FMessage6[0])
			require.Equal(t, "a", data.FMessage6[0].FString1)
			require.Equal(t, "b", data.FMessage6[0].FString2)
			require.Equal(t, "c", data.FMessage6[0].FString3)
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[1])
			require.NotNil(t, data.FMessage6[2])
			require.Equal(t, "a", data.FMessage6[2].FString1)
			require.Equal(t, "b", data.FMessage6[2].FString2)
			require.Equal(t, "c", data.FMessage6[2].FString3)
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[3])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[4])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[5])
		}
	})

	t.Run("ANY", func(t *testing.T) {
		require.Equal(t, 11, len(data.FAny1))
		{
			require.NotNil(t, data.FAny1[0])
			anyVal := &pbtypes.MessageRepeated1{}
			err := data.FAny1[0].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1[1])
			anyVal := &pbtypes.MessageRepeated1_Embed1{}
			err := data.FAny1[1].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "x", anyVal.FString1)
			require.Equal(t, "y", anyVal.FString2)
			require.Equal(t, "z", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1[2])
			anyVal := &pbtypes.MessageRepeated1_Embed1_Embed2{}
			err := data.FAny1[2].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "1", anyVal.FString1)
			require.Equal(t, "2", anyVal.FString2)
			require.Equal(t, "3", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1[3])
			anyVal := &pbexternal.External1{}
			err := data.FAny1[3].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString)
			require.Equal(t, (*string)(nil), anyVal.PString)
		}
		{
			require.NotNil(t, data.FAny1[4])
			anyVal := &pbexternal.External2_Embed1{}
			err := data.FAny1[4].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString1)
			require.Equal(t, "", anyVal.TString2)
		}
		{
			require.NotNil(t, data.FAny1[5])
			anyVal := &timestamppb.Timestamp{}
			err := data.FAny1[5].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, int64(0), anyVal.Seconds)
			require.Equal(t, int32(0), anyVal.Nanos)
		}
		{
			require.NotNil(t, data.FAny1[6])
			anyVal := &anypb.Any{}
			err := data.FAny1[6].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TypeUrl)
			require.Equal(t, ([]byte)(nil), anyVal.Value)
		}
		{
			require.NotNil(t, data.FAny1[7])
			anyVal := &pbtypes.MessageCommon1{}
			err := data.FAny1[7].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.Equal(t, (*anypb.Any)(nil), data.FAny1[8])
		}
		{
			require.Equal(t, (*anypb.Any)(nil), data.FAny1[9])
		}
		{
			require.Equal(t, (*anypb.Any)(nil), data.FAny1[10])
		}
	})
}

// Test cases the field value is empty.
func Test_TypeRepeated1_Empty(t *testing.T) {
	data := pbtypes.TypeRepeated1{
		FString1:    []string{},
		FInt32:      []int32{},
		FInt64:      []int64{},
		FUint32:     []uint32{},
		FUint64:     []uint64{},
		FSint32:     []int32{},
		FSint64:     []int64{},
		FSfixed32:   []int32{},
		FSfixed64:   []int64{},
		FFixed32:    []uint32{},
		FFixed64:    []uint64{},
		FFloat:      []float32{},
		FDouble:     []float64{},
		FBool1:      []bool{},
		FBytes1:     [][]byte{},
		FEnum1:      []pbtypes.EnumRepeated1{},
		FEnum2:      []pbtypes.EnumRepeated1{},
		FEnum3:      []pbexternal.Month1{},
		FEnum4:      []pbexternal.Month2{},
		FEnum5:      []pbexternal.EnumWeek1_Week{},
		FEnum6:      []pbexternal.EnumWeek2_Embed1_Week{},
		FEnum7:      []pbtypes.EnumCommon1{},
		FDuration1:  []*durationpb.Duration{},
		FTimestamp1: []*timestamppb.Timestamp{},
		FMessage1:   []*pbtypes.MessageRepeated1{},
		FMessage2:   []*pbtypes.MessageRepeated1_Embed1{},
		FMessage3:   []*pbtypes.MessageRepeated1_Embed1_Embed2{},
		FMessage4:   []*pbexternal.External1{},
		FMessage5:   []*pbexternal.External2_Embed1{},
		FMessage6:   []*pbtypes.MessageCommon1{},
		FAny1:       []*anypb.Any{},
	}
	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, 11, len(data.FString1))
		require.Equal(t, "ts1", data.FString1[0])
		require.Equal(t, "", data.FString1[1])
		require.Equal(t, `""`, data.FString1[2])
		require.Equal(t, `"ts4"`, data.FString1[3])
		require.Equal(t, `"ts5"`, data.FString1[4])
		require.Equal(t, `"ts"6"`, data.FString1[5])
		require.Equal(t, `"ts"7"`, data.FString1[6])
		require.Equal(t, `[ts8]`, data.FString1[7])
		require.Equal(t, `{ts9}`, data.FString1[8])
		require.Equal(t, "", data.FString1[9])
		require.Equal(t, "", data.FString1[10])
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, 5, len(data.FInt32))
		require.Equal(t, int32(11), data.FInt32[0])
		require.Equal(t, int32(12), data.FInt32[1])
		require.Equal(t, int32(13), data.FInt32[2])
		require.Equal(t, int32(0), data.FInt32[3])
		require.Equal(t, int32(0), data.FInt32[4])

		require.Equal(t, 5, len(data.FInt64))
		require.Equal(t, int64(21), data.FInt64[0])
		require.Equal(t, int64(22), data.FInt64[1])
		require.Equal(t, int64(23), data.FInt64[2])
		require.Equal(t, int64(0), data.FInt64[3])
		require.Equal(t, int64(0), data.FInt64[4])

		require.Equal(t, 5, len(data.FUint32))
		require.Equal(t, uint32(31), data.FUint32[0])
		require.Equal(t, uint32(32), data.FUint32[1])
		require.Equal(t, uint32(33), data.FUint32[2])
		require.Equal(t, uint32(0), data.FUint32[3])
		require.Equal(t, uint32(0), data.FUint32[4])

		require.Equal(t, 5, len(data.FUint64))
		require.Equal(t, uint64(41), data.FUint64[0])
		require.Equal(t, uint64(42), data.FUint64[1])
		require.Equal(t, uint64(43), data.FUint64[2])
		require.Equal(t, uint64(0), data.FUint64[3])
		require.Equal(t, uint64(0), data.FUint64[4])

		require.Equal(t, 5, len(data.FSint32))
		require.Equal(t, int32(51), data.FSint32[0])
		require.Equal(t, int32(52), data.FSint32[1])
		require.Equal(t, int32(53), data.FSint32[2])
		require.Equal(t, int32(0), data.FSint32[3])
		require.Equal(t, int32(0), data.FSint32[4])

		require.Equal(t, 5, len(data.FSint64))
		require.Equal(t, int64(61), data.FSint64[0])
		require.Equal(t, int64(62), data.FSint64[1])
		require.Equal(t, int64(63), data.FSint64[2])
		require.Equal(t, int64(0), data.FSint64[3])
		require.Equal(t, int64(0), data.FSint64[4])

		require.Equal(t, 5, len(data.FSfixed32))
		require.Equal(t, int32(71), data.FSfixed32[0])
		require.Equal(t, int32(72), data.FSfixed32[1])
		require.Equal(t, int32(73), data.FSfixed32[2])
		require.Equal(t, int32(0), data.FSfixed32[3])
		require.Equal(t, int32(0), data.FSfixed32[4])

		require.Equal(t, 5, len(data.FSfixed64))
		require.Equal(t, int64(81), data.FSfixed64[0])
		require.Equal(t, int64(82), data.FSfixed64[1])
		require.Equal(t, int64(83), data.FSfixed64[2])
		require.Equal(t, int64(0), data.FSfixed64[3])
		require.Equal(t, int64(0), data.FSfixed64[4])

		require.Equal(t, 5, len(data.FFixed32))
		require.Equal(t, uint32(91), data.FFixed32[0])
		require.Equal(t, uint32(92), data.FFixed32[1])
		require.Equal(t, uint32(93), data.FFixed32[2])
		require.Equal(t, uint32(0), data.FFixed32[3])
		require.Equal(t, uint32(0), data.FFixed32[4])

		require.Equal(t, 5, len(data.FFixed64))
		require.Equal(t, uint64(101), data.FFixed64[0])
		require.Equal(t, uint64(102), data.FFixed64[1])
		require.Equal(t, uint64(103), data.FFixed64[2])
		require.Equal(t, uint64(0), data.FFixed64[3])
		require.Equal(t, uint64(0), data.FFixed64[4])

		require.Equal(t, 5, len(data.FFloat))
		require.Equal(t, float32(111.111), data.FFloat[0])
		require.Equal(t, float32(112.112), data.FFloat[1])
		require.Equal(t, float32(113.113), data.FFloat[2])
		require.Equal(t, float32(0), data.FFloat[3])
		require.Equal(t, float32(0), data.FFloat[4])

		require.Equal(t, 5, len(data.FDouble))
		require.Equal(t, float64(121.121), data.FDouble[0])
		require.Equal(t, float64(122.122), data.FDouble[1])
		require.Equal(t, float64(123.123), data.FDouble[2])
		require.Equal(t, float64(0), data.FDouble[3])
		require.Equal(t, float64(0), data.FDouble[4])
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, 5, len(data.FBool1))
		require.Equal(t, true, data.FBool1[0])
		require.Equal(t, false, data.FBool1[1])
		require.Equal(t, true, data.FBool1[2])
		require.Equal(t, false, data.FBool1[3])
		require.Equal(t, false, data.FBool1[4])
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, 5, len(data.FBytes1))
		require.Equal(t, []byte("abc"), data.FBytes1[0])
		require.Equal(t, []byte(""), data.FBytes1[1])
		require.Equal(t, []byte("123"), data.FBytes1[2])
		require.Equal(t, []byte(""), data.FBytes1[3])
		require.Equal(t, []byte(""), data.FBytes1[4])
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, 5, len(data.FEnum1))
		require.Equal(t, pbtypes.EnumRepeated1(1), data.FEnum1[0])
		require.Equal(t, pbtypes.EnumRepeated1(3), data.FEnum1[1])
		require.Equal(t, pbtypes.EnumRepeated1(5), data.FEnum1[2])
		require.Equal(t, pbtypes.EnumRepeated1(0), data.FEnum1[3])
		require.Equal(t, pbtypes.EnumRepeated1(0), data.FEnum1[4])

		require.Equal(t, 5, len(data.FEnum2))
		require.Equal(t, pbtypes.EnumRepeated1(2), data.FEnum2[0])
		require.Equal(t, pbtypes.EnumRepeated1(4), data.FEnum2[1])
		require.Equal(t, pbtypes.EnumRepeated1(6), data.FEnum2[2])
		require.Equal(t, pbtypes.EnumRepeated1(0), data.FEnum2[3])
		require.Equal(t, pbtypes.EnumRepeated1(0), data.FEnum2[4])

		require.Equal(t, 5, len(data.FEnum3))
		require.Equal(t, pbexternal.Month1(1), data.FEnum3[0])
		require.Equal(t, pbexternal.Month1(3), data.FEnum3[1])
		require.Equal(t, pbexternal.Month1(5), data.FEnum3[2])
		require.Equal(t, pbexternal.Month1(0), data.FEnum3[3])
		require.Equal(t, pbexternal.Month1(0), data.FEnum3[4])

		require.Equal(t, 5, len(data.FEnum4))
		require.Equal(t, pbexternal.Month2(2), data.FEnum4[0])
		require.Equal(t, pbexternal.Month2(4), data.FEnum4[1])
		require.Equal(t, pbexternal.Month2(6), data.FEnum4[2])
		require.Equal(t, pbexternal.Month2(0), data.FEnum4[3])
		require.Equal(t, pbexternal.Month2(0), data.FEnum4[4])

		require.Equal(t, 5, len(data.FEnum5))
		require.Equal(t, pbexternal.EnumWeek1_Week(1), data.FEnum5[0])
		require.Equal(t, pbexternal.EnumWeek1_Week(3), data.FEnum5[1])
		require.Equal(t, pbexternal.EnumWeek1_Week(5), data.FEnum5[2])
		require.Equal(t, pbexternal.EnumWeek1_Week(0), data.FEnum5[3])
		require.Equal(t, pbexternal.EnumWeek1_Week(0), data.FEnum5[4])

		require.Equal(t, 5, len(data.FEnum6))
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(2), data.FEnum6[0])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(4), data.FEnum6[1])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(6), data.FEnum6[2])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(0), data.FEnum6[3])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(0), data.FEnum6[4])

		require.Equal(t, 5, len(data.FEnum7))
		require.Equal(t, pbtypes.EnumCommon1(2), data.FEnum7[0])
		require.Equal(t, pbtypes.EnumCommon1(4), data.FEnum7[1])
		require.Equal(t, pbtypes.EnumCommon1(6), data.FEnum7[2])
		require.Equal(t, pbtypes.EnumCommon1(0), data.FEnum7[3])
		require.Equal(t, pbtypes.EnumCommon1(0), data.FEnum7[4])
	})

	t.Run("WKT", func(t *testing.T) {
		{
			require.Equal(t, 6, len(data.FDuration1))
			require.Equal(t, int64(100), data.FDuration1[0].Seconds)
			require.Equal(t, int32(200), data.FDuration1[0].Nanos)

			require.Equal(t, int64(0), data.FDuration1[1].Seconds)
			require.Equal(t, int32(0), data.FDuration1[1].Nanos)

			require.Equal(t, int64(102), data.FDuration1[2].Seconds)
			require.Equal(t, int32(202), data.FDuration1[2].Nanos)

			require.Equal(t, int64(0), data.FDuration1[3].Seconds)
			require.Equal(t, int32(0), data.FDuration1[3].Nanos)

			require.Equal(t, (*durationpb.Duration)(nil), data.FDuration1[4])
			require.Equal(t, (*durationpb.Duration)(nil), data.FDuration1[5])
		}
		{
			require.Equal(t, 6, len(data.FTimestamp1))
			require.Equal(t, int64(200), data.FTimestamp1[0].Seconds)
			require.Equal(t, int32(300), data.FTimestamp1[0].Nanos)

			require.Equal(t, int64(0), data.FTimestamp1[1].Seconds)
			require.Equal(t, int32(0), data.FTimestamp1[1].Nanos)

			require.Equal(t, int64(202), data.FTimestamp1[2].Seconds)
			require.Equal(t, int32(302), data.FTimestamp1[2].Nanos)

			require.Equal(t, int64(0), data.FTimestamp1[3].Seconds)
			require.Equal(t, int32(0), data.FTimestamp1[3].Nanos)

			require.Equal(t, (*timestamppb.Timestamp)(nil), data.FTimestamp1[4])
			require.Equal(t, (*timestamppb.Timestamp)(nil), data.FTimestamp1[5])
		}
	})

	t.Run("MESSAGE", func(t *testing.T) {
		{
			require.Equal(t, 6, len(data.FMessage1))
			require.NotNil(t, data.FMessage1[0])
			require.Equal(t, "a", data.FMessage1[0].FString1)
			require.Equal(t, "b", data.FMessage1[0].FString2)
			require.Equal(t, "c", data.FMessage1[0].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1)(nil), data.FMessage1[1])
			require.NotNil(t, data.FMessage1[2])
			require.Equal(t, "a", data.FMessage1[2].FString1)
			require.Equal(t, "b", data.FMessage1[2].FString2)
			require.Equal(t, "c", data.FMessage1[2].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1)(nil), data.FMessage1[3])
			require.Equal(t, (*pbtypes.MessageRepeated1)(nil), data.FMessage1[4])
			require.Equal(t, (*pbtypes.MessageRepeated1)(nil), data.FMessage1[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage2))
			require.NotNil(t, data.FMessage2[0])
			require.Equal(t, "", data.FMessage2[0].FString1)
			require.Equal(t, "", data.FMessage2[0].FString2)
			require.Equal(t, "", data.FMessage2[0].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1)(nil), data.FMessage2[1])
			require.NotNil(t, data.FMessage2[2])
			require.Equal(t, "", data.FMessage2[2].FString1)
			require.Equal(t, "", data.FMessage2[2].FString2)
			require.Equal(t, "", data.FMessage2[2].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1)(nil), data.FMessage2[3])
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1)(nil), data.FMessage2[4])
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1)(nil), data.FMessage2[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage3))
			require.NotNil(t, data.FMessage3[0])
			require.Equal(t, "1", data.FMessage3[0].FString1)
			require.Equal(t, "2", data.FMessage3[0].FString2)
			require.Equal(t, "3", data.FMessage3[0].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1_Embed2)(nil), data.FMessage3[1])
			require.NotNil(t, data.FMessage3[2])
			require.Equal(t, "1", data.FMessage3[2].FString1)
			require.Equal(t, "2", data.FMessage3[2].FString2)
			require.Equal(t, "3", data.FMessage3[2].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1_Embed2)(nil), data.FMessage3[3])
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1_Embed2)(nil), data.FMessage3[4])
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1_Embed2)(nil), data.FMessage3[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage4))
			require.NotNil(t, data.FMessage4[0])
			require.Equal(t, "", data.FMessage4[0].TString)
			require.Equal(t, (*string)(nil), data.FMessage4[0].PString)
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[1])
			require.NotNil(t, data.FMessage4[2])
			require.Equal(t, "", data.FMessage4[2].TString)
			require.Equal(t, (*string)(nil), data.FMessage4[2].PString)
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[3])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[4])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage5))
			require.NotNil(t, data.FMessage5[0])
			require.Equal(t, "", data.FMessage5[0].TString1)
			require.Equal(t, "", data.FMessage5[0].TString2)
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[1])
			require.NotNil(t, data.FMessage5[2])
			require.Equal(t, "", data.FMessage5[2].TString1)
			require.Equal(t, "", data.FMessage5[2].TString2)
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[3])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[4])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage6))
			require.NotNil(t, data.FMessage6[0])
			require.Equal(t, "a", data.FMessage6[0].FString1)
			require.Equal(t, "b", data.FMessage6[0].FString2)
			require.Equal(t, "c", data.FMessage6[0].FString3)
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[1])
			require.NotNil(t, data.FMessage6[2])
			require.Equal(t, "a", data.FMessage6[2].FString1)
			require.Equal(t, "b", data.FMessage6[2].FString2)
			require.Equal(t, "c", data.FMessage6[2].FString3)
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[3])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[4])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[5])
		}
	})

	t.Run("ANY", func(t *testing.T) {
		require.Equal(t, 11, len(data.FAny1))
		{
			require.NotNil(t, data.FAny1[0])
			anyVal := &pbtypes.MessageRepeated1{}
			err := data.FAny1[0].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1[1])
			anyVal := &pbtypes.MessageRepeated1_Embed1{}
			err := data.FAny1[1].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "x", anyVal.FString1)
			require.Equal(t, "y", anyVal.FString2)
			require.Equal(t, "z", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1[2])
			anyVal := &pbtypes.MessageRepeated1_Embed1_Embed2{}
			err := data.FAny1[2].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "1", anyVal.FString1)
			require.Equal(t, "2", anyVal.FString2)
			require.Equal(t, "3", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1[3])
			anyVal := &pbexternal.External1{}
			err := data.FAny1[3].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString)
			require.Equal(t, (*string)(nil), anyVal.PString)
		}
		{
			require.NotNil(t, data.FAny1[4])
			anyVal := &pbexternal.External2_Embed1{}
			err := data.FAny1[4].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString1)
			require.Equal(t, "", anyVal.TString2)
		}
		{
			require.NotNil(t, data.FAny1[5])
			anyVal := &timestamppb.Timestamp{}
			err := data.FAny1[5].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, int64(0), anyVal.Seconds)
			require.Equal(t, int32(0), anyVal.Nanos)
		}
		{
			require.NotNil(t, data.FAny1[6])
			anyVal := &anypb.Any{}
			err := data.FAny1[6].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TypeUrl)
			require.Equal(t, ([]byte)(nil), anyVal.Value)
		}
		{
			require.NotNil(t, data.FAny1[7])
			anyVal := &pbtypes.MessageCommon1{}
			err := data.FAny1[7].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.Equal(t, (*anypb.Any)(nil), data.FAny1[8])
		}
		{
			require.Equal(t, (*anypb.Any)(nil), data.FAny1[9])
		}
		{
			require.Equal(t, (*anypb.Any)(nil), data.FAny1[10])
		}
	})
}

// Test cases the field value are have valid value, expected ignore default value.
func Test_TypeRepeated1_Ignore(t *testing.T) {
	data := pbtypes.TypeRepeated1{
		FString1:    []string{"101"},
		FInt32:      []int32{101},
		FInt64:      []int64{102},
		FUint32:     []uint32{103},
		FUint64:     []uint64{104},
		FSint32:     []int32{105},
		FSint64:     []int64{106},
		FSfixed32:   []int32{107},
		FSfixed64:   []int64{108},
		FFixed32:    []uint32{109},
		FFixed64:    []uint64{110},
		FFloat:      []float32{111.111},
		FDouble:     []float64{112.112},
		FBool1:      []bool{true},
		FBytes1:     [][]byte{[]byte("bytes1")},
		FEnum1:      []pbtypes.EnumRepeated1{201},
		FEnum2:      []pbtypes.EnumRepeated1{202},
		FEnum3:      []pbexternal.Month1{203},
		FEnum4:      []pbexternal.Month2{204},
		FEnum5:      []pbexternal.EnumWeek1_Week{205},
		FEnum6:      []pbexternal.EnumWeek2_Embed1_Week{206},
		FEnum7:      []pbtypes.EnumCommon1{207},
		FDuration1:  []*durationpb.Duration{{Seconds: 300, Nanos: 301}},
		FTimestamp1: []*timestamppb.Timestamp{{Seconds: 500, Nanos: 501}},
		FMessage1:   []*pbtypes.MessageRepeated1{{FString1: "1101", FString2: "1102", FString3: "1103"}},
		FMessage2:   []*pbtypes.MessageRepeated1_Embed1{{FString1: "1201", FString2: "1202", FString3: "1203"}},
		FMessage3:   []*pbtypes.MessageRepeated1_Embed1_Embed2{{FString1: "1301", FString2: "1302", FString3: "1303"}},
		FMessage4:   []*pbexternal.External1{{}},
		FMessage5:   []*pbexternal.External2_Embed1{{TString1: "1501", TString2: "1502"}},
		FMessage6:   []*pbtypes.MessageCommon1{{FString1: "1601", FString2: "1602", FString3: "1603"}},
		FAny1:       []*anypb.Any{mustNewAny(&pbexternal.External3{TString: "any1"})},
	}
	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, 1, len(data.FString1))
		require.Equal(t, "101", data.FString1[0])
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, 1, len(data.FInt32))
		require.Equal(t, int32(101), data.FInt32[0])

		require.Equal(t, 1, len(data.FInt64))
		require.Equal(t, int64(102), data.FInt64[0])

		require.Equal(t, 1, len(data.FUint32))
		require.Equal(t, uint32(103), data.FUint32[0])

		require.Equal(t, 1, len(data.FUint64))
		require.Equal(t, uint64(104), data.FUint64[0])

		require.Equal(t, 1, len(data.FSint32))
		require.Equal(t, int32(105), data.FSint32[0])

		require.Equal(t, 1, len(data.FSint64))
		require.Equal(t, int64(106), data.FSint64[0])

		require.Equal(t, 1, len(data.FSfixed32))
		require.Equal(t, int32(107), data.FSfixed32[0])

		require.Equal(t, 1, len(data.FSfixed64))
		require.Equal(t, int64(108), data.FSfixed64[0])

		require.Equal(t, 1, len(data.FFixed32))
		require.Equal(t, uint32(109), data.FFixed32[0])

		require.Equal(t, 1, len(data.FFixed64))
		require.Equal(t, uint64(110), data.FFixed64[0])

		require.Equal(t, 1, len(data.FFloat))
		require.Equal(t, float32(111.111), data.FFloat[0])

		require.Equal(t, 1, len(data.FDouble))
		require.Equal(t, float64(112.112), data.FDouble[0])
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, 1, len(data.FBool1))
		require.Equal(t, true, data.FBool1[0])
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, 1, len(data.FBytes1))
		require.Equal(t, []byte("bytes1"), data.FBytes1[0])
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, 1, len(data.FEnum1))
		require.Equal(t, pbtypes.EnumRepeated1(201), data.FEnum1[0])

		require.Equal(t, 1, len(data.FEnum2))
		require.Equal(t, pbtypes.EnumRepeated1(202), data.FEnum2[0])

		require.Equal(t, 1, len(data.FEnum3))
		require.Equal(t, pbexternal.Month1(203), data.FEnum3[0])

		require.Equal(t, 1, len(data.FEnum4))
		require.Equal(t, pbexternal.Month2(204), data.FEnum4[0])

		require.Equal(t, 1, len(data.FEnum5))
		require.Equal(t, pbexternal.EnumWeek1_Week(205), data.FEnum5[0])

		require.Equal(t, 1, len(data.FEnum6))
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(206), data.FEnum6[0])

		require.Equal(t, 1, len(data.FEnum7))
		require.Equal(t, pbtypes.EnumCommon1(207), data.FEnum7[0])
	})

	t.Run("WKT", func(t *testing.T) {
		require.Equal(t, 1, len(data.FDuration1))
		require.Equal(t, int64(300), data.FDuration1[0].Seconds)
		require.Equal(t, int32(301), data.FDuration1[0].Nanos)

		require.Equal(t, 1, len(data.FTimestamp1))
		require.Equal(t, int64(500), data.FTimestamp1[0].Seconds)
		require.Equal(t, int32(501), data.FTimestamp1[0].Nanos)
	})

	t.Run("MESSAGE", func(t *testing.T) {
		{
			require.Equal(t, 1, len(data.FMessage1))
			require.NotNil(t, data.FMessage1[0])
			require.Equal(t, "1101", data.FMessage1[0].FString1)
			require.Equal(t, "1102", data.FMessage1[0].FString2)
			require.Equal(t, "1103", data.FMessage1[0].FString3)
		}
		{
			require.Equal(t, 1, len(data.FMessage2))
			require.NotNil(t, data.FMessage2[0])
			require.Equal(t, "1201", data.FMessage2[0].FString1)
			require.Equal(t, "1202", data.FMessage2[0].FString2)
			require.Equal(t, "1203", data.FMessage2[0].FString3)
		}
		{
			require.Equal(t, 1, len(data.FMessage3))
			require.NotNil(t, data.FMessage3[0])
			require.Equal(t, "1301", data.FMessage3[0].FString1)
			require.Equal(t, "1302", data.FMessage3[0].FString2)
			require.Equal(t, "1303", data.FMessage3[0].FString3)
		}
		{
			require.Equal(t, 1, len(data.FMessage4))
			require.NotNil(t, data.FMessage4[0])
			require.Equal(t, "", data.FMessage4[0].TString)
			require.Equal(t, (*string)(nil), data.FMessage4[0].PString)
		}
		{
			require.Equal(t, 1, len(data.FMessage5))
			require.NotNil(t, data.FMessage5[0])
			require.Equal(t, "1501", data.FMessage5[0].TString1)
			require.Equal(t, "1502", data.FMessage5[0].TString2)
		}
		{
			require.Equal(t, 1, len(data.FMessage6))
			require.NotNil(t, data.FMessage6[0])
			require.Equal(t, "1601", data.FMessage6[0].FString1)
			require.Equal(t, "1602", data.FMessage6[0].FString2)
			require.Equal(t, "1603", data.FMessage6[0].FString3)
		}
	})

	t.Run("ANY", func(t *testing.T) {
		require.Equal(t, 1, len(data.FAny1))
		{
			require.NotNil(t, data.FAny1[0])
			anyVal := &pbexternal.External3{}
			err := data.FAny1[0].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any1", anyVal.TString)
		}
	})
}

// Test cases the field value is nil pointer.
func Test_TypeRepeated2_NIL(t *testing.T) {
	data := pbtypes.TypeRepeated2{}
	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, 11, len(data.FString1))
		require.Equal(t, "ts1", data.FString1[0])
		require.Equal(t, "", data.FString1[1])
		require.Equal(t, `""`, data.FString1[2])
		require.Equal(t, `"ts4"`, data.FString1[3])
		require.Equal(t, `"ts5"`, data.FString1[4])
		require.Equal(t, `"ts"6"`, data.FString1[5])
		require.Equal(t, `"ts"7"`, data.FString1[6])
		require.Equal(t, `[ts8]`, data.FString1[7])
		require.Equal(t, `{ts9}`, data.FString1[8])
		require.Equal(t, "", data.FString1[9])
		require.Equal(t, "", data.FString1[10])
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, 5, len(data.FInt32))
		require.Equal(t, int32(11), data.FInt32[0])
		require.Equal(t, int32(12), data.FInt32[1])
		require.Equal(t, int32(13), data.FInt32[2])
		require.Equal(t, int32(0), data.FInt32[3])
		require.Equal(t, int32(0), data.FInt32[4])

		require.Equal(t, 5, len(data.FInt64))
		require.Equal(t, int64(21), data.FInt64[0])
		require.Equal(t, int64(22), data.FInt64[1])
		require.Equal(t, int64(23), data.FInt64[2])
		require.Equal(t, int64(0), data.FInt64[3])
		require.Equal(t, int64(0), data.FInt64[4])

		require.Equal(t, 5, len(data.FUint32))
		require.Equal(t, uint32(31), data.FUint32[0])
		require.Equal(t, uint32(32), data.FUint32[1])
		require.Equal(t, uint32(33), data.FUint32[2])
		require.Equal(t, uint32(0), data.FUint32[3])
		require.Equal(t, uint32(0), data.FUint32[4])

		require.Equal(t, 5, len(data.FUint64))
		require.Equal(t, uint64(41), data.FUint64[0])
		require.Equal(t, uint64(42), data.FUint64[1])
		require.Equal(t, uint64(43), data.FUint64[2])
		require.Equal(t, uint64(0), data.FUint64[3])
		require.Equal(t, uint64(0), data.FUint64[4])

		require.Equal(t, 5, len(data.FSint32))
		require.Equal(t, int32(51), data.FSint32[0])
		require.Equal(t, int32(52), data.FSint32[1])
		require.Equal(t, int32(53), data.FSint32[2])
		require.Equal(t, int32(0), data.FSint32[3])
		require.Equal(t, int32(0), data.FSint32[4])

		require.Equal(t, 5, len(data.FSint64))
		require.Equal(t, int64(61), data.FSint64[0])
		require.Equal(t, int64(62), data.FSint64[1])
		require.Equal(t, int64(63), data.FSint64[2])
		require.Equal(t, int64(0), data.FSint64[3])
		require.Equal(t, int64(0), data.FSint64[4])

		require.Equal(t, 5, len(data.FSfixed32))
		require.Equal(t, int32(71), data.FSfixed32[0])
		require.Equal(t, int32(72), data.FSfixed32[1])
		require.Equal(t, int32(73), data.FSfixed32[2])
		require.Equal(t, int32(0), data.FSfixed32[3])
		require.Equal(t, int32(0), data.FSfixed32[4])

		require.Equal(t, 5, len(data.FSfixed64))
		require.Equal(t, int64(81), data.FSfixed64[0])
		require.Equal(t, int64(82), data.FSfixed64[1])
		require.Equal(t, int64(83), data.FSfixed64[2])
		require.Equal(t, int64(0), data.FSfixed64[3])
		require.Equal(t, int64(0), data.FSfixed64[4])

		require.Equal(t, 5, len(data.FFixed32))
		require.Equal(t, uint32(91), data.FFixed32[0])
		require.Equal(t, uint32(92), data.FFixed32[1])
		require.Equal(t, uint32(93), data.FFixed32[2])
		require.Equal(t, uint32(0), data.FFixed32[3])
		require.Equal(t, uint32(0), data.FFixed32[4])

		require.Equal(t, 5, len(data.FFixed64))
		require.Equal(t, uint64(101), data.FFixed64[0])
		require.Equal(t, uint64(102), data.FFixed64[1])
		require.Equal(t, uint64(103), data.FFixed64[2])
		require.Equal(t, uint64(0), data.FFixed64[3])
		require.Equal(t, uint64(0), data.FFixed64[4])

		require.Equal(t, 5, len(data.FFloat))
		require.Equal(t, float32(111.111), data.FFloat[0])
		require.Equal(t, float32(112.112), data.FFloat[1])
		require.Equal(t, float32(113.113), data.FFloat[2])
		require.Equal(t, float32(0), data.FFloat[3])
		require.Equal(t, float32(0), data.FFloat[4])

		require.Equal(t, 5, len(data.FDouble))
		require.Equal(t, float64(121.121), data.FDouble[0])
		require.Equal(t, float64(122.122), data.FDouble[1])
		require.Equal(t, float64(123.123), data.FDouble[2])
		require.Equal(t, float64(0), data.FDouble[3])
		require.Equal(t, float64(0), data.FDouble[4])
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, 5, len(data.FBool1))
		require.Equal(t, true, data.FBool1[0])
		require.Equal(t, false, data.FBool1[1])
		require.Equal(t, true, data.FBool1[2])
		require.Equal(t, false, data.FBool1[3])
		require.Equal(t, false, data.FBool1[4])
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, 5, len(data.FBytes1))
		require.Equal(t, []byte("abc"), data.FBytes1[0])
		require.Equal(t, []byte(""), data.FBytes1[1])
		require.Equal(t, []byte("123"), data.FBytes1[2])
		require.Equal(t, []byte(""), data.FBytes1[3])
		require.Equal(t, []byte(""), data.FBytes1[4])
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, 5, len(data.FEnum1))
		require.Equal(t, pbtypes.EnumRepeated1(1), data.FEnum1[0])
		require.Equal(t, pbtypes.EnumRepeated1(3), data.FEnum1[1])
		require.Equal(t, pbtypes.EnumRepeated1(5), data.FEnum1[2])
		require.Equal(t, pbtypes.EnumRepeated1(0), data.FEnum1[3])
		require.Equal(t, pbtypes.EnumRepeated1(0), data.FEnum1[4])

		require.Equal(t, 5, len(data.FEnum2))
		require.Equal(t, pbtypes.EnumRepeated1(2), data.FEnum2[0])
		require.Equal(t, pbtypes.EnumRepeated1(4), data.FEnum2[1])
		require.Equal(t, pbtypes.EnumRepeated1(6), data.FEnum2[2])
		require.Equal(t, pbtypes.EnumRepeated1(0), data.FEnum2[3])
		require.Equal(t, pbtypes.EnumRepeated1(0), data.FEnum2[4])

		require.Equal(t, 5, len(data.FEnum3))
		require.Equal(t, pbexternal.Month1(1), data.FEnum3[0])
		require.Equal(t, pbexternal.Month1(3), data.FEnum3[1])
		require.Equal(t, pbexternal.Month1(5), data.FEnum3[2])
		require.Equal(t, pbexternal.Month1(0), data.FEnum3[3])
		require.Equal(t, pbexternal.Month1(0), data.FEnum3[4])

		require.Equal(t, 5, len(data.FEnum4))
		require.Equal(t, pbexternal.Month2(2), data.FEnum4[0])
		require.Equal(t, pbexternal.Month2(4), data.FEnum4[1])
		require.Equal(t, pbexternal.Month2(6), data.FEnum4[2])
		require.Equal(t, pbexternal.Month2(0), data.FEnum4[3])
		require.Equal(t, pbexternal.Month2(0), data.FEnum4[4])

		require.Equal(t, 5, len(data.FEnum5))
		require.Equal(t, pbexternal.EnumWeek1_Week(1), data.FEnum5[0])
		require.Equal(t, pbexternal.EnumWeek1_Week(3), data.FEnum5[1])
		require.Equal(t, pbexternal.EnumWeek1_Week(5), data.FEnum5[2])
		require.Equal(t, pbexternal.EnumWeek1_Week(0), data.FEnum5[3])
		require.Equal(t, pbexternal.EnumWeek1_Week(0), data.FEnum5[4])

		require.Equal(t, 5, len(data.FEnum6))
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(2), data.FEnum6[0])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(4), data.FEnum6[1])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(6), data.FEnum6[2])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(0), data.FEnum6[3])
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(0), data.FEnum6[4])

		require.Equal(t, 5, len(data.FEnum7))
		require.Equal(t, pbtypes.EnumCommon1(2), data.FEnum7[0])
		require.Equal(t, pbtypes.EnumCommon1(4), data.FEnum7[1])
		require.Equal(t, pbtypes.EnumCommon1(6), data.FEnum7[2])
		require.Equal(t, pbtypes.EnumCommon1(0), data.FEnum7[3])
		require.Equal(t, pbtypes.EnumCommon1(0), data.FEnum7[4])
	})

	t.Run("WKT", func(t *testing.T) {
		{
			require.Equal(t, 6, len(data.FDuration1))
			require.Equal(t, int64(100), data.FDuration1[0].Seconds)
			require.Equal(t, int32(200), data.FDuration1[0].Nanos)

			require.Equal(t, int64(0), data.FDuration1[1].Seconds)
			require.Equal(t, int32(0), data.FDuration1[1].Nanos)

			require.Equal(t, int64(102), data.FDuration1[2].Seconds)
			require.Equal(t, int32(202), data.FDuration1[2].Nanos)

			require.Equal(t, int64(0), data.FDuration1[3].Seconds)
			require.Equal(t, int32(0), data.FDuration1[3].Nanos)

			require.Equal(t, (*durationpb.Duration)(nil), data.FDuration1[4])
			require.Equal(t, (*durationpb.Duration)(nil), data.FDuration1[5])
		}
		{
			require.Equal(t, 6, len(data.FTimestamp1))
			require.Equal(t, int64(200), data.FTimestamp1[0].Seconds)
			require.Equal(t, int32(300), data.FTimestamp1[0].Nanos)

			require.Equal(t, int64(0), data.FTimestamp1[1].Seconds)
			require.Equal(t, int32(0), data.FTimestamp1[1].Nanos)

			require.Equal(t, int64(202), data.FTimestamp1[2].Seconds)
			require.Equal(t, int32(302), data.FTimestamp1[2].Nanos)

			require.Equal(t, int64(0), data.FTimestamp1[3].Seconds)
			require.Equal(t, int32(0), data.FTimestamp1[3].Nanos)

			require.Equal(t, (*timestamppb.Timestamp)(nil), data.FTimestamp1[4])
			require.Equal(t, (*timestamppb.Timestamp)(nil), data.FTimestamp1[5])
		}
	})

	t.Run("MESSAGE", func(t *testing.T) {
		{
			require.Equal(t, 6, len(data.FMessage1))
			require.NotNil(t, data.FMessage1[0])
			require.Equal(t, "a", data.FMessage1[0].FString1)
			require.Equal(t, "b", data.FMessage1[0].FString2)
			require.Equal(t, "c", data.FMessage1[0].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1)(nil), data.FMessage1[1])
			require.NotNil(t, data.FMessage1[2])
			require.Equal(t, "a", data.FMessage1[2].FString1)
			require.Equal(t, "b", data.FMessage1[2].FString2)
			require.Equal(t, "c", data.FMessage1[2].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1)(nil), data.FMessage1[3])
			require.Equal(t, (*pbtypes.MessageRepeated1)(nil), data.FMessage1[4])
			require.Equal(t, (*pbtypes.MessageRepeated1)(nil), data.FMessage1[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage2))
			require.NotNil(t, data.FMessage2[0])
			require.Equal(t, "", data.FMessage2[0].FString1)
			require.Equal(t, "", data.FMessage2[0].FString2)
			require.Equal(t, "", data.FMessage2[0].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1)(nil), data.FMessage2[1])
			require.NotNil(t, data.FMessage2[2])
			require.Equal(t, "", data.FMessage2[2].FString1)
			require.Equal(t, "", data.FMessage2[2].FString2)
			require.Equal(t, "", data.FMessage2[2].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1)(nil), data.FMessage2[3])
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1)(nil), data.FMessage2[4])
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1)(nil), data.FMessage2[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage3))
			require.NotNil(t, data.FMessage3[0])
			require.Equal(t, "1", data.FMessage3[0].FString1)
			require.Equal(t, "2", data.FMessage3[0].FString2)
			require.Equal(t, "3", data.FMessage3[0].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1_Embed2)(nil), data.FMessage3[1])
			require.NotNil(t, data.FMessage3[2])
			require.Equal(t, "1", data.FMessage3[2].FString1)
			require.Equal(t, "2", data.FMessage3[2].FString2)
			require.Equal(t, "3", data.FMessage3[2].FString3)
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1_Embed2)(nil), data.FMessage3[3])
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1_Embed2)(nil), data.FMessage3[4])
			require.Equal(t, (*pbtypes.MessageRepeated1_Embed1_Embed2)(nil), data.FMessage3[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage4))
			require.NotNil(t, data.FMessage4[0])
			require.Equal(t, "", data.FMessage4[0].TString)
			require.Equal(t, (*string)(nil), data.FMessage4[0].PString)
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[1])
			require.NotNil(t, data.FMessage4[2])
			require.Equal(t, "", data.FMessage4[2].TString)
			require.Equal(t, (*string)(nil), data.FMessage4[2].PString)
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[3])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[4])
			require.Equal(t, (*pbexternal.External1)(nil), data.FMessage4[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage5))
			require.NotNil(t, data.FMessage5[0])
			require.Equal(t, "", data.FMessage5[0].TString1)
			require.Equal(t, "", data.FMessage5[0].TString2)
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[1])
			require.NotNil(t, data.FMessage5[2])
			require.Equal(t, "", data.FMessage5[2].TString1)
			require.Equal(t, "", data.FMessage5[2].TString2)
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[3])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[4])
			require.Equal(t, (*pbexternal.External2_Embed1)(nil), data.FMessage5[5])
		}
		{
			require.Equal(t, 6, len(data.FMessage6))
			require.NotNil(t, data.FMessage6[0])
			require.Equal(t, "a", data.FMessage6[0].FString1)
			require.Equal(t, "b", data.FMessage6[0].FString2)
			require.Equal(t, "c", data.FMessage6[0].FString3)
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[1])
			require.NotNil(t, data.FMessage6[2])
			require.Equal(t, "a", data.FMessage6[2].FString1)
			require.Equal(t, "b", data.FMessage6[2].FString2)
			require.Equal(t, "c", data.FMessage6[2].FString3)
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[3])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[4])
			require.Equal(t, (*pbtypes.MessageCommon1)(nil), data.FMessage6[5])
		}
	})

	t.Run("ANY", func(t *testing.T) {
		require.Equal(t, 11, len(data.FAny1))
		{
			require.NotNil(t, data.FAny1[0])
			anyVal := &pbtypes.MessageRepeated1{}
			err := data.FAny1[0].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1[1])
			anyVal := &pbtypes.MessageRepeated1_Embed1{}
			err := data.FAny1[1].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "x", anyVal.FString1)
			require.Equal(t, "y", anyVal.FString2)
			require.Equal(t, "z", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1[2])
			anyVal := &pbtypes.MessageRepeated1_Embed1_Embed2{}
			err := data.FAny1[2].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "1", anyVal.FString1)
			require.Equal(t, "2", anyVal.FString2)
			require.Equal(t, "3", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny1[3])
			anyVal := &pbexternal.External1{}
			err := data.FAny1[3].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString)
			require.Equal(t, (*string)(nil), anyVal.PString)
		}
		{
			require.NotNil(t, data.FAny1[4])
			anyVal := &pbexternal.External2_Embed1{}
			err := data.FAny1[4].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString1)
			require.Equal(t, "", anyVal.TString2)
		}
		{
			require.NotNil(t, data.FAny1[5])
			anyVal := &timestamppb.Timestamp{}
			err := data.FAny1[5].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, int64(0), anyVal.Seconds)
			require.Equal(t, int32(0), anyVal.Nanos)
		}
		{
			require.NotNil(t, data.FAny1[6])
			anyVal := &anypb.Any{}
			err := data.FAny1[6].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TypeUrl)
			require.Equal(t, ([]byte)(nil), anyVal.Value)
		}
		{
			require.NotNil(t, data.FAny1[7])
			anyVal := &pbtypes.MessageCommon1{}
			err := data.FAny1[7].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.Equal(t, (*anypb.Any)(nil), data.FAny1[8])
		}
		{
			require.Equal(t, (*anypb.Any)(nil), data.FAny1[9])
		}
		{
			require.Equal(t, (*anypb.Any)(nil), data.FAny1[10])
		}
	})
}

// Test cases the field value is empty.
func Test_TypeRepeated2_Empty(t *testing.T) {
	data := pbtypes.TypeRepeated2{
		FString1:    []string{},
		FInt32:      []int32{},
		FInt64:      []int64{},
		FUint32:     []uint32{},
		FUint64:     []uint64{},
		FSint32:     []int32{},
		FSint64:     []int64{},
		FSfixed32:   []int32{},
		FSfixed64:   []int64{},
		FFixed32:    []uint32{},
		FFixed64:    []uint64{},
		FFloat:      []float32{},
		FDouble:     []float64{},
		FBool1:      []bool{},
		FBytes1:     [][]byte{},
		FEnum1:      []pbtypes.EnumRepeated1{},
		FEnum2:      []pbtypes.EnumRepeated1{},
		FEnum3:      []pbexternal.Month1{},
		FEnum4:      []pbexternal.Month2{},
		FEnum5:      []pbexternal.EnumWeek1_Week{},
		FEnum6:      []pbexternal.EnumWeek2_Embed1_Week{},
		FEnum7:      []pbtypes.EnumCommon1{},
		FDuration1:  []*durationpb.Duration{},
		FTimestamp1: []*timestamppb.Timestamp{},
		FMessage1:   []*pbtypes.MessageRepeated1{},
		FMessage2:   []*pbtypes.MessageRepeated1_Embed1{},
		FMessage3:   []*pbtypes.MessageRepeated1_Embed1_Embed2{},
		FMessage4:   []*pbexternal.External1{},
		FMessage5:   []*pbexternal.External2_Embed1{},
		FMessage6:   []*pbtypes.MessageCommon1{},
		FAny1:       []*anypb.Any{},
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
func Test_TypeRepeated2_Ignore(t *testing.T) {
	data := pbtypes.TypeRepeated2{
		FString1:    []string{"101"},
		FInt32:      []int32{101},
		FInt64:      []int64{102},
		FUint32:     []uint32{103},
		FUint64:     []uint64{104},
		FSint32:     []int32{105},
		FSint64:     []int64{106},
		FSfixed32:   []int32{107},
		FSfixed64:   []int64{108},
		FFixed32:    []uint32{109},
		FFixed64:    []uint64{110},
		FFloat:      []float32{111.111},
		FDouble:     []float64{112.112},
		FBool1:      []bool{true},
		FBytes1:     [][]byte{[]byte("bytes1")},
		FEnum1:      []pbtypes.EnumRepeated1{201},
		FEnum2:      []pbtypes.EnumRepeated1{202},
		FEnum3:      []pbexternal.Month1{203},
		FEnum4:      []pbexternal.Month2{204},
		FEnum5:      []pbexternal.EnumWeek1_Week{205},
		FEnum6:      []pbexternal.EnumWeek2_Embed1_Week{206},
		FEnum7:      []pbtypes.EnumCommon1{207},
		FDuration1:  []*durationpb.Duration{{Seconds: 300, Nanos: 301}},
		FTimestamp1: []*timestamppb.Timestamp{{Seconds: 500, Nanos: 501}},
		FMessage1:   []*pbtypes.MessageRepeated1{{FString1: "1101", FString2: "1102", FString3: "1103"}},
		FMessage2:   []*pbtypes.MessageRepeated1_Embed1{{FString1: "1201", FString2: "1202", FString3: "1203"}},
		FMessage3:   []*pbtypes.MessageRepeated1_Embed1_Embed2{{FString1: "1301", FString2: "1302", FString3: "1303"}},
		FMessage4:   []*pbexternal.External1{{}},
		FMessage5:   []*pbexternal.External2_Embed1{{TString1: "1501", TString2: "1502"}},
		FMessage6:   []*pbtypes.MessageCommon1{{FString1: "1601", FString2: "1602", FString3: "1603"}},
		FAny1:       []*anypb.Any{mustNewAny(&pbexternal.External3{TString: "any1"})},
	}
	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, 1, len(data.FString1))
		require.Equal(t, "101", data.FString1[0])
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, 1, len(data.FInt32))
		require.Equal(t, int32(101), data.FInt32[0])

		require.Equal(t, 1, len(data.FInt64))
		require.Equal(t, int64(102), data.FInt64[0])

		require.Equal(t, 1, len(data.FUint32))
		require.Equal(t, uint32(103), data.FUint32[0])

		require.Equal(t, 1, len(data.FUint64))
		require.Equal(t, uint64(104), data.FUint64[0])

		require.Equal(t, 1, len(data.FSint32))
		require.Equal(t, int32(105), data.FSint32[0])

		require.Equal(t, 1, len(data.FSint64))
		require.Equal(t, int64(106), data.FSint64[0])

		require.Equal(t, 1, len(data.FSfixed32))
		require.Equal(t, int32(107), data.FSfixed32[0])

		require.Equal(t, 1, len(data.FSfixed64))
		require.Equal(t, int64(108), data.FSfixed64[0])

		require.Equal(t, 1, len(data.FFixed32))
		require.Equal(t, uint32(109), data.FFixed32[0])

		require.Equal(t, 1, len(data.FFixed64))
		require.Equal(t, uint64(110), data.FFixed64[0])

		require.Equal(t, 1, len(data.FFloat))
		require.Equal(t, float32(111.111), data.FFloat[0])

		require.Equal(t, 1, len(data.FDouble))
		require.Equal(t, float64(112.112), data.FDouble[0])
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, 1, len(data.FBool1))
		require.Equal(t, true, data.FBool1[0])
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, 1, len(data.FBytes1))
		require.Equal(t, []byte("bytes1"), data.FBytes1[0])
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, 1, len(data.FEnum1))
		require.Equal(t, pbtypes.EnumRepeated1(201), data.FEnum1[0])

		require.Equal(t, 1, len(data.FEnum2))
		require.Equal(t, pbtypes.EnumRepeated1(202), data.FEnum2[0])

		require.Equal(t, 1, len(data.FEnum3))
		require.Equal(t, pbexternal.Month1(203), data.FEnum3[0])

		require.Equal(t, 1, len(data.FEnum4))
		require.Equal(t, pbexternal.Month2(204), data.FEnum4[0])

		require.Equal(t, 1, len(data.FEnum5))
		require.Equal(t, pbexternal.EnumWeek1_Week(205), data.FEnum5[0])

		require.Equal(t, 1, len(data.FEnum6))
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(206), data.FEnum6[0])

		require.Equal(t, 1, len(data.FEnum7))
		require.Equal(t, pbtypes.EnumCommon1(207), data.FEnum7[0])
	})

	t.Run("WKT", func(t *testing.T) {
		require.Equal(t, 1, len(data.FDuration1))
		require.Equal(t, int64(300), data.FDuration1[0].Seconds)
		require.Equal(t, int32(301), data.FDuration1[0].Nanos)

		require.Equal(t, 1, len(data.FTimestamp1))
		require.Equal(t, int64(500), data.FTimestamp1[0].Seconds)
		require.Equal(t, int32(501), data.FTimestamp1[0].Nanos)
	})

	t.Run("MESSAGE", func(t *testing.T) {
		{
			require.Equal(t, 1, len(data.FMessage1))
			require.NotNil(t, data.FMessage1[0])
			require.Equal(t, "1101", data.FMessage1[0].FString1)
			require.Equal(t, "1102", data.FMessage1[0].FString2)
			require.Equal(t, "1103", data.FMessage1[0].FString3)
		}
		{
			require.Equal(t, 1, len(data.FMessage2))
			require.NotNil(t, data.FMessage2[0])
			require.Equal(t, "1201", data.FMessage2[0].FString1)
			require.Equal(t, "1202", data.FMessage2[0].FString2)
			require.Equal(t, "1203", data.FMessage2[0].FString3)
		}
		{
			require.Equal(t, 1, len(data.FMessage3))
			require.NotNil(t, data.FMessage3[0])
			require.Equal(t, "1301", data.FMessage3[0].FString1)
			require.Equal(t, "1302", data.FMessage3[0].FString2)
			require.Equal(t, "1303", data.FMessage3[0].FString3)
		}
		{
			require.Equal(t, 1, len(data.FMessage4))
			require.NotNil(t, data.FMessage4[0])
			require.Equal(t, "", data.FMessage4[0].TString)
			require.Equal(t, (*string)(nil), data.FMessage4[0].PString)
		}
		{
			require.Equal(t, 1, len(data.FMessage5))
			require.NotNil(t, data.FMessage5[0])
			require.Equal(t, "1501", data.FMessage5[0].TString1)
			require.Equal(t, "1502", data.FMessage5[0].TString2)
		}
		{
			require.Equal(t, 1, len(data.FMessage6))
			require.NotNil(t, data.FMessage6[0])
			require.Equal(t, "1601", data.FMessage6[0].FString1)
			require.Equal(t, "1602", data.FMessage6[0].FString2)
			require.Equal(t, "1603", data.FMessage6[0].FString3)
		}
	})

	t.Run("ANY", func(t *testing.T) {
		require.Equal(t, 1, len(data.FAny1))
		{
			require.NotNil(t, data.FAny1[0])
			anyVal := &pbexternal.External3{}
			err := data.FAny1[0].UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any1", anyVal.TString)
		}
	})
}
