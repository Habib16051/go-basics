package main

import "fmt"

func addNumbers(num1 int, num2 int) {
	result := num1 + num2
	fmt.Println("Addition  of two numbers =", result)
}

func main() {
	addNumbers(10, 15)

	func(a, b int) {
		sum := a + b
		fmt.Println("Summation of two numbers using anonymous function =", sum)
	}(10, 15)
}

func init() {
	fmt.Println("Hello, I will be executed first")
}

// An anonymous function is a function without any name
// IIFE - Immediate Invoke Function Expression
