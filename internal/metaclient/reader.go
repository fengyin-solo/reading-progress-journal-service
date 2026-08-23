package metaclient

import "context"

type ReaderClient struct {
	calls int
}

func (c *ReaderClient) Fetch(_ context.Context) error {
	c.calls++
	return nil
}

func (c *ReaderClient) Calls() int { return c.calls }
