package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName   string
	lastName    string
	dateOfBirth string
	createdAt   time.Time
}

func (u user) outputUserDetails() {
	// fmt.Println("User Details:")
	// fmt.Println("First Name:", u.firstName)
	// fmt.Println("Last Name:", u.lastName)
	// fmt.Println("Date of Birth:", u.dateOfBirth)
	// fmt.Println("Account Created At:", u.createdAt.Format("02/01/2006 15:04:05"))
	fmt.Println(u.firstName, u.lastName, u.dateOfBirth)
}

// Mutation Methods - You should pass the pointer address and not just 'u user'. That will create a new copy of user.
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userDateOfBirth := getUserData("Enter your date of birth (DD/MM/YYYY): ")

	// var appUser user
	// appuser = user{}

	var appUser = user{
		firstName:   userFirstName,
		lastName:    userLastName,
		dateOfBirth: userDateOfBirth,
		createdAt:   time.Now(),
	}

	// outputUserDetails(appUser)

	appUser.outputUserDetails()
	appUser.clearUserName()
	fmt.Println("User name cleared.")
	appUser.outputUserDetails() // This will show empty first and last names after clearing
}

// func outputUserDetails(u user) {
// 	fmt.Println("User Details:")
// 	fmt.Println("First Name:", u.firstName)
// 	fmt.Println("Last Name:", u.lastName)
// 	fmt.Println("Date of Birth:", u.dateOfBirth)
// 	fmt.Println("Account Created At:", u.createdAt.Format("02/01/2006 15:04:05"))
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
