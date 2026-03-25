package main

import (
	"fmt"
)

func UserWithdrawals() {
	accountBalance := 5000
	var withdrawalAmount int
	var button int
	var depositAmount int

	for {

		fmt.Println("Welcome to Ekwueme Bank for Africa.")
		fmt.Println("Enter 1: to view balance, Enter 2: top-up your account, Enter 0: to exit, Enter 9: to withdraw cash!!")

		fmt.Scan(&button)

		if button != 1 && button != 0 && button != 9 && button != 2 {
			fmt.Println("invalid code, pls try again!!")
			continue
		}

		if button == 1 {
			fmt.Printf("Your account balance is: %d\n", accountBalance)
			continue
		}

		if button == 0 {
			fmt.Printf("Thank you for banking with us, Your current balance is: %d\n", accountBalance)
			break
		}
		if button == 2 {
			fmt.Println("Please Enter deposit amount:")
			fmt.Scan(&depositAmount)

			if depositAmount > 0 {
				accountBalance += depositAmount
				fmt.Printf("Deposit successfull your new account balance is now: %d \n", accountBalance)
			} else {
				fmt.Println("Error, Enter a valid deposite Amount.")
				continue
			}
		}

		if button == 9 {
			fmt.Println("Pls enter Withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount > 0 && withdrawalAmount <= accountBalance {
				accountBalance -= withdrawalAmount
				fmt.Printf("withdrawal successfull, your new balance is now: %d \n", accountBalance)
			} else {
				fmt.Printf("Your account balance is: %d please recharge \n", accountBalance)
			}
			if accountBalance <= 0 {
				fmt.Printf("Your account balance is now %d please recharge \n", accountBalance)
				continue
			}
		}
	}
}


func main() {
	UserWithdrawals()
}