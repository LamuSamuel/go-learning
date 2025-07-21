package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var cruizeName = "Sam's Cruize" // declaring a variable in go and can be updated anytime
const total_Tickets int = 50    // constant , will not be able to change this later.
var remaining_Tickets uint = 50
var bookings = make([]map[string]string, 0) // creating empty list of maps
var wg = sync.WaitGroup{}                   // a Waitgroup is a synchronization primitive used to wait for a group of goroutines to finish their execution.

// Some imp notes
// Go discovers the errors at compile time not at the run time.
// Arrays in go have a fixed size
// to ignore variables we can use blank identifiers such as _
// waitgroup{} - add , wait and done.
func main() {
	//remaining_Tickets := 50         // variable can also be declared this way , another way from line 7, this way go infer's what kind of variable it is. but we dont want negatives so let's declare it using uint.
	//Println - new line at end , Printf - formats , Print - no formatting,
	// fmt.Printf("%T", remaining_Tickets) - to print a type of the data
	//fmt.Println("Hello Welcome to ", cruizeName) //to print we need to use fmt library and then print a statement

	greetUser()

	// var bookings [50]string   Array declaration , fixed of 50
	//var bookings []string // slide declaration , slices are dynamic and flexible arrays , which can grow or shrink according to the usecase
	//bookings := []string{} // declaring a slice in another way
	//for { // the only for loop in go .

	// it keeps track of how many goroutines are running and lets the main goroutine wait until all of them are done.
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remaining_Tickets)

	if isValidName && isValidEmail && isValidTicket {
		bookTickets(userTickets, firstName, lastName, email)
		wg.Add(1) // adding counter 1

		go sendTicket(userTickets, firstName, lastName, email) // very crucial keyword go helps in multithreading and concurrency
		//if there is another goroutine then we add the count say wg.Add(2), increment counter to two ,so keeps track of how many goroutines are in present
		firstNames := getFirstnames()
		fmt.Printf("The first names of people who booked tickets are:  %v\n", firstNames)

		//ticketsSold := remaining_Tickets == 0 // checking via bool and storing in ticketsSold
		if remaining_Tickets == 0 { // if statement in go
			fmt.Printf("Sorry we are sold out, Come back for the spring season")
			//break // usage of break , so as to cease the for loop when the count of remaining tickets = 0
		}

		//} else if userTickets > remaining_Tickets {
		//	fmt.Printf("Sorry we dont have %v tickets, we only have %v ,either book the available or be unlucky\n", userTickets, remaining_Tickets)

	} else {
		if !isValidName {
			fmt.Printf("The name is too short , please enter more than 2 characters\n")
		}
		if !isValidEmail {
			fmt.Printf("The Email you entered is %v, please check if it is correct\n", email)
		}
		if !isValidTicket {
			if userTickets == 0 {
				fmt.Printf("You entered 0 tickets, please enter a number greater than zero\n")

			} else {
				fmt.Printf("we only have %v , but you have entered %v .Please enter the limit with in range\n", remaining_Tickets, userTickets)
			}

		}

	}
	wg.Wait() // Wait for all goroutines to complete
	//}

}
func greetUser() {
	fmt.Printf("Hello, welcome to %v \n", cruizeName)
	fmt.Printf("We have total of %v tickets ,out of which %v are remaining \n", total_Tickets, remaining_Tickets)
}

func getFirstnames() []string { // input booking and output " " parameters
	// lets store and display only the first names.
	firstNames := []string{}
	for _, booking_firstName := range bookings {
		var names = strings.Fields(booking_firstName["firstName"])
		firstNames = append(firstNames, names[0])
	}
	//fmt.Printf("The first names of people who booked tickets are:  %v\n", firstNames)
	return firstNames

}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var userTickets uint
	var lastName string
	var email string

	fmt.Print("Enter your first name : ")
	fmt.Scan(&firstName) // this function will get the user input. '&' pointer
	fmt.Print("Enter your last name :")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address :")
	fmt.Scan(&email)
	fmt.Print("Enter the number of tickets to be purchased: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets

}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {

	remaining_Tickets = remaining_Tickets - userTickets
	// bookings[0] = firstName + " " + lastName // adding element in a list

	var userDetails = make(map[string]string)

	bookings = append(bookings, userDetails)
	userDetails["firstName"] = firstName
	userDetails["lastName"] = lastName
	userDetails["email"] = email
	userDetails["UserTickets"] = strconv.FormatUint(uint64(userTickets), 10) // we have created a map that holds key as string and value also as string , so if we give usertickets as int it wont be accepting it .
	// in order to overcome this issue we convert it to string format and we use base 10 such that it is a regular human-readable decimal number.
	// to overcoe the above error we can use struct keyword and save any kind of datatypes to a map

	//fmt.Print(userDetails)

	//fmt.Printf("the whole slice is %v \n", bookings)
	//fmt.Printf("The first value is %v \n", bookings[0])
	//fmt.Printf("The slice type is %T \n", bookings)
	//fmt.Printf("the slice lenght is %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets . You will receive a confirmation email shortly at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remaining_Tickets, cruizeName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tikets booked for %v %v", userTickets, firstName, lastName)
	fmt.Print("***************************\n")
	fmt.Printf("sent %v to email %v\n", ticket, email)
	wg.Done() // decrement counter and acknowledge that the thread is done.
}
