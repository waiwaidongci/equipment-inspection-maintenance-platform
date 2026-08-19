package inspectioncache

import (
	"sync"
	"testing"
)

func TestCacheReadWriteKeepsHistoricalInspectionView(t *testing.T) {
	initial := Snapshot{"temperature": 4, "pressure": 6}
	store := NewStore(initial)
	initial["temperature"] = 999
	if got := store.Snapshot()["temperature"]; got != 4 {
		t.Errorf("constructor retained caller map: %d", got)
	}

	history := store.Snapshot()
	start := make(chan struct{})
	var workers sync.WaitGroup
	workers.Add(2)
	go func() { defer workers.Done(); <-start; NewWorker(store).Merge(Snapshot{"temperature": 8}) }()
	go func() { defer workers.Done(); <-start; _ = NewService(store).Summarize() }()
	close(start)
	workers.Wait()
	if history["temperature"] != 4 {
		t.Errorf("published history changed: %#v", history)
	}
	copyView := store.Snapshot().Copy()
	copyView["pressure"] = 100
	if store.Snapshot()["pressure"] != 6 {
		t.Error("view copy still shares store map")
	}
}
