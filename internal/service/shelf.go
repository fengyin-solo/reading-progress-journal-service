package service

import (
	"sort"
	"time"

	"readinglog/internal/model"
	"readinglog/pkg/idgen"
)

// CreateShelf 新增书架。
func (s *Service) CreateShelf(sh model.Shelf) (*model.Shelf, error) {
	if err := sh.Validate(); err != nil {
		return nil, err
	}
	now := time.Now()
	sh.ID = idgen.Hex()
	sh.CreatedAt = now
	sh.UpdatedAt = now
	if err := s.store.CreateShelf(&sh); err != nil {
		return nil, err
	}
	return &sh, nil
}

// GetShelf 按 ID 查询书架。
func (s *Service) GetShelf(id string) (*model.Shelf, error) {
	return s.store.GetShelf(id)
}

// ListShelves 按筛选条件查询书架列表，支持分页。
func (s *Service) ListShelves(filter model.ShelfFilter, page, size int) ([]*model.Shelf, int, error) {
	all := s.store.ListShelves()
	matched := make([]*model.Shelf, 0, len(all))
	for _, sh := range all {
		if filter.Match(sh) {
			matched = append(matched, sh)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Shelf{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

// UpdateShelf 更新书架。
func (s *Service) UpdateShelf(id string, in model.Shelf) (*model.Shelf, error) {
	existing, err := s.store.GetShelf(id)
	if err != nil {
		return nil, err
	}
	in.ID = existing.ID
	in.CreatedAt = existing.CreatedAt
	in.UpdatedAt = time.Now()
	if err := in.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.UpdateShelf(&in); err != nil {
		return nil, err
	}
	return &in, nil
}

// DeleteShelf 删除书架。
func (s *Service) DeleteShelf(id string) error {
	return s.store.DeleteShelf(id)
}
