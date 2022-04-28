package main

import (
	sap_api_caller "sap-api-integrations-campaign-reads/SAP_API_Caller"
	"sap-api-integrations-campaign-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Campaign_Campaign_Name_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/sap/c4c/odata/v1/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"CampaignCollection", "CampaignName",
		}
	}

	caller.AsyncGetCampaign(
		inoutSDC.CampaignCollection.CampaignID,
		inoutSDC.CampaignCollection.CampaignName,
		accepter,
	)
}
