/*
Write a function that adds all numbers
from 1 to 100 together and prints the total.

*/

package main

import "fmt"

func accumulator() {
	total := 0
	for number := 1; number <= 100; number++ {
		total += number
	}
	fmt.Println(total)
}
