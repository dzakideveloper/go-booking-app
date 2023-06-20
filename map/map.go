package helper

import(
	// "strconv"
	"fmt"
)

var firstName string
var lastName string
var userName string
var email string
var userTickets int

var bookings = make([]UserData, 0)

type UserData struct{
	firstName string
	lastName string
	email string
	numberTickets int
}

func main(){
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

	fmt.Print("masukkan nama depan anda:")
	fmt.Scan( & firstName)
	fmt.Print("masukkan nama belakang anda:")
	fmt.Scan( & lastName)
	fmt.Print("masukkan email anda:")
	fmt.Scan( & email)
	fmt.Print("jumlah tiket yang anda beli:")
	fmt.Scan( & userTickets)

	print(bookings,"\n")
}