package service

import (
	"fmt"

	"readinglog/internal/model"
	"readinglog/internal/store"
)

type progressSyncStore interface {
	SaveProgressSync(*model.ProgressSync)
	CacheProgressSync(*model.ProgressSync)
	ProgressSync(string) (*model.ProgressSync, error)
	ProgressSyncFeed(string) (*model.ProgressSync, error)
	RecordSyncNotice(string) int
	SyncNoticeCount(string) int
}

func (s *Service) syncStore() progressSyncStore { return s.store.(*store.MemoryStore) }

func (s *Service) StartProgressSync(bookID string) *model.ProgressSync {
	job := &model.ProgressSync{ID: fmt.Sprintf("sync-%s", bookID), BookID: bookID, Version: 1, State: model.SyncProcessing}
	s.syncStore().SaveProgressSync(job)
	s.syncStore().CacheProgressSync(job)
	s.syncStore().RecordSyncNotice(fmt.Sprintf("%s:%d", job.ID, job.Version))
	return job
}

func (s *Service) RetryProgressSync(id string) (*model.ProgressSync, error) {
	job, err := s.syncStore().ProgressSync(id)
	if err != nil {
		return nil, err
	}
	job.Version++
	job.State = model.SyncSucceeded
	s.syncStore().SaveProgressSync(job)
	s.syncStore().CacheProgressSync(job)
	s.syncStore().RecordSyncNotice(fmt.Sprintf("%s:%d", job.ID, job.Version))
	return job, nil
}

func (s *Service) CompleteProgressSync(id string, version int, state string) (*model.ProgressSync, error) {
	job, err := s.syncStore().ProgressSync(id)
	if err != nil {
		return nil, err
	}
	job.ApplyCallback(version, state)
	s.syncStore().SaveProgressSync(job)
	return job, nil
}

func (s *Service) ProgressSyncViews(id string) (detail, feed *model.ProgressSync, notices int, err error) {
	detail, err = s.syncStore().ProgressSync(id)
	if err != nil {
		return nil, nil, 0, err
	}
	feed, err = s.syncStore().ProgressSyncFeed(id)
	if err != nil {
		return nil, nil, 0, err
	}
	notices = s.syncStore().SyncNoticeCount(fmt.Sprintf("%s:%d", id, 1)) + s.syncStore().SyncNoticeCount(fmt.Sprintf("%s:%d", id, 2))
	return detail, feed, notices, nil
}
