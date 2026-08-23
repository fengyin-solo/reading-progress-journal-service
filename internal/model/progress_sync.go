package model

const (
	SyncProcessing = "processing"
	SyncSucceeded  = "succeeded"
)

// ProgressSync tracks retries of one reader-visible progress synchronization.
type ProgressSync struct {
	ID      string `json:"id"`
	BookID  string `json:"book_id"`
	Version int    `json:"version"`
	State   string `json:"state"`
}

// ApplyCallback applies the state reported by an asynchronous attempt.
func (s *ProgressSync) ApplyCallback(version int, state string) {
	s.Version = version
	s.State = state
}
