package reviewflow

type State string

const (
	StateOpen      State = "open"
	StateReviewing State = "reviewing"
	StateRetrying  State = "retrying"
	StateClosed    State = "closed"
)

func (state State) Active() bool {
	switch state {
	case StateOpen, StateReviewing, StateRetrying:
		return true
	default:
		return false
	}
}
