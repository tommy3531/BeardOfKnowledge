package router

import (
	"github.com/tommarler/Beard_Of_knowledge/system/app/routes"
	"github.com/tommarler/Beard_Of_knowledge/system/app/controllers/propublica"
	"github.com/tommarler/Beard_Of_knowledge/system/app/controllers/page"
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside main middleware.")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() routes.Routes {

	// AuthHandler.Init(db)
	// HomeHandler.Init(db)

	// r.Router.HandleFunc("/home", page.HomePage)
	// r.HandleFunc("/contact", page.ContactPage).Methods("GET")

	// r.Router.HandleFunc("/propublica/senators", propublica.GetCurrentSenators)
	// propublicaRouter.Path("/new").HandlerFunc(propublica.GetNewMembers).Methods("GET")
	// propublicaRouter.Path("/{id}").HandlerFunc(propublica.GetSenatorDetails).Methods("GET")
	// // propublicaRouter.Path("/propublica/house").HandlerFunc(propublica.GetCurrentHouseMembers).Methods("GET")

	return routes.Routes{
		routes.Route{"Home", "GET", "/", page.HomePage},
		routes.Route{"Senators", "GET", "/propublica/senators", propublica.GetCurrentSenators},
		// routes.Route{"AuthCheck", "GET", "/auth/check", AuthHandler.Check},
	}
}