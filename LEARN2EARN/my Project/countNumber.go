/*
Write a function that uses a loop to count up from 0
and stops when the count reaches 10, then prints the final count.
*/

package main

import "fmt"

func countNbr() {
	count := 0
	for count < 10 {
		count++
	}
	fmt.Println(count)
}
