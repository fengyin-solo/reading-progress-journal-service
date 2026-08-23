package model

import "errors"

var (
	ErrJournalWrite    = errors.New("journal storage rejected entry")
	ErrRollbackCleanup = errors.New("rollback cleanup failed")
)

type SaveEvent struct {
	EntryID string `json:"entry_id"`
	Attempt int    `json:"attempt"`
	State   string `json:"state"`
}

type SaveScenarioResult struct {
	EntrySaved    bool        `json:"entry_saved"`
	FirstError    string      `json:"first_error"`
	Notifications []SaveEvent `json:"notifications"`
	AuditTrail    []SaveEvent `json:"audit_trail"`
}
