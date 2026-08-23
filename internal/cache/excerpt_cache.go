package cache

import "readinglog/internal/model"

type ExcerptCache struct {
	items map[string]model.ExcerptBatch
}

func NewExcerptCache() *ExcerptCache {
	return &ExcerptCache{items: make(map[string]model.ExcerptBatch)}
}

// Put 写入批次。对 Payload 做独立拷贝，调用方事后修改传入的 Payload 不会污染缓存。
func (c *ExcerptCache) Put(batch model.ExcerptBatch) {
	c.items[batch.ID] = batch.Clone()
}

// Get 读取批次。返回的 Payload 是独立拷贝，调用方修改返回值不会污染缓存，也不会影响其它 Get 的结果。
func (c *ExcerptCache) Get(id string) (model.ExcerptBatch, bool) {
	batch, ok := c.items[id]
	if !ok {
		return model.ExcerptBatch{}, false
	}
	return batch.Clone(), true
}
