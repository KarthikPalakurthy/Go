// package main

// import (
// 	"fmt"
// 	"path"
// )

// func main() {

// 	_, file := path.Split("google.com/learning.go")

// 	fmt.Println("The name of the file is:", file)
// }

// // func main() {

// // 	var file string
// // 	// dir, file = path.Split("css/main.css")
// // 	//Over here, I declared an underscore and it helps me to get the value of the file, ignoring the directory
// // 	_, file = path.Split("css/main.css")

// // 	// fmt.Println("dir:", dir)
// // 	fmt.Println("path:", file)
// // }

package main

import "fmt"

func main() {

	// normal variable declaration

	var firstName, lastName string

	firstName, lastName = "Karthik", "Sharma"

	fmt.Println("The name is:", firstName, lastName)

	// Short Variable declaration

	finance := 350

	fmt.Println("My salary per month is:", finance*4)
}
