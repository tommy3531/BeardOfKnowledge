package main

import (
	"fmt"
	"net/http"
	"log"
	// propublicaAPI "github.com/tommarler/Beard_Of_knowledge/controllers/propublica"
)

func main() {
	fmt.Println("Hello from main")
	// propublicaAPI.GetCurrentSenators()
	// propublicaAPI.GetCurrentHouseMembers()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	
}




