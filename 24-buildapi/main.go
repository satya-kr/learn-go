package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for courses - go inside saprate file
type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"price"`
	Author      *Author
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB
var courses []Course

//middleware, helper - go inside saprate file

func (c *Course) IsEmpty() bool {
	return c.CourseName == "" && c.CoursePrice == 0
}

func main() {
	fmt.Println("Hello, Welcome to go lang API")

	router := mux.NewRouter()

	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "React JS",
		CoursePrice: 500,
		Author: &Author{
			Fullname: "Satyajit Kumar",
			Website:  "https://astergo.in",
		},
	})

	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Laravel",
		CoursePrice: 2999,
		Author: &Author{
			Fullname: "Satyajit Kumar",
			Website:  "https://astergo.in",
		},
	})

	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourse).Methods("GET")
	router.HandleFunc("/course/{id}", getSingleCourse).Methods("GET")
	router.HandleFunc("/course", createCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteSingleCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", router))
}

//controllers - file

//server home
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, Welcome to go lang API</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Single Courses")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// loop througr the course and match the id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found with this id")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Single Courses")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please fill some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside JSON")
		return
	}

	//generate unique id
	rand.Seed(time.Now().UnixNano())               //make new seed of rand
	course.CourseId = strconv.Itoa(rand.Intn(100)) // convert int to string
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Single Courses")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO : send a response when id not found
}

func deleteSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Single Courses")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted successfully")
			break
		}
	}
	//TODO : send a response when id not found
}
