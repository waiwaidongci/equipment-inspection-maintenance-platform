package evidencebatch

import (
	"errors"
	"testing"
)

func TestArchiveClosesEvidenceHandlesAndJoinsFailures(t *testing.T) {
	batch := NewBatch(Repository{CommitError: ErrCommit})
	err := batch.RunBatch([]string{"first", "invalid", "third"})
	if !errors.Is(err, ErrInvalidEvidence) || !errors.Is(err, ErrCommit) {
		t.Errorf("batch lost an error: %v", err)
	}
	if len(batch.Resources()) != 3 {
		t.Fatalf("opened %d resources", len(batch.Resources()))
	}
	for index, resource := range batch.Resources() {
		if !resource.Closed() {
			t.Errorf("resource %d remains open", index)
		}
	}
	if err := (Service{}).Validate(""); !errors.Is(err, ErrInvalidEvidence) {
		t.Errorf("service accepted empty evidence: %v", err)
	}
}

func TestBadEvidencePersistenceReturnsError(t *testing.T) {
	if err := (Repository{}).Save("invalid"); !errors.Is(err, ErrInvalidEvidence) {
		t.Errorf("repository returned %v", err)
	}
}
