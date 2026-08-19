package h_reviewgrade

import (
	core "github.com/example/inspection-platform/internal/reviewflow"
	"testing"
)

func TestReviewRecoveryEndsClosedAndQueryable(t *testing.T) {
	if !core.StateReviewing.Active() || !core.StateRetrying.Active() || core.StateClosed.Active() {
		t.Error("activity contract")
	}
	s := core.Service{}
	if !s.Transition(core.StateOpen, core.StateReviewing) || !s.Transition(core.StateReviewing, core.StateRetrying) || !s.Transition(core.StateRetrying, core.StateClosed) {
		t.Error("transition graph")
	}
	if core.RetrySucceeded(core.StateRetrying) != core.StateClosed {
		t.Error("retry not closed")
	}
	active := core.Active([]core.State{core.StateOpen, core.StateReviewing, core.StateRetrying, core.StateClosed})
	if len(active) != 3 || active[2] != core.StateRetrying {
		t.Errorf("active=%v", active)
	}
}
