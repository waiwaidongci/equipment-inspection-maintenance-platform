package taskpipeline

import "context"

func Consume(ctx context.Context, in <-chan int) ([]int, error) {
	var out []int
	for {
		select {
		case <-ctx.Done():
			return out, ctx.Err()
		case v, ok := <-in:
			if !ok {
				return out, nil
			}
			out = append(out, v)
		}
	}
}
