package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the User Input Programm")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n\nGive us some ratings")

	input, _ := reader.ReadString('\n')
	fmt.Println("Your rating is: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to you rating: ", numRating+1)
	}
}
