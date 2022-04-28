package responses

type ToCampaignInboundBizTxDocRef struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID             string `json:"ObjectID"`
			ParentObjectID       string `json:"ParentObjectID"`
			CampaignID           string `json:"CampaignID"`
			AccountID            string `json:"AccountID"`
			ContactID            string `json:"ContactID"`
			EmployeeID           string `json:"EmployeeID"`
			ReactionTypeCode     string `json:"ReactionTypeCode"`
			ReactionTypeCodeText string `json:"ReactionTypeCodeText"`
			ActivityID           string `json:"ActivityID"`
			OpportunityID        string `json:"OpportunityID"`
			LeadID               string `json:"LeadID"`
			QuoteID              string `json:"QuoteID"`
			CreationDateTime     string `json:"CreationDateTime"`
			Campaign             struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"Campaign"`
		} `json:"results"`
	} `json:"d"`
}
