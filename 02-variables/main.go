package main

import "fmt"

/*
const first letter should be capital which denoted that the variable is PUBLIC
, which can be accessable from any file
*/

const LoginToken string = "dwadijdadijadwa"

var myLang string = "GO Language"

func main() {
	fmt.Println("\nMyLanguage : ", myLang)
	fmt.Printf("\nVariable is type of : %T", myLang)

	var name string = "Satya"
	fmt.Println("\n\nName : ", name)
	fmt.Printf("Variable is type of : %T", name)

	var isLoggedin bool = false
	fmt.Println("\n\nName : ", isLoggedin)
	fmt.Printf("Variable is type of : %T", isLoggedin)

	var smallVal int = 255
	fmt.Println("\n\nName : ", smallVal)
	fmt.Printf("Variable is type of : %T", smallVal)

	var f32 float32 = 255.4516548546
	fmt.Println("\n\nName : ", f32)
	fmt.Printf("Variable is type of : %T", f32)

	var f64 float64 = 255.4516548546
	fmt.Println("\n\nName : ", f64)
	fmt.Printf("Variable is type of : %T", f64)

	//implicity type

	var website = "http://astergo.in"
	fmt.Println("\nWebsite : ", website)
	// website = 10 // its throw an error

	//dynamic variables
	numOfUsers := 1000
	fmt.Println("\nnumOfUsers : ", numOfUsers)

	fmt.Println("\nLoginToken : ", LoginToken)

}
