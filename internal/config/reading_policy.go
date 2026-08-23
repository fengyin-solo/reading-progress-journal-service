package config

import "readinglog/internal/model"

type ReadingPolicy struct {
	Rules     map[string]int
	Validator model.ReadingPolicyValidator
}

func LoadReadingPolicy() *ReadingPolicy {
	var validator *model.PagePolicyValidator
	return &ReadingPolicy{Validator: validator}
}
