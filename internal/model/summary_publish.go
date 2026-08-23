package model

type PublishState string

const (
	PublishPending   PublishState = "pending"
	PublishRetrying  PublishState = "retrying"
	PublishCompleted PublishState = "completed"
)

type SummaryPublishJob struct {
	ID      string
	State   PublishState
	Version int
}

func (j *SummaryPublishJob) Apply(state PublishState, version int) {
	j.State = state
	j.Version = version
}

type SummaryPublishResult struct {
	Attempts     int          `json:"attempts"`
	Deliveries   int          `json:"deliveries"`
	StoreState   PublishState `json:"store_state"`
	CacheState   PublishState `json:"cache_state"`
	StoreVersion int          `json:"store_version"`
	CacheVersion int          `json:"cache_version"`
}
