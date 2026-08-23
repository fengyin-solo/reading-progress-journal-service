package service

import (
	"readinglog/internal/cache"
	"readinglog/internal/exporter"
	"readinglog/internal/model"
	"readinglog/internal/parser"
)

// delayedExcerptReceipt 返回一个延迟回执闭包。
// 捕获时即对 Payload 做独立拷贝，之后调用方对原 batch.Payload 的修改不会影响回执返回的值。
func delayedExcerptReceipt(batch model.ExcerptBatch) func() string {
	snapshot := batch.Clone()
	return func() string { return string(snapshot.Payload) }
}

func RunSliceScenario() model.SliceScenarioResult {
	p := &parser.ExcerptParser{}
	first := p.Parse("batch-alpha", "alpha")
	_ = p.Parse("batch-bravo", "bravo")

	receiptSource := model.ExcerptBatch{ID: "receipt", Payload: []byte("cedar")}
	receipt := delayedExcerptReceipt(receiptSource)
	receiptSource.Payload[0] = 'X'

	c := cache.NewExcerptCache()
	c.Put(model.ExcerptBatch{ID: "cached", Payload: []byte("delta")})
	exposed, _ := c.Get("cached")
	exposed.Payload[0] = 'X'
	replayed, _ := c.Get("cached")

	e := &exporter.ExcerptExporter{}
	exportSource := model.ExcerptBatch{ID: "export", Payload: []byte("echo")}
	e.Queue(exportSource)
	exportSource.Payload[0] = 'X'
	exported := e.Flush()

	return model.SliceScenarioResult{
		ParsedFirst:    string(first.Payload),
		DelayedReceipt: receipt(),
		CachedReplay:   string(replayed.Payload),
		Exported:       string(exported[0].Payload),
	}
}
