package main

import "fmt"

func main() {

	// A variable can be initialised as follows:

	var num1 int = 23                     //Normal Declaration
	num2, num3, num4, _ := 24, 25, 26, 27 //short declaration where compilier identifies the type of vaiable and we are also declaring a blank variable.

	var num5 int //declaring a zero-value variable

	fmt.Printf("The values of all the numbers are %d, %d, %d and %d\n", num1, num2, num3, num4)
	fmt.Printf("The sum of all the numbers are %d\n", num1+num2+num3+num4)
	fmt.Printf("The product of all the numbers are %d\n", num1*num2*num3*num4)

	num5 = 28

	fmt.Printf("The values of all the numbers are %d, %d, %d,%d and %d\n", num1, num2, num3, num4, num5)
	fmt.Printf("The sum of all the numbers are %d\n", num1+num2+num3+num4+num5)
	fmt.Printf("The product of all the numbers are %d\n", num1*num2*num3*num4*num5)
}
