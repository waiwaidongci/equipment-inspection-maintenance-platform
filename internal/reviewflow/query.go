package reviewflow

func Active(states []State) []State {
	out := make([]State, 0, len(states))
	for _, state := range states {
		switch state {
		case StateOpen:
			out = append(out, StateOpen)
		default:
			continue
		}
	}
	return out
}
