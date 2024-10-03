package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return ("Welcome to the Tech Palace, " + strings.ToUpper(customer))
	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	b := strings.Repeat("*", numStarsPerLine)
	s := b + "\n" + welcomeMsg + "\n" + b
	return s
	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	s := strings.Split(oldMsg, "\n")[1]
	s = strings.Trim(s, "*")
	s = strings.TrimSpace(s)
	return s
	panic("Please implement the CleanupMessage() function")
}