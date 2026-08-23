package requestscope

import "context"

func Downstream(_ context.Context) context.Context {
	return context.Background()
}
