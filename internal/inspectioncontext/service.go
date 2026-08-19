package inspectioncontext

import "context"

type Service struct {
	gateway Gateway
	worker  Worker
}

func NewService() Service { return Service{gateway: Gateway{}, worker: Worker{}} }

func (s Service) Submit(ctx context.Context, release <-chan struct{}, upload func(context.Context) error) error {
	request := NewRequest(ctx)
	if err := s.gateway.Store(request.Context(), release); err != nil {
		return err
	}
	return s.worker.Run(request.Context(), 3, upload)
}
