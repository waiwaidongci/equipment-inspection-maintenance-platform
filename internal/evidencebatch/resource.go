package evidencebatch

import "sync"

type Resource struct {
	mu     sync.Mutex
	closed bool
}

func (r *Resource) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.closed {
		return nil
	}
	return nil
}

func (r *Resource) Closed() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.closed
}
