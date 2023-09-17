package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string
	fmt.Print("hello, world")

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	// var userName string
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		fmt.Println("Enter FirstName")
		fmt.Scan(&firstName)

		fmt.Println("Enter LastName")
		fmt.Scan(&lastName)

		fmt.Println("Enter Email")
		fmt.Scan(&email)

		fmt.Println("Enter Number of Tickets")
		fmt.Scan(&userTickets)

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail = strings.Contains(email, "@")
		var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v : \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining : \n", remainingTickets)
			fmt.Printf("%v . These are all the bookings \n", bookings)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("sold out.")
				break
			}
		} else {
			fmt.Printf("Your input is invalid")
			continue
		}
	}

}
