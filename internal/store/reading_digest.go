package store

import "readinglog/internal/model"

func (s *MemoryStore) CacheReadingDigest(digest *model.ReadingDigest) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.digests[digest.ID] = digest
}

func (s *MemoryStore) ReadingDigest(id string) (*model.ReadingDigest, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	digest, ok := s.digests[id]
	if !ok {
		return nil, ErrNotFound
	}
	return digest, nil
}

func (s *MemoryStore) ReadingDigests() []*model.ReadingDigest {
	s.mu.RLock()
	defer s.mu.RUnlock()
	items := make([]*model.ReadingDigest, 0, len(s.digests))
	for _, digest := range s.digests {
		items = append(items, digest)
	}
	return items
}
