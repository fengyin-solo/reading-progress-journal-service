package service

import (
	"sort"
	"time"

	"readinglog/internal/model"
	"readinglog/pkg/idgen"
)

// CreateBook 新增书籍。
func (s *Service) CreateBook(b model.Book) (*model.Book, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}
	now := time.Now()
	b.ID = idgen.Hex()
	b.CreatedAt = now
	b.UpdatedAt = now
	if err := s.store.CreateBook(&b); err != nil {
		return nil, err
	}
	return &b, nil
}

// GetBook 按 ID 查询书籍。
func (s *Service) GetBook(id string) (*model.Book, error) {
	return s.store.GetBook(id)
}

// ListBooks 按筛选条件查询书籍列表，支持分页。
func (s *Service) ListBooks(filter model.BookFilter, page, size int) ([]*model.Book, int, error) {
	all := s.store.ListBooks()
	matched := make([]*model.Book, 0, len(all))
	for _, b := range all {
		if filter.Match(b) {
			matched = append(matched, b)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Book{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

// UpdateBook 更新书籍。
func (s *Service) UpdateBook(id string, in model.Book) (*model.Book, error) {
	existing, err := s.store.GetBook(id)
	if err != nil {
		return nil, err
	}
	in.ID = existing.ID
	in.CreatedAt = existing.CreatedAt
	in.Status = existing.Status
	in.StartedAt = existing.StartedAt
	in.FinishedAt = existing.FinishedAt
	in.UpdatedAt = time.Now()
	if err := in.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.UpdateBook(&in); err != nil {
		return nil, err
	}
	return &in, nil
}

// DeleteBook 删除书籍。
func (s *Service) DeleteBook(id string) error {
	return s.store.DeleteBook(id)
}

// StartReading 开始阅读（wishlist→reading），写入开始时间。
func (s *Service) StartReading(id string) (*model.Book, error) {
	existing, err := s.store.GetBook(id)
	if err != nil {
		return nil, err
	}
	if !model.CanTransitionBookStatus(existing.Status, model.BookReading) {
		return nil, model.NewValidationError("status", "当前状态不允许开始阅读")
	}
	now := time.Now()
	existing.Status = model.BookReading
	existing.StartedAt = &now
	existing.UpdatedAt = now
	if err := s.store.UpdateBook(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

// FinishReading 标记读完（reading→finished），写入完成时间。
func (s *Service) FinishReading(id string) (*model.Book, error) {
	existing, err := s.store.GetBook(id)
	if err != nil {
		return nil, err
	}
	if !model.CanTransitionBookStatus(existing.Status, model.BookFinished) {
		return nil, model.NewValidationError("status", "当前状态不允许标记读完")
	}
	now := time.Now()
	existing.Status = model.BookFinished
	existing.FinishedAt = &now
	existing.UpdatedAt = now
	if err := s.store.UpdateBook(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

// DropReading 弃读（wishlist/reading→dropped）。
func (s *Service) DropReading(id string) (*model.Book, error) {
	existing, err := s.store.GetBook(id)
	if err != nil {
		return nil, err
	}
	if !model.CanTransitionBookStatus(existing.Status, model.BookDropped) {
		return nil, model.NewValidationError("status", "当前状态不允许弃读")
	}
	existing.Status = model.BookDropped
	existing.UpdatedAt = time.Now()
	if err := s.store.UpdateBook(existing); err != nil {
		return nil, err
	}
	return existing, nil
}
