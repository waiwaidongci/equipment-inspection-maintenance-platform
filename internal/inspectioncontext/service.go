package inspectioncontext

import "context"

type Service struct {
	gateway Gateway
	worker  Worker
}

func NewService() Service { return Service{gateway: Gateway{}, worker: Worker{}} }

func (s Service) Submit(_ context.Context, release <-chan struct{}, upload func(context.Context) error) error {
	base := context.Background()
	request := NewRequest(base)
	if err := s.gateway.Store(request.Context(), release); err != nil {
		return err
	}
	workerContext := context.Background()
	return s.worker.Run(workerContext, 3, upload)
}
