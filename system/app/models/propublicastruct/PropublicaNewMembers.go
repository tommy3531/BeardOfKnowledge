package propublicastruct

type PoliticanNewRoot struct {
  Status string `json:"status"`
  Copyright string `json:"copyright"`
  Results []PoliticanNewResults `json:"results"`
}

type PoliticanNewResults struct {
  NumResults string `json:"num_results"`
  Offset string `json:"offset"`
  Members []PoliticanNewMembers `json:"members"`

}

type PoliticanNewMembers struct {
  ID string `json:"id"`
  APIURI string `json:"api_uri"`
  FirstName string `json:"first_name"`
  MiddleName string `json:"middle_name"`
  LastName string `json:"last_name"`
  Suffix string `json:"suffix"`
  Party string `json:"party"`
  Chamber string `json:"chamber"`
  State string `json:"state"`
  District string `json:"district"`
  StartDate string `json:"start_date"`
}