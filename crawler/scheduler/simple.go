// 每个worker共用一个channel
package scheduler

import "google/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

// 要workerchan
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

// 不干活，但是得实现，方便接口
func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

// 建workerchan
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(
	r engine.Request) {
	go func() { s.workerChan <- r }() //解决阻塞，每个request建立一个goroutine

}
