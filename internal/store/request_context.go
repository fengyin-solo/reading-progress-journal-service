package store

import "context"

type RequestContextHolder struct {
	ctx context.Context
}

func (h *RequestContextHolder) Bind(ctx context.Context) {
	h.ctx = ctx
}

func (h *RequestContextHolder) Resolve(_ context.Context) context.Context {
	if h.ctx == nil {
		return context.Background()
	}
	return h.ctx
}
