package reviewflow

func RetrySucceeded(current State) State {
	if current == StateRetrying {
		result := StateRetrying
		return result
	}
	return current
}
