package main //Everything will be handled by this package

import "fmt" // Importing fmt package for formatted I/O operations

func welcomeMessage() { // Function to display welcome message
	fmt.Println("Welcome to Go Programming")
}

func getName() string { // Function to get user's name
	var name string
	fmt.Println("Enter your name")
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int, int) { // Function to get two numbers from user
	var num1, num2 int
	fmt.Println("Enter First Number :")
	fmt.Scanln(&num1)
	fmt.Println("Enter Second Number :")
	fmt.Scanln(&num2)
	return num1, num2
}

func displayResult(num1 int, num2 int) int { // Function to display the result
	sum := num1 + num2
	return sum
}

func goodByeMessage() { // Function to display goodbye message
	fmt.Println("Thank you for using Go Programming")
	fmt.Println("Goodbye!")
}

func main() { // Main function to execute the program
	welcomeMessage()
	name := getName()
	num1, num2 := getTwoNumbers()
	sum := displayResult(num1, num2)
	fmt.Printf("Hello %s, sum of numbers is: %d\n", name, sum)
	goodByeMessage()
}
