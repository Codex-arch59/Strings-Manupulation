package main

import (
	"regexp"
)

/*
🚀 Phase 4 Begins — Pattern Matching!
Welcome to Advanced skills Innocent!
First lesson — checking patterns inside strings using regexp package!
Regular expressions let us search for patterns not just exact words!


✏️ Your Task Now
Write a program that checks:

Is "Innocent123" all letters only? → false
Is "Innocent" all letters only? → true
Is "12345" all digits only? → true
Is "hello@gmail.com" a valid email pattern? → true


Pattern symbols:
```
^        = start of string
$        = end of string
[a-z]    = any lowercase letter
[A-Z]    = any uppercase letter
[0-9]    = any digit
+        = one or more
*        = zero or more
?        = zero or one
\d       = any digit
\w       = any word character
.        = any character

*/

func re(s string) bool {

	matched, _ := regexp.MatchString("^[a-zA-Z]+$", s)
	return matched
}
func reAlnum(s string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]+$", s)
	return matched
}

func regNum(s string) bool {
	matched, _ := regexp.MatchString("^[0-9]+$", s)
	return matched
}

func reEmail(s string) bool {
	matched, _ := regexp.MatchString("^[a-z0-9]+@[a-z]+\\.[a-z]+$", s)
	return matched
}

/*func main() {
	fmt.Println(re("Innocent"))
	fmt.Println(reAlnum("INnocent123"))
	fmt.Println(regNum("11234"))
	fmt.Println(reEmail("test@gmail.com"))
	fmt.Println(reEmail("test@gmailcom"))
}
*/
