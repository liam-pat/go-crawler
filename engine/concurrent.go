package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkCount        int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkReady(chan Request)
}

func (engine *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)
	out := make(chan ParseResult)
	//engine.Scheduler.ConfigureMasterWorkerChan(in)
	engine.Scheduler.Run()

	for i := 0; i < engine.WorkCount; i++ {
		engine.createWork(engine.Scheduler.WorkerChan(), out, engine.Scheduler)
	}

	for _, request := range seeds {
		engine.Scheduler.Submit(request)
	}

	for {
		result := <-out

		for _, item := range result.Items {
			go func() {
				engine.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			engine.Scheduler.Submit(request)
		}
	}
}

func (engine *ConcurrentEngine) createWork(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkReady(in)
			request := <-in
			result, err := engine.RequestProcessor(request)

			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
