package service

func (s *Service) CheckReadingPolicy(pages int) error {
	if s.policy.Validator != nil {
		return s.policy.Validator.ValidatePages(pages)
	}
	return nil
}

func (s *Service) SetReadingRule(name string, pages int) {
	s.policy.Rules[name] = pages
}

func (s *Service) ReadingRule(name string) int {
	return s.policy.Rules[name]
}
