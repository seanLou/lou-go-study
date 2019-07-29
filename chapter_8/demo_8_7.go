package main

import "fmt"

type Demo7 struct {
	ID      int
	Name    string
	Phone   string
	Address string
}

func main() {
	var lou Demo7

	lou.ID = 1000
	lou.Name = "Guanyang"
	lou.Phone = "13800000000"
	lou.Address = "Hangzhou"

	fmt.Printf("值传递前 Guest ID: %d\n", lou.ID)
	operateDemo(lou)
	fmt.Printf("值传递后 Guest ID: %d\n", lou.ID)

	fmt.Printf("\n指针传递前 Guest ID: %d\n", lou.ID)
	operateDemo2(&lou)
	fmt.Printf("指针传递后 Guest ID: %d\n", lou.ID)
}

// 值传递
func operateDemo(demo7 Demo7) {
	demo7.ID = 10001
}

// 指针传递
func operateDemo2(demo7 *Demo7) {
	demo7.ID = 10002
}
