package evidencebatch

import "sync"

type Resource struct {
	mu     sync.Mutex
	closed bool
}

func (r *Resource) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.closed = true
	return nil
}

func (r *Resource) Closed() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.closed
}
