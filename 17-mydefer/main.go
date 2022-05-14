package main

import "fmt"

func main() {
	defer fmt.Println("Execute Line One....")
	fmt.Println("Execute Line Two....")
	defer fmt.Println("Execute Line Three....")
	fmt.Println("Execute Line Four....")
	fmt.Println("Execute Line Five....")
}
