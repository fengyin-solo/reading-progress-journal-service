package store

import "readinglog/internal/model"

// CreateHighlight 新增书摘。
func (s *MemoryStore) CreateHighlight(h *model.Highlight) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.highlights[h.ID] = h
	return nil
}

// GetHighlight 按 ID 查询书摘。
func (s *MemoryStore) GetHighlight(id string) (*model.Highlight, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	h, ok := s.highlights[id]
	if !ok {
		return nil, ErrNotFound
	}
	return h, nil
}

// ListHighlights 返回全部书摘。
func (s *MemoryStore) ListHighlights() []*model.Highlight {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Highlight, 0, len(s.highlights))
	for _, h := range s.highlights {
		list = append(list, h)
	}
	return list
}

// UpdateHighlight 更新书摘。
func (s *MemoryStore) UpdateHighlight(h *model.Highlight) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.highlights[h.ID]; !ok {
		return ErrNotFound
	}
	s.highlights[h.ID] = h
	return nil
}

// DeleteHighlight 删除书摘。
func (s *MemoryStore) DeleteHighlight(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.highlights[id]; !ok {
		return ErrNotFound
	}
	delete(s.highlights, id)
	return nil
}
