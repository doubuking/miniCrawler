package secheduler

import "miniCrawler/engine"

type SimpleScheduler struct {
	wokerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.wokerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.wokerChan = make(chan engine.Request)
}


func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan

	go func() { s.wokerChan <- r }()

}


