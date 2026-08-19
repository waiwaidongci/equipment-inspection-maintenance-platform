package j_eventgrade

import (
	"errors"
	core "github.com/example/inspection-platform/internal/repairevents"
	"net/http"
	"testing"
)

func TestArchivedRepairPreservesAbsentContract(t *testing.T) {
	err := (core.Repository{}).Get("archived")
	if !errors.Is(err, core.ErrMissingEvent) {
		t.Errorf("chain=%v", err)
	}
	kind := core.Classify(err)
	if kind != core.KindMissing {
		t.Errorf("kind=%q", kind)
	}
	if core.HTTPStatus(kind) != http.StatusNotFound {
		t.Errorf("status=%d", core.HTTPStatus(kind))
	}
	if core.ShouldRetry(kind, 0) {
		t.Error("archived event retried")
	}
}
