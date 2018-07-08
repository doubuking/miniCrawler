package fetcher

import (
	"net/http"
	"fmt"
	"bufio"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func Fetch(url string)([]byte,error){
	//获取网页
	resp,err := http.Get(url)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return nil,fmt.Errorf("wrong status code : %d",resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	//检测编码
	e := determineRncoding(bodyReader)
	//转换编码
	utf8Reader := transform.NewReader(bodyReader,e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

//检测编码类型
func determineRncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil{
		log.Printf("Fetcher error: %v",err)
		//遇到问题返回utf8
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
