//	source code: https://stackoverflow.com/questions/19549199/golang-implementing-a-cron-executing-tasks-at-a-specific-time
package fixed_time_ticker

import (
	"time"
)

type Ticker struct {
	*time.Timer
	interval    time.Duration
	hour2Tick   int
	minute2Tick int
	second2Tick int
}

func (t Ticker) getNextTickDuration() time.Duration {
	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), t.hour2Tick, t.minute2Tick, t.second2Tick, 0, time.UTC)
	if nextTick.Before(now) {
		nextTick = nextTick.Add(t.interval)
	}
	return nextTick.Sub(time.Now())
}

func NewTicker(interval time.Duration, hour2Tick int, minute2Tick int, second2Tick int) *Ticker {
	t := &Ticker{
		interval:    interval,
		hour2Tick:   hour2Tick,
		minute2Tick: minute2Tick,
		second2Tick: second2Tick,
	}
	t.Timer = time.NewTimer(t.getNextTickDuration())
	return t
}

func (t Ticker) SetNextTime2Tick() {
	t.Reset(t.getNextTickDuration())
}
