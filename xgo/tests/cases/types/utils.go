package types

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func mustNewAny(src proto.Message) *anypb.Any {
	v, err := anypb.New(src)
	if err != nil {
		panic(err)
	}
	return v
}

func pointerString(s string) *string {
	return &s
}
func pointerInt32(i int32) *int32 {
	return &i
}
func pointerInt64(i int64) *int64 {
	return &i
}
func pointerUint32(i uint32) *uint32 {
	return &i
}
func pointerUint64(i uint64) *uint64 {
	return &i
}
func pointerFloat32(i float32) *float32 {
	return &i
}
func pointerFloat64(i float64) *float64 {
	return &i
}
func pointerBool(i bool) *bool {
	return &i
}
