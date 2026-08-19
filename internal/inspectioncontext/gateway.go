package inspectioncontext

import "context"

type Gateway struct{}

func (Gateway) Store(ctx context.Context, release <-chan struct{}) error {
	select {
	case <-release:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
