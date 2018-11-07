package router

import (
	"github.com/gorilla/mux"
	"github.com/tommarler/Beard_Of_knowledge/system/app/controllers/page"
	"github.com/tommarler/Beard_Of_knowledge/system/app/controllers/propublica"

)




type Router struct {
	Router *mux.Router
}

func NewRouter() (r Router) {

	r.Router = mux.NewRouter()

	// // propublicaNew := mux.NewRouter()
	r.Router.HandleFunc("/home", page.HomePage)
	// r.HandleFunc("/contact", page.ContactPage).Methods("GET")

	propublicaRouter := r.Router.PathPrefix("/propublica").Subrouter()
	propublicaRouter.Path("/senators").HandlerFunc(propublica.GetCurrentSenators)
	// propublicaRouter.Path("/new").HandlerFunc(propublica.GetNewMembers).Methods("GET")
	// propublicaRouter.Path("/{id}").HandlerFunc(propublica.GetSenatorDetails).Methods("GET")
	// // propublicaRouter.Path("/propublica/house").HandlerFunc(propublica.GetCurrentHouseMembers).Methods("GET")
	// log.Fatal(http.ListenAndServe(":8080", r.Router))

	return
}