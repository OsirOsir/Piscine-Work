package main

import "github.com/01-edu/z01"

func main() {
	for num := '0'; num <= '9'; num++ {
		z01.PrintRune(rune(0 + num))
	}
	z01.PrintRune('\n')
}
