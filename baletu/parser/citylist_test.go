package parser

import (
	"testing"
	"fmt"
	"io/ioutil"
)

func TestParserCityList(t *testing.T) {
	contents ,err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil{
		panic(err)
	}
	result := ParserCityList(contents)

	fmt.Println(result)

}