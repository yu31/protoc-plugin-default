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
func Test_TypePlain1_NIL(t *testing.T) {
	data := &pbtypes.TypePlain1{}
	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, "ts1", data.FString1)
		require.Equal(t, "", data.FString2)
		require.Equal(t, `""`, data.FString3)
		require.Equal(t, `"ts4"`, data.FString4)
		require.Equal(t, `"ts5"`, data.FString5)
		require.Equal(t, `"ts"6"`, data.FString6)
		require.Equal(t, `"ts"7"`, data.FString7)
		require.Equal(t, `[ts8]`, data.FString8)
		require.Equal(t, `{ts9}`, data.FString9)
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, int32(1), data.FInt32)
		require.Equal(t, int64(2), data.FInt64)
		require.Equal(t, uint32(3), data.FUint32)
		require.Equal(t, uint64(4), data.FUint64)
		require.Equal(t, int32(5), data.FSint32)
		require.Equal(t, int64(6), data.FSint64)
		require.Equal(t, int32(7), data.FSfixed32)
		require.Equal(t, int64(8), data.FSfixed64)
		require.Equal(t, uint32(9), data.FFixed32)
		require.Equal(t, uint64(10), data.FFixed64)
		require.Equal(t, float32(11.11), data.FFloat)
		require.Equal(t, float64(12.12), data.FDouble)
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, true, data.FBool1)
		require.Equal(t, false, data.FBool2)
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, []byte(""), data.FBytes1)
		require.Equal(t, []byte("abc"), data.FBytes2)
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, pbtypes.EnumPlain1(2), data.FEnum1)
		require.Equal(t, pbtypes.EnumPlain1(10), data.FEnum2)
		require.Equal(t, pbexternal.Month1(1), data.FEnum3)
		require.Equal(t, pbexternal.Month2(7), data.FEnum4)
		require.Equal(t, pbexternal.EnumWeek1_Week(3), data.FEnum5)
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(13), data.FEnum6)
		require.Equal(t, pbtypes.EnumCommon1(9), data.FEnum7)
	})

	t.Run("WKT", func(t *testing.T) {
		require.NotNil(t, data.FDuration1)
		require.Equal(t, int64(100), data.FDuration1.Seconds)
		require.Equal(t, int32(200), data.FDuration1.Nanos)

		require.NotNil(t, data.FDuration2)
		require.Equal(t, int64(0), data.FDuration2.Seconds)
		require.Equal(t, int32(0), data.FDuration2.Nanos)

		require.NotNil(t, data.FTimestamp1)
		require.Equal(t, int64(200), data.FTimestamp1.Seconds)
		require.Equal(t, int32(300), data.FTimestamp1.Nanos)

		require.NotNil(t, data.FTimestamp2)
		require.Equal(t, int64(0), data.FTimestamp2.Seconds)
		require.Equal(t, int32(0), data.FTimestamp2.Nanos)
	})

	t.Run("MESSAGE1", func(t *testing.T) {
		require.NotNil(t, data.FMessage11)
		require.Equal(t, "a", data.FMessage11.FString1)
		require.Equal(t, "b", data.FMessage11.FString2)
		require.Equal(t, "c", data.FMessage11.FString3)

		require.NotNil(t, data.FMessage12)
		require.Equal(t, "", data.FMessage12.FString1)
		require.Equal(t, "", data.FMessage12.FString2)
		require.Equal(t, "", data.FMessage12.FString3)

		require.NotNil(t, data.FMessage13)
		require.Equal(t, "1", data.FMessage13.FString1)
		require.Equal(t, "2", data.FMessage13.FString2)
		require.Equal(t, "3", data.FMessage13.FString3)

		require.NotNil(t, data.FMessage14)
		require.Equal(t, "", data.FMessage14.TString)
		require.Equal(t, (*string)(nil), data.FMessage14.PString)

		require.NotNil(t, data.FMessage15)
		require.Equal(t, "", data.FMessage15.TString1)
		require.Equal(t, "", data.FMessage15.TString2)

		require.NotNil(t, data.FMessage16)
		require.Equal(t, "", data.FMessage16.FString1)
		require.Equal(t, "", data.FMessage16.FString2)
		require.Equal(t, "", data.FMessage16.FString3)

		require.NotNil(t, data.FMessage17)
		require.Equal(t, "x", data.FMessage17.FString1)
		require.Equal(t, "y", data.FMessage17.FString2)
		require.Equal(t, "z", data.FMessage17.FString3)

		require.NotNil(t, data.FMessage18)
		require.Equal(t, "", data.FMessage18.FString1)
		require.Equal(t, "", data.FMessage18.FString2)
		require.Equal(t, "", data.FMessage18.FString3)
	})

	t.Run("MESSAGE2", func(t *testing.T) {
		require.Nil(t, data.FMessage21)
		require.Nil(t, data.FMessage22)
		require.Nil(t, data.FMessage23)
		require.Nil(t, data.FMessage24)
		require.Nil(t, data.FMessage25)
		require.Nil(t, data.FMessage26)
		require.Nil(t, data.FMessage27)
		require.Nil(t, data.FMessage28)
	})

	t.Run("ANY", func(t *testing.T) {
		{
			require.NotNil(t, data.FAny1)
			anyVal := &pbtypes.MessagePlain1{}
			err := data.FAny1.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny2)
			anyVal := &pbtypes.MessagePlain1{}
			err := data.FAny2.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.FString1)
			require.Equal(t, "", anyVal.FString2)
			require.Equal(t, "", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny3)
			anyVal := &pbtypes.MessagePlain1_Embed1{}
			err := data.FAny3.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.FString1)
			require.Equal(t, "", anyVal.FString2)
			require.Equal(t, "", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny4)
			anyVal := &pbtypes.MessagePlain1_Embed1_Embed2{}
			err := data.FAny4.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "1", anyVal.FString1)
			require.Equal(t, "2", anyVal.FString2)
			require.Equal(t, "3", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny5)
			anyVal := &pbexternal.External1{}
			err := data.FAny5.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString)
			require.Equal(t, (*string)(nil), anyVal.PString)
		}
		{
			require.NotNil(t, data.FAny6)
			anyVal := &pbexternal.External2_Embed1{}
			err := data.FAny6.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString1)
			require.Equal(t, "", anyVal.TString2)
		}
		{
			require.NotNil(t, data.FAny7)
			anyVal := &timestamppb.Timestamp{Seconds: 0, Nanos: 0}
			err := data.FAny7.UnmarshalTo(anyVal)
			require.Nil(t, err)
		}
		{
			require.NotNil(t, data.FAny8)
			anyVal := &anypb.Any{}
			err := data.FAny8.UnmarshalTo(anyVal)
			require.Nil(t, err)
		}
		{
			require.NotNil(t, data.FAny9)
			anyVal := &pbtypes.MessageCommon1{}
			err := data.FAny9.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
	})
}

