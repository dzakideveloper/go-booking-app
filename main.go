package main

import (
	// "strconv"
	"fmt"
	"sync"
	"time"	
)

var firstName string
var lastName string
var userName string
var email string
var userTickets int

var wg = sync.WaitGroup{}

var bookings = make([]UserData, 0)

type UserData struct{
	firstName string
	lastName string
	email string
	numberTickets int
}

func main(){
		fmt.Print("masukkan nama depan anda:")
		fmt.Scan( & firstName)
		fmt.Print("masukkan nama belakang anda:")
		fmt.Scan( & lastName)
		fmt.Print("masukkan email anda:")
		fmt.Scan( & email)
		fmt.Print("jumlah tiket yang anda beli:")
		fmt.Scan( & userTickets)

		wg.Add(1)
		go sendEmail()

		var userData = UserData{
			firstName: firstName,
			lastName: lastName,
			email: email,
			numberTickets: userTickets,
		}

		// var userData = make(map[string]string)
		// userData["firstName"] = firstName
		// userData["lastName"] = lastName
		// userData["email"] = email
		// userData["numberTickets"] = strconv.FormatInt(int64(userTickets), 10)

		// userName = firstName + " " + lastName
		// bookings = append(bookings, userName)

		bookings = append(bookings, userData)

		// var bookings = append(bookings, userData)

		for _, booking := range bookings {
			fmt.Println(booking)
		}
		wg.Wait()

	// print(bookings)
}

func sendEmail(){
	time.Sleep(10 * time.Second)
	fmt.Println(firstName, lastName, "DONE")
	wg.Done()
}

// import (
// 	"fmt"
// 	"strings"
// 	"strconv"
// )

// var slice = make([]map[string]string, 0)

// func main() {
// 	for{
// 		firstName, lastName, email, userName, userTickets, remainingTickets := input()
// 		greet(firstName)
// 		mauApaEngga()

// 		var userData = make(map[string]string)
// 		userData["firstName"] = firstName
// 		userData["lastName"] = lastName
// 		userData["email"] = email
// 		userData["numberTickets"] = strconv.FormatInt(int64(userTickets), 10)

// 		bookings := append(slice, userData)
// 		print(bookings,"\n")

// 		fmt.Print("jumlah tiket yang anda beli:")
// 		fmt.Scan( & userTickets)

// 		// validateNama, validateEmail, validateTicketNumber := validate(firstName, lastName, email, userTickets, remainingTickets)

// 		// mainFunc(validateEmail, validateNama, validateTicketNumber, userName, firstName, lastName, email, userTickets, remainingTickets)

// 		cityCase()
// 	}
// }

// func greet(nama string){
// 	fmt.Printf("halo %v\n", nama)
// }

// func getFirstName(slice []string) []string{
// 	firstNames := []string{}
// 	for _, apa := range slice {
// 		var nama = strings.Fields(apa)
// 		firstNames = append(firstNames, nama[0])
// 	}
// 	return firstNames
// }

// func validate(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool){
// 	validateNama := len(firstName) >= 2 && len(lastName) >= 2
// 	validateEmail := strings.Contains(email, "@")
// 	validateTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)
// 	return validateNama, validateEmail, validateTicketNumber
// }

// func input() (string, string, string, string, int, uint){
// 	var firstName string
// 	var lastName string
// 	var userName string
// 	var email string
// 	var userTickets int = 50
// 	var remainingTickets uint = 50

// 	fmt.Print("masukkan nama depan anda:")
// 	fmt.Scan( & firstName)
// 	fmt.Print("masukkan nama belakang anda:")
// 	fmt.Scan( & lastName)
// 	fmt.Print("masukkan email anda:")
// 	fmt.Scan( & email)
// 	return firstName, lastName, userName, email, userTickets, remainingTickets
// }

// func mauApaEngga(){
// 	var mau bool

// 	fmt.Print("mau beli?")
// 	fmt.Scan( & mau)

// 	if mau == false {
// 		fmt.Print("yakin gamau beli?")
// 		fmt.Scan( & mau)
// 		if mau == true {
// 			print("apa coba!!")
// 			return
// 		}
// 	}
// }

// func cityCase(){
// 	city := "makassar"

// 	switch city{
// 		case "jakarta", "depok":
// 			fmt.Println("jakarta", "depok")
// 		case "makassar":
// 			fmt.Println("makassar")
// 		default:
// 			fmt.Println("apasih")
// 	}
// }

// func mainFunc(validateEmail bool, validateNama bool, validateTicketNumber bool, userName string, firstName string, lastName string, email string, userTickets int, remainingTickets uint){
// 	if  validateNama && validateTicketNumber {
// 		userName = firstName + " " + lastName
// 		remainingTickets = remainingTickets - uint(userTickets)
// 		slice = append(slice, userName)

// 		firstNames := getFirstName(slice)
// 		fmt.Printf("ini adalah nama depan:%v\n", firstNames)

// 		if remainingTickets == 0 {
// 			print("SOLD OUT")
// 			return
// 		}

// 		fmt.Printf("hello %v, goodmorning %v, have a nice day %v\n", firstName, firstName, firstName)
// 		fmt.Printf("tiket tersisa tinggal %v \n", remainingTickets)
// 		fmt.Println(userName)
// 	} else {
// 		if !validateNama{
// 			fmt.Println("your name is too short")
// 		}
// 		if !validateTicketNumber{
// 			fmt.Printf("We only have %v tickets\n", remainingTickets)
// 		}
// 	}
// }