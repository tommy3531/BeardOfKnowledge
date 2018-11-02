package api

import (
	"net/http"
	"os"
	"encoding/json"
	"log"
	"fmt"
	propublicaModel "github.com/tommarler/Beard_Of_knowledge/models/propublica"
)

const BASEURL = "https://api.propublica.org/congress/v1/115/senator/members.json"

func PropublicaClient() {

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

