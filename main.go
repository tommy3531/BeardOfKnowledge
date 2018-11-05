package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/tommarler/Beard_Of_knowledge/controllers/page"
	"github.com/tommarler/Beard_Of_knowledge/controllers/propublica"
)

func main() {
	r := mux.NewRouter()
	propublicaNew := mux.NewRouter()
	r.HandleFunc("/home", page.HomePage).Methods("GET")
	r.HandleFunc("/contact", page.ContactPage).Methods("GET")

	propublicaNew.HandleFunc("/propublica/senators", propublica.GetCurrentSenators).Methods("GET")
	propublicaNew.HandleFunc("/propublica/new", propublica.GetNewMembers).Methods("GET")
	propublicaNew.HandleFunc("/propublica/{id}", propublica.GetSenatorDetails).Methods("GET")
	// http.HandleFunc("/propublica/house", propublica.GetCurrentHouseMembers)
	log.Fatal(http.ListenAndServe(":8080", propublicaNew))
}








