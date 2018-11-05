package propublicastruct


type PoliticanDetailRoot struct {
	Status string `json:"status"`
	Copyright string `json:"copyright"`
	Results []PoliticanDetailResult `json:"results"`
}

type PoliticanDetailResult struct {
	MemberID string `json:"member_id"`
	FirstName string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName string `json:"last_name"`
	Suffix string `json:"suffix"`
	DateOfBirth string `json:"date_of_birth"`
	Gender string `json:"gender"`
	Url string `json:"url"`
	TimeTopicsURL string `json:"times_topics_url"`
	TimesTag string `json:"times_tag"`
	GovTrackID string `json:"govtrack_id"`
	CspanID string `json:"cspan_id"`
	VotesmartID string `json:"votesmart_id"`
	IcpsrID string `json:"icpsr_id"`
	TwitterAccount string `json:"twitter_account"`
	FacebookAccount string `json:"facebook_account"`
	YoutubeAccount string `json:"youtube_account"`
	CrpID string `json:"crp_id"`
	GoogleEntityID string `json:"google_entity_id"`
	RssUrl string `json:"rss_url"`
	InOffice bool `json:"in_office"`
	CurrentParty string `json:"current_party"`
	MostRecentVote string `json:"most_recent_vote"`
	LastUpdated string `json:"current_party"`
	Roles []PoliticanDetailRoles `json:"roles"` 
}

type PoliticanDetailRoles struct {
	Congress string `json:"congress"`
	Chamber string `json:"chamber"`
	Title string `json:"title"`
	ShortTitle string `json:"short_title"`
	State string `json:"state"`
	Party string `json:"party"`
	LeadershipRole string `json:leadership_role`
	FecID string `json:"fec_candidate_id"`
	Seniority string `json:"seniority"`
	District string `json:"district"`
	AtLarge bool `json:"at_large"`
	OcdID string `json:"ocd_id"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	Office string `json:"office"`
	Phone string `json:"phone"`
	Fax string `json:"fax"`
	ContactForm string `json:"contact_form"`
	BillsSponsored int `json:"bills_sponsored"`
	BillsCosponsored int `json:"bills_cosponsored"`
	MissedVotesPCT float64 `json:"missed_votes_pct"`
	VotesWithPartyPCT float64 `json:"votes_with_party_pct"`
	Committee []PoliticanCommittees `json:"committees"`
	Subcommittee []PoliticanSubCommittees `json:"subcommittees"`
}

type PoliticanCommittees struct {
	Name string `json:"name"`
	Code string `json:"code"`
	ApiURL string `json:"api_uri"`
	Side string `json:"side"`
	Title string `json:"title"`
	RankInParty int `json:"rank_in_party"`
	BeginDate string `json:"begin_date"`
	EndDate string `json:"end_date"`
}

type PoliticanSubCommittees struct {
	Name string `json:"name"`
	Code string `json:"code"`
	ParentCommitteeID string `json:"parent_committee_id"`
	ApiURL string `json:"api_uri"`
	Side string `json:"side"`
	Title string `json:"title"`
	RankInParty int `json:"rank_in_party"`
	BeginDate string `json:"begin_date"`
	EndDate string `json:"end_date"`

}
		
	