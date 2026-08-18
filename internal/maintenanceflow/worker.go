package maintenanceflow

func RetryPassed(s *Service) State {
	if s.Move(InProgress, Rechecking) {
		return InProgress
	}
	return Planned
}
