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

func Palindarame(s string) {
	var p string = ""
	for i := len(s) - 1; i >= 0; i-- {
		p += string(s[i])
		// fmt.Printf("%c\n", s[i])

	}
	if p == s {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}

}

func Findlag(a [5]int) {
	var maxi int = int(a[0])
	for i := 0; i <= len(a)-1; i++ {
		if int(a[i]) > maxi {
			maxi = int(a[i])
		}
	}
	fmt.Println(maxi)
}
