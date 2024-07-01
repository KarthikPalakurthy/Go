package main

import (
	"fmt"
	"math/rand"
)

func main() {

	/*
		We used a new package called the math and rand package here.

		Also remember, the scope of the random variable generated will be 0-99. If you also want to include 100, update the number to 101.

	*/

	fmt.Println("The number I generated is ", rand.Intn(100))
	fmt.Println("The number I generated is ", rand.Intn(100))
	fmt.Println("The number I generated is ", rand.Intn(100))
	fmt.Println("The number I generated is ", rand.Intn(100))
	fmt.Println("The number I generated is ", rand.Intn(100))
	fmt.Println("The number I generated is ", rand.Intn(100))
	fmt.Println("The number I generated is ", rand.Intn(100))
	fmt.Println("The number I generated is ", rand.Intn(100))
}
