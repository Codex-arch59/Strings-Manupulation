package main

import (
	"fmt"
	"strings"
)

/*

## 🔥 Mini Challenge!

This is slightly harder — think carefully!

Write a program that:
1. Takes this string:
```
"  innocent,is,learning,golang,in,lagos  "
```
2. **Trims** the spaces from both sides
3. **Splits** it by `","`
4. **Capitalizes** the first letter of each word
5. **Joins** them back with `" | "`
6. Prints the final result

Expected output:
```
Innocent | Is | Learning | Golang | In | Lagos

*/

func cleaner(s string) string {
	word := strings.Split(strings.TrimSpace(s), ",")

	for i := 0; i < len(word); i++ {
		word[i] = strings.ToUpper(string(word[i][0])) + strings.ToLower(word[i][1:])
	}
	return strings.Join(word, " | ")
}

func main() {
	fmt.Println(cleaner("  innocent,is,learning,golang,in,lagos  "))
}
