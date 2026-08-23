package worker

import (
	"context"

	"readinglog/internal/metaclient"
)

type ContextReader struct {
	client *metaclient.ReaderClient
}

func NewContextReader(client *metaclient.ReaderClient) *ContextReader {
	return &ContextReader{client: client}
}

func (r *ContextReader) Run(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	var lastErr error
	for attempt := 0; attempt < 2; attempt++ {
		lastErr = r.client.Fetch(context.Background())
	}
	return lastErr
}
