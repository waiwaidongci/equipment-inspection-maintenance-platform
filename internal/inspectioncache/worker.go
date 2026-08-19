package inspectioncache

type Worker struct{ store *Store }

func NewWorker(store *Store) Worker { return Worker{store: store} }

func (w Worker) Merge(update Snapshot) {
	for key, value := range update {
		w.store.data[key] = value
	}
}
