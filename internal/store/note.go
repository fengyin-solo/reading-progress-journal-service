package store

import "readinglog/internal/model"

// CreateNote 新增笔记。
func (s *MemoryStore) CreateNote(n *model.Note) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.notes[n.ID] = n
	return nil
}

// GetNote 按 ID 查询笔记。
func (s *MemoryStore) GetNote(id string) (*model.Note, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	n, ok := s.notes[id]
	if !ok {
		return nil, ErrNotFound
	}
	return n, nil
}

// ListNotes 返回全部笔记。
func (s *MemoryStore) ListNotes() []*model.Note {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Note, 0, len(s.notes))
	for _, n := range s.notes {
		list = append(list, n)
	}
	return list
}

// UpdateNote 更新笔记。
func (s *MemoryStore) UpdateNote(n *model.Note) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.notes[n.ID]; !ok {
		return ErrNotFound
	}
	s.notes[n.ID] = n
	return nil
}

// DeleteNote 删除笔记。
func (s *MemoryStore) DeleteNote(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.notes[id]; !ok {
		return ErrNotFound
	}
	delete(s.notes, id)
	return nil
}
