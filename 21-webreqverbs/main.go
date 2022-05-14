package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Wlcm to web verb")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "https://smartstudy.astergo.in/api/question_sets.php?subject_id=27"

	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Status Code", res.StatusCode)
	fmt.Println("Content Length", res.ContentLength)

	var resString strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("Content : ", string(content))

	byteCount, _ := resString.Write(content)
	fmt.Println("Byte Count: ", byteCount)
	fmt.Println("Content : ", resString.String())
}

func PerformPostRequest() {
	const myPostUrl = "http://144.24.114.183:3000/post/save"
	//fake json payload

	reqBody := strings.NewReader(`
		{
			"PostId": 1015,
			"Title": "learn go language",
			"Description": "Now i m learning go languaga and its quite easy.",
			"PostDate": "13-05-2022"
		}
	`)

	response, err := http.Post(myPostUrl, "application/json", reqBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Content :", string(content))
}

func PerformPostFormRequest() {
	const formPotUrl = "http://144.24.114.183:3000/post/save"

	//formData

	data := url.Values{}
	data.Add("PostId", "1055")
	data.Add("Title", "Go Lang Form Data")
	data.Add("Description", "Learn Form Data post in go language")
	data.Add("PostDate", "13-05-2022")

	response, err := http.PostForm(formPotUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content :", string(content))
}
