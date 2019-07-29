package main

import "fmt"

type Customer struct {
	ID      int
	Name    string
	Phone   string
	Address string
}

func main() {
	var lou Customer

	lou.ID = 1000
	lou.Name = "Guanyang"
	lou.Phone = "13800000000"
	lou.Address = "Hangzhou"

	fmt.Printf("lou ID: %d\n", lou.ID)
	fmt.Printf("lou Name: %s\n", lou.Name)
	fmt.Printf("lou Phone: %s\n", lou.Phone)
	fmt.Printf("lou Address: %s\n", lou.Address)
}
