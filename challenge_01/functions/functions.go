package main

import "fmt"

/*
func add(x int, y int) int {
	return x + y
}
*/

// The above add function ca also be wrtten as follows as it shares the same data type.

func add(x, y int) int {
	return x + y
}

func main() {

	fmt.Println(add(58678, 96))
	sayHello()
}

func sayHello() {
	fmt.Println("Hello Gophers!")
}
