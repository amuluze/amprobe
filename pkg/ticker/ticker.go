// Package ticker
// Date: 2024/3/6 10:55
// Author: Amu
// Description:
package ticker

import "time"

type Ticker interface {
	Chan() <-chan time.Time
	Stop()
}

type realTicker struct {
	*time.Ticker
}

// NewTicker returns a Ticker.
func NewTicker(d time.Duration) Ticker {
	return &realTicker{
		Ticker: time.NewTicker(d),
	}
}

func (r *realTicker) Chan() <-chan time.Time {
	return r.C
}

func (r *realTicker) Stop() {
	r.Ticker.Stop()
}
