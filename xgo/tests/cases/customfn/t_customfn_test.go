package customfn

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-default/xgo/tests/pb/pbcustomfn"
)

func Test_CustomMethodFuncName1(t *testing.T) {
	data := pbcustomfn.CustomMethodFuncName1{F1: "a"}

	require.NotPanics(t, func() {
		data.InitField()
	})
}

func Test_CustomMethodFuncName2(t *testing.T) {
	data := pbcustomfn.CustomMethodFuncName2{
		F1: "123",
		M1: nil,
	}
	require.NotPanics(t, func() {
		data.InitField()
	})
}
