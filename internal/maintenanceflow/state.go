package maintenanceflow

type State string

const (
	Planned    State = "planned"
	InProgress State = "in_progress"
	Rechecking State = "rechecking"
	Closed     State = "closed"
)

func (s State) Open() bool { return s == Planned || s == InProgress }
