package cxtcreate

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"testpod/network"

	openapi "github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
)

func buildSessCreateData(supi string, pduSessionId int32) *models.PostSmContextsRequest {
	var ueDataStr string
	var request models.PostSmContextsRequest
	ueDataJson := `{
		"supi":"imsi-208930000000003",
		"pei":"imeisv-9902267900440000", "gpsi":"msisdn-0900000000",
		"pduSessionId":5,
		"dnn":"internet",
		"sNssai":{
		   "sst":1,
		   "sd":"010203"
		},
		"servingNfId":"5c599907-63d0-4fcc-90c9-923ac5deb9fd",
		"guami":{
		   "plmnId":{
			  "mcc":"208",
			  "mnc":"93"
		   },
		   "amfId":"f511b2"
		},
		"servingNetwork":{
		   "mcc":"208",
		   "mnc":"93"
		},
		"n1SmMsg":{
		   "contentId":"n1SmMsg"
		},
		"anType":"3GPP_ACCESS",
		"smContextStatusUri":"http://amf:29518/namf-callback/v1/smContextStatus/20893f511b200000003/5"
	 }`
	ueDataStr = strings.ReplaceAll(ueDataJson, " ", "")
	request.JsonData = new(models.SmContextCreateData)
	json.Unmarshal([]byte(ueDataStr), request.JsonData)

	//N1 Data
	n1SmMsg := "2e0501c10000917b006280c223230101002310eca390003edbf917becfa8148acdde56554d54535f434841505f53525652c223150201001510b6faadc56a436b2f0f9f82356e07d9d980211c0100001c810600000000820600000000830600000000840600000000001a0105"
	decoded, err := hex.DecodeString(n1SmMsg)
	if err != nil {
		log.Fatal(err)
	}
	request.BinaryDataN1SmMessage = []byte(decoded)

	//Update user provided data
	if supi != "" {
		request.JsonData.Supi = supi
	}

	if pduSessionId != 0 {
		request.JsonData.PduSessionId = pduSessionId
	}

	return &request
}

func makeHeader(req *http.Request, contentType string) {
	req.Header.Set("Accept", "application/json")
	//req.Header.Set("Content-Type", "application/json")
	log.Println("Content-Type: ", contentType)
	req.Header.Set("Content-Type", contentType)
}

func buildMsgBody(supi string, pduSessionId int32) (*bytes.Buffer, string) {

	body := &bytes.Buffer{}

	//replace IMSI, PDUID as received
	request := buildSessCreateData(supi, pduSessionId)

	contentType, err := openapi.MultipartEncode(request, body)
	if err != nil {
		log.Fatalln(err)
		return nil, ""
	}

	return body, contentType
}

func makePduSessCreateRequest(supi string, pduSessionId int32) *http.Request {
	r, contentType := buildMsgBody(supi, pduSessionId)
	req, err := http.NewRequest("POST", "http://smf:29502/nsmf-pdusession/v1/sm-contexts", r)
	if err != nil {
		log.Fatalln(err)
	}

	//Make header of http Req
	makeHeader(req, contentType)

	return req
}

func SendPduSessCreateRequest(supi string, pduSessionId int32) bool {

	httpReq := makePduSessCreateRequest(supi, pduSessionId)
	resp := network.Send(httpReq)
	defer resp.Body.Close()

	//get UE resource ID(Location) from header
	subscriberUuid := resp.Header["Location"][0]
	log.Println("Subsriber Location: ", subscriberUuid)
	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Println(err)
		return false
	}

	SubsImsiToUuidTable[supi] = subscriberUuid

	log.Println(string(bytes))
	return true
}
