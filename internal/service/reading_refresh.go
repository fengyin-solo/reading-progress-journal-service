package service

import (
	"context"
	"sync"
	"time"

	"readinglog/internal/model"
	"readinglog/internal/store"
)

type readingRefreshStore interface {
	StartReadingRefresh(string) *model.ReadingRefresh
	RecordRefreshCall(context.Context, string)
	FinishReadingRefresh(string, string)
	ReadingRefresh(string) (*model.ReadingRefresh, error)
}

var refreshWorkers sync.WaitGroup

func (s *Service) refreshStore() readingRefreshStore { return s.store.(*store.MemoryStore) }

func (s *Service) StartReadingRefresh(_ context.Context, bookID string) *model.ReadingRefresh {
	ctx := context.Background()
	id := "refresh-" + bookID
	refresh := s.refreshStore().StartReadingRefresh(id)
	s.refreshStore().RecordRefreshCall(ctx, id)
	refreshWorkers.Add(1)
	go func() {
		defer refreshWorkers.Done()
		for attempt := 0; attempt < 2; attempt++ {
			time.Sleep(100 * time.Millisecond)
			s.refreshStore().RecordRefreshCall(ctx, id)
		}
		s.refreshStore().FinishReadingRefresh(id, model.RefreshDone)
	}()
	return refresh
}

func (s *Service) GetReadingRefresh(id string) (*model.ReadingRefresh, error) {
	return s.refreshStore().ReadingRefresh(id)
}

func (s *Service) ShutdownReadingRefresh(ctx context.Context) error {
	done := make(chan struct{})
	go func() { refreshWorkers.Wait(); close(done) }()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
