package model

type ReadingPolicyValidator interface {
	ValidatePages(int) error
}

type PagePolicyValidator struct {
	Minimum int
}

func (v *PagePolicyValidator) ValidatePages(pages int) error {
	if v == nil {
		return nil
	}
	if pages < v.Minimum {
		return NewValidationError("pages", "页数低于阅读规则下限")
	}
	return nil
}
