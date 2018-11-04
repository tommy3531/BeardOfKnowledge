package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/tommarler/Beard_Of_knowledge/routes/page"
)



func main() {
	fmt.Println("Hello from main")
	// propublicaAPI.GetCurrentSenators()
	// propublicaAPI.GetCurrentHouseMembers()
	http.HandleFunc("/home", page.HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}








