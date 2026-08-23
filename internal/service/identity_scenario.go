package service

import "readinglog/internal/model"

func delayedReceipt(identity *model.RequestIdentity) func() string {
	return func() string { return identity.ReaderID }
}

func delayedAudit(identity *model.RequestIdentity) func() string {
	return func() string { return identity.ReaderID }
}

func (s *Service) RunIdentityScenario() model.IdentityScenario {
	first := s.identities.Acquire("reader-A", []string{"book-A"})
	receipt := delayedReceipt(first)
	audit := delayedAudit(first)
	s.identities.Release(first)

	second := s.identities.Acquire("reader-B", []string{"book-B"})
	result := model.IdentityScenario{
		FirstReceipt: receipt(),
		AuditReader:  audit(),
		SecondReader: second.ReaderID,
	}
	s.identities.Release(second)
	return result
}
