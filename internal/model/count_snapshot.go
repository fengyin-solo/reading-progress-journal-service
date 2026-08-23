package model

type ReadCount struct {
	BookID string
	Value  int
}

type CountScenarioResult struct {
	InitialSnapshot int `json:"initial_snapshot"`
	CachedAfterWork int `json:"cached_after_work"`
	CachedTotal     int `json:"cached_total"`
}
