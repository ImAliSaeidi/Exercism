package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	println(strings.Repeat("-", 20))
	println("Old:\n" + oldMsg)
	newMessage := strings.ReplaceAll(oldMsg, "*", "")
	newMessage = strings.ReplaceAll(newMessage, "\n", "")
	newMessage = strings.ReplaceAll(newMessage, "\t", "")
	println("Without Stars:\n" + newMessage)
	newMessage = strings.TrimLeft(newMessage, " ")
	println("Trim Left: " + newMessage)
	newMessage = strings.TrimRight(newMessage, " ")
	println("Trim Right: " + newMessage)
	println(strings.Repeat("-", 20))
	return newMessage
}
