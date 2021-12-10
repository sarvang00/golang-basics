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

// Model for Course - file
type Course struct {
	CourseId    string  `json:"id"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return ((c.CourseId == "") && (c.CourseName == ""))
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - LCO")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "React",
		CoursePrice: 299, Author: &Author{FullName: "Sarvang Dave",
			Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack",
		CoursePrice: 199, Author: &Author{FullName: "Sarvang Dave",
			Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting one course")
	w.Header().Set("Content-Type", "application/json")

	// Get id from request
	params := mux.Vars(r)

	// loop through Courses, find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one course")
	w.Header().Set("Content-Type", "application/json")

	// What if enire body of response is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	// what if data sent is empty {}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data in json")
		return
	}

	// Check if the course already exists
	for _, existingCourse := range courses {
		if existingCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already exists in DB")
			return
		}
	}

	// generate unique id, string
	// append new course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with my id
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
	// TODO: Send a response when id is not found
	json.NewEncoder(w).Encode("Course not found!")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	// loop, id, remove on on id match
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: Send a confirm or deny response
			json.NewEncoder(w).Encode("Course deleted!")
			return
		}
	}
	// Send a response when id is not found
	json.NewEncoder(w).Encode("Course not found!")
	return
}

// TODO: Delete all courses
func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete all courses")
	w.Header().Set("Content-Type", "application/json")

	// loop, remove
	for index := range courses {
		courses = append(courses[:index], courses[index+1:]...)
	}

	json.NewEncoder(w).Encode("Courses deleted!")
	return
}
