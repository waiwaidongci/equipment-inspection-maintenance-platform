package inspectioncontext

import "context"

type Request struct{ context context.Context }

func NewRequest(ctx context.Context) Request {
	if ctx == nil {
		ctx = context.Background()
	}
	return Request{context: ctx}
}

func (r Request) Context() context.Context { return r.context }
