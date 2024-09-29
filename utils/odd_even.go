package utils // Define a package for utility functions

import "fmt"

// Function to check if a number is odd or even
func CheckOddOrEven(a int) {
	if a%2 == 0 {
		fmt.Println("It is an even number")
	} else {
		fmt.Println("It is an odd number")
	}
}
