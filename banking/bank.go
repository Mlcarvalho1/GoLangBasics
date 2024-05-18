package main

import (
	"example/banking/files"
	"fmt"
)

const balanceFile = "balance.txt"

func main() {
	var choice int
	var running bool = true
	var balance float64 = files.GetFloatFromFile(balanceFile)

	fmt.Println("Welcome to the bank!")

	for running {
		fmt.Println("Choose an option:")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is: %v \n", balance)
		case 2:
			var deposit float64

			fmt.Printf("How much do you want to deposit? ")
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Println("The value inserted is invalid. Please insert a value greater than 0.")
				break
			}

			balance += deposit

			files.WriteFloatInFile(balanceFile, balance)
			fmt.Printf("Your new balance is: %v \n", balance)
		case 3:
			var withdraw float64
			fmt.Printf("How much do you want to withdraw? ")

			fmt.Scan(&withdraw)

			if withdraw <= 0 {
				fmt.Println("The value inserted is invalid. Please insert a value greater than 0.")
				break
			}

			if withdraw > balance {
				fmt.Println("You don't have enough money to withdraw this amount.")
				break
			}

			balance -= withdraw

			files.WriteFloatInFile(balanceFile, balance)
			fmt.Printf("Your new balance is: %v \n", balance)
		case 4:
			running = false
		}
	}
}
