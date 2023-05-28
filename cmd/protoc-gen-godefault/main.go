package main

import (
	"github.com/yu31/protoc-kit-go/pkgenerator"

	"github.com/yu31/protoc-plugin-default/cmd/internal/godefault"
)

func main() {
	pkgenerator.Run(godefault.New())
}
