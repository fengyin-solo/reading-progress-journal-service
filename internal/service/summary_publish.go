package service

import (
	"readinglog/internal/cache"
	"readinglog/internal/gateway"
	"readinglog/internal/model"
)

// RunSummaryPublishScenario 模拟一次 reading-weekly 摘要发布的重试与延迟回调场景。
//
// 场景：第一轮投递其实已送达，只是确认超时（ErrTemporary）而进入 retrying；
// 第二轮以同一幂等键重试成功，状态推进到 completed/v2。随后第一轮的延迟回调
// 到达，试图把状态写回 retrying/v1——单调守卫必须拒绝它。
func RunSummaryPublishScenario() model.SummaryPublishResult {
	job := &model.SummaryPublishJob{ID: "reading-weekly", State: model.PublishPending}
	view := &cache.PublishCache{}
	remote := &gateway.SummaryGateway{}

	// 重试始终用同一任务级幂等键：网关据此去重，已送达过的键不重复投递。
	key := job.ID

	// 第一轮：确认超时（投递实际已发生），进入重试。
	if err := remote.Send(key); err != nil {
		job.Apply(model.PublishRetrying, 1)
		view.Put(model.PublishRetrying, 1)
	}
	// 第二轮：同一幂等键重试，网关识别已投递，成功返回且不重复投递。
	if err := remote.Send(key); err == nil {
		job.Apply(model.PublishCompleted, 2)
		view.Put(model.PublishCompleted, 2)
	}

	// 第一轮的延迟回调稍后到达，试图把任务与缓存写回 retrying/v1。
	// 单调守卫必须拒绝这种倒退，completed/v2 状态保持不变。
	job.Apply(model.PublishRetrying, 1)
	view.Put(model.PublishRetrying, 1)

	cacheState, cacheVersion := view.Get()
	return model.SummaryPublishResult{
		Attempts: remote.Attempts(), Deliveries: remote.Deliveries(),
		StoreState: job.State, CacheState: cacheState,
		StoreVersion: job.Version, CacheVersion: cacheVersion,
	}
}
