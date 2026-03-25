package main

import "fmt"

func trasactionSystem() {
	accountBalance := 10000

	var (
		button           int
		withdrawalAmount int
		depositAmout     int
	)
	fmt.Println("Welcome to Ekwueme's Banking System")

	for {

		fmt.Println("Enter -1: to withdraw, Enter 2: to top-up Account Balance, Enter 3: to check Balance, Enter 0 to exit.")
		fmt.Scan(&button)

		if button != -1 && button != 2 && button != 0 && button != 3 {
			fmt.Println("Error, you have entered an invalid number pls try again")
			continue
		}

		if button == 0 {
			fmt.Printf("Thank you for banking with us, Your account balance is: %d\n", accountBalance)
			break
		}

		if button == 3 {
			fmt.Printf("Your Acoount Balance is %d \n", accountBalance)
		}

		if button == -1 {
			fmt.Println("Enter withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount < 0 {
				fmt.Println("Error, withdrwal amount cannot be negative, pls try again")
				continue
			}
			projectedBalance := accountBalance - withdrawalAmount
			if projectedBalance < 1000 {
				fmt.Println("Sorry, you cannot go below minimum balance of 1000")
			} else {
				accountBalance -= withdrawalAmount
				fmt.Printf("Withdrawal successfull. Your new balance is now; %d\n", accountBalance)
			}
			if projectedBalance == 1000 {
				fmt.Println("You have reach your minimum balance pls top-up.")
			}
		}
		if button == 2 {
			fmt.Println("Enter deposite amout")
			fmt.Scan(&depositAmout)

			if depositAmout < 0 {
				fmt.Println("Error, you have entered and invalid amount, Please try again.")
				continue
			} else {
				accountBalance += depositAmout
				fmt.Printf("deposoit successfull, your new balance is now %d \n", accountBalance)
			}
		}
	}
}


func main() {
	trasactionSystem()
}