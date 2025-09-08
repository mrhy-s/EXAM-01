package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

func CountAlpha(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) {
			count++
		}
	}
	return count
}
