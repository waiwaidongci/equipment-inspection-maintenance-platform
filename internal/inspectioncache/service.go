package inspectioncache

type Service struct{ store *Store }

func NewService(store *Store) Service { return Service{store: store} }

func (s Service) Summarize() int {
	snapshot := s.store.Snapshot()
	total := 0
	for _, value := range snapshot {
		total += value
	}
	return total
}
