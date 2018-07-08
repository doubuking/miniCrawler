package parser

import (
	"regexp"
	"miniCrawler/model"
	"strconv"
	"miniCrawler/engine"
	"strings"
)

var nameRe = regexp.MustCompile(`<div class="basic-title basic-title-margin">[^<]*<h1>[^<]*<a href="http://.*?html"[^>]+>([^<]+)</a>`)

var priceRe = regexp.MustCompile(`<li class="price">([^<]*)<span>.*?</span></li>`)

var areaRe = regexp.MustCompile(`<li class="cent">([^<]+)<span>M²</span></li>`)

var directionRe = regexp.MustCompile(`<li class="cent">.*?</li>[^<]*<li>([^<]+)</li>`)

var trafficRe = regexp.MustCompile(`<dt>交通：</dt>[^<]*<dd>([^<]+)</dd>`)

var houseTypeRe = regexp.MustCompile(`<dt>户型：</dt>[^<]*<dd>([^<]+)</dd>`)

var floorRe = regexp.MustCompile(`<dt>楼层：</dt>[^<]*<dd>([^<]+)</dd>`)

var lifeTypeRe = regexp.MustCompile(`<dt>类型：</dt>[^<]*<dd>[^<]*<a href="http://.*?html">([^<]*)</a>([^<]*)</dd>`)

var paymentMethodRe = regexp.MustCompile(`<dt>付款：</dt>[^<]*<dd>([^<]+)</dd>`)

var regionRe = regexp.MustCompile(`<dt>区域：</dt>[^<]*<dd>[^<]*<a href="http://.*?html">([^<]+)</a>[^<]*<a href="http://.*?html">([^<]+)</a>[^<]*</dd>`)

var addressRe = regexp.MustCompile(`<dt>地址：</dt>[^<]*<dd>([^<]+)</dd>`)

var idUrlRe = regexp.MustCompile(`http://sh.baletu.com/house/([\d]+).html`)

func ParseHouse(contents []byte,url string) engine.ParseResult  {

	house := model.House{}

	house.Name = extractString(contents,nameRe)

	price,err := strconv.Atoi(extractString(contents,priceRe))

	if err == nil{
		house.Price = price
	}

	house.Area = extractString(contents,areaRe)

	house.Direction = extractString(contents,directionRe)

	house.Traffic = extractString(contents,trafficRe)

	house.HouseType = extractString(contents,houseTypeRe)

	house.Floor = extractString(contents,floorRe)

	house.LifeType = extractString(contents,lifeTypeRe)

	house.PaymentMethod = extractString(contents,paymentMethodRe)

	house.Region = extractString(contents,regionRe)

	house.Address = extractString(contents,addressRe)

	result := engine.ParseResult{
		Items:[]engine.Item{
			{
				Url:url,
				Type:"baletu",
				Id:extractString([]byte(url),idUrlRe),
				Payload:house,
			},
		},
	}

	return result

}


//处理重复的访问

func extractString(contents []byte,re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)

	if len(match) >= 3 {
		return delStingSpace(string(match[1])+string(match[2]))
	}
	if len(match) >= 2 {
		return  delStingSpace(string(match[1]))
	}



	return ""

}

func delStingSpace(str string)string  {
	//去除空格
	str = strings.Replace(str," ","",-1)
	//去除换行符
	return strings.Replace(str,"\n","",-1)
}