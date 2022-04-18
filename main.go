package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)


type Student struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
	Marks   int    `json:"Marks"`
}

type Students []Student

func allStudents(w http.ResponseWriter, r *http.Request) {
     students:=Students{
		 Student{Name :"Riya", Address:"Ram nagar",Marks:80},
		 {Name :"Sham", Address:"ashok nagar",Marks:85},
		 {Name :"Radha", Address:"Raj nagar",Marks:90},
		 {Name :"Rekha", Address:"Rajput colony",Marks:85},
		 {Name :"Jay", Address:"Green socity",Marks:83},
		 {Name :"Mayur", Address:"ashok nagar",Marks:94},
		
	 }

	fmt.Println("Endpoint hit :all student endpoint")
	json.NewEncoder(w).Encode(students)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"HomePage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/",homePage)
    http.HandleFunc("/students",allStudents)
	log.Fatal(http.ListenAndServe(":8084",nil))
}

func main() {
handleRequests()
}
 