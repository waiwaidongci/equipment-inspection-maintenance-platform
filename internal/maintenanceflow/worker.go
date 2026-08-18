package maintenanceflow

func RetryPassed(s *Service) State {
	if s.Move(InProgress, Rechecking) {
		return Rechecking
	}
	return InProgress
}
