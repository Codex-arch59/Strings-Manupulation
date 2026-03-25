package main

import "fmt"

func ViewBalnce(balance int) {
	if balance >= 3000 {
		balance -= 1000
		fmt.Println("Balance after withdrawals:", balance)
	} else {
		fmt.Println("insufficient funds")
	}
}
