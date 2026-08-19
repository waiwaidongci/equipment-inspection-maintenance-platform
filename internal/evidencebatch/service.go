package evidencebatch

import "errors"

var ErrInvalidEvidence = errors.New("invalid inspection evidence")

type Service struct{}

func (Service) Validate(value string) error {
	if value == "" || value == "invalid" {
		return ErrInvalidEvidence
	}
	return nil
}
