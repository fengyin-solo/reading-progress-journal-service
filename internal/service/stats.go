package service

import (
	"time"

	"readinglog/internal/model"
)

// StatsOverview 阅读记录概览统计。
type StatsOverview struct {
	TotalBooks           int            `json:"total_books"`
	BooksByCategory      map[string]int `json:"books_by_category"`
	BooksByStatus        map[string]int `json:"books_by_status"`
	TotalShelves         int            `json:"total_shelves"`
	TotalNotes           int            `json:"total_notes"`
	TotalHighlights      int            `json:"total_highlights"`
	TotalProgressRecords int            `json:"total_progress_records"`
	TotalReadingPages    int            `json:"total_reading_pages"`
}

// Overview 计算整体概览统计。
func (s *Service) Overview() *StatsOverview {
	o := &StatsOverview{
		BooksByCategory: make(map[string]int),
		BooksByStatus:   make(map[string]int),
	}

	bookMap := make(map[string]*model.Book)
	for _, b := range s.store.ListBooks() {
		o.TotalBooks++
		o.BooksByCategory[b.Category]++
		o.BooksByStatus[b.Status]++
		bookMap[b.ID] = b
	}

	o.TotalShelves = len(s.store.ListShelves())
	o.TotalNotes = len(s.store.ListNotes())
	o.TotalHighlights = len(s.store.ListHighlights())

	// 每本书取最新进度页数求和，作为总阅读页数。
	latestPages := make(map[string]int)
	latestTime := make(map[string]time.Time)
	for _, p := range s.store.ListProgresses() {
		o.TotalProgressRecords++
		if _, ok := bookMap[p.BookID]; !ok {
			continue
		}
		if t, ok := latestTime[p.BookID]; !ok || p.RecordedAt.After(t) {
			latestTime[p.BookID] = p.RecordedAt
			latestPages[p.BookID] = p.CurrentPage
		}
	}
	for _, page := range latestPages {
		o.TotalReadingPages += page
	}

	return o
}
