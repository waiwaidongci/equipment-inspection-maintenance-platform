package reviewflow

import "testing"

func TestReviewRetryClosesAndRemainsQueryable(t *testing.T) {
	if !StateReviewing.Active() || !StateRetrying.Active() || StateClosed.Active() {
		t.Error("state activity contract is inconsistent")
	}
	service := Service{}
	if !service.Transition(StateOpen, StateReviewing) || !service.Transition(StateReviewing, StateRetrying) || !service.Transition(StateRetrying, StateClosed) {
		t.Error("review transition graph is incomplete")
	}
	if got := RetrySucceeded(StateRetrying); got != StateClosed {
		t.Errorf("retry success remained %q", got)
	}
	active := Active([]State{StateOpen, StateReviewing, StateRetrying, StateClosed})
	if len(active) != 3 || active[2] != StateRetrying {
		t.Errorf("active query lost retrying: %#v", active)
	}
}
