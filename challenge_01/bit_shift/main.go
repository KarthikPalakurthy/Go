package main

import "fmt"

var (
	a = 2
)

func main() {

	fmt.Printf("%d \t %b\n", a, a)
	fmt.Printf("%d \t %b\n", a<<1, a<<1)
	fmt.Printf("%d \t %b\n", a<<2, a<<2)
	fmt.Printf("%d \t %b\n", a<<3, a<<3)
	fmt.Printf("%d \t %b\n", a<<4, a<<4)
	fmt.Printf("%d \t %b\n", a<<5, a<<5)
}
