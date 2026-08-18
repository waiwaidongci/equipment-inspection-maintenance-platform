package taskpipeline

import (
	"context"
)

func Consume(ctx context.Context, in <-chan int) ([]int, error) {
	var out []int
	for v := range in {
		out = append(out, v)
	}
	return out, nil
}
