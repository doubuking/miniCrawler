package parser

import (
	"testing"
	"miniCrawler/fetcher"
	"fmt"
)

func TestParseCity(t *testing.T) {
	contents,err :=  fetcher.Fetch("http://sh.baletu.com/zhaofang/")
	//fmt.Printf("%s",contents)
	if err != nil {
		panic(err)
	}
	result := ParseCity(contents,"sh.baletu.com")
	fmt.Println("=================")
	for _,aa := range result.Items {
		fmt.Println(aa)
	}
}
