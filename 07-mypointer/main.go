package main

import "fmt"

func main() {
	fmt.Println("Learning Pointer...")

	//create a pointer
	var myPtr *int
	fmt.Println("\n\nValue of pointer : ", myPtr)

	//create a pointer to existing variable

	mynumber := 24
	var nPtr = &mynumber

	fmt.Println("Address of pointer : ", nPtr)
	fmt.Println("Actual value of pointer : ", *nPtr)

	*nPtr = *nPtr * 2
	fmt.Println("Value of pointer : ", mynumber)
}
