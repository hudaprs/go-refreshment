package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Conference Go"
	const totalTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking app \n", conferenceName)
	fmt.Printf("We have total of %v and have still available about %v\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName, lastName, email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address:")
		fmt.Scan(&email)
		fmt.Println("How many ticket you want to book?")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booked %v ticket(s), you will receive an email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		users := []string{}
		for _, booking := range bookings {
			names := strings.Fields(booking)
			users = append(users, names[0])
		}

		fmt.Printf("These are all our bookings %v\n", users)
	}
}
