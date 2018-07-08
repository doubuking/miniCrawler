package main

import (
	"miniCrawler/baletu/parser"
	"miniCrawler/engine"
	"miniCrawler/secheduler"
	"miniCrawler/persist"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:"http://sh.baletu.com/zhaofang/",
	//	ParserFunc:parser.ParserCityList,
	//})

	itemChan,err := persist.ItemSaver("dating_profile")

	if err != nil{
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:&secheduler.QueueScheduler{},
		WorkerCount:100,
		ItemChan:itemChan,
	}
	e.Run(engine.Request{
		Url:"http://sh.baletu.com/zhaofang/",
		ParserFunc: func(b []byte) engine.ParseResult {
				return parser.ParseCity(b,"sh.baletu.com")
		},
	})



}





