package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/satya-kr/apimongo/router"
)

func main() {
	fmt.Println("Mongodb API")
	fmt.Println("Server is getting started ...")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server listening at port 4000 ...")
}
