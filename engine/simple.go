package engine

import (
	"log"
)
type SimpleEngine struct {

}

func (s SimpleEngine)Run(seeds ...Request)  {
	var requests []Request

	for _,r := range seeds{
		requests = append(requests,r)
	}

	repeatData := make(map[string]bool)
	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]
		//去重
		if _,ok := repeatData[r.Url];ok  {
			continue
		}

		repeatData[r.Url] = true

		parseResult, err := worker(r)

		if err != nil{
			continue
		}

		//三个点就是展开一个个加进去
		requests = append(requests,parseResult.Requests...)

		for _,item := range parseResult.Items{
			log.Printf("Got item %v  ",item)


		}


	}


}

