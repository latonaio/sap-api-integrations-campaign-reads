# sap-api-integrations-campaign-reads  
sap-api-integrations-campaign-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API キャンペーンデータを取得するマイクロサービスです。  
sap-api-integrations-campaign-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-campaign-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/campaign/overview  

## 動作環境
sap-api-integrations-campaign-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-campaign-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-campaign-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/campaign/overview 
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-campaign-reads には、次の API をコールするためのリソースが含まれています。  

* CampaignCollection（キャンペーン - キャンペーン）※キャンペーンの詳細データを取得するために、ToCampaignInboundBizTxDocRef、と合わせて利用されます。
* ToCampaignInboundBizTxDocRef（キャンペーン - キャンペーン取引文書 ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-campaign-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.CampaignCollection.CampaignID（キャンペーンID）  
* inoutSDC.CampaignCollection.CampaignName（キャンペーン名）  


## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"CampaignCollection" が指定されています。    
  
```
	"api_schema": "CampaignCampaignCollection",
	"accepter": ["CampaignCollection"],
	"campaign_code": "1",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "CampaignCampaignCollection",
	"accepter": ["All"],
	"campaign_code": "1",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
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
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP キャンペーン  の キャンペーンデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "CampaignInboundBusinessTransactionDocumentReference" は、/SAP_API_Output_Formatter/type.go 内 の Type CampaignCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-campaign-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-campaign-reads/SAP_API_Caller.(*SAPAPICaller).CampaignCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E03A0701ED28B9EB2A63189D226",
			"CampaignType": "4",
			"CampaignTypeText": "Phone Call",
			"CampaignID": "1",
			"CampaignName": "Campaign for CN Customers",
			"EndDate": "2012-09-30T09:00:00+09:00",
			"StartDate": "2012-05-01T09:00:00+09:00",
			"Status": "2",
			"StatusText": "Active",
			"ChannelTypeCode": "",
			"ChannelTypeCodeText": "",
			"TargetGroupID": "1",
			"SalesOrganization": "",
			"EmployeeResponsibleID": "",
			"ReferenceID": "",
			"ReferenceBusinessSystemID": "",
			"EntityLastChangedOn": "2012-11-13T03:20:06+09:00",
			"CampaignInboundBusinessTransactionDocumentReference": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CampaignCollection('00163E03A0701ED28B9EB2A63189D226')/CampaignInboundBusinessTransactionDocumentReference"
		},
		{
			"ObjectID": "00163E03A0701ED28B9EB2A63189F226",
			"CampaignType": "4",
			"CampaignTypeText": "Phone Call",
			"CampaignID": "2",
			"CampaignName": "Campaign for IT Customers",
			"EndDate": "2012-09-30T09:00:00+09:00",
			"StartDate": "2012-05-01T09:00:00+09:00",
			"Status": "2",
			"StatusText": "Active",
			"ChannelTypeCode": "",
			"ChannelTypeCodeText": "",
			"TargetGroupID": "2",
			"SalesOrganization": "",
			"EmployeeResponsibleID": "",
			"ReferenceID": "",
			"ReferenceBusinessSystemID": "",
			"EntityLastChangedOn": "2012-11-13T03:20:06+09:00",
			"CampaignInboundBusinessTransactionDocumentReference": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CampaignCollection('00163E03A0701ED28B9EB2A63189F226')/CampaignInboundBusinessTransactionDocumentReference"
		},
		{
			"ObjectID": "00163E03A0701ED28B9EB2A6318A1226",
			"CampaignType": "4",
			"CampaignTypeText": "Phone Call",
			"CampaignID": "3",
			"CampaignName": "Campaña para clientes de MX",
			"EndDate": "2012-09-30T09:00:00+09:00",
			"StartDate": "2012-05-01T09:00:00+09:00",
			"Status": "2",
			"StatusText": "Active",
			"ChannelTypeCode": "",
			"ChannelTypeCodeText": "",
			"TargetGroupID": "3",
			"SalesOrganization": "",
			"EmployeeResponsibleID": "",
			"ReferenceID": "",
			"ReferenceBusinessSystemID": "",
			"EntityLastChangedOn": "2012-11-13T03:20:06+09:00",
			"CampaignInboundBusinessTransactionDocumentReference": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CampaignCollection('00163E03A0701ED28B9EB2A6318A1226')/CampaignInboundBusinessTransactionDocumentReference"
		},
		{
			"ObjectID": "00163E03A0701ED28B9EB2A6318A3226",
			"CampaignType": "4",
			"CampaignTypeText": "Phone Call",
			"CampaignID": "4",
			"CampaignName": "Campaña para clientes de España",
			"EndDate": "2012-09-30T09:00:00+09:00",
			"StartDate": "2012-05-01T09:00:00+09:00",
			"Status": "2",
			"StatusText": "Active",
			"ChannelTypeCode": "",
			"ChannelTypeCodeText": "",
			"TargetGroupID": "4",
			"SalesOrganization": "",
			"EmployeeResponsibleID": "",
			"ReferenceID": "",
			"ReferenceBusinessSystemID": "",
			"EntityLastChangedOn": "2012-11-13T03:20:06+09:00",
			"CampaignInboundBusinessTransactionDocumentReference": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CampaignCollection('00163E03A0701ED28B9EB2A6318A3226')/CampaignInboundBusinessTransactionDocumentReference"
		},
		{
			"ObjectID": "00163E03A0701ED28B9EB2A6318A5226",
			"CampaignType": "4",
			"CampaignTypeText": "Phone Call",
			"CampaignID": "5",
			"CampaignName": "Campagne voor NL klanten",
			"EndDate": "2012-09-30T09:00:00+09:00",
			"StartDate": "2012-05-01T09:00:00+09:00",
			"Status": "2",
			"StatusText": "Active",
			"ChannelTypeCode": "",
			"ChannelTypeCodeText": "",
			"TargetGroupID": "5",
			"SalesOrganization": "",
			"EmployeeResponsibleID": "",
			"ReferenceID": "",
			"ReferenceBusinessSystemID": "",
			"EntityLastChangedOn": "2012-11-13T03:20:06+09:00",
			"CampaignInboundBusinessTransactionDocumentReference": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CampaignCollection('00163E03A0701ED28B9EB2A6318A5226')/CampaignInboundBusinessTransactionDocumentReference"
		},
		{
			"ObjectID": "00163E03A0701ED28B9EB2A6318A7226",
			"CampaignType": "4",
			"CampaignTypeText": "Phone Call",
			"CampaignID": "6",
			"CampaignName": "Campagne pour les Clients Français",
			"EndDate": "2012-09-30T09:00:00+09:00",
			"StartDate": "2012-05-01T09:00:00+09:00",
			"Status": "2",
			"StatusText": "Active",
			"ChannelTypeCode": "",
			"ChannelTypeCodeText": "",
			"TargetGroupID": "6",
			"SalesOrganization": "",
			"EmployeeResponsibleID": "",
			"ReferenceID": "",
			"ReferenceBusinessSystemID": "",
			"EntityLastChangedOn": "2012-11-13T03:20:06+09:00",
			"CampaignInboundBusinessTransactionDocumentReference": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CampaignCollection('00163E03A0701ED28B9EB2A6318A7226')/CampaignInboundBusinessTransactionDocumentReference"
		},
		{
			"ObjectID": "00163E03A0701ED28B9EB2A6318A9226",
			"CampaignType": "4",
			"CampaignTypeText": "Phone Call",
			"CampaignID": "7",
			"CampaignName": "Kampagne für Kunden in CH",
			"EndDate": "2012-09-30T09:00:00+09:00",
			"StartDate": "2012-05-01T09:00:00+09:00",
			"Status": "2",
			"StatusText": "Active",
			"ChannelTypeCode": "",
			"ChannelTypeCodeText": "",
			"TargetGroupID": "7",
			"SalesOrganization": "",
			"EmployeeResponsibleID": "",
			"ReferenceID": "",
			"ReferenceBusinessSystemID": "",
			"EntityLastChangedOn": "2012-11-13T03:20:06+09:00",
			"CampaignInboundBusinessTransactionDocumentReference": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CampaignCollection('00163E03A0701ED28B9EB2A6318A9226')/CampaignInboundBusinessTransactionDocumentReference"
		},
		{
			"ObjectID": "00163E03A0701ED28B9EB2A6318AB226",
			"CampaignType": "4",
			"CampaignTypeText": "Phone Call",
			"CampaignID": "8",
			"CampaignName": "Kampagne für Kunden in AT",
			"EndDate": "2012-09-30T09:00:00+09:00",
			"StartDate": "2012-05-01T09:00:00+09:00",
			"Status": "2",
			"StatusText": "Active",
			"ChannelTypeCode": "",
			"ChannelTypeCodeText": "",
			"TargetGroupID": "8",
			"SalesOrganization": "",
			"EmployeeResponsibleID": "",
			"ReferenceID": "",
			"ReferenceBusinessSystemID": "",
			"EntityLastChangedOn": "2012-11-13T03:20:06+09:00",
			"CampaignInboundBusinessTransactionDocumentReference": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CampaignCollection('00163E03A0701ED28B9EB2A6318AB226')/CampaignInboundBusinessTransactionDocumentReference"
		},
		{
			"ObjectID": "00163E03A0701ED28B9EB2A6318AD226",
			"CampaignType": "4",
			"CampaignTypeText": "Phone Call",
			"CampaignID": "9",
			"CampaignName": "Kampagne für Kunden in DE",
			"EndDate": "2012-09-30T09:00:00+09:00",
			"StartDate": "2012-05-01T09:00:00+09:00",
			"Status": "2",
			"StatusText": "Active",
			"ChannelTypeCode": "",
			"ChannelTypeCodeText": "",
			"TargetGroupID": "9",
			"SalesOrganization": "",
			"EmployeeResponsibleID": "",
			"ReferenceID": "",
			"ReferenceBusinessSystemID": "",
			"EntityLastChangedOn": "2012-11-13T03:20:06+09:00",
			"CampaignInboundBusinessTransactionDocumentReference": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CampaignCollection('00163E03A0701ED28B9EB2A6318AD226')/CampaignInboundBusinessTransactionDocumentReference"
		},
		{
			"ObjectID": "00163E03A0701ED28B9EB2A6318AF226",
			"CampaignType": "4",
			"CampaignTypeText": "Phone Call",
			"CampaignID": "10",
			"CampaignName": "Campaign for AU Customers",
			"EndDate": "2012-09-30T09:00:00+09:00",
			"StartDate": "2012-05-01T09:00:00+09:00",
			"Status": "2",
			"StatusText": "Active",
			"ChannelTypeCode": "",
			"ChannelTypeCodeText": "",
			"TargetGroupID": "11",
			"SalesOrganization": "",
			"EmployeeResponsibleID": "",
			"ReferenceID": "",
			"ReferenceBusinessSystemID": "",
			"EntityLastChangedOn": "2012-11-13T03:20:06+09:00",
			"CampaignInboundBusinessTransactionDocumentReference": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CampaignCollection('00163E03A0701ED28B9EB2A6318AF226')/CampaignInboundBusinessTransactionDocumentReference"
		}
	],
	"time": "2022-04-28T21:39:34+09:00"
}

```