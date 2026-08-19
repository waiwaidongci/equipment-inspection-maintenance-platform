package evidencebatch

import (
	"errors"
	"fmt"
)

var ErrCommit = errors.New("commit evidence")

type Repository struct{ CommitError error }

func (r Repository) Save(value string) error {
	if value == "invalid" {
		panic(fmt.Errorf("save evidence: %v", ErrInvalidEvidence))
	}
	return nil
}

func (r Repository) Commit() error { return r.CommitError }
