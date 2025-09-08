package main

import "github.com/01-edu/z01"

func main() {
	ascii := 'z'
	for i := 0; i < 26; i++ {
		z01.PrintRune(rune(ascii))
		ascii--
	}
	z01.PrintRune('\n')
}
