package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*

## ✏️ Your Task Now

Build a `checkPassword` function that:
1. Checks minimum **8 characters**
2. Checks for **uppercase** letter
3. Checks for **lowercase** letter
4. Checks for **digit**
5. Checks for **special character** `[!@#$%^&*]`
6. Prints each check result
7. Prints final verdict — **Strong** or **Weak**

Test with these passwords:
```
"hello"          → Weak
"Hello1!"        → Strong
"onlylowercase"  → Weak
```

Expected output for `"Hello1!"`:
```
Length OK    : true
Has Upper    : true
Has Lower    : true
Has Digit    : true
Has Special  : true
Verdict      : Strong! ✅

*/

func checkPassword(passWord string) {
	fmt.Printf("\nPassWord: %s\n", passWord)
	hasLength := len(passWord) >= 8
	fmt.Printf("%-12s : %v\n", "Length OK", hasLength)

	hasUpper, _ := regexp.MatchString("[A-Z]+", passWord)
	fmt.Printf("%-12s : %v\n", "Has Upper", hasUpper)

	hasLower, _ := regexp.MatchString("[a-z]+", passWord)
	fmt.Printf("%-12s : %v\n", "Has Lower", hasLower)

	hasDIgit, _ := regexp.MatchString("[0-9]+", passWord)
	fmt.Printf("%-12s : %v\n", "Has Digit", hasDIgit)

	hasSpecial, _ := regexp.MatchString("[!@#$%^&*]", passWord)
	fmt.Printf("%-12s : %v\n", "Has Special", hasSpecial)

	hasNoSpaces := !strings.Contains(passWord, " ")
	fmt.Printf("%-12s : %v\n", "Has No Space", hasNoSpaces)

	result := hasLength && hasUpper && hasLower && hasDIgit && hasSpecial && hasNoSpaces

	if result == true {
		fmt.Printf("%-12s : %s\n", "Verdict", "Strong")
	} else {
		fmt.Printf("%-12s : %s\n", "Verdict", "Weak ❌")
	}
}

func main() {
	checkPassword("Innocent")
	checkPassword("Hello1!y")
	checkPassword("Hello1!")
	checkPassword("onlylowercase")
}
