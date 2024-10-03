package main

import (
	"fmt"
	"golanglearn/utils"
)

func main() {
	// var number int
	var word string
	fmt.Print("Enter a word: ")
	fmt.Scan(&word)
	utils.Palindarame(word)
	// utils.CheckOddOrEven(number) // Call the function from the utils package
	// utils.Factorial(number)
}
