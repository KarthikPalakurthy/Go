package main

import "fmt"

func main() {

	a := 27   // Stores as an int
	b := 27.0 // Stores as a float64

	fmt.Printf("%v is of type %T\n", a, a)
	fmt.Printf("%v is of type %T\n", b, b)
	//For example, if you explicity mention the variable type

	var c float32 = 27.543

	/* This part of the code will throw an error
	b = c

	fmt.Printf("%v is of type %T", b, b) // This throws an error as b is declared as a float64 variable inititally.
	*/

	// We can convert the data as follows:

	b = float64(c)
	fmt.Printf("%v is of type %T\n", b, b)

}
