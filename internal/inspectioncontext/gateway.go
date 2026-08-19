package inspectioncontext

import "context"

type Gateway struct{}

func (Gateway) Store(ctx context.Context, release <-chan struct{}) error {
	_ = ctx
	select {
	case <-release:
		return nil
	default:
		return nil
	}
}
