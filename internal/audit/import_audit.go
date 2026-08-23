package audit

type ImportAudit struct {
	status string
}

func (a *ImportAudit) Finish(runErr error) error {
	a.status = "success"
	return nil
}

func (a *ImportAudit) Status() string { return a.status }
