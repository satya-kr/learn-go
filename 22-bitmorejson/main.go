package main

import (
	"encoding/json"
	"fmt"
)

//define structure with name course and first letter is lower obviously case to make this private
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`              // hide in json
	Tags     []string `json:"tags,omitempty"` // empty when nil
}

func main() {
	fmt.Println("Wlcm more about JSON")
	// EncodeJson()
	// data := DecodeJson()
	// fmt.Println(data)
	// fmt.Println("Course Name : ", data.Name)
	// fmt.Println("Price : ", data.Price)
	DecodeJson()
}

func EncodeJson() {

	//we specify type course
	myCourses := []course{
		{"React JS Bootcamp", 499, "https://astergo.in", "00000", []string{"web-dev", "font-end", "javascript"}},
		{"Go with Go Lang", 1499, "https://astergo.in", "00000", []string{"web-dev", "back-end", "api"}},
		{"Laravel Bootcamp", 2499, "https://astergo.in", "00000", nil},
	}

	//package this data as a JSON data

	// finalJson, err := json.Marshal(myCourses)
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

// func DecodeJson() course {
func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "React JS Bootcamp",
			"Price": 499,
			"Platform": "https://astergo.in",
			"tags": ["web-dev", "font-end", "javascript" ]
		}
	`)
	var myCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Valid JSON format")
		//now decode the json
		json.Unmarshal(jsonDataFromWeb, &myCourse) // we pass referance of course as myCourse because we don't want any copy of myCourse
		fmt.Printf("%#v\n", myCourse)
		fmt.Println("-------------------------------------------")
		// fmt.Println("Name : ", myCourse.Name)
		// fmt.Println("Price : ", myCourse.Price)
		// fmt.Println("Platform : ", myCourse.Platform)
		// fmt.Println("Tags : ", myCourse.Tags)

	} else {
		fmt.Println("Invalid JSON Format")
	}
	// return myCourse

	var myOnlineData map[string]interface{} // use interface because of we don't know the vale,
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("\nKey is %v and value is %v and type is %T:", k, v, v)
	}
}
