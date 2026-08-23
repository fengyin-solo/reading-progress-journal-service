package parser

import "readinglog/internal/model"

// ExcerptParser 解析文本为阅读摘录批次。
type ExcerptParser struct{}

// Parse 将 text 解析为 ExcerptBatch。
// 返回的 Payload 是独立分配的副本（[]byte(text) 每次都新建底层数组），
// 因此后续对同一 parser 的调用、或对外部 text 的修改都不会影响已返回的批次。
func (p *ExcerptParser) Parse(id, text string) model.ExcerptBatch {
	return model.ExcerptBatch{ID: id, Payload: []byte(text)}
}
