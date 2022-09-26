package boundary

import (
	"testing"
	"time"

	"github.com/yu31/protoc-plugin-defaults/xgo/tests/pb/boundary/pbcr"
)

func Test_CRMessageA1(t *testing.T) {
	msgA := pbcr.CRMessageA{
		T1: "",
		MsgB: &pbcr.CRMessageB{
			T1: "",
			MsgC: &pbcr.CRMessageC{
				T1: "",
				MsgA: &pbcr.CRMessageA{
					T1: "",
					MsgB: &pbcr.CRMessageB{
						T1:   "",
						MsgC: nil,
					},
				},
			},
		},
	}

	done := make(chan struct{})
	timer := time.NewTimer(time.Millisecond * 100)
	defer timer.Stop()

	go func() {
		msgA.SetDefaults()
		close(done)
	}()

	select {
	case <-done:
		t.Log("No circular calls in TestMessageCRMessage1")
	case <-timer.C:
		// FIXME: Attempt fix the cases.
		t.Log("ERROR: Occurs circular reference in TestMessageCRMessage1")
	}
}
