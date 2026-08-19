package inspectioncache

type Store struct {
	data Snapshot
}

func NewStore(initial Snapshot) *Store { return &Store{data: initial.Copy()} }

func (s *Store) Snapshot() Snapshot {
	return s.data
}

func (s *Store) Set(key string, value int) {
	s.data[key] = value
}
