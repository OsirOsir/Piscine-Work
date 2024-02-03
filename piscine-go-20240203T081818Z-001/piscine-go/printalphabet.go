package main

import "fmt"

func printAltenative() {
	for i := 'a'; i <= 'z'; i++ {
		if i%2 == 0 {
			fmt.Print(string(i))
		} else {
			fmt.Print(string(i - ('a' - 'A')))
		}
	}
	fmt.Println('\n')
}

func main() {
	printAltenative()
}
