package main

import "fmt"

type myStruct struct {
	intValue   int
	floatValue float32
	str        string
}

func main() {
	// ms 是一个指向 myStruct 的指针
	ms := new(myStruct)
	ms.intValue = 10
	ms.floatValue = 1.28
	ms.str = "Google"
	fmt.Printf("ms int:%d\n", ms.intValue)
	fmt.Printf("ms float:%f\n", ms.floatValue)
	fmt.Printf("ms string:%s\n", ms.str)
	fmt.Println(ms)

	// 混合字面量语法， 会调用 new() 初始化
	ms2 := &myStruct{100, 2.5, "Hello"}
	fmt.Printf("ms2 int:%d\n", ms2.intValue)
	fmt.Printf("ms2 float:%f\n", ms2.floatValue)
	fmt.Printf("ms2 string:%s\n", ms2.str)
	fmt.Println(ms2)
}
