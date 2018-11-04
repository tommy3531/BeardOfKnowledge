package propublica

import (
	"fmt"
	"net/http"
	"html/template"
)

func GetSenatorDetails(w http.ResponseWriter, r *http.Request) {
	// Need Legislator ID
	fmt.Println("Get Senator Details: Need leg_id")
	title := "Contact Page"
	t, _ := template.ParseFiles("template/showSenatorDetails.html")
	t.Execute(w, title)
}

// func GetHouseDetails() {
// 	// Need Legislator ID
// 	fmt.Println("Get House Details: Need leg_id")
// }