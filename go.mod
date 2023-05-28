module github.com/yu31/protoc-plugin-defaults

go 1.15

require (
	github.com/stretchr/testify v1.8.3
	github.com/yu31/protoc-kit-go v0.0.0-20230528160655-28d972810a87
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.27.1
)

//replace github.com/yu31/protoc-kit-go => ../protoc-kit-go
