package service

import (
	"runtime"

	"readinglog/internal/cache"
	"readinglog/internal/model"
	"readinglog/internal/store"
	"readinglog/internal/worker"
)

func RunCountSnapshotScenario() model.CountScenarioResult {
	index := store.NewReadIndex()
	cached := &cache.CountCache{}
	cached.Put(index.Snapshot())
	view := cached.Get()
	initial := view["book-A"].Value

	start := make(chan struct{})
	ready := make(chan struct{})
	done := make(chan struct{})
	go worker.NewCountWriter(index).ApplyMany(start, ready, done)
	<-ready
	close(start)
	for i := 0; i < 100000; i++ {
		_ = view["book-A"].Value
		runtime.Gosched()
	}
	<-done

	return model.CountScenarioResult{
		InitialSnapshot: initial,
		CachedAfterWork: view["book-A"].Value,
		CachedTotal:     cached.Total(),
	}
}
