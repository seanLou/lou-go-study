package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq := new(Square)
	sq.side = 10

	areaIntf := sq

	fmt.Printf("正方形面积为: %f\n", areaIntf.Area())
}
