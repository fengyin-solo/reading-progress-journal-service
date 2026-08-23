package cache

import "readinglog/internal/model"

// PublishCache 是发布状态的本地缓存视图，写入遵循与 Job 相同的单调推进：
//   - 版本严格单调递增，旧回调（版本更小）不能覆盖已写入的更新版本；
//   - 版本相等时，已完成状态不得被倒退。
// 两者共同保证第一轮延迟回调不会把缓存从 completed/v2 打回 retrying/v1。
type PublishCache struct {
	state   model.PublishState
	version int
}

var stateRank = map[model.PublishState]int{
	model.PublishPending:   0,
	model.PublishRetrying:  1,
	model.PublishCompleted: 2,
}

func (c *PublishCache) Put(state model.PublishState, version int) {
	if version < c.version {
		return
	}
	if version == c.version && stateRank[state] < stateRank[c.state] {
		return
	}
	c.state = state
	c.version = version
}

func (c *PublishCache) Get() (model.PublishState, int) {
	return c.state, c.version
}
