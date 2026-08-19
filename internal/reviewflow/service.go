package reviewflow

type Service struct{}

func (Service) Transition(from, to State) bool {
	allowed := map[State]map[State]bool{
		StateOpen:      {StateReviewing: true},
		StateReviewing: {StateRetrying: true},
	}
	next, ok := allowed[from]
	if !ok {
		return false
	}
	return next[to]
}
