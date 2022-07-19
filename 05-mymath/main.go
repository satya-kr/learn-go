package main

import (
	"fmt"
	// "math/rand"
	"crypto/rand"
	"math/big"
	// "time"
)

func main() {
	fmt.Println("Ramdom, Crypto and Math in go")

	// var numOne int = 5
	// var numTwo float64 = 4.5
	// fmt.Println("Som of these:", numOne+int(numTwo))

	//random number
	// rand.Seed(700) //we use seed here but the problem is that we alwasy get same rand number everytime
	//so we use time as seed value
	// rand.Seed(time.Now().UnixNano()) //now its give us true rand number
	// fmt.Println(rand.Intn(5))

	//random from crypto

	myRandNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandNum)
}
