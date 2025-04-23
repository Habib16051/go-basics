package main

import (
	"example.com/mathlib"
	"fmt"
)

func main() {

	x := 20
	fmt.Println("Result =", x) // Print a message

	// Package scope
	// Tries to run:
	// go init mod example.com
	// Now we are using our own custom package
	fmt.Println("Our Custom Package :")
	resultOfMultiplication := mathlib.Multiplication(10, 15)
	fmt.Println("Result =", resultOfMultiplication)

}
