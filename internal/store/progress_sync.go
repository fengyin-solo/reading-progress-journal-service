package store

import "readinglog/internal/model"

func cloneProgressSync(job *model.ProgressSync) *model.ProgressSync {
	copy := *job
	return &copy
}

func (s *MemoryStore) SaveProgressSync(job *model.ProgressSync) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncJobs[job.ID] = cloneProgressSync(job)
}

func (s *MemoryStore) CacheProgressSync(job *model.ProgressSync) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncFeed[job.ID] = cloneProgressSync(job)
}

func (s *MemoryStore) ProgressSync(id string) (*model.ProgressSync, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	job, ok := s.syncJobs[id]
	if !ok {
		return nil, ErrNotFound
	}
	return cloneProgressSync(job), nil
}

func (s *MemoryStore) ProgressSyncFeed(id string) (*model.ProgressSync, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	job, ok := s.syncFeed[id]
	if !ok {
		return nil, ErrNotFound
	}
	return cloneProgressSync(job), nil
}

func (s *MemoryStore) RecordSyncNotice(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncNotices[key]++
	return s.syncNotices[key]
}

func (s *MemoryStore) SyncNoticeCount(key string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.syncNotices[key]
}
