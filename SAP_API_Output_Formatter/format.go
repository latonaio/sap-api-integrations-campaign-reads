package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-campaign-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToCampaignCollection(raw []byte, l *logger.Logger) ([]CampaignCollection, error) {
	pm := &responses.CampaignCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CampaignCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	campaignCollection := make([]CampaignCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		campaignCollection = append(campaignCollection, CampaignCollection{
			ObjectID:                     data.ObjectID,
			CampaignType:                 data.CampaignType,
			CampaignTypeText:             data.CampaignTypeText,
			CampaignID:                   data.CampaignID,
			CampaignName:                 data.CampaignName,
			EndDate:                      data.EndDate,
			StartDate:                    data.StartDate,
			Status:                       data.Status,
			StatusText:                   data.StatusText,
			ChannelTypeCode:              data.ChannelTypeCode,
			ChannelTypeCodeText:          data.ChannelTypeCodeText,
			TargetGroupID:                data.TargetGroupID,
			SalesOrganization:            data.SalesOrganization,
			EmployeeResponsibleID:        data.EmployeeResponsibleID,
			ReferenceID:                  data.ReferenceID,
			ReferenceBusinessSystemID:    data.ReferenceBusinessSystemID,
			EntityLastChangedOn:          data.EntityLastChangedOn,
			ToCampaignInboundBizTxDocRef: data.CampaignInboundBusinessTransactionDocumentReference.Deferred.URI,
		})
	}

	return campaignCollection, nil
}

func ConvertToToCampaignInboundBizTxDocRef(raw []byte, l *logger.Logger) ([]ToCampaignInboundBizTxDocRef, error) {
	pm := &responses.ToCampaignInboundBizTxDocRef{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToCampaignInboundBizTxDocRef. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toCampaignInboundBizTxDocRef := make([]ToCampaignInboundBizTxDocRef, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toCampaignInboundBizTxDocRef = append(toCampaignInboundBizTxDocRef, ToCampaignInboundBizTxDocRef{
			ObjectID:             data.ObjectID,
			ParentObjectID:       data.ParentObjectID,
			CampaignID:           data.CampaignID,
			AccountID:            data.AccountID,
			ContactID:            data.ContactID,
			EmployeeID:           data.EmployeeID,
			ReactionTypeCode:     data.ReactionTypeCode,
			ReactionTypeCodeText: data.ReactionTypeCodeText,
			ActivityID:           data.ActivityID,
			OpportunityID:        data.OpportunityID,
			LeadID:               data.LeadID,
			QuoteID:              data.QuoteID,
			CreationDateTime:     data.CreationDateTime,
		})
	}

	return toCampaignInboundBizTxDocRef, nil
}
