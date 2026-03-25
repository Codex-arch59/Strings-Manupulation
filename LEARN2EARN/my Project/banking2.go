package main

import "fmt"

func bankingSystem() {
	accountBalance := 10000
	var userInput int

	for {
		fmt.Println("welcome to Globe Banking System, please note: negative = withdrawal, positive = deposit & 0 = exit app.")
		fmt.Scan(&userInput)

		if userInput == 0 {
			fmt.Println("Thank you for banking with us.")
			break
		}
		if userInput > 0 {
			accountBalance += userInput
			fmt.Printf("deposit of %d is succesfull, your new balance is now: %d\n", userInput, accountBalance)
		}
		if userInput < 0 {
			converter := -userInput
			projectedBalance := accountBalance - converter

			if projectedBalance < 1000 {
				fmt.Println("You cannot go below minimum withdrwal amount of 1000")
			} else {
				accountBalance = projectedBalance
				fmt.Printf("withdrawal successfull, you new balane is now %d\n", accountBalance)
			}

		}
	}
}

func main() {
	bankingSystem()
}
