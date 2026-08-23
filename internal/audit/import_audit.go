package audit

type ImportAudit struct {
	status string
}

// Finish 根据 runErr 记录审计状态，并把错误原样返回，避免吞掉失败。
func (a *ImportAudit) Finish(runErr error) error {
	if runErr != nil {
		a.status = "failed"
		return runErr
	}
	a.status = "success"
	return nil
}

func (a *ImportAudit) Status() string { return a.status }
