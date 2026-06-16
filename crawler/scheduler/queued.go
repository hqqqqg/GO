// 控制request和worker，建两个队列
// 每个worker有自己的channel
package scheduler

import "google/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request      //submit往这里
	workerChan  chan chan engine.Request //workerready往这里
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request) //给每个worker自己的channel
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}
func (s *QueuedScheduler) WorkerReady( //有一个worker准备好接收request
	w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	//生成,要变为指针接收者
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		//声明队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-s.requestChan: //有新的request提交
				requestQ = append(requestQ, r)
			case w := <-s.workerChan: //有新的worker提交
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: //都在排队
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
