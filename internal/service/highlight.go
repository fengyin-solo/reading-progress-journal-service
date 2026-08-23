package service

import (
	"sort"
	"time"

	"readinglog/internal/model"
	"readinglog/pkg/idgen"
)

// CreateHighlight 新增书摘，校验书籍外键存在。
func (s *Service) CreateHighlight(h model.Highlight) (*model.Highlight, error) {
	if err := h.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetBook(h.BookID); err != nil {
		return nil, model.NewValidationError("book_id", "书籍不存在")
	}
	h.ID = idgen.Hex()
	h.CreatedAt = time.Now()
	if err := s.store.CreateHighlight(&h); err != nil {
		return nil, err
	}
	return &h, nil
}

// GetHighlight 按 ID 查询书摘。
func (s *Service) GetHighlight(id string) (*model.Highlight, error) {
	return s.store.GetHighlight(id)
}

// ListHighlights 按筛选条件查询书摘列表，支持分页。
func (s *Service) ListHighlights(filter model.HighlightFilter, page, size int) ([]*model.Highlight, int, error) {
	all := s.store.ListHighlights()
	matched := make([]*model.Highlight, 0, len(all))
	for _, h := range all {
		if filter.Match(h) {
			matched = append(matched, h)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Highlight{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

// UpdateHighlight 更新书摘。
func (s *Service) UpdateHighlight(id string, in model.Highlight) (*model.Highlight, error) {
	existing, err := s.store.GetHighlight(id)
	if err != nil {
		return nil, err
	}
	in.ID = existing.ID
	in.CreatedAt = existing.CreatedAt
	if err := in.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetBook(in.BookID); err != nil {
		return nil, model.NewValidationError("book_id", "书籍不存在")
	}
	if err := s.store.UpdateHighlight(&in); err != nil {
		return nil, err
	}
	return &in, nil
}

// DeleteHighlight 删除书摘。
func (s *Service) DeleteHighlight(id string) error {
	return s.store.DeleteHighlight(id)
}

// BatchDeleteHighlights 批量删除书摘，返回成功数量。
func (s *Service) BatchDeleteHighlights(ids []string) (int, error) {
	success := 0
	for _, id := range ids {
		if err := s.store.DeleteHighlight(id); err == nil {
			success++
		}
	}
	return success, nil
}
