package store

import "readinglog/internal/model"

// CreateBook 新增书籍。
func (s *MemoryStore) CreateBook(b *model.Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.books[b.ID] = b
	return nil
}

// GetBook 按 ID 查询书籍。
func (s *MemoryStore) GetBook(id string) (*model.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.books[id]
	if !ok {
		return nil, ErrNotFound
	}
	return b, nil
}

// ListBooks 返回全部书籍。
func (s *MemoryStore) ListBooks() []*model.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Book, 0, len(s.books))
	for _, b := range s.books {
		list = append(list, b)
	}
	return list
}

// UpdateBook 更新书籍。
func (s *MemoryStore) UpdateBook(b *model.Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.books[b.ID]; !ok {
		return ErrNotFound
	}
	s.books[b.ID] = b
	return nil
}

// DeleteBook 删除书籍。
func (s *MemoryStore) DeleteBook(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.books[id]; !ok {
		return ErrNotFound
	}
	delete(s.books, id)
	return nil
}
