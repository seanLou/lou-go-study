package main

import "fmt"

type firstS struct {
	in1 int
	in2 int
}

type secondS struct {
	b      int
	c      float32
	int    //匿名字段
	firstS //匿名字段
}

// 在一个结构体中对于每种数据类型只能有一个匿名字段
func main() {
	sec := new(secondS)

	sec.b = 6
	sec.c = 1.6
	sec.int = 7
	sec.in1 = 8
	sec.in2 = 9

	fmt.Printf("sec.b %d\n", sec.b)
	fmt.Printf("sec.c %f\n", sec.c)
	fmt.Printf("sec.int %d\n", sec.int)
	fmt.Printf("sec.in1 %d\n", sec.in1)
	fmt.Printf("sec.in2 %d\n", sec.in2)

	// 使用结构体字面量
	sec2 := secondS{1, 2.3, 8, firstS{1, 2}}
	fmt.Println("sec2 is:", sec2)
}
