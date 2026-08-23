package store

import "readinglog/internal/model"

// CreateShelf 新增书架，名称唯一性校验。
func (s *MemoryStore) CreateShelf(sh *model.Shelf) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.shelves {
		if exist.Name == sh.Name {
			return ErrConflict
		}
	}
	s.shelves[sh.ID] = sh
	return nil
}

// GetShelf 按 ID 查询书架。
func (s *MemoryStore) GetShelf(id string) (*model.Shelf, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sh, ok := s.shelves[id]
	if !ok {
		return nil, ErrNotFound
	}
	return sh, nil
}

// ListShelves 返回全部书架。
func (s *MemoryStore) ListShelves() []*model.Shelf {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Shelf, 0, len(s.shelves))
	for _, sh := range s.shelves {
		list = append(list, sh)
	}
	return list
}

// UpdateShelf 更新书架，名称唯一性校验。
func (s *MemoryStore) UpdateShelf(sh *model.Shelf) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.shelves[sh.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.shelves {
		if exist.ID != sh.ID && exist.Name == sh.Name {
			return ErrConflict
		}
	}
	s.shelves[sh.ID] = sh
	return nil
}

// DeleteShelf 删除书架。
func (s *MemoryStore) DeleteShelf(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.shelves[id]; !ok {
		return ErrNotFound
	}
	delete(s.shelves, id)
	return nil
}
