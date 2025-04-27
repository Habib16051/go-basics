package main

import "fmt"

var a = 10

func main() {
	fmt.Println(a)
}

func init() {
	fmt.Println("initial result :", a)
	a = 25

}
