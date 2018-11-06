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
	// propublicaNew := mux.NewRouter()
	r.HandleFunc("/home", page.HomePage).Methods("GET")
	r.HandleFunc("/contact", page.ContactPage).Methods("GET")

	propublicaRouter := r.PathPrefix("/propublica").Subrouter()
	propublicaRouter.Path("/senators").HandlerFunc(propublica.GetCurrentSenators).Methods("GET")
	propublicaRouter.Path("/new").HandlerFunc(propublica.GetNewMembers).Methods("GET")
	propublicaRouter.Path("/{id}").HandlerFunc(propublica.GetSenatorDetails).Methods("GET")
	// propublicaRouter.Path("/propublica/house").HandlerFunc(propublica.GetCurrentHouseMembers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}








