package main

import (
	"fmt"
	"golanglearn/utils"
)

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	utils.CheckOddOrEven(number) // Call the function from the utils package
	utils.Factorial(number)
}
