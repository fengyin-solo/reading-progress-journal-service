package service

import (
	"errors"

	"readinglog/internal/audit"
	"readinglog/internal/model"
	"readinglog/internal/publisher"
	"readinglog/internal/store"
)

func RunSaveRetryScenario(entryID string) model.SaveScenarioResult {
	pub := publisher.NewSavePublisher()
	log := audit.NewSaveAudit()
	result := model.SaveScenarioResult{}

	for attempt := 1; attempt <= 2; attempt++ {
		tx := store.BeginSave(attempt)
		success := model.SaveEvent{EntryID: entryID, Attempt: attempt, State: "saved"}
		pub.Publish(success)
		log.Record(success)

		err := tx.Finish()
		if err == nil {
			result.EntrySaved = tx.Committed()
			break
		}
		if attempt == 1 {
			result.FirstError = err.Error()
		}
		log.Record(model.SaveEvent{EntryID: entryID, Attempt: attempt, State: "failed"})
		if !errors.Is(err, model.ErrRollbackCleanup) {
			break
		}
	}

	result.Notifications = pub.Events()
	result.AuditTrail = log.Events()
	return result
}
