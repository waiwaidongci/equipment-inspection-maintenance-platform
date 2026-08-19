package i_uploadgrade

import (
	"context"
	"errors"
	core "github.com/example/inspection-platform/internal/inspectioncontext"
	"testing"
	"time"
)

func TestUploadCancellationStaysRequestScoped(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if !errors.Is(core.NewRequest(ctx).Context().Err(), context.Canceled) {
		t.Error("request lost cancel")
	}
	started := time.Now()
	if err := (core.Gateway{}).Store(ctx, make(chan struct{})); !errors.Is(err, context.Canceled) || time.Since(started) > 100*time.Millisecond {
		t.Errorf("gateway=%v", err)
	}
	calls := 0
	if err := (core.Worker{}).Run(ctx, 3, func(context.Context) error { calls++; return errors.New("down") }); !errors.Is(err, context.Canceled) || calls != 0 {
		t.Errorf("calls=%d err=%v", calls, err)
	}
	type key struct{}
	request := context.WithValue(context.Background(), key{}, "kept")
	release := make(chan struct{})
	close(release)
	seen := false
	if err := core.NewService().Submit(request, release, func(got context.Context) error { seen = got.Value(key{}) == "kept"; return nil }); err != nil || !seen {
		t.Errorf("seen=%v err=%v", seen, err)
	}
}
