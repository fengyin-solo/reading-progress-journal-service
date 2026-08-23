package model

type PublishState string

const (
	PublishPending   PublishState = "pending"
	PublishRetrying  PublishState = "retrying"
	PublishCompleted PublishState = "completed"
)

// stateRank 反映发布状态的单调进度：pending < retrying < completed。
// 用于阻止旧回调把已完成的状态倒退回重试中。
var stateRank = map[PublishState]int{
	PublishPending:   0,
	PublishRetrying:  1,
	PublishCompleted: 2,
}

type SummaryPublishJob struct {
	ID      string
	State   PublishState
	Version int
}

// Apply 更新任务状态，遵循单调推进：
//   - 版本严格单调递增，旧版本的回调一律拒绝（防止第一轮延迟回调倒退第二轮结果）；
//   - 即便版本相等（同一回调等幂回放），也不得把已完成状态倒退。
func (j *SummaryPublishJob) Apply(state PublishState, version int) {
	if version < j.Version {
		return
	}
	if version == j.Version && stateRank[state] < stateRank[j.State] {
		return
	}
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
