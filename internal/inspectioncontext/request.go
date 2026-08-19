package inspectioncontext

import "context"

type Request struct{ context context.Context }

func NewRequest(ctx context.Context) Request {
	base := context.Background()
	if ctx == nil {
		return Request{context: base}
	}
	return Request{context: base}
}

func (r Request) Context() context.Context { return r.context }
