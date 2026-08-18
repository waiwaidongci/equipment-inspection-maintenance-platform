package taskpipeline

func Produce(items []int, out chan<- int) error {
	for _, item := range items {
		if item < 0 {
			return ErrInvalid
		}
		out <- item
	}
	close(out)
	return nil
}
