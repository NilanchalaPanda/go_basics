package main

import (
	"fmt"

	"exmaple.com/struct/user"
)

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userDateOfBirth := getUserData("Enter your date of birth (DD/MM/YYYY): ")

	// var appUser user
	// appuser = user{}

	// Old way of initializing the user struct
	// -- Way 1
	// var appUser = user{
	// 	firstName:   userFirstName,
	// 	lastName:    userLastName,
	// 	dateOfBirth: userDateOfBirth,
	// 	createdAt:   time.Now(),
	// }
	// -- Way 1
	// outputUserDetails(appUser)

	// New way of initializing the user struct
	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userDateOfBirth)

	if err != nil {
		fmt.Println("Error creating user:", err)
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	fmt.Println("User name cleared.")
	appUser.OutputUserDetails() // This will show empty first and last names after clearing
}

// -- Way 1
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
