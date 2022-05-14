package main

import (
	"fmt"
	"net/url"
)

// const myUrl = "https://smartstudy.astergo.in:443/api/question_sets.php?subject_id=27"
const myUrl = "https://datausa.io/api/data?drilldowns=Nation&measures=Population"

func main() {
	fmt.Println("Wlcm to handling urls in go lang")
	fmt.Println(myUrl)

	//parsing
	res, _ := url.Parse(myUrl) //underscor _ means we ignore the error here

	// fmt.Println("Schema:", res.Scheme)
	// fmt.Println("Host:", res.Host)
	// fmt.Println("Port:", res.Port()) //its not property its function
	// fmt.Println("Path:", res.Path)
	// fmt.Println("Raw Query:", res.RawQuery)
	// fmt.Println("Req URI:", res.RequestURI())

	qparams := res.Query()

	fmt.Printf("Type of Query params are  %T:", qparams)

	fmt.Println("\nValue : ", qparams["drilldowns"])

	for _, v := range qparams {
		fmt.Println("Param is :", v)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Location in memory", &anotherUrl)
	fmt.Println(anotherUrl)
}
