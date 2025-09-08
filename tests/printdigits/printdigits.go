package main

import "github.com/01-edu/z01"

func main() {
	ascii := '0'
	for range 10 {
		z01.PrintRune(rune(ascii))
		ascii++
	}
	z01.PrintRune('\n')
}
