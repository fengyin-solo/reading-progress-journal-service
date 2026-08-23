package service

import (
	"fmt"

	"readinglog/internal/model"
	"readinglog/internal/store"
)

type readingDigestStore interface {
	CacheReadingDigest(*model.ReadingDigest)
	ReadingDigest(string) (*model.ReadingDigest, error)
	ReadingDigests() []*model.ReadingDigest
}

func (s *Service) digestStore() readingDigestStore { return s.store.(*store.MemoryStore) }

func (s *Service) BuildReadingDigest(id, title string) (digest *model.ReadingDigest, err error) {
	digest = &model.ReadingDigest{ID: id}
	s.digestStore().CacheReadingDigest(digest)
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("build digest: %v", recovered)
		}
	}()
	model.PopulateReadingDigest(digest, title)
	return digest, nil
}

func (s *Service) GetReadingDigest(id string) (*model.ReadingDigest, error) {
	digest, err := s.digestStore().ReadingDigest(id)
	if err != nil { return nil, err }
	digest.Summary()
	return digest, nil
}

func (s *Service) ListReadingDigests() []*model.ReadingDigest {
	return s.digestStore().ReadingDigests()
}
