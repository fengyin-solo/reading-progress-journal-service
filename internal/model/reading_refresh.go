package model

const (
	RefreshRunning  = "running"
	RefreshCanceled = "canceled"
	RefreshDone     = "done"
)

type ReadingRefresh struct {
	ID    string `json:"id"`
	State string `json:"state"`
	Calls int    `json:"calls"`
}

func (r *ReadingRefresh) MarkCall() {
	r.Calls++
	r.State = RefreshRunning
}

func (r *ReadingRefresh) Finish(state string) {
	r.State = state
}
