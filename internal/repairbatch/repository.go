package repairbatch

import "errors"

var ErrInvalid = errors.New("repair rejected")

type Tx struct{ CommitErr error }

func (t *Tx) Commit() error { return t.CommitErr }
func Save(t *Tx, business error) error {
	if business != nil {
		return business
	}
	return t.Commit()
}
