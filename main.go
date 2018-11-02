package main

import (
	"fmt"
	// propublicaBasic "github.com/tommarler/Beard_Of_knowledge/controllers/propublica"
	// propubicaDetail "github.com/tommarler/Beard_Of_knowledge/controllers/propublica"
	propublicaAPI "github.com/tommarler/Beard_Of_knowledge/api"
)

func main() {
	fmt.Println("Hello from main")
	propublicaAPI.PropublicaClient()
	// propublicaBasic.GetCurrentSenators()
	// propublicaBasic.GetCurrentHouseMembers()

	// propubicaDetail.GetSenatorDetails()
	// propubicaDetail.GetHouseDetails()
}