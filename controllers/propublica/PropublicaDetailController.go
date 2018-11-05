package propublica

import (
	"net/http"
	"net/url"
	"html/template"
	// "encoding/json"
	// "log"
	// "html/template"
	"github.com/gorilla/mux"
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

	// if resp.StatusCode == http.StatusOK {
	// 	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// 	bodyString := string(bodyBytes)
	// 	fmt.Println(bodyString)
	// }

	// var responseObject propublicaModel.Politican
	// if err := json.NewDecoder(resp.Body).Decode(&responseObject); err != nil {
	// 	log.Println(err)
	// }

	t, _ := template.ParseFiles("template/showSenatorDetails.html")
	t.Execute(w, "")
}

// func GetHouseDetails() {
// 	// Need Legislator ID
// 	fmt.Println("Get House Details: Need leg_id")
// }