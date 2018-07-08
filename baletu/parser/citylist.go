package parser

import (
	"regexp"
	"miniCrawler/engine"
	"strings"
	"fmt"
)

const CityListRe  =`<a id="navM2" mark="I2" href="([^"]+)">([^<]+)</a>`
func ParserCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(CityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _,m := range matches{
		baseUrl := string(m[1])

		baseUrls := strings.Split(baseUrl,"/")

		fmt.Println(baseUrls)
		//result.Items = append(result.Items,string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseCity(c,baseUrls[2])
			},
		})

	}
	return result
}
