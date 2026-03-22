package main

/*
Write a program that:
1. Creates a function called `reverse`
2. Takes a string as input
3. Returns the reversed string
4. Tests it with these three words:
   - `"Innocent"`
   - `"Golang"`
   - `"Lagos"`

Expected output:
```
tneconni
gnaloG
soGaL

*/

func reveres(s string) string {
	chars := []rune(s)

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
