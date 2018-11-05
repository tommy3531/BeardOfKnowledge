package propublica

import (
	"net/http"
	"net/url"
	"html/template"
	"os"
	"encoding/json"
	"log"
	"fmt"
	"io/ioutil"
	"github.com/gorilla/mux"
	propublicaModel "github.com/tommarler/Beard_Of_knowledge/models/propublicastruct"
)

func GetSenatorDetails(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	client := &http.Client{}
	url := url.URL{
		Scheme: "https",
		Host:   "api.propublica.org",
		Path:   "/congress/v1/members/" + id + ".json",
	}
	println(url.String())
	req, _ := http.NewRequest("GET", url.String(), nil)

	req.Header.Add("X-API-KEY", "SpzjlPZlkMlPKKGCLQS1OqZtCN96lPl7sszOTKra")
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	var responseObject propublicaModel.PoliticanDetailRoot
	if err := json.NewDecoder(resp.Body).Decode(&responseObject); err != nil {
		log.Println(err)
	}

	t, _ := template.ParseFiles("template/showSenatorDetails.html")
	t.Execute(w, responseObject)
}

func GetCurrentSenators(w http.ResponseWriter, r *http.Request) {

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

	t, _ := template.ParseFiles("template/showAllSenators.html")
	t.Execute(w, responseObject)
}

func GetNewMembers(w http.ResponseWriter, r *http.Request) {


	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.propublica.org/congress/v1/members/new.json", nil)
	if err != nil {
		os.Exit(1)
	}

	req.Header.Add("X-API-KEY", "SpzjlPZlkMlPKKGCLQS1OqZtCN96lPl7sszOTKra")
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	// var responseObject propublicaModel.PoliticanNewRoot
	// if err := json.NewDecoder(resp.Body).Decode(&responseObject); err != nil {
	// 	log.Println(err)
	// }

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	t, _ := template.ParseFiles("template/showNewMember.html")
	t.Execute(w, "title")
}



