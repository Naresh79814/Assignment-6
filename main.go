package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Emp struct {
	Id        string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Age       int    `json:"age,omitempty"`
}

var emps []Emp

func GetEmps(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(emps)
}

func GetEmp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, emp := range emps {
		if emp.Id == params["id"] {
			json.NewEncoder(w).Encode(emp)
			return
		}
	}
	json.NewEncoder(w).Encode(&Emp{})
}

func CreateEmp(w http.ResponseWriter, r *http.Request) {
	var emp Emp
	_ = json.NewDecoder(r.Body).Decode(&emp)
	emps = append(emps, emp)
	json.NewEncoder(w).Encode(emp)
}

func DeleteEmp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, emp := range emps {
		if emp.Id == params["id"] {
			emps = append(emps[:index], emps[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(emps)
	}
}

func main() {
	router := mux.NewRouter()
	emps = append(emps, Emp{"E-1", "Anupam", "Raj", 20})
	emps = append(emps, Emp{"E-2", "Rahul", "Anand", 30})

	router.HandleFunc("/emp", GetEmps).Methods("GET")
	router.HandleFunc("/emp/{id}", GetEmp).Methods("GET")
	router.HandleFunc("/emp", CreateEmp).Methods("POST")
	router.HandleFunc("/emp/{id}", DeleteEmp).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}