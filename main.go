package main

import (
	"fmt"
	"strings"
)

func main() {
	var eventName string = "Go Event"
	const eventTickets uint = 50
	var remainingTickets uint = 50
	var customers []string
	fmt.Printf("Welcome to this %v ", eventName)
	fmt.Printf("we have a total of %v and available tickets are %v", eventTickets, remainingTickets)
	fmt.Println("Get your tickets here")

	for {
		var firstName string
		var lastName string
		var emailAddress string
		var userTickets uint
		fmt.Println("Enter you first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter you last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&emailAddress)
		fmt.Println("Enter amount of tickets to book: ")
		fmt.Scan(&userTickets)

		var validName bool = len(firstName) >= 2 && len(lastName) >= 2
		var validEmail bool = strings.Contains(emailAddress, "@")
		var validTicketAmount bool = userTickets > 0 && userTickets <= remainingTickets

		if validName && validEmail && validTicketAmount {
			remainingTickets = remainingTickets - userTickets
			customers = append(customers, firstName+" "+lastName)

			fmt.Printf("%v %v you have booked %v ticket, we will send confirmation to %v\n", firstName, lastName, userTickets, emailAddress)
			fmt.Printf("The amount of remaining tickets are %v\n", remainingTickets)

			var firstNamesSlice []string
			for _, customer := range customers {
				var name = strings.Fields(customer)
				var firstName = name[0]
				firstNamesSlice = append(firstNamesSlice, firstName)
			}
			fmt.Printf("People with tickets are: %v\n", firstNamesSlice)

			var noTicketsRemaining bool = remainingTickets == 0
			if noTicketsRemaining {
				fmt.Print("The tickets are sold out!")
				break
			}
		} else {
			if !validName {
				fmt.Println("Invalid name, try again")
			}
			if !validEmail {
				fmt.Println("Invalid email, try again")
			}
			if !validTicketAmount {
				fmt.Println("Invalid ticket input, try again")
			}
		}
	}
}
