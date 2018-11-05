package propublica

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

func GetSenatorDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	t, _ := template.ParseFiles("template/showSenatorDetails.html")
	t.Execute(w, id)
}

// func GetHouseDetails() {
// 	// Need Legislator ID
// 	fmt.Println("Get House Details: Need leg_id")
// }