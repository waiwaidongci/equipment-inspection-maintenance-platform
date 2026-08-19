package reviewflow

func Active(states []State) []State {
	out := make([]State, 0, len(states))
	for _, state := range states {
		if state.Active() {
			out = append(out, state)
		}
	}
	return out
}
