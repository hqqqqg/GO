package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler //调度器接口
	WorkerCount int       //10个worker
}

type Scheduler interface {
	Submit(Request) //submit?
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//建worker
	//先创建两个channel 无缓冲
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in) //crawler/scheduler/simple.go in和workerchan是同一channel

	//启动10个worker goroutine
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
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
	in chan Request, out chan ParseResult) {
	go func() {
		for { //死循环
			request := <-in                //阻塞，等in管道有数据
			result, err := worker(request) //拿到就干活
			if err != nil {
				continue
			}
			out <- result //干完活就把结果发到out管道
		}
	}()
}
