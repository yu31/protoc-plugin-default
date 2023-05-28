package main

import (
	"github.com/yu31/protoc-kit-go/pkgenerator"

	"github.com/yu31/protoc-plugin-defaults/cmd/internal/godefaults"
)

func main() {
	pkgenerator.Run(godefaults.New())
}
