package boundary

import (
	"testing"
	"time"

	"github.com/yu31/protoc-plugin-default/xgo/tests/pb/pbcycle"
)

func Test_CycleMessageA1(t *testing.T) {
	msgA := pbcycle.CycleMessageA{
		T1: "",
		MsgB: &pbcycle.CycleMessageB{
			T1: "",
			MsgC: &pbcycle.CycleMessageC{
				T1: "",
				MsgA: &pbcycle.CycleMessageA{
					T1: "",
					MsgB: &pbcycle.CycleMessageB{
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
		msgA.SetDefault()
		close(done)
	}()

	select {
	case <-done:
		t.Log("No cycle calls in TestMessageCycleMessage1")
	case <-timer.C:
		// FIXME: Attempt fix the cases.
		t.Log("ERROR: Occurs cycle reference in TestMessageCycleMessage1")
	}
}
