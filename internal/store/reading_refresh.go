package store

import (
	"context"

	"readinglog/internal/model"
)

func cloneReadingRefresh(refresh *model.ReadingRefresh) *model.ReadingRefresh {
	copy := *refresh
	return &copy
}

func (s *MemoryStore) StartReadingRefresh(id string) *model.ReadingRefresh {
	s.mu.Lock()
	defer s.mu.Unlock()
	refresh := &model.ReadingRefresh{ID: id, State: model.RefreshRunning}
	s.refreshes[id] = refresh
	return cloneReadingRefresh(refresh)
}

func (s *MemoryStore) RecordRefreshCall(_ context.Context, id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if refresh := s.refreshes[id]; refresh != nil {
		refresh.MarkCall()
	}
}

func (s *MemoryStore) FinishReadingRefresh(id, state string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if refresh := s.refreshes[id]; refresh != nil {
		refresh.Finish(state)
	}
}

func (s *MemoryStore) ReadingRefresh(id string) (*model.ReadingRefresh, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	refresh, ok := s.refreshes[id]
	if !ok {
		return nil, ErrNotFound
	}
	return cloneReadingRefresh(refresh), nil
}
