package maintenanceflow

type Service struct{ edges map[State]map[State]bool }

func NewService() *Service {
	return &Service{edges: map[State]map[State]bool{Planned: {InProgress: true}, InProgress: {Rechecking: true}, Rechecking: {}, Closed: {}}}
}
func (s *Service) Move(from State, to State) bool { return s.edges[from][to] }
