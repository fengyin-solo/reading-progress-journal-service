package model

type ChapterInput struct {
	ID   string
	Body string
}

type ChapterImportResult struct {
	Imported    []string `json:"imported"`
	Error       string   `json:"error"`
	AuditStatus string   `json:"audit_status"`
	OpenReaders int      `json:"open_readers"`
	PeakReaders int      `json:"peak_readers"`
}
