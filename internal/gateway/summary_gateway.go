package gateway

import "errors"

var ErrTemporary = errors.New("delivery acknowledgement timed out")

type SummaryGateway struct {
	attempts   int
	deliveries int
}

func (g *SummaryGateway) Send(idempotencyKey string) error {
	g.attempts++
	g.deliveries++
	if g.attempts == 1 {
		return ErrTemporary
	}
	return nil
}

func (g *SummaryGateway) Attempts() int   { return g.attempts }
func (g *SummaryGateway) Deliveries() int { return g.deliveries }
