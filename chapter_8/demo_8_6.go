package main

import "fmt"

type Guest struct {
	ID      int
	Name    string
	Phone   string
	Address string
}

func main() {
	var tom Guest

	tom.ID = 1001
	tom.Name = "Tom"
	tom.Phone = "13800000001"
	tom.Address = "Hangzhou"

	printGuest(&tom)

	guest := new(Guest)

	guest.ID = 1002
	guest.Name = "张三"
	guest.Phone = "13800000002"
	guest.Address = "Hangzhou"

	printGuest(guest)
}

func printGuest(c *Guest) {
	fmt.Printf("Guest ID: %d\n", c.ID)
	fmt.Printf("Guest Name: %s\n", c.Name)
	fmt.Printf("Guest Phone: %s\n", c.Phone)
	fmt.Printf("Guest Address: %s\n", c.Address)
}
