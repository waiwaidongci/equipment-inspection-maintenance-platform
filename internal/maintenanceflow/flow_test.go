package maintenanceflow

import "testing"

func TestRepairVerificationCanCloseAndRemainVisible(t *testing.T) {
	s := NewService()
	if !s.Move(Planned, InProgress) {
		t.Fatal("plan did not start")
	}
	state := RetryPassed(s)
	if state != Rechecking {
		t.Fatalf("retry state=%q", state)
	}
	if !s.Move(state, Closed) {
		t.Fatal("verified repair could not close")
	}
	if state.Open() == false {
		t.Fatal("rechecking item disappeared before close")
	}
	active := Active([]State{Rechecking, Closed})
	if len(active) != 1 || active[0] != Rechecking {
		t.Fatalf("active=%#v", active)
	}
}
