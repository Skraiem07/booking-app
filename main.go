package main

import (
	"fmt"
)

var conferenceName = "Go Conference" // In Go, variables are containers for storing data values, and they are created using the var
const conferenceTickets = 50         //the value can not be changed
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData = struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	//unit takes only positive ints
	fmt.Printf("conferenceName is %T, conferenceTicket is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		//input validation
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("These first names of bookings are: %v\n", firstNames)

			noTicketRemaining := remainingTickets == 0
			if noTicketRemaining {
				//end program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of ticket you entered is invalid")
			}
		}
	}
	/*city := "London"
	switch city{
	case "New York":
		//execute code
	case "Singapore","Hong Kong":
		//execute 2nd code
	case "London", "Berlin":
		//
	case "Mexico City":
	default:
		fmt.Println("No valid city selected")
	}*/
}
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	/*Data types in Go
	--> most basic ones:
		- String
		- Integer
		- Array: have a fixed size
		- Slice is an array with a dynamic size
	*/

	// ask the use for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName) //this & is a pointer : the value is stored in Memory , a pointer is a variable that references the memory address of another variable

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets

}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	/*fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))*/

	fmt.Printf("thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n ", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining ticket for %v\n", remainingTickets, conferenceName)

}
