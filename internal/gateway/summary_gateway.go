package gateway

import "errors"

// ErrTemporary 表示投递已送达，但确认超时：调用方应按幂等键重试，
// 但重试不会产生新的投递（同一幂等键只投递一次）。
var ErrTemporary = errors.New("delivery acknowledgement timed out")

// SummaryGateway 是摘要投递的远程网关。
//
// 投递按幂等键去重：同一幂等键首次 Send 会真正投递并计入 deliveries；
// 若该次确认超时返回 ErrTemporary，调用方以相同幂等键重试时，网关识别出
// 该键已投递过，直接返回 nil 且不再重复投递——避免收件人收到两份。
type SummaryGateway struct {
	attempts   int
	deliveries int
	delivered  map[string]bool
}

func (g *SummaryGateway) Send(idempotencyKey string) error {
	g.attempts++
	if g.delivered == nil {
		g.delivered = make(map[string]bool)
	}

	// 首次见到该幂等键才真正投递。
	if !g.delivered[idempotencyKey] {
		g.delivered[idempotencyKey] = true
		g.deliveries++
		if g.attempts == 1 {
			// 首次投递已送达，仅确认超时，触发调用方重试。
			return ErrTemporary
		}
	}
	// 已投递过的幂等键：幂等回放，不重复投递。
	return nil
}

func (g *SummaryGateway) Attempts() int   { return g.attempts }
func (g *SummaryGateway) Deliveries() int { return g.deliveries }
