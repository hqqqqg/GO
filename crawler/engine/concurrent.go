package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler //调度器接口
	WorkerCount int       //10个worker
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)           //放到in里面
	WorkerChan() chan Request //我有worker，给我哪种channel
	Run()
}

// workerready拿出来了
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run() //queued.go

	//启动n个worker goroutine
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), //问scheduler要channel，对共用和各一个channel都实现，
			out, e.Scheduler) //scheduler有readynotifier不用改
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

func createWorker(in chan Request,
	out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in                //QueuedScheduler的workerChan选择了in,类型一样
			result, err := worker(request) //拿到就干活
			if err != nil {
				continue
			}
			out <- result //干完活就把结果发到out管道
		}
	}()
}
