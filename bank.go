package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func readBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	// 0644 is the permission mode for the file to READ & WRITE by the owner, and READ by others.
}

func main() {
	var accountBalance float64 = readBalanceFromFile()

	fmt.Println("Welcome to the Bank Application!")

	for {
		fmt.Println("What do you want to do today?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your current balance is:", accountBalance)

		case 2:
			fmt.Print("Enter the amount to deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Deposit successful!")
			writeBalanceToFile(accountBalance)

		case 3:
			fmt.Print("Enter the amount to withdraw:")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Withdrawal amount must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds! You cannot withdraw more than your current balance.")
				continue
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Withdrawn successful")
				writeBalanceToFile(accountBalance)
			}

		default:
			fmt.Print("GoodBye!")
			return
		}
	}
}

/*
	if choice == 1 {
		fmt.Println("Your current balance is: $", accountBalance)
	} else if choice == 2 {
		fmt.Print("Enter the amount to deposit: $")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Deposit amount must be greater than 0")
			continue
		}

		accountBalance += depositAmount
		fmt.Println("Deposit successful!")
	} else if choice == 3 {
		fmt.Print("Enter the amount to withdraw: $")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Withdrawal amount must be greater than 0")
			continue
		}

		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds! You cannot withdraw more than your current balance.")
			continue
		} else {
			accountBalance -= withdrawAmount
			fmt.Println("Withdrawn successful")
		}
	} else fmt.Print("GoodBye!") break
*/
