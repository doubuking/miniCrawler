package parser

import (
	"miniCrawler/engine"
	"regexp"
	"strconv"
)

const cityRe  =`<a target="_blank" href="(http://.*?html)[^>]+>([^<]+)</a>`
const totalRe =`var totalPage = ([0-9]*);`
func ParseCity(contents []byte,baseUrl string) engine.ParseResult  {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}


	for _,m := range matches{
		//result.Items = append(result.Items,strings.TrimSpace(string(m[2])))
		Url := string(m[1])
		result.Requests = append(
			result.Requests,engine.Request{
				Url:Url,
				ParserFunc: func(b []byte) engine.ParseResult {
					return ParseHouse(b,Url)
				},
			})

		//log.Printf("%s",m)
	}


	pageTotalre := regexp.MustCompile(totalRe)
	pageMatches := pageTotalre.FindSubmatch(contents)
	pageTotal,err := strconv.Atoi(string(pageMatches[1]))

	if err != nil{
		pageTotal = 1
	}

	for i := 0;i <= pageTotal;i++ {
		//result.Items = append(result.Items,baseUrl+"/zhaofang/p"+strconv.Itoa(i)+"/")
		result.Requests = append(
			result.Requests,engine.Request{
				Url:"http://"+baseUrl+"/zhaofang/p"+strconv.Itoa(i)+"/",
				ParserFunc: func(c []byte) engine.ParseResult {
					return ParseCity(c,baseUrl)
				},
			})
	}

	return result
}
