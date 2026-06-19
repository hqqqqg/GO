package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler //调度器接口
	WorkerCount      int       //10个worker
	ItemChan         chan Item
	RequestProcessor Processor
}
type Processor func(Request) (ParseResult, error)
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
		e.createWorker(e.Scheduler.WorkerChan(), //问scheduler要channel，对共用和各一个channel都实现，
			out, e.Scheduler) //scheduler有readynotifier不用改
	}
	//投喂种子任务
	for _, r := range seeds {
		e.Scheduler.Submit(r) //scheduler/simple.go  往in里面塞数据
	}
	for { //没有退出会死掉
		result := <-out //接收结果
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request) //跳到createworker里面的request := <-in 新任务又塞进in
		}
	}
}

func (e *ConcurrentEngine) createWorker(
	in chan Request,
	out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in //QueuedScheduler的workerChan选择了in,类型一样
			result, err := e.RequestProcessor(
				request)
			if err != nil {
				continue
			}
			out <- result //干完活就把结果发到out管道
		}
	}()
}

// 去重
var visitedUrls = make(map[string]bool) //记录已经访问过的url

func isDuplicate(url string) bool {
	if visitedUrls[url] { //这个url之前是否已经存在
		return true //存在
	}
	visitedUrls[url] = true //已访问
	return false            //不是重复的
}
