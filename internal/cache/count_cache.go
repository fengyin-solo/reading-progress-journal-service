package cache

import "readinglog/internal/model"

type CountCache struct {
	snapshot map[string]*model.ReadCount
}

func (c *CountCache) Put(snapshot map[string]*model.ReadCount) {
	c.snapshot = snapshot
}

func (c *CountCache) Get() map[string]*model.ReadCount {
	return c.snapshot
}

func (c *CountCache) Total() int {
	total := 0
	for _, count := range c.snapshot {
		total += count.Value
	}
	return total
}
