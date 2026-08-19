package inspectioncontext

import (
	"context"
	"errors"
	"testing"
	"time"
)

var errUpload = errors.New("upload unavailable")

func TestUploadContextCancellationIsRequestScoped(t *testing.T) {
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if !errors.Is(NewRequest(cancelled).Context().Err(), context.Canceled) {
		t.Error("request discarded cancellation")
	}
	if NewRequest(context.Background()).Context().Err() != nil {
		t.Error("cancelled request polluted next request")
	}
	started := time.Now()
	if err := (Gateway{}).Store(cancelled, make(chan struct{})); !errors.Is(err, context.Canceled) || time.Since(started) > 100*time.Millisecond {
		t.Errorf("gateway did not stop: %v", err)
	}
	calls := 0
	if err := (Worker{}).Run(cancelled, 3, func(context.Context) error { calls++; return errUpload }); !errors.Is(err, context.Canceled) || calls != 0 {
		t.Errorf("worker retried cancelled upload: calls=%d err=%v", calls, err)
	}
	type requestKey struct{}
	requestContext := context.WithValue(context.Background(), requestKey{}, "inspection-17")
	release := make(chan struct{})
	close(release)
	seen := false
	err := NewService().Submit(requestContext, release, func(ctx context.Context) error {
		seen = ctx.Value(requestKey{}) == "inspection-17"
		return nil
	})
	if err != nil || !seen {
		t.Errorf("service changed context: seen=%v err=%v", seen, err)
	}
}
