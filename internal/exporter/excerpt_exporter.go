package exporter

import "readinglog/internal/model"

type ExcerptExporter struct {
	queue []model.ExcerptBatch
}

func (e *ExcerptExporter) Queue(batch model.ExcerptBatch) {
	e.queue = append(e.queue, batch)
}

func (e *ExcerptExporter) Flush() []model.ExcerptBatch {
	result := e.queue
	e.queue = nil
	return result
}
