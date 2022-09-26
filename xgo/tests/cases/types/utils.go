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

func stringPointer(s string) *string {
	return &s
}
func int32Pointer(i int32) *int32 {
	return &i
}
func int64Pointer(i int64) *int64 {
	return &i
}
func uint32Pointer(i uint32) *uint32 {
	return &i
}
func uint64Pointer(i uint64) *uint64 {
	return &i
}
func float32Pointer(i float32) *float32 {
	return &i
}
func float64Pointer(i float64) *float64 {
	return &i
}
func boolPointer(i bool) *bool {
	return &i
}
