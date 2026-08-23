package store

import (
	"sync"

	"readinglog/internal/model"
)

type ReadIndex struct {
	mu    sync.RWMutex
	items map[string]*model.ReadCount
}

func NewReadIndex() *ReadIndex {
	return &ReadIndex{items: map[string]*model.ReadCount{"book-A": {BookID: "book-A", Value: 1}}}
}

func (s *ReadIndex) Snapshot() map[string]*model.ReadCount {
	return s.items
}

func (s *ReadIndex) Increment(bookID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[bookID].Value++
}
