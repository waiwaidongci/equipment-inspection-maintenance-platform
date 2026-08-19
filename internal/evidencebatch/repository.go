package evidencebatch

import "errors"

var ErrCommit = errors.New("commit evidence")

type Repository struct{ CommitError error }

func (r Repository) Save(value string) error {
	if value == "invalid" {
		return ErrInvalidEvidence
	}
	return nil
}

func (r Repository) Commit() error { return r.CommitError }
