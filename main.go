package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tommarler/Beard_Of_knowledge/controllers/page"
	"github.com/tommarler/Beard_Of_knowledge/controllers/propublica"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Hello from main")
	r.HandleFunc("/home", page.HomePage).Methods("GET")
	r.HandleFunc("/contact", page.ContactPage).Methods("GET")
	r.HandleFunc("/propublica/senators", propublica.GetCurrentSenators).Methods("GET")
	r.HandleFunc("/propublica/{id}", propublica.GetSenatorDetails).Methods("GET")
	// http.HandleFunc("/propublica/house", propublica.GetCurrentHouseMembers)
	// http.HandleFunc("/propublica/newMembers", propublica.GetNewMembers)
	log.Fatal(http.ListenAndServe(":8080", r))

	
}








