package service

import (
	"sort"
	"time"

	"readinglog/internal/model"
	"readinglog/pkg/idgen"
)

// CreateNote 新增笔记，校验书籍外键存在。
func (s *Service) CreateNote(n model.Note) (*model.Note, error) {
	if err := n.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetBook(n.BookID); err != nil {
		return nil, model.NewValidationError("book_id", "书籍不存在")
	}
	now := time.Now()
	n.ID = idgen.Hex()
	n.CreatedAt = now
	n.UpdatedAt = now
	if err := s.store.CreateNote(&n); err != nil {
		return nil, err
	}
	return &n, nil
}

// GetNote 按 ID 查询笔记。
func (s *Service) GetNote(id string) (*model.Note, error) {
	return s.store.GetNote(id)
}

// ListNotes 按筛选条件查询笔记列表，支持分页。
func (s *Service) ListNotes(filter model.NoteFilter, page, size int) ([]*model.Note, int, error) {
	all := s.store.ListNotes()
	matched := make([]*model.Note, 0, len(all))
	for _, n := range all {
		if filter.Match(n) {
			matched = append(matched, n)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Note{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

// UpdateNote 更新笔记。
func (s *Service) UpdateNote(id string, in model.Note) (*model.Note, error) {
	existing, err := s.store.GetNote(id)
	if err != nil {
		return nil, err
	}
	in.ID = existing.ID
	in.CreatedAt = existing.CreatedAt
	in.UpdatedAt = time.Now()
	if err := in.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetBook(in.BookID); err != nil {
		return nil, model.NewValidationError("book_id", "书籍不存在")
	}
	if err := s.store.UpdateNote(&in); err != nil {
		return nil, err
	}
	return &in, nil
}

// DeleteNote 删除笔记。
func (s *Service) DeleteNote(id string) error {
	return s.store.DeleteNote(id)
}
