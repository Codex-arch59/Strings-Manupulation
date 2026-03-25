package main

func checkAge(age int) string {
	if age <= 0 {
		return "Invalid age"
	} else if age <= 18 {
		return " Minor"
	} else {
		return "Adult"
	}
}
