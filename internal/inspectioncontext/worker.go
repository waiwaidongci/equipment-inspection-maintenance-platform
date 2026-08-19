package inspectioncontext

import "context"

type Worker struct{}

func (Worker) Run(ctx context.Context, attempts int, upload func(context.Context) error) error {
	var last error
	for index := 0; index < attempts; index++ {
		last = upload(ctx)
		if last == nil {
			return nil
		}
	}
	return last
}
