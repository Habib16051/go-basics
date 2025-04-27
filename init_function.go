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

//In Go, the init function is a special function that gets called automatically when a package
//is initialized—before the main function runs (in the main package) or before the package is used (in imported packages).
//In Go, global variable initialization happens before the init() function is called
