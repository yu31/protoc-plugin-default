// Code generated by protoc-gen-godefault. DO NOT EDIT.
// versions:
// 		protoc-gen-godefault 0.0.1
// source: tests/proto/cases/empty/empty_optional.proto

package pbempty

import (
	_ "github.com/yu31/protoc-plugin-default/xgo/pb/pbdefault"
	_ "github.com/yu31/protoc-plugin-default/xgo/tests/pb/pbexternal"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Set default value for message MessageOptional1 that in file tests/proto/cases/empty/empty_optional.proto
func (x *MessageOptional1) SetDefault() {
	if x == nil {
		return
	}
}

// Set default value for message Embed1 that in file tests/proto/cases/empty/empty_optional.proto
func (x *MessageOptional1_Embed1) SetDefault() {
	if x == nil {
		return
	}
}

// Set default value for message Embed2 that in file tests/proto/cases/empty/empty_optional.proto
func (x *MessageOptional1_Embed1_Embed2) SetDefault() {
	if x == nil {
		return
	}
}

// Set default value for message TypeOptional1 that in file tests/proto/cases/empty/empty_optional.proto
func (x *TypeOptional1) SetDefault() {
	if x == nil {
		return
	}
	if x.FEnum1 != nil {
		x.FEnum1.SetDefault()
	}
	if x.FEnum2 != nil {
		x.FEnum2.SetDefault()
	}
	if x.FMessage11 != nil {
		x.FMessage11.SetDefault()
	}
	if x.FMessage12 != nil {
		x.FMessage12.SetDefault()
	}
	if x.FMessage13 != nil {
		x.FMessage13.SetDefault()
	}
	if x.FMessage14 != nil {
		x.FMessage14.SetDefault()
	}
	if x.FMessage15 != nil {
		x.FMessage15.SetDefault()
	}
	if x.FMessage16 != nil {
		x.FMessage16.SetDefault()
	}
	if x.FMessage17 != nil {
		x.FMessage17.SetDefault()
	}
	if x.FMessage18 != nil {
		x.FMessage18.SetDefault()
	}
	if x.FMessage21 != nil {
		x.FMessage21.SetDefault()
	}
	if x.FMessage22 != nil {
		x.FMessage22.SetDefault()
	}
	if x.FMessage23 != nil {
		x.FMessage23.SetDefault()
	}
	if x.FMessage24 != nil {
		x.FMessage24.SetDefault()
	}
	if x.FMessage25 != nil {
		x.FMessage25.SetDefault()
	}
	if x.FMessage26 != nil {
		x.FMessage26.SetDefault()
	}
	if x.FMessage27 != nil {
		x.FMessage27.SetDefault()
	}
	if x.FMessage28 != nil {
		x.FMessage28.SetDefault()
	}
}
