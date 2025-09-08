package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

func CountAlpha(s string) int {
	count := 0
	for _, char := range s {
		if IsAlpha(char) {
			count++
		}
	}
	return count
}

func IsAlpha(s rune) bool {
	if (s < 'a' || s > 'z') && (s < '0' || s > '9') && (s < 'A' || s > 'Z') {
		return false
	}
	return true
}
