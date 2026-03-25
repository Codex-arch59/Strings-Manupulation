/*
Write a function that finds the first number between 1 and 200 that is divisible by
both 5 and 7, stops searching immediately once found, and prints that number.
*/
package main

import "fmt"

func hard() {
	finder := 0

	for i := 1; i <= 200; i++ {
		if i%5 == 0 && i%7 == 0 {
			finder = i
			break
		}
	}
	fmt.Println(finder)
}
