package equipmentlookup

type Result string

const (
	Missing Result = "missing"
	Broken  Result = "broken"
)

func Decide(err error) Result {
	if err == ErrMissing {
		return Missing
	}
	return Broken
}
