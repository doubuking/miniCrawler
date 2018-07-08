package model

import "encoding/json"

type House struct {
	Name string	//小区名称
	Price int	//价格
	Area string	//面积
	Direction string //朝向
	Traffic string //交通
	HouseType string //户型
	Floor string //楼层
	LifeType string //类型
	PaymentMethod string //付款方式
	Region string //区域
	Address string //地址
}

func FromJsonObj(o interface{}) (House,error) {
	var house House
	s, err := json.Marshal(o)

	if err != nil{
		return house,err
	}

	err = json.Unmarshal(s,&house)

	return house,err
}
