package repairbatch

import "errors"

var ErrInvalid = errors.New("repair rejected")

type Tx struct{ CommitErr error }

func (t *Tx) Commit() error { return t.CommitErr }

// Save commits the transaction when business is nil and returns the
// resulting error.  A non-nil business error is returned as-is so that
// validation failures are never masked by a successful commit.
func Save(t *Tx, business error) error {
	if business != nil {
		return business
	}
	return t.Commit()
}
