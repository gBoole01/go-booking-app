package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const conferenceName = "Go Conference 2024"
	const conferenceTickets = 500
	var remainingTickets uint = conferenceTickets
	var bookings []string

	greetUsers(conferenceName, remainingTickets, conferenceTickets)

	for {
		userFirstName, userLastName, userEmail, userTickets := getUserInput()
		isValid := validateUserInput(userFirstName, userLastName, userEmail, userTickets)
		if !isValid {
			fmt.Println("Sorry, you have entered invalid details")
			fmt.Println("Please try again")
			continue
		}

		if userTickets > remainingTickets {
			fmt.Printf("Sorry %v, we only have %v tickets remaining\n", userFirstName, remainingTickets)
			fmt.Printf("Please try again\n")
			continue
		}

		if userTickets <= remainingTickets {
			remainingTickets -= userTickets
			bookings = append(bookings, userFirstName+" "+userLastName+" - "+userEmail+" ("+strconv.Itoa(int(userTickets))+" tickets)")
			confirmUserPurchase(userFirstName, userTickets, userEmail, remainingTickets, conferenceName)
		}

		displayCurrentBookings(bookings)

		if remainingTickets == 0 {
			fmt.Println("Sorry, we have sold out of tickets")
			break
		}
	}
}

func greetUsers(conferenceName string, remainingTickets uint, conferenceTickets int) {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets\n", conferenceTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
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
	return userFirstName, userLastName, userEmail, userTickets
}

func validateUserInput(userFirstName string, userLastName string, userEmail string, userTickets uint) bool {
	return len(userFirstName) > 0 && len(userLastName) > 0 && strings.Contains(userEmail, "@") && userTickets > 0
}

func confirmUserPurchase(userFirstName string, userTickets uint, userEmail string, remainingTickets uint, conferenceName string) {
	fmt.Printf("Thank you %v, you have purchased %v tickets\n", userFirstName, userTickets)
	fmt.Printf("You will receive an email to %v confirming your purchase\n", userEmail)
	fmt.Printf("There are now %v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func displayCurrentBookings(bookings []string) {
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
}
