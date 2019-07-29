package main

import (
	"fmt"

	"github.com/seanLou/lou-go-study/struct_pack"
)

func main() {
	louStruct := new(struct_pack.LouStruct)
	louStruct.Id = 100
	louStruct.Name = "Guanyang"

	fmt.Printf("id is %d\n", louStruct.Id)
	fmt.Printf("name is %s\n", louStruct.Name)
}
