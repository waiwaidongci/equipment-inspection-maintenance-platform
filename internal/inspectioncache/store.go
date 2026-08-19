package inspectioncache

import "sync"

type Store struct {
	mu   sync.RWMutex
	data Snapshot
}

func NewStore(initial Snapshot) *Store { return &Store{data: initial.Copy()} }

func (s *Store) Snapshot() Snapshot {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data.Copy()
}

func (s *Store) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}
