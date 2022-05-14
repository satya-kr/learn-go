package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Learn Web Request")

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response is %T\n", res)

	defer res.Body.Close()

	dataInBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	content := string(dataInBytes)
	fmt.Println("Response body after parsing :", content)
}
