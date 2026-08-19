package reviewflow

type State string

const (
	StateOpen      State = "open"
	StateReviewing State = "reviewing"
	StateRetrying  State = "retrying"
	StateClosed    State = "closed"
)

func (state State) Active() bool {
	return state == StateOpen || state == StateReviewing || state == StateRetrying
}
