package g_evidencegrade

import (
	"errors"
	core "github.com/example/inspection-platform/internal/evidencebatch"
	"testing"
)

func TestEvidenceResourcesCloseAndFailuresJoin(t *testing.T) {
	batch := core.NewBatch(core.Repository{CommitError: core.ErrCommit})
	err := batch.RunBatch([]string{"first", "invalid", "third"})
	if !errors.Is(err, core.ErrInvalidEvidence) || !errors.Is(err, core.ErrCommit) {
		t.Errorf("errors=%v", err)
	}
	if len(batch.Resources()) != 3 {
		t.Fatalf("resources=%d", len(batch.Resources()))
	}
	for _, r := range batch.Resources() {
		if !r.Closed() {
			t.Error("resource open")
		}
	}
	if !errors.Is((core.Service{}).Validate(""), core.ErrInvalidEvidence) {
		t.Error("empty evidence accepted")
	}
}

func TestBadEvidenceReturnsAnErrorValue(t *testing.T) {
	if !errors.Is((core.Repository{}).Save("invalid"), core.ErrInvalidEvidence) {
		t.Error("invalid identity lost")
	}
}
