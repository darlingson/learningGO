package main

import "fmt"

func main() {
	fmt.Print("hello, world")

	var userName string
	var userTickets int

	fmt.Scan(&userName)

	userTickets = 2

	fmt.Printf("Hello%v you have %v", userName, userTickets)
}
