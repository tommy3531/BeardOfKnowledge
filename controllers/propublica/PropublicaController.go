package propublica

import (
	"net/http"
	"os"
	"encoding/json"
	"log"
	"fmt"
	"html/template"
	propublicaModel "github.com/tommarler/Beard_Of_knowledge/models/propublica"
)

func GetCurrentSenators(w http.ResponseWriter, r *http.Request) {

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
	t, _ := template.ParseFiles("template/showAllSenators.html")
	t.Execute(w, responseObject)
}

// func GetCurrentHouseMembers((w http.ResponseWriter, r *http.Request) {

// 	// 102-115
// 	fmt.Println("Get Current House Memebers: (102-115)")
// }

// func GetNewMembers((w http.ResponseWriter, r *http.Request) {

// 	fmt.Println("Display New Members of Congress")
// }

// func GetCurrentMemberByState((w http.ResponseWriter, r *http.Request) {

// 	fmt.Println("Get Member by State: Need Two Digit State")
// }

// func GetMembersLeavingOffice((w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Get Member Leaving Office: ")
// }

// func GetMemberExpenses((w http.ResponseWriter, r *http.Request) {

// 	// leg_id, year, quarter
// 	fmt.Println("Get Member Expenses by year and quarter ")
// }

// func GetMemberExpensesByCategory((w http.ResponseWriter, r *http.Request) {
	
// 	// leg_id
// 	// Category -> travel, personnel, rent-utilities, other-services, supplies, franked-mail, printing, equipment, total
// 	fmt.Println("Get Expenses By Category")
// }

