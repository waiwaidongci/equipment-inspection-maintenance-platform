package taskpipeline

func Produce(items []int, out chan<- int) error {
	defer close(out)
	for _, item := range items {
		if item < 0 {
			return ErrInvalid
		}
		out <- item
	}
	return nil
}
