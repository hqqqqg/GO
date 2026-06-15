package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler //调度器接口
	WorkerCount int       //10个worker
}

type Scheduler interface {
	Submit(Request) //submit?
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run() //queued.go

	//启动n个worker goroutine
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}
	//投喂种子任务
	for _, r := range seeds {
		e.Scheduler.Submit(r) //scheduler/simple.go  往in里面塞数据
	}
	itemCount := 0 //计数
	for {          //没有退出会死掉
		result := <-out //接收结果
		for _, item := range result.Items {
			log.Printf("Got item #%d:%v",
				itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request) //跳到createworker里面的request := <-in 新任务又塞进in
		}
	}
}

func createWorker(
	out chan ParseResult, s Scheduler) {
	in := make(chan Request) //私有的
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in                //QueuedScheduler的workerChan选择了in,类型一样
			result, err := worker(request) //拿到就干活
			if err != nil {
				continue
			}
			out <- result //干完活就把结果发到out管道
		}
	}()
}
