package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	remainingTicket := 50
	
	
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %d conference tickets, %d are still available\n", conferenceTickets, remainingTicket)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Print("Enter your first Name: ")
	fmt.Scan(&firstName)
	
	fmt.Print("Enter your last Name: ")
	fmt.Scan(&lastName)
	
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You email is %v.\n", firstName, lastName, userTickets, email)
}