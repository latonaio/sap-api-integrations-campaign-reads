package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-campaign-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetCampaign(campaignID, campaignName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "CampaignCollection":
			func() {
				c.CampaignCollection(campaignID)
				wg.Done()
			}()
		case "CampaignName":
			func() {
				c.CampaignName(campaignName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) CampaignCollection(campaignID string) {
	campaignCollectionData, err := c.callCampaignSrvAPIRequirementCampaignCollection("CampaignCollection", campaignID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(campaignCollectionData)

	campaignInboundBizTxDocRefData, err := c.callToCampaignInboundBizTxDocRef(campaignCollectionData[0].ToCampaignInboundBizTxDocRef)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(campaignInboundBizTxDocRefData)

}

func (c *SAPAPICaller) callCampaignSrvAPIRequirementCampaignCollection(api, campaignID string) ([]sap_api_output_formatter.CampaignCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithCampaignCollection(req, campaignID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCampaignCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToCampaignInboundBizTxDocRef(url string) ([]sap_api_output_formatter.ToCampaignInboundBizTxDocRef, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToCampaignInboundBizTxDocRef(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) CampaignName(campaignName string) {
	campaignNameData, err := c.callCampaignSrvAPIRequirementCampaignName("CampaignCollection", campaignName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(campaignNameData)

	campaignInboundBizTxDocRefData, err := c.callToCampaignInboundBizTxDocRef(campaignNameData[0].ToCampaignInboundBizTxDocRef)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(campaignInboundBizTxDocRefData)

}

func (c *SAPAPICaller) callCampaignSrvAPIRequirementCampaignName(api, campaignName string) ([]sap_api_output_formatter.CampaignCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithCampaignName(req, campaignName)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCampaignCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithCampaignCollection(req *http.Request, campaignID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("CampaignID eq '%s'", campaignID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithCampaignName(req *http.Request, campaignName string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("substringof('%s', CampaignName)", campaignName))
	req.URL.RawQuery = params.Encode()
}
