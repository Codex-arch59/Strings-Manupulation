/*
Write a function that loops through numbers 1 to 50,
adds only the numbers that are divisible by 5, and prints the total.
*/

package main

import "fmt"

func modlus5() {
	sum := 0

	for i := 1; i <= 50; i++ {
		if i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
