package inspectioncache

type Worker struct{ store *Store }

func NewWorker(store *Store) Worker { return Worker{store: store} }

func (w Worker) Merge(update Snapshot) {
	w.store.mu.Lock()
	defer w.store.mu.Unlock()
	for key, value := range update {
		w.store.data[key] = value
	}
}