// Test cases the field value is empty.
func Test_TypePlain1_Empty(t *testing.T) {
	data := &pbtypes.TypePlain1{
		FBytes1:     []byte(""),
		FBytes2:     []byte(""),
		FDuration1:  &durationpb.Duration{},
		FDuration2:  &durationpb.Duration{},
		FTimestamp1: &timestamppb.Timestamp{},
		FTimestamp2: &timestamppb.Timestamp{},
		FMessage11:  &pbtypes.MessagePlain1{},
		FMessage12:  &pbtypes.MessagePlain1_Embed1{},
		FMessage13:  &pbtypes.MessagePlain1_Embed1_Embed2{},
		FMessage14:  &pbexternal.External1{},
		FMessage15:  &pbexternal.External2_Embed1{},
		FMessage16:  &pbtypes.MessageCommon1{},
		FMessage17:  &pbtypes.MessageCommon1_Embed1{},
		FMessage18:  &pbtypes.MessageCommon1_Embed1_Embed2{},
		FMessage21:  &pbtypes.MessagePlain1{},
		FMessage22:  &pbtypes.MessagePlain1_Embed1{},
		FMessage23:  &pbtypes.MessagePlain1_Embed1_Embed2{},
		FMessage24:  &pbexternal.External1{},
		FMessage25:  &pbexternal.External2_Embed1{},
		FMessage26:  &pbtypes.MessageCommon1{},
		FMessage27:  &pbtypes.MessageCommon1_Embed1{},
		FMessage28:  &pbtypes.MessageCommon1_Embed1_Embed2{},
		FAny1:       &anypb.Any{},
		FAny2:       &anypb.Any{},
		FAny3:       &anypb.Any{},
		FAny4:       &anypb.Any{},
		FAny5:       &anypb.Any{},
		FAny6:       &anypb.Any{},
		FAny7:       &anypb.Any{},
		FAny8:       &anypb.Any{},
		FAny9:       &anypb.Any{},
	}
	data.SetDefault()

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, []byte(""), data.FBytes1)
		require.Equal(t, []byte("abc"), data.FBytes2)
	})

	t.Run("WKT", func(t *testing.T) {
		require.NotNil(t, data.FDuration1)
		require.Equal(t, int64(100), data.FDuration1.Seconds)
		require.Equal(t, int32(200), data.FDuration1.Nanos)

		require.NotNil(t, data.FDuration2)
		require.Equal(t, int64(0), data.FDuration2.Seconds)
		require.Equal(t, int32(0), data.FDuration2.Nanos)

		require.NotNil(t, data.FTimestamp1)
		require.Equal(t, int64(200), data.FTimestamp1.Seconds)
		require.Equal(t, int32(300), data.FTimestamp1.Nanos)

		require.NotNil(t, data.FTimestamp2)
		require.Equal(t, int64(0), data.FTimestamp2.Seconds)
		require.Equal(t, int32(0), data.FTimestamp2.Nanos)
	})

	t.Run("MESSAGE1", func(t *testing.T) {
		require.NotNil(t, data.FMessage11)
		require.Equal(t, "a", data.FMessage11.FString1)
		require.Equal(t, "b", data.FMessage11.FString2)
		require.Equal(t, "c", data.FMessage11.FString3)

		require.NotNil(t, data.FMessage12)
		require.Equal(t, "", data.FMessage12.FString1)
		require.Equal(t, "", data.FMessage12.FString2)
		require.Equal(t, "", data.FMessage12.FString3)

		require.NotNil(t, data.FMessage13)
		require.Equal(t, "1", data.FMessage13.FString1)
		require.Equal(t, "2", data.FMessage13.FString2)
		require.Equal(t, "3", data.FMessage13.FString3)

		require.NotNil(t, data.FMessage14)
		require.Equal(t, "", data.FMessage14.TString)
		require.Equal(t, (*string)(nil), data.FMessage14.PString)

		require.NotNil(t, data.FMessage15)
		require.Equal(t, "", data.FMessage15.TString1)
		require.Equal(t, "", data.FMessage15.TString2)

		require.NotNil(t, data.FMessage16)
		require.Equal(t, "", data.FMessage16.FString1)
		require.Equal(t, "", data.FMessage16.FString2)
		require.Equal(t, "", data.FMessage16.FString3)

		require.NotNil(t, data.FMessage17)
		require.Equal(t, "x", data.FMessage17.FString1)
		require.Equal(t, "y", data.FMessage17.FString2)
		require.Equal(t, "z", data.FMessage17.FString3)

		require.NotNil(t, data.FMessage18)
		require.Equal(t, "", data.FMessage18.FString1)
		require.Equal(t, "", data.FMessage18.FString2)
		require.Equal(t, "", data.FMessage18.FString3)
	})

	t.Run("MESSAGE2", func(t *testing.T) {
		require.NotNil(t, data.FMessage21)
		require.Equal(t, "a", data.FMessage21.FString1)
		require.Equal(t, "b", data.FMessage21.FString2)
		require.Equal(t, "c", data.FMessage21.FString3)

		require.NotNil(t, data.FMessage22)
		require.Equal(t, "", data.FMessage22.FString1)
		require.Equal(t, "", data.FMessage22.FString2)
		require.Equal(t, "", data.FMessage22.FString3)

		require.NotNil(t, data.FMessage23)
		require.Equal(t, "1", data.FMessage23.FString1)
		require.Equal(t, "2", data.FMessage23.FString2)
		require.Equal(t, "3", data.FMessage23.FString3)

		require.NotNil(t, data.FMessage24)
		require.Equal(t, "", data.FMessage24.TString)
		require.Equal(t, (*string)(nil), data.FMessage24.PString)

		require.NotNil(t, data.FMessage25)
		require.Equal(t, "", data.FMessage25.TString1)
		require.Equal(t, "", data.FMessage25.TString2)

		require.NotNil(t, data.FMessage26)
		require.Equal(t, "", data.FMessage26.FString1)
		require.Equal(t, "", data.FMessage26.FString2)
		require.Equal(t, "", data.FMessage26.FString3)

		require.NotNil(t, data.FMessage27)
		require.Equal(t, "x", data.FMessage27.FString1)
		require.Equal(t, "y", data.FMessage27.FString2)
		require.Equal(t, "z", data.FMessage27.FString3)

		require.NotNil(t, data.FMessage28)
		require.Equal(t, "", data.FMessage28.FString1)
		require.Equal(t, "", data.FMessage28.FString2)
		require.Equal(t, "", data.FMessage28.FString3)
	})

	t.Run("ANY", func(t *testing.T) {
		{
			require.NotNil(t, data.FAny1)
			anyVal := &pbtypes.MessagePlain1{}
			err := data.FAny1.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny2)
			anyVal := &pbtypes.MessagePlain1{}
			err := data.FAny2.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.FString1)
			require.Equal(t, "", anyVal.FString2)
			require.Equal(t, "", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny3)
			anyVal := &pbtypes.MessagePlain1_Embed1{}
			err := data.FAny3.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.FString1)
			require.Equal(t, "", anyVal.FString2)
			require.Equal(t, "", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny4)
			anyVal := &pbtypes.MessagePlain1_Embed1_Embed2{}
			err := data.FAny4.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "1", anyVal.FString1)
			require.Equal(t, "2", anyVal.FString2)
			require.Equal(t, "3", anyVal.FString3)
		}
		{
			require.NotNil(t, data.FAny5)
			anyVal := &pbexternal.External1{}
			err := data.FAny5.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString)
			require.Equal(t, (*string)(nil), anyVal.PString)
		}
		{
			require.NotNil(t, data.FAny6)
			anyVal := &pbexternal.External2_Embed1{}
			err := data.FAny6.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "", anyVal.TString1)
			require.Equal(t, "", anyVal.TString2)
		}
		{
			require.NotNil(t, data.FAny7)
			anyVal := &timestamppb.Timestamp{Seconds: 0, Nanos: 0}
			err := data.FAny7.UnmarshalTo(anyVal)
			require.Nil(t, err)
		}
		{
			require.NotNil(t, data.FAny8)
			anyVal := &anypb.Any{}
			err := data.FAny8.UnmarshalTo(anyVal)
			require.Nil(t, err)
		}
		{
			require.NotNil(t, data.FAny9)
			anyVal := &pbtypes.MessageCommon1{}
			err := data.FAny9.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "a", anyVal.FString1)
			require.Equal(t, "b", anyVal.FString2)
			require.Equal(t, "c", anyVal.FString3)
		}
	})
}

