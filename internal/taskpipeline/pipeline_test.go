package taskpipeline

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func TestPipelineLifecycleContracts(t *testing.T) {
	t.Run("producer closes on error", func(t *testing.T) {
		ch := make(chan int)
		done := make(chan error, 1)
		go func() { done <- Produce([]int{1, -1}, ch) }()
		<-ch
		select {
		case _, ok := <-ch:
			if ok {
				t.Fatal("channel stayed open")
			}
		case <-time.After(20 * time.Millisecond):
			t.Fatal("producer did not close output")
		}
		if !errors.Is(<-done, ErrInvalid) {
			t.Fatal("missing producer error")
		}
	})
	t.Run("wait registers before launch", func(t *testing.T) {
		var wg sync.WaitGroup
		gate := make(chan struct{})
		Launch(&wg, gate, func() {})
		waited := make(chan struct{})
		go func() { wg.Wait(); close(waited) }()
		select {
		case <-waited:
			t.Error("wait returned before worker registration")
		case <-time.After(10 * time.Millisecond):
		}
		close(gate)
		wg.Wait()
	})
	t.Run("consumer honors cancellation", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()
		done := make(chan error, 1)
		go func() { _, err := Consume(ctx, make(chan int)); done <- err }()
		select {
		case err := <-done:
			if !errors.Is(err, context.DeadlineExceeded) {
				t.Fatalf("err=%v", err)
			}
		case <-time.After(30 * time.Millisecond):
			t.Fatal("consumer ignored cancellation")
		}
	})
	if cap(ErrorChannel()) != 1 {
		t.Fatalf("error channel capacity=%d", cap(ErrorChannel()))
	}
}
