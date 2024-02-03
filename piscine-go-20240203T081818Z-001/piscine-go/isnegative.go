package main

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	var result rune

	if nb >= 0 {
		result = 'F'
	} else {
		result = 'T'
	}
	z01.PrintRune(result)
	z01.PrintRune('\n')
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}
