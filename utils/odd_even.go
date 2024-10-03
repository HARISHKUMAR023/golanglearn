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

func Factorial(a int) {
	var ans int = 1
	for i := 1; i <= a; i++ {
		ans = ans * i
	}
	fmt.Println(ans)
}
