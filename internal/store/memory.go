package store

import (
	"sync"

	"readinglog/internal/model"
)

// MemoryStore 基于内存的 Store 实现，线程安全。
type MemoryStore struct {
	mu         sync.RWMutex
	books      map[string]*model.Book
	shelves    map[string]*model.Shelf
	notes      map[string]*model.Note
	highlights map[string]*model.Highlight
	progresses map[string]*model.Progress
	syncJobs   map[string]*model.ProgressSync
	syncFeed   map[string]*model.ProgressSync
	syncNotices map[string]int
}

// NewMemoryStore 创建空的内存 Store。
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		books:      make(map[string]*model.Book),
		shelves:    make(map[string]*model.Shelf),
		notes:      make(map[string]*model.Note),
		highlights: make(map[string]*model.Highlight),
		progresses: make(map[string]*model.Progress),
		syncJobs:   make(map[string]*model.ProgressSync),
		syncFeed:   make(map[string]*model.ProgressSync),
		syncNotices: make(map[string]int),
	}
}

var _ Store = (*MemoryStore)(nil)
