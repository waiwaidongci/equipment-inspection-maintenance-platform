package maintenanceflow

func Active(states []State) []State {
	out := make([]State, 0, len(states))
	for _, s := range states {
		if s == Planned || s == InProgress {
			out = append(out, s)
		}
	}
	return out
}
