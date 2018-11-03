package controllers

import (
	"net/http"
	"os"
	"encoding/json"
	"log"
	"fmt"
	propublicaModel "github.com/tommarler/Beard_Of_knowledge/models/propublica"
)

func GetCurrentSenators() {

	fmt.Println("Get ALL Senators")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.propublica.org/congress/v1/115/senate/members.json", nil)
	if err != nil {
		os.Exit(1)
	}

	req.Header.Add("X-API-KEY", "SpzjlPZlkMlPKKGCLQS1OqZtCN96lPl7sszOTKra")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	var responseObject propublicaModel.Politican
	if err := json.NewDecoder(resp.Body).Decode(&responseObject); err != nil {
		log.Println(err)
	}

	fmt.Println(responseObject)
}

func GetCurrentHouseMembers() {

	// 102-115
	fmt.Println("Get Current House Memebers: (102-115)")
}

func GetNewMembers() {

	fmt.Println("Display New Members of Congress")
}

func GetCurrentMemberByState() {

	fmt.Println("Get Member by State: Need Two Digit State")
}

func GetMembersLeavingOffice() {
	fmt.Println("Get Member Leaving Office: ")
}

func GetMemberExpenses() {

	// leg_id, year, quarter
	fmt.Println("Get Member Expenses by year and quarter ")
}

func GetMemberExpensesByCategory() {
	
	// leg_id
	// Category -> travel, personnel, rent-utilities, other-services, supplies, franked-mail, printing, equipment, total
	fmt.Println("Get Expenses By Category")
}

