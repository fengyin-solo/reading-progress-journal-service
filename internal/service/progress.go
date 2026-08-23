package service

import (
	"sort"
	"time"

	"readinglog/internal/model"
	"readinglog/pkg/idgen"
)

// CreateProgress 新增阅读进度，校验书籍外键存在并做页数越界校验。
func (s *Service) CreateProgress(p model.Progress) (*model.Progress, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}
	book, err := s.store.GetBook(p.BookID)
	if err != nil {
		return nil, model.NewValidationError("book_id", "书籍不存在")
	}
	if p.CurrentPage > book.TotalPages {
		return nil, model.NewValidationError("current_page", "当前页码超过书籍总页数")
	}
	now := time.Now()
	p.ID = idgen.Hex()
	p.CreatedAt = now
	if p.RecordedAt.IsZero() {
		p.RecordedAt = now
	}
	// 未显式指定百分比时自动计算。
	if p.Percentage == 0 {
		p.Percentage = float64(p.CurrentPage) / float64(book.TotalPages) * 100
	}
	if err := s.store.CreateProgress(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

// GetProgress 按 ID 查询进度。
func (s *Service) GetProgress(id string) (*model.Progress, error) {
	return s.store.GetProgress(id)
}

// ListProgresses 按筛选条件查询进度列表，支持分页。
func (s *Service) ListProgresses(filter model.ProgressFilter, page, size int) ([]*model.Progress, int, error) {
	all := s.store.ListProgresses()
	matched := make([]*model.Progress, 0, len(all))
	for _, p := range all {
		if filter.Match(p) {
			matched = append(matched, p)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].RecordedAt.After(matched[j].RecordedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Progress{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

// UpdateProgress 更新进度。
func (s *Service) UpdateProgress(id string, in model.Progress) (*model.Progress, error) {
	existing, err := s.store.GetProgress(id)
	if err != nil {
		return nil, err
	}
	in.ID = existing.ID
	in.CreatedAt = existing.CreatedAt
	if err := in.Validate(); err != nil {
		return nil, err
	}
	book, err := s.store.GetBook(in.BookID)
	if err != nil {
		return nil, model.NewValidationError("book_id", "书籍不存在")
	}
	if in.CurrentPage > book.TotalPages {
		return nil, model.NewValidationError("current_page", "当前页码超过书籍总页数")
	}
	if err := s.store.UpdateProgress(&in); err != nil {
		return nil, err
	}
	return &in, nil
}

// DeleteProgress 删除进度。
func (s *Service) DeleteProgress(id string) error {
	return s.store.DeleteProgress(id)
}
