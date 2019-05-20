package scheduler

import "go-crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (simpleScheduler *SimpleScheduler) ConfigureMasterWorkerChan(requests chan engine.Request) {
	simpleScheduler.workerChan = requests
}

func (simpleScheduler *SimpleScheduler) Submit(request engine.Request) {
	go func() { simpleScheduler.workerChan <- request }()
}
