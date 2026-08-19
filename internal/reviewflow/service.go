package reviewflow

type Service struct{}

func (Service) Transition(from, to State) bool {
	allowed := map[State]map[State]bool{
		StateOpen:      {StateReviewing: true},
		StateReviewing: {StateRetrying: true, StateClosed: true},
		StateRetrying:  {StateClosed: true},
	}
	return allowed[from][to]
}
