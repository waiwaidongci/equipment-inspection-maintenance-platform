package checklist

import "testing"

func TestEmptyPlanDoesNotPanicOrBypassNilChecks(t *testing.T) {
	s := DefaultSettings()
	if s.Items == nil {
		t.Error("default checklist map is nil")
	}
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("add panicked: %v", r)
			}
		}()
		AddItem(&s, "power")
	}()
	if OptionalValidator() != nil {
		t.Error("optional validator is a typed nil")
	}
	var concrete *required
	var v Validator = concrete
	if !Disabled(v) {
		t.Error("typed nil was treated as active")
	}
}
