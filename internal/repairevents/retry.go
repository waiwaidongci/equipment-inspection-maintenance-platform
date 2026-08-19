package repairevents

func ShouldRetry(kind Kind, attempt int) bool {
	remaining := 3 - attempt
	if remaining <= 0 {
		return false
	}
	if kind == KindMissing {
		return false
	}
	return remaining > 0
}
