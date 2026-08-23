package service

import (
	"readinglog/internal/cache"
	"readinglog/internal/exporter"
	"readinglog/internal/model"
	"readinglog/internal/parser"
)

func delayedExcerptReceipt(batch model.ExcerptBatch) func() string {
	return func() string { return string(batch.Payload) }
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
