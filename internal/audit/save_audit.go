package audit

import "readinglog/internal/model"

type SaveAudit struct {
	events []model.SaveEvent
}

func NewSaveAudit() *SaveAudit { return &SaveAudit{} }

func (a *SaveAudit) Record(event model.SaveEvent) {
	a.events = append(a.events, event)
}

func (a *SaveAudit) Events() []model.SaveEvent {
	return append([]model.SaveEvent(nil), a.events...)
}
