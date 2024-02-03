package main

import "github.com/01-edu/z01"

func main() {
	var char rune
	for char = 'z'; char >= 'a'; char-- {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
