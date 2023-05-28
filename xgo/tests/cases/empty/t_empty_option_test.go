package empty

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-default/xgo/tests/pb/pbempty"
)

func Test_Empty_TypeMap1(t *testing.T) {
	data := &pbempty.TypeMap1{}

	require.NotPanics(t, func() {
		data.SetDefault()
	})
}

func Test_Empty_TypeOneOf1(t *testing.T) {
	data := &pbempty.TypeOneOf1{}

	require.NotPanics(t, func() {
		data.SetDefault()
	})
}

func Test_Empty_TypeOneOf2(t *testing.T) {
	data := &pbempty.TypeOneOf2{}

	require.NotPanics(t, func() {
		data.SetDefault()
	})
}

func Test_Empty_TypeOptional1(t *testing.T) {
	data := &pbempty.TypeOptional1{}

	require.NotPanics(t, func() {
		data.SetDefault()
	})
}

func Test_Empty_TypePlain1(t *testing.T) {
	data := &pbempty.TypePlain1{}

	require.NotPanics(t, func() {
		data.SetDefault()
	})
}

func Test_Empty_TypeRepeated1(t *testing.T) {
	data := &pbempty.TypeRepeated1{}

	require.NotPanics(t, func() {
		data.SetDefault()
	})
}
