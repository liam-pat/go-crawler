package scheduler

import "go-crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (simpleScheduler *SimpleScheduler) WorkerChan() chan engine.Request {
	return simpleScheduler.workerChan
}

func (simpleScheduler *SimpleScheduler) WorkReady(chan engine.Request) {
}

func (simpleScheduler *SimpleScheduler) Run() {
	simpleScheduler.workerChan = make(chan engine.Request)
}

func (simpleScheduler *SimpleScheduler) Submit(request engine.Request) {
	go func() { simpleScheduler.workerChan <- request }()
}
