package exporter

import "readinglog/internal/model"

type ExcerptExporter struct {
	queue []model.ExcerptBatch
}

// Queue 将批次加入导出队列。对 Payload 做独立拷贝，调用方事后修改传入的 Payload 不会污染队列。
func (e *ExcerptExporter) Queue(batch model.ExcerptBatch) {
	e.queue = append(e.queue, batch.Clone())
}

// Flush 导出并清空队列。返回的每个批次都是独立拷贝：调用方修改返回值不会波及队列，
// 队列清空后可继续复用该 exporter 而不影响已导出的批次。
func (e *ExcerptExporter) Flush() []model.ExcerptBatch {
	result := make([]model.ExcerptBatch, len(e.queue))
	for i, b := range e.queue {
		result[i] = b.Clone()
	}
	e.queue = nil
	return result
}
