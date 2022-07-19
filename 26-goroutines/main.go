package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //useuly these are ointers
var mut sync.Mutex    //useuly these are ointers

func main() {
	fmt.Println("Welcome to GO Concurrency and go")

	websites := []string{
		"https://lco.dev",
		"https://astergo.in",
		"https://google.com",
		"https://dealzup.in",
		"http://symic.in",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1) //its something like status which is 1 and after all done is change it to 0 and befor call wait
	}

	wg.Wait()
	fmt.Println(signals)
	// go greteer("Hello")
	// greteer("World")
}

// func greteer(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("\n%d status code for %s", res.StatusCode, endpoint)
	}
}
