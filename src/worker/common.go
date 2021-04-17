package worker

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Task ...
type Task = func()

// AppWorker ...
type AppWorker struct {
	Task   Task
	period int
}

// SetTask ..
func (worker *AppWorker) SetTask(fn Task) *AppWorker {
	worker.Task = fn
	return worker
}

// SetRepeatPeriod ...
func (worker *AppWorker) SetRepeatPeriod(seconds int) *AppWorker {
	worker.period = seconds
	return worker
}

// Execute ...
func (worker *AppWorker) Execute() {
	tick := time.NewTicker(time.Second * time.Duration(worker.period))
	go worker.scheduler(tick)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	tick.Stop()
	os.Exit(0)
}

func (worker *AppWorker) scheduler(tick *time.Ticker) {
	for range tick.C {
		worker.Task()
	}
}
