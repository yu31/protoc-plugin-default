// Code generated by protoc-gen-godefault. DO NOT EDIT.
// versions:
// 		protoc-gen-godefault 0.0.1
// source: tests/proto/cases/empty/empty_common.proto

package pbempty

import (
	_ "github.com/yu31/protoc-plugin-default/xgo/pb/pbdefault"
)

// Set default value for message MessageCommon1 that in file tests/proto/cases/empty/empty_common.proto
func (x *MessageCommon1) SetDefault() {
	if x == nil {
		return
	}
}

// Set default value for message Embed1 that in file tests/proto/cases/empty/empty_common.proto
func (x *MessageCommon1_Embed1) SetDefault() {
	if x == nil {
		return
	}
}

// Set default value for message Embed2 that in file tests/proto/cases/empty/empty_common.proto
func (x *MessageCommon1_Embed1_Embed2) SetDefault() {
	if x == nil {
		return
	}
}