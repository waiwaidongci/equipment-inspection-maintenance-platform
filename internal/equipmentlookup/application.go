package equipmentlookup

import "errors"

type Result string

const (
	Missing Result = "missing"
	Broken  Result = "broken"
)

func Decide(err error) Result {
	if errors.Is(err, ErrMissing) {
		return Missing
	}
	return Broken
}
