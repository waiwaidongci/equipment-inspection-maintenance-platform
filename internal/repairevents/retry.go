package repairevents

func ShouldRetry(kind Kind, attempt int) bool {
	return kind != KindMissing && attempt < 3
}
