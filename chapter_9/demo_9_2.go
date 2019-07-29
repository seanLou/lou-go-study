package main

import "fmt"

type Shaper2 interface {
	Area() float32
}

type Square2 struct {
	side float32
}

func (sq *Square2) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{20, 30}
	square2 := new(Square2)
	square2.side = 15

	shaper2s := []Shaper2{r, square2}

	for n, _ := range shaper2s {
		fmt.Println("形状参数：", shaper2s[n])
		fmt.Println("形状面积：", shaper2s[n].Area())
	}

}