package main

import "fmt"

const (
	a1 = iota // a1 = 0
	a2 = iota // a2 = 1
	a3 = iota // a3 = 2
)

const (
	a4 = iota // a4 = 0
	a5        // a5 = 1
	a6        // a6 = 2
)

func main() {

	fmt.Println(a1, a2, a3)
	fmt.Println(a4, a5, a6)
}
