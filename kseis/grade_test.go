package f_cachegrade

import (
	core "github.com/example/inspection-platform/internal/inspectioncache"
	"sync"
	"testing"
)

func TestConcurrentCacheSnapshotSurvivesWrites(t *testing.T) {
	initial := core.Snapshot{"temperature": 4, "pressure": 6}
	store := core.NewStore(initial)
	initial["temperature"] = 999
	if store.Snapshot()["temperature"] != 4 {
		t.Error("constructor alias")
	}
	history := store.Snapshot()
	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); <-start; core.NewWorker(store).Merge(core.Snapshot{"temperature": 8}) }()
	go func() { defer wg.Done(); <-start; _ = core.NewService(store).Summarize() }()
	close(start)
	wg.Wait()
	if history["temperature"] != 4 {
		t.Error("history changed")
	}
	view := store.Snapshot().Copy()
	view["pressure"] = 100
	if store.Snapshot()["pressure"] != 6 {
		t.Error("copy aliases store")
	}
}
