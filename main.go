package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceName = "Go Conference 2024"
const conferenceTickets = 500

var remainingTickets uint = conferenceTickets
var bookings = make([]map[string]string, 0)

func main() {
	greetUsers()

	for {
		if remainingTickets == 0 {
			fmt.Println("Sorry, we have sold out of tickets")
			break
		}

		userFirstName, userLastName, userEmail, userTickets := getUserInput()
		if !helper.ValidateUserInput(userFirstName, userLastName, userEmail, userTickets) {
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
			bookTickets(userTickets, userFirstName, userLastName, userEmail)
		}

		displayCurrentBookings()
	}
}

func bookTickets(userTickets uint, userFirstName string, userLastName string, userEmail string) {
	var userData = make(map[string]string)
	userData["firstName"] = userFirstName
	userData["lastName"] = userLastName
	userData["email"] = userEmail
	userData["nbOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	remainingTickets -= userTickets
	bookings = append(bookings, userData)
	confirmUserPurchase(userFirstName, userTickets, userEmail, remainingTickets, conferenceName)
}

func greetUsers() {
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

func confirmUserPurchase(userFirstName string, userTickets uint, userEmail string, remainingTickets uint, conferenceName string) {
	fmt.Printf("Thank you %v, you have purchased %v tickets\n", userFirstName, userTickets)
	fmt.Printf("You will receive an email to %v confirming your purchase\n", userEmail)
	fmt.Printf("There are now %v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func displayCurrentBookings() {
	fmt.Print("\n\n")
	fmt.Println("Current bookings:")
	fmt.Println("-----------------")
	for _, booking := range bookings {
		fmt.Printf("%s %s (%s) - %s tickets\n", booking["firstName"], booking["lastName"], booking["email"], booking["nbOfTickets"])
	}
	fmt.Println("-----------------")
	fmt.Print("\n\n")
}
