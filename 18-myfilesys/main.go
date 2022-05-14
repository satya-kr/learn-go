package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to File and System in go")
	content := "this text needs to be store in a file"

	filepath := "./myfile.txt"

	file, err := os.Create(filepath)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	fmt.Println("Length is : ", length)
	defer file.Close()

	readFile(filepath)

}

func readFile(filename string) {
	dataInByte, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	fmt.Println("Text data inside the file is bytes \n", dataInByte)
	fmt.Println("Text data inside the file in string :\n", string(dataInByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
