package cache

import "readinglog/internal/model"

type ExcerptCache struct {
	items map[string]model.ExcerptBatch
}

func NewExcerptCache() *ExcerptCache {
	return &ExcerptCache{items: make(map[string]model.ExcerptBatch)}
}

func (c *ExcerptCache) Put(batch model.ExcerptBatch) {
	c.items[batch.ID] = batch
}

func (c *ExcerptCache) Get(id string) (model.ExcerptBatch, bool) {
	batch, ok := c.items[id]
	return batch, ok
}
