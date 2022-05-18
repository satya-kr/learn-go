package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Wlcm to go mod and Gorila/Mux")
	greeter()

	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/articles/{category}/{id:[0-9]+}", myArtical)

	//starting server
	log.Fatal(http.ListenAndServe(":4000", router))
}

func greeter() {
	fmt.Println("Hello i m a new go lang dev")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Wow go is super easy and handy!</h1>"))
}

func myArtical(w http.ResponseWriter, r *http.Request) {
	reqData := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Article %v ID : %v\n", reqData["category"], reqData["id"])
}
