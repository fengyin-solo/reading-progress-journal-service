package cache

import "readinglog/internal/model"

type PublishCache struct {
	state   model.PublishState
	version int
}

func (c *PublishCache) Put(state model.PublishState, version int) {
	c.state = state
	c.version = version
}

func (c *PublishCache) Get() (model.PublishState, int) {
	return c.state, c.version
}
