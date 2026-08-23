package model

// ExcerptBatch 表示一批阅读摘录的导入载荷。
//
// Payload 是可变的 []byte，存在切片别名风险：若两个 ExcerptBatch 共享同一底层数组，
// 修改一方会波及另一方。因此跨组件边界（解析、缓存、导出、延迟回执）传递时，
// 必须通过 Clone 生成独立副本，以保证“已交出去的批次在后续复用和修改后仍保持原值”。
type ExcerptBatch struct {
	ID      string
	Payload []byte
}

// Clone 返回与原批次互不影响的副本。
// ID 为不可变字符串，随结构按值复制即可；Payload 则分配新的底层数组并拷贝内容，
// 使得副本与原批次在修改 Payload 时互不波及。
func (b ExcerptBatch) Clone() ExcerptBatch {
	payload := make([]byte, len(b.Payload))
	copy(payload, b.Payload)
	return ExcerptBatch{ID: b.ID, Payload: payload}
}

type SliceScenarioResult struct {
	ParsedFirst    string `json:"parsed_first"`
	DelayedReceipt string `json:"delayed_receipt"`
	CachedReplay   string `json:"cached_replay"`
	Exported       string `json:"exported"`
}
