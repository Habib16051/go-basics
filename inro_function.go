package main

import "fmt"

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println("The summation of number is  =", sum)

}

func main() {
	a := 20
	b := 30
	add(a, b)

}
