package model

type ContextScenarioResult struct {
	FirstError  string `json:"first_error"`
	FirstCalls  int    `json:"first_calls"`
	SecondError string `json:"second_error"`
	SecondCalls int    `json:"second_calls"`
}
