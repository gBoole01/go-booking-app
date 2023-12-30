package main

import (
	"fmt"
	"strconv"
)

func main() {
	const conferenceName = "Go Conference 2024"
	const conferenceTickets = 500
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We currently have %v tickets remaining out of %v\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string

	for {
		var userFirstName string
		var userLastName string
		var userEmail string
		var userTickets uint

		fmt.Print("Please enter your first name: ")
		fmt.Scanln(&userFirstName)
		fmt.Print("Please enter your last name: ")
		fmt.Scanln(&userLastName)
		fmt.Print("Please enter your email address: ")
		fmt.Scanln(&userEmail)

		fmt.Print("How many tickets would you like to purchase? ")
		fmt.Scanln(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Sorry %v, we only have %v tickets remaining\n", userFirstName, remainingTickets)
		} else {
			remainingTickets -= userTickets
			bookings = append(bookings, userFirstName+" "+userLastName+" - "+userEmail+" ("+strconv.Itoa(int(userTickets))+" tickets)")
			fmt.Printf("Thank you %v, you have purchased %v tickets\n", userFirstName, userTickets)
			fmt.Printf("You will receive an email to %v confirming your purchase\n", userEmail)
			fmt.Printf("There are now %v tickets remaining for %v\n", remainingTickets, conferenceName)
		}

		fmt.Print("\n\n")
		fmt.Println("Current bookings:")
		fmt.Println("-----------------")
		for _, booking := range bookings {
			if booking != "" {
				fmt.Println(booking)
			}
		}
		fmt.Println("-----------------")
		fmt.Print("\n\n")

		if remainingTickets == 0 {
			fmt.Println("Sorry, we have sold out of tickets")
			break
		}
	}
}
