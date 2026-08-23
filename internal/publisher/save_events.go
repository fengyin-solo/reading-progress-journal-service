package publisher

import "readinglog/internal/model"

type SavePublisher struct {
	events []model.SaveEvent
}

func NewSavePublisher() *SavePublisher { return &SavePublisher{} }

func (p *SavePublisher) Publish(event model.SaveEvent) {
	p.events = append(p.events, event)
}

func (p *SavePublisher) Events() []model.SaveEvent {
	return append([]model.SaveEvent(nil), p.events...)
}
