package store

import "readinglog/internal/model"

// CreateProgress 新增阅读进度。
func (s *MemoryStore) CreateProgress(p *model.Progress) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.progresses[p.ID] = p
	return nil
}

// GetProgress 按 ID 查询进度。
func (s *MemoryStore) GetProgress(id string) (*model.Progress, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	p, ok := s.progresses[id]
	if !ok {
		return nil, ErrNotFound
	}
	return p, nil
}

// ListProgresses 返回全部进度记录。
func (s *MemoryStore) ListProgresses() []*model.Progress {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Progress, 0, len(s.progresses))
	for _, p := range s.progresses {
		list = append(list, p)
	}
	return list
}

// UpdateProgress 更新进度。
func (s *MemoryStore) UpdateProgress(p *model.Progress) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.progresses[p.ID]; !ok {
		return ErrNotFound
	}
	s.progresses[p.ID] = p
	return nil
}

// DeleteProgress 删除进度。
func (s *MemoryStore) DeleteProgress(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.progresses[id]; !ok {
		return ErrNotFound
	}
	delete(s.progresses, id)
	return nil
}
