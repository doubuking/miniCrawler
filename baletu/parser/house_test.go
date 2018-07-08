package parser

import (
	"testing"
	"miniCrawler/fetcher"
)

func TestParseHouse(t *testing.T) {
	contents,err :=  fetcher.Fetch("http://bj.baletu.com/house/1778609.html")
	//fmt.Printf("%s",contents)
	if err != nil {
		panic(err)
	}
	ParseHouse(contents)

}