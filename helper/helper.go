package helper

import "strings"

func ValidateUserInput(userFirstName string, userLastName string, userEmail string, userTickets uint) bool {
	return len(userFirstName) > 0 && len(userLastName) > 0 && strings.Contains(userEmail, "@") && userTickets > 0
}
