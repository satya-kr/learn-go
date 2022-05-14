package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcom to UserInput Program")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n\nEnter name : ")

	// comma ok || error ok syntax
	// input, err := reader.ReadString("\n")
	// or
	input, _ := reader.ReadString('\n')
	fmt.Println("Name is : ", input)
}
