package main

import "fmt"

func main() {
	fmt.Println("Welcome to the loop program")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	fmt.Println("days: ", days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for i, v := range days {
	// 	fmt.Printf("index is %v and value is %v\n", i+1, v)
	// }
	rougueValue := 1
	for rougueValue < 10 {

		// if rougueValue == 5 {
		// 	break
		// }

		if rougueValue == 2 {
			goto astergo
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is :", rougueValue)
		rougueValue++
	}

astergo:
	fmt.Println("jumped to the right place")
}
