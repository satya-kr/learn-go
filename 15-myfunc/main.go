package main

import "fmt"

func main() {
	fmt.Println("Welcome to the functions in golang")

	greeting()

	results := addNum(40, 66)
	fmt.Println("Results:", results)

	proRes, myMsg := proAdder(2, 5, 6, 10)
	fmt.Println("Pro results is :", proRes)
	fmt.Println("Pro message is :", myMsg)
}

func addNum(x int, y int) int {
	return x + y
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, v := range values {
		total += v
	}

	return total, "Hi Pro result function"
}

func greeting() {
	fmt.Println("This is a greeting")
}
