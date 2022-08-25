package tickerTask

import (
	"context"
	"fmt"
	"time"

	"worktime_server/internal/user"
)

type TickerTask struct {
	user      *user.User
	store     interface{}
	intervals []*IntervalTime

	ctx    context.Context
	cancel context.CancelFunc
}

type IntervalTime struct {
	startTime time.Time
	endTime   time.Time
}

func NewTask(user *user.User) *TickerTask {
	return &TickerTask{
		user:      user,
		intervals: make([]*IntervalTime, 1),
	}
}

func (task *TickerTask) Start() {

	if task.user.ActiveTicker {
		return
	}

	interval := IntervalTime{
		startTime: time.Now(),
		endTime:   time.Now(),
	}

	task.ctx, task.cancel = context.WithCancel(task.user.Ctx)
	task.intervals = append(task.intervals, &interval)

	task.user.ActiveTicker = true

	go func() {

		ticker := time.NewTicker(1 * time.Second)

		for {
			select {
			case <-ticker.C:

				interval.endTime = time.Now()
				fmt.Println(interval.endTime.String())

			case <-task.ctx.Done():
				ticker.Stop()
				return
			}
		}

	}()
}

func (task *TickerTask) Stop() {
	task.user.ActiveTicker = false
	task.cancel()
}
