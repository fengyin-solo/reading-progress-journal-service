package service

import (
	"readinglog/internal/cache"
	"readinglog/internal/gateway"
	"readinglog/internal/model"
)

func RunSummaryPublishScenario() model.SummaryPublishResult {
	job := &model.SummaryPublishJob{ID: "reading-weekly", State: model.PublishPending}
	view := &cache.PublishCache{}
	remote := &gateway.SummaryGateway{}

	if err := remote.Send(job.ID + "-attempt-1"); err != nil {
		job.Apply(model.PublishRetrying, 1)
		view.Put(model.PublishRetrying, 1)
	}
	if err := remote.Send(job.ID + "-attempt-2"); err == nil {
		job.Apply(model.PublishCompleted, 2)
		view.Put(model.PublishCompleted, 2)
	}

	job.Apply(model.PublishRetrying, 1)
	view.Put(model.PublishRetrying, 1)
	cacheState, cacheVersion := view.Get()
	return model.SummaryPublishResult{
		Attempts: remote.Attempts(), Deliveries: remote.Deliveries(),
		StoreState: job.State, CacheState: cacheState,
		StoreVersion: job.Version, CacheVersion: cacheVersion,
	}
}
