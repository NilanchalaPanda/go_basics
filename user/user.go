package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	dateOfBirth string
	createdAt   time.Time
}

func (u User) OutputUserDetails() {
	// fmt.Println("User Details:")
	// fmt.Println("First Name:", u.firstName)
	// fmt.Println("Last Name:", u.lastName)
	// fmt.Println("Date of Birth:", u.dateOfBirth)
	// fmt.Println("Account Created At:", u.createdAt.Format("02/01/2006 15:04:05"))
	fmt.Println(u.firstName, u.lastName, u.dateOfBirth)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, dateOfBirth string) (*User, error) {
	if firstName == "" || lastName == "" || dateOfBirth == "" {
		return nil, errors.New("ALL FIELDS ARE REQUIRED")
	}

	return &User{
		firstName:   firstName,
		lastName:    lastName,
		dateOfBirth: dateOfBirth,
		createdAt:   time.Now(),
	}, nil
}
