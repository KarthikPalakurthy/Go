package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {

	fmt.Printf("The file size of KB is %v bytes\n", KB)
	fmt.Printf("The file size of MB is %v bytes\n", MB)
	fmt.Printf("The file size of GB is %v bytes\n", GB)
	fmt.Printf("The file size of TB is %v bytes\n", TB)
	fmt.Printf("The file size of PB is %v bytes\n", PB)
	fmt.Printf("The file size of EB is %v bytes\n", EB)
}
