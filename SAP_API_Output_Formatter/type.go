package sap_api_output_formatter

type Campaign struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	CampaignCode  string `json:"campaign_code"`
	Deleted       bool   `json:"deleted"`
}

type CampaignCollection struct {
	ObjectID                     string `json:"ObjectID"`
	CampaignType                 string `json:"CampaignType"`
	CampaignTypeText             string `json:"CampaignTypeText"`
	CampaignID                   string `json:"CampaignID"`
	CampaignName                 string `json:"CampaignName"`
	EndDate                      string `json:"EndDate"`
	StartDate                    string `json:"StartDate"`
	Status                       string `json:"Status"`
	StatusText                   string `json:"StatusText"`
	ChannelTypeCode              string `json:"ChannelTypeCode"`
	ChannelTypeCodeText          string `json:"ChannelTypeCodeText"`
	TargetGroupID                string `json:"TargetGroupID"`
	SalesOrganization            string `json:"SalesOrganization"`
	EmployeeResponsibleID        string `json:"EmployeeResponsibleID"`
	ReferenceID                  string `json:"ReferenceID"`
	ReferenceBusinessSystemID    string `json:"ReferenceBusinessSystemID"`
	EntityLastChangedOn          string `json:"EntityLastChangedOn"`
	ToCampaignInboundBizTxDocRef string `json:"CampaignInboundBusinessTransactionDocumentReference"`
}

type ToCampaignInboundBizTxDocRef struct {
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
}