// Test cases the field value are have valid value, expected ignore default value.
func Test_TypePlain1_Ignore(t *testing.T) {
	data := &pbtypes.TypePlain1{
		FString1:    "101",
		FString2:    "102",
		FString3:    "103",
		FString4:    "104",
		FString5:    "105",
		FString6:    "106",
		FString7:    "107",
		FString8:    "108",
		FString9:    "109",
		FInt32:      101,
		FInt64:      102,
		FUint32:     103,
		FUint64:     104,
		FSint32:     105,
		FSint64:     106,
		FSfixed32:   107,
		FSfixed64:   108,
		FFixed32:    109,
		FFixed64:    110,
		FFloat:      111.111,
		FDouble:     112.112,
		FBool1:      true,
		FBool2:      true,
		FBytes1:     []byte("bytes1"),
		FBytes2:     []byte("bytes2"),
		FEnum1:      201,
		FEnum2:      202,
		FEnum3:      203,
		FEnum4:      204,
		FEnum5:      205,
		FEnum6:      206,
		FEnum7:      207,
		FDuration1:  &durationpb.Duration{Seconds: 300, Nanos: 301},
		FDuration2:  &durationpb.Duration{Seconds: 400, Nanos: 401},
		FTimestamp1: &timestamppb.Timestamp{Seconds: 500, Nanos: 501},
		FTimestamp2: &timestamppb.Timestamp{Seconds: 600, Nanos: 601},
		FMessage11:  &pbtypes.MessagePlain1{FString1: "1101", FString2: "1102", FString3: "1103"},
		FMessage12:  &pbtypes.MessagePlain1_Embed1{FString1: "1201", FString2: "1202", FString3: "1203"},
		FMessage13:  &pbtypes.MessagePlain1_Embed1_Embed2{FString1: "1301", FString2: "1302", FString3: "1303"},
		FMessage14:  &pbexternal.External1{},
		FMessage15:  &pbexternal.External2_Embed1{TString1: "1501", TString2: "1502"},
		FMessage16:  &pbtypes.MessageCommon1{FString1: "1601", FString2: "1602", FString3: "1603"},
		FMessage17:  &pbtypes.MessageCommon1_Embed1{FString1: "1701", FString2: "1702", FString3: "1703"},
		FMessage18:  &pbtypes.MessageCommon1_Embed1_Embed2{FString1: "1801", FString2: "1802", FString3: "1803"},
		FMessage21:  &pbtypes.MessagePlain1{FString1: "2101", FString2: "2102", FString3: "2103"},
		FMessage22:  &pbtypes.MessagePlain1_Embed1{FString1: "2201", FString2: "2202", FString3: "2203"},
		FMessage23:  &pbtypes.MessagePlain1_Embed1_Embed2{FString1: "2301", FString2: "2302", FString3: "2303"},
		FMessage24:  &pbexternal.External1{},
		FMessage25:  &pbexternal.External2_Embed1{TString1: "2501", TString2: "2502"},
		FMessage26:  &pbtypes.MessageCommon1{FString1: "2601", FString2: "2602", FString3: "2603"},
		FMessage27:  &pbtypes.MessageCommon1_Embed1{FString1: "2701", FString2: "2702", FString3: "2703"},
		FMessage28:  &pbtypes.MessageCommon1_Embed1_Embed2{FString1: "2801", FString2: "2802", FString3: "2803"},
		FAny1:       mustNewAny(&pbexternal.External3{TString: "any1"}),
		FAny2:       mustNewAny(&pbexternal.External3{TString: "any2"}),
		FAny3:       mustNewAny(&pbexternal.External3{TString: "any3"}),
		FAny4:       mustNewAny(&pbexternal.External3{TString: "any4"}),
		FAny5:       mustNewAny(&pbexternal.External3{TString: "any5"}),
		FAny6:       mustNewAny(&pbexternal.External3{TString: "any6"}),
		FAny7:       mustNewAny(&pbexternal.External3{TString: "any7"}),
		FAny8:       mustNewAny(&pbexternal.External3{TString: "any8"}),
		FAny9:       mustNewAny(&pbexternal.External3{TString: "any9"}),
	}

	data.SetDefault()

	t.Run("STRING", func(t *testing.T) {
		require.Equal(t, "101", data.FString1)
		require.Equal(t, "102", data.FString2)
		require.Equal(t, "103", data.FString3)
		require.Equal(t, "104", data.FString4)
		require.Equal(t, "105", data.FString5)
		require.Equal(t, "106", data.FString6)
		require.Equal(t, "107", data.FString7)
		require.Equal(t, "108", data.FString8)
		require.Equal(t, "109", data.FString9)
	})

	t.Run("NUMBER", func(t *testing.T) {
		require.Equal(t, int32(101), data.FInt32)
		require.Equal(t, int64(102), data.FInt64)
		require.Equal(t, uint32(103), data.FUint32)
		require.Equal(t, uint64(104), data.FUint64)
		require.Equal(t, int32(105), data.FSint32)
		require.Equal(t, int64(106), data.FSint64)
		require.Equal(t, int32(107), data.FSfixed32)
		require.Equal(t, int64(108), data.FSfixed64)
		require.Equal(t, uint32(109), data.FFixed32)
		require.Equal(t, uint64(110), data.FFixed64)
		require.Equal(t, float32(111.111), data.FFloat)
		require.Equal(t, float64(112.112), data.FDouble)
	})

	t.Run("BOOL", func(t *testing.T) {
		require.Equal(t, true, data.FBool1)
		require.Equal(t, true, data.FBool2)
	})

	t.Run("BYTES", func(t *testing.T) {
		require.Equal(t, []byte("bytes1"), data.FBytes1)
		require.Equal(t, []byte("bytes2"), data.FBytes2)
	})

	t.Run("ENUM", func(t *testing.T) {
		require.Equal(t, pbtypes.EnumPlain1(201), data.FEnum1)
		require.Equal(t, pbtypes.EnumPlain1(202), data.FEnum2)
		require.Equal(t, pbexternal.Month1(203), data.FEnum3)
		require.Equal(t, pbexternal.Month2(204), data.FEnum4)
		require.Equal(t, pbexternal.EnumWeek1_Week(205), data.FEnum5)
		require.Equal(t, pbexternal.EnumWeek2_Embed1_Week(206), data.FEnum6)
		require.Equal(t, pbtypes.EnumCommon1(207), data.FEnum7)
	})

	t.Run("WKT", func(t *testing.T) {
		require.NotNil(t, data.FDuration1)
		require.Equal(t, int64(300), data.FDuration1.Seconds)
		require.Equal(t, int32(301), data.FDuration1.Nanos)

		require.NotNil(t, data.FDuration2)
		require.Equal(t, int64(400), data.FDuration2.Seconds)
		require.Equal(t, int32(401), data.FDuration2.Nanos)

		require.NotNil(t, data.FTimestamp1)
		require.Equal(t, int64(500), data.FTimestamp1.Seconds)
		require.Equal(t, int32(501), data.FTimestamp1.Nanos)

		require.NotNil(t, data.FTimestamp2)
		require.Equal(t, int64(600), data.FTimestamp2.Seconds)
		require.Equal(t, int32(601), data.FTimestamp2.Nanos)
	})

	t.Run("MESSAGE1", func(t *testing.T) {
		require.NotNil(t, data.FMessage11)
		require.Equal(t, "1101", data.FMessage11.FString1)
		require.Equal(t, "1102", data.FMessage11.FString2)
		require.Equal(t, "1103", data.FMessage11.FString3)

		require.NotNil(t, data.FMessage12)
		require.Equal(t, "1201", data.FMessage12.FString1)
		require.Equal(t, "1202", data.FMessage12.FString2)
		require.Equal(t, "1203", data.FMessage12.FString3)

		require.NotNil(t, data.FMessage13)
		require.Equal(t, "1301", data.FMessage13.FString1)
		require.Equal(t, "1302", data.FMessage13.FString2)
		require.Equal(t, "1303", data.FMessage13.FString3)

		require.NotNil(t, data.FMessage14)
		require.Equal(t, "", data.FMessage14.TString)
		require.Equal(t, (*string)(nil), data.FMessage14.PString)

		require.NotNil(t, data.FMessage15)
		require.Equal(t, "1501", data.FMessage15.TString1)
		require.Equal(t, "1502", data.FMessage15.TString2)

		require.NotNil(t, data.FMessage16)
		require.Equal(t, "1601", data.FMessage16.FString1)
		require.Equal(t, "1602", data.FMessage16.FString2)
		require.Equal(t, "1603", data.FMessage16.FString3)

		require.NotNil(t, data.FMessage17)
		require.Equal(t, "1701", data.FMessage17.FString1)
		require.Equal(t, "1702", data.FMessage17.FString2)
		require.Equal(t, "1703", data.FMessage17.FString3)

		require.NotNil(t, data.FMessage18)
		require.Equal(t, "1801", data.FMessage18.FString1)
		require.Equal(t, "1802", data.FMessage18.FString2)
		require.Equal(t, "1803", data.FMessage18.FString3)
	})

	t.Run("MESSAGE2", func(t *testing.T) {
		require.NotNil(t, data.FMessage21)
		require.Equal(t, "2101", data.FMessage21.FString1)
		require.Equal(t, "2102", data.FMessage21.FString2)
		require.Equal(t, "2103", data.FMessage21.FString3)

		require.NotNil(t, data.FMessage22)
		require.Equal(t, "2201", data.FMessage22.FString1)
		require.Equal(t, "2202", data.FMessage22.FString2)
		require.Equal(t, "2203", data.FMessage22.FString3)

		require.NotNil(t, data.FMessage23)
		require.Equal(t, "2301", data.FMessage23.FString1)
		require.Equal(t, "2302", data.FMessage23.FString2)
		require.Equal(t, "2303", data.FMessage23.FString3)

		require.NotNil(t, data.FMessage24)
		require.Equal(t, "", data.FMessage24.TString)
		require.Equal(t, (*string)(nil), data.FMessage24.PString)

		require.NotNil(t, data.FMessage25)
		require.Equal(t, "2501", data.FMessage25.TString1)
		require.Equal(t, "2502", data.FMessage25.TString2)

		require.NotNil(t, data.FMessage26)
		require.Equal(t, "2601", data.FMessage26.FString1)
		require.Equal(t, "2602", data.FMessage26.FString2)
		require.Equal(t, "2603", data.FMessage26.FString3)

		require.NotNil(t, data.FMessage27)
		require.Equal(t, "2701", data.FMessage27.FString1)
		require.Equal(t, "2702", data.FMessage27.FString2)
		require.Equal(t, "2703", data.FMessage27.FString3)

		require.NotNil(t, data.FMessage28)
		require.Equal(t, "2801", data.FMessage28.FString1)
		require.Equal(t, "2802", data.FMessage28.FString2)
		require.Equal(t, "2803", data.FMessage28.FString3)
	})

	t.Run("ANY", func(t *testing.T) {
		{
			require.NotNil(t, data.FAny1)
			anyVal := &pbexternal.External3{}
			err := data.FAny1.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any1", anyVal.TString)
		}
		{
			require.NotNil(t, data.FAny2)
			anyVal := &pbexternal.External3{}
			err := data.FAny2.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any2", anyVal.TString)
		}
		{
			require.NotNil(t, data.FAny3)
			anyVal := &pbexternal.External3{}
			err := data.FAny3.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any3", anyVal.TString)
		}
		{
			require.NotNil(t, data.FAny4)
			anyVal := &pbexternal.External3{}
			err := data.FAny4.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any4", anyVal.TString)
		}
		{
			require.NotNil(t, data.FAny5)
			anyVal := &pbexternal.External3{}
			err := data.FAny5.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any5", anyVal.TString)
		}
		{
			require.NotNil(t, data.FAny6)
			anyVal := &pbexternal.External3{}
			err := data.FAny6.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any6", anyVal.TString)
		}
		{
			require.NotNil(t, data.FAny7)
			anyVal := &pbexternal.External3{}
			err := data.FAny7.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any7", anyVal.TString)
		}
		{
			require.NotNil(t, data.FAny8)
			anyVal := &pbexternal.External3{}
			err := data.FAny8.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any8", anyVal.TString)
		}
		{
			require.NotNil(t, data.FAny9)
			anyVal := &pbexternal.External3{}
			err := data.FAny9.UnmarshalTo(anyVal)
			require.Nil(t, err)
			require.Equal(t, "any9", anyVal.TString)
		}
	})
}
