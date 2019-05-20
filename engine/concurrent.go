package engine

import "log"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	//WorkerReady(chan Request)
	//Run()
}

func (engine *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	engine.Scheduler.ConfigureMasterWorkerChan(in)
	//engine.Scheduler.Run()

	for i := 0; i < engine.WorkCount; i++ {
		createWork(in, out)
	}

	for _, request := range seeds {
		engine.Scheduler.Submit(request)
	}

	for {
		result := <-out

		for _, item := range result.Items {
			log.Printf("Got item %+v", item)
		}

		for _, request := range result.Requests {
			engine.Scheduler.Submit(request)
		}
	}
}

func createWork(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := Worker(request)

			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
