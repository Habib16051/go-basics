package main

import "fmt"

func twoOperations(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2

	return sum, mul

}

func main() {
	a := 20
	b := 30
	sum, mul := twoOperations(a, b)
	fmt.Println("Summation of two values =", sum)
	fmt.Println("Multiplication of two values =", mul)
	//fmt.Println(twoOperations(10, 20))
}
