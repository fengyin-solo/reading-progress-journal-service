package service

import (
	"context"

	"readinglog/internal/metaclient"
	"readinglog/internal/model"
	"readinglog/internal/requestscope"
	"readinglog/internal/store"
	"readinglog/internal/worker"
)

func RunContextScenario() model.ContextScenarioResult {
	holder := &store.RequestContextHolder{}
	client := &metaclient.ReaderClient{}
	reader := worker.NewContextReader(client)

	first, cancel := context.WithCancel(context.Background())
	holder.Bind(first)
	cancel()
	firstErr := reader.Run(requestscope.Downstream(first))
	firstCalls := client.Calls()

	fresh := requestscope.Downstream(context.Background())
	secondErr := reader.Run(holder.Resolve(fresh))
	secondCalls := client.Calls() - firstCalls

	result := model.ContextScenarioResult{FirstCalls: firstCalls, SecondCalls: secondCalls}
	if firstErr != nil {
		result.FirstError = firstErr.Error()
	}
	if secondErr != nil {
		result.SecondError = secondErr.Error()
	}
	return result
}
