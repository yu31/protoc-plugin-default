// Code generated by protoc-gen-godefault. DO NOT EDIT.
// versions:
// 		protoc-gen-godefault 0.0.1
// source: tests/proto/cases/empty/empty_map.proto

package pbempty

import (
	_ "github.com/yu31/protoc-plugin-default/xgo/pb/pbdefault"
	_ "github.com/yu31/protoc-plugin-default/xgo/tests/pb/pbexternal"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Set default value for message MessageMap1 that in file tests/proto/cases/empty/empty_map.proto
func (x *MessageMap1) SetDefault() {
	if x == nil {
		return
	}
}

// Set default value for message Embed1 that in file tests/proto/cases/empty/empty_map.proto
func (x *MessageMap1_Embed1) SetDefault() {
	if x == nil {
		return
	}
}

// Set default value for message Embed2 that in file tests/proto/cases/empty/empty_map.proto
func (x *MessageMap1_Embed1_Embed2) SetDefault() {
	if x == nil {
		return
	}
}

// Set default value for message TypeMap1 that in file tests/proto/cases/empty/empty_map.proto
func (x *TypeMap1) SetDefault() {
	if x == nil {
		return
	}
	for _, mv := range x.FMessage3 {
		if mv != nil {
			mv.SetDefault()
		}
	}
	for _, mv := range x.FMessage5 {
		if mv != nil {
			mv.SetDefault()
		}
	}
	for _, mv := range x.FMessage6 {
		if mv != nil {
			mv.SetDefault()
		}
	}
}

// Set default value for message TypeMap3 that in file tests/proto/cases/empty/empty_map.proto
func (x *TypeMap3) SetDefault() {
	if x == nil {
		return
	}
}
