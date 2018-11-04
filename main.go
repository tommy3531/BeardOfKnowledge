package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/tommarler/Beard_Of_knowledge/routes/page"
	"github.com/tommarler/Beard_Of_knowledge/controllers/propublica"
)

func main() {
	fmt.Println("Hello from main")
	http.HandleFunc("/home", page.HomePage)
	http.HandleFunc("/contact", page.ContactPage)
	http.HandleFunc("/propublica/senators", propublica.GetCurrentSenators)
	// http.HandleFunc("/propublica/house", propublica.GetCurrentHouseMembers)
	// http.HandleFunc("/propublica/newMembers", propublica.GetNewMembers)
	log.Fatal(http.ListenAndServe(":8080", nil))

	
}








