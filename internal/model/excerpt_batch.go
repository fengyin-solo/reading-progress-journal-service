package model

type ExcerptBatch struct {
	ID      string
	Payload []byte
}

type SliceScenarioResult struct {
	ParsedFirst    string `json:"parsed_first"`
	DelayedReceipt string `json:"delayed_receipt"`
	CachedReplay   string `json:"cached_replay"`
	Exported       string `json:"exported"`
}
