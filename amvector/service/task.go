// Package service
// Date: 2024/3/6 13:34
// Author: Amu
// Description:
package service

import (
	"fmt"
	"time"
	
	"amvector/service/task"
	
	"github.com/amuluze/amutool/timex"
	"github.com/amuluze/docker"
)

type TimedTask struct {
	task   *task.Task
	ticker timex.Ticker
	stopCh chan struct{}
}

func NewTimedTask(conf *Config) *TimedTask {
	interval := conf.Task.Interval
	tk := timex.NewTicker(time.Duration(interval) * time.Second)
	manager, err := docker.NewManager()
	if err != nil {
		return nil
	}
	
	newTask := task.NewTask(interval, manager)
	
	return &TimedTask{
		task:   newTask,
		ticker: tk,
		stopCh: make(chan struct{}),
	}
}

func (a *TimedTask) Execute() {

}

func (a *TimedTask) Run() {
	for {
		select {
		case <-a.ticker.Chan():
			go a.Execute()
		case <-a.stopCh:
			fmt.Println("task exit")
			return
		}
	}
}

func (a *TimedTask) Stop() {
	close(a.stopCh)
}
