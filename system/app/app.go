package app

import (
	"log"
	"net/http"
	// "github.com/gorilla/mux"
	"github.com/tommarler/Beard_Of_knowledge/system/app/controllers/page"
	"github.com/tommarler/Beard_Of_knowledge/system/app/router"

)

type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Init(port string) {
	log.Println("Initializing server....")
	s.port = ":" + port
}

func (s *Server) Start() {
	log.Println("Starting server on port!" + s.port)

	r := router.NewRouter() 
	// // propublicaNew := mux.NewRouter()
	r.Router.HandleFunc("/home", page.HomePage).Methods("GET")
	// r.HandleFunc("/contact", page.ContactPage).Methods("GET")

	// propublicaRouter := r.PathPrefix("/propublica").Subrouter()
	// propublicaRouter.Path("/senators").HandlerFunc(propublica.GetCurrentSenators).Methods("GET")
	// propublicaRouter.Path("/new").HandlerFunc(propublica.GetNewMembers).Methods("GET")
	// propublicaRouter.Path("/{id}").HandlerFunc(propublica.GetSenatorDetails).Methods("GET")
	// // propublicaRouter.Path("/propublica/house").HandlerFunc(propublica.GetCurrentHouseMembers).Methods("GET")
	
	http.ListenAndServe(s.port, r.Router)
}