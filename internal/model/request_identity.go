package model

type RequestIdentity struct {
	ReaderID string
	BookIDs  []string
}

type IdentityScenario struct {
	FirstReceipt string `json:"first_receipt"`
	AuditReader  string `json:"audit_reader"`
	SecondReader string `json:"second_reader"`
}
