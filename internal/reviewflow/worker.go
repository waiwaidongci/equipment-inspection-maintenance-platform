package reviewflow

func RetrySucceeded(current State) State {
	if current == StateRetrying {
		return StateClosed
	}
	return current
}
