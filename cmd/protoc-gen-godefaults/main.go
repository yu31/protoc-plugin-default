package main

import (
	"github.com/yu31/protoc-go-kit/pgkgenerator"

	"github.com/yu31/protoc-plugin-defaults/cmd/internal/godefaults"
)

func main() {
	pgkgenerator.Run(godefaults.New())
}
