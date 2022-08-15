package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	request1()
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the Go Web API!")
	fmt.Println("Endpoint Hit: Home Page")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "Caio Veloso"
	fmt.Fprintf(response, "About me...")
	fmt.Println("Endpoint Hit: About", who)
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func bookingRegister() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 // uint does not accept negative
	fmt.Printf("ConferenceTickets is %T, RemainingTicket is %T, ConferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, but only %v remaining.\n", conferenceTickets, remainingTickets)

	var bookings []string

	for { // for is the only loop command in go
		var firstName string
		var LastName string
		var userTickets uint
		// gets user name
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName) // pointer are also called "special variable" in go
		fmt.Println("Enter your last name")
		fmt.Scan(&LastName)
		fmt.Printf("Hello %v. How many ticket do you want?\n", firstName)
		fmt.Scan(&userTickets)
		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets available. You cannot book %v\n", remainingTickets, userTickets)
			continue
		}

		//bookings[0] = userName
		bookings = append(bookings, firstName+" "+LastName)
		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you! %v, you will received the %v tickets required\n", firstName, userTickets)
		fmt.Printf("The remaining tickets now are: %v\n", remainingTickets)

		firstNames := []string{}
		for _, booking := range bookings { // The _ is used here that we will not need the index return for the range function
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The list of bookings is %v\n", firstNames)

		var noTickertsRemainng bool = remainingTickets == 0
		if noTickertsRemainng {
			fmt.Println("Our conference is booked out. Come back next year")
			break
		}
	}
}
