package main

import(
	"fmt")


func euclides(a, b int) int {
	//If the numbers are negative, we convert them to positive before calculating the MCD
	//The for loop continues as long as b is not zero. At each iteration, the value of b becomes the remainder of a/b, and a takes the previous value of b.
	//Return MCD
	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func main(){

	var num1, num2 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	result := euclides(num1, num2)

	fmt.Printf("MCD ~ %d and %d: %d\n", num1, num2, result)
}
