package scheduler

import "go-crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQueue []engine.Request
		var workQueue []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 && len(workQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workQueue[0]
			}
			select {
			case request := <-s.requestChan:
				requestQueue = append(requestQueue, request)
			case work := <-s.workerChan:
				workQueue = append(workQueue, work)
			case activeWorker <- activeRequest:
				workQueue = workQueue[1:]
				requestQueue = requestQueue[1:]
			}
		}
	}()
}
