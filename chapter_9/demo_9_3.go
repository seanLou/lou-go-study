package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable)  {
	fmt.Printf("资产的价值是: %f\n", asset.getValue())
}

func main() {
	var o valuable = stockPosition{"APPLE", 578.92, 23456}
	showValue(o)

	o = car{"BMW", "X7", 899999}
	showValue(o)
}