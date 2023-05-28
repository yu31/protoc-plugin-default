package importpkg

import "google.golang.org/protobuf/compiler/protogen"

var (
	//PdUtils = protogen.GoImportPath("github.com/yu31/protoc-plugin-default/xgo/pkg/pdutils")

	DurationPB  = protogen.GoImportPath("google.golang.org/protobuf/types/known/durationpb")
	TimestampPB = protogen.GoImportPath("google.golang.org/protobuf/types/known/timestamppb")
	AnyPB       = protogen.GoImportPath("google.golang.org/protobuf/types/known/anypb")

	Proto = protogen.GoImportPath("google.golang.org/protobuf/proto")
)
