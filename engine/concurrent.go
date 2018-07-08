package engine

import (
	"github.com/mediocregopher/radix.v2/pool"
	"time"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine)Run(seeds ...Request)  {

	//in := make(chan Request)
	out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	//redis pool
	//pool redis
	redisPool,err := pool.New("tcp","localhost:6379",10)
	if err != nil{
		//log.Fatalln("Redis pool created failed.")
		panic(err)
	}else {
		//redis 有保持时间 在一定时间内没有动作就会断开 所以这个要每隔一段时间ping 一下保持活跃
		go func() {
			for {
				redisPool.Cmd("PING")
				time.Sleep(3*time.Second)
			}
		}()
	}

	for i:=0;i<e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler,redisPool)
	}

	for _,r := range seeds{

		e.Scheduler.Submit(r)

	}

	for {
		result := <- out
		for _,item := range result.Items {

			go func() {e.ItemChan <- item}()
		}

		for _,request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult,ready ReadyNotifier,redisPool *pool.Pool){
	go func() {

		hyperLogLogKey := "quchong_hpll"


		for {
			ready.WorkerReady(in)
			request := <- in


			ret ,err := redisPool.Cmd("PFADD",hyperLogLogKey,request.Url,"EX",86400).Int()
			if err != nil{
				panic("出错了")
			}
			if ret != 1{
				continue
			}

			result,err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}
