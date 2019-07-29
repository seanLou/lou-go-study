package main

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

func main() {
	b := B{A{100, 50}, 20.8, 30.9}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}
