package responses

type CampaignCollection struct {
	D struct {
		Count   string `json:"__count"`
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                                            string `json:"ObjectID"`
			CampaignType                                        string `json:"CampaignType"`
			CampaignTypeText                                    string `json:"CampaignTypeText"`
			CampaignID                                          string `json:"CampaignID"`
			CampaignName                                        string `json:"CampaignName"`
			EndDate                                             string `json:"EndDate"`
			StartDate                                           string `json:"StartDate"`
			Status                                              string `json:"Status"`
			StatusText                                          string `json:"StatusText"`
			ChannelTypeCode                                     string `json:"ChannelTypeCode"`
			ChannelTypeCodeText                                 string `json:"ChannelTypeCodeText"`
			TargetGroupID                                       string `json:"TargetGroupID"`
			SalesOrganization                                   string `json:"SalesOrganization"`
			EmployeeResponsibleID                               string `json:"EmployeeResponsibleID"`
			ReferenceID                                         string `json:"ReferenceID"`
			ReferenceBusinessSystemID                           string `json:"ReferenceBusinessSystemID"`
			EntityLastChangedOn                                 string `json:"EntityLastChangedOn"`
			CampaignInboundBusinessTransactionDocumentReference struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"CampaignInboundBusinessTransactionDocumentReference"`
		} `json:"results"`
	} `json:"d"`
}
