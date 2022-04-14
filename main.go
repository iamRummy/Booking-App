package main

import "fmt"

func main() {
	var conferencename = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v powered by Rummy Solutions\n", conferencename)

	fmt.Printf("We have a total of %v tickets and %v are still availible.\n", conferenceTickets, remainingTickets)

	fmt.Println("Recieve your tickets to attend the function.")
}
