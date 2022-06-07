// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package cxtdelete

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"testpod/network"
	"testpod/testnodes/amf/testcases/cxtcreate"

	openapi "github.com/omec-project/openapi"
	"github.com/omec-project/openapi/models"
)

func buildSessUpCnxDeactivateData() *models.UpdateSmContextRequest {
	var ueDataStr string
	var request models.UpdateSmContextRequest
	ueDataJson := `{
		"ueLocation":{
		   "nrLocation":{
			  "tai":{
				 "plmnId":{
					"mcc":"208",
					"mnc":"93"
				 },
				 "tac":"000001"
			  },
			  "ncgi":{
				 "plmnId":{
					"mcc":"208",
					"mnc":"93"
				 },
				 "nrCellId":"000004000"
			  },
			  "ueLocationTimestamp":"2021-06-25T09:43:56.26582706Z"
		   }
		},
		"upCnxState":"DEACTIVATED",
		"ngApCause":{
		   "group":3,
		   "value":0
		}
	 }`
	ueDataStr = strings.ReplaceAll(ueDataJson, " ", "")
	request.JsonData = new(models.SmContextUpdateData)
	json.Unmarshal([]byte(ueDataStr), request.JsonData)

	return &request
}

func makeHeader(req *http.Request, contentType string) {
	req.Header.Set("Accept", "application/json")
	//req.Header.Set("Content-Type", "application/json")
	log.Println("Content-Type: ", contentType)
	req.Header.Set("Content-Type", contentType)
}

func buildMsgBody() (*bytes.Buffer, string) {

	body := &bytes.Buffer{}

	//replace IMSI, PDUID as received
	request := buildSessUpCnxDeactivateData()

	contentType, err := openapi.MultipartEncode(request, body)
	if err != nil {
		log.Fatalln(err)
		return nil, ""
	}

	return body, contentType
}

func makePduSessUpCnxDeactivateRequest(supi string) *http.Request {
	r, contentType := buildMsgBody()
	uuid := cxtcreate.SubsImsiToUuidTable[supi]
	log.Println("Modify Req, Subs: ", uuid)
	url := "/nsmf-pdusession/v1/sm-contexts/" + uuid + "/modify"
	//req, err := http.NewRequest("POST", "http://smf:29502/nsmf-pdusession/v1/sm-contexts/urn:uuid:553162e0-3ecb-43b8-b74e-04c21ca6f12a/modify", r)
	req, err := http.NewRequest("POST", "http://smf:29502"+url, r)
	if err != nil {
		log.Fatalln(err)
	}

	//Make header of http Req
	makeHeader(req, contentType)

	return req
}

func SendPduSessUpCnxDeactivateRequest(supi string) bool {

	httpReq := makePduSessUpCnxDeactivateRequest(supi)
	resp := network.Send(httpReq)
	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(string(bytes))
	return true
}

//Release
func SendPduSessReleaseRequest(supi string) bool {

	httpReq := makePduSessReleaseRequest(supi)
	resp := network.Send(httpReq)
	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(string(bytes))
	return true
}

func makePduSessReleaseRequest(supi string) *http.Request {
	r, contentType := buildRelMsgBody()
	uuid := cxtcreate.SubsImsiToUuidTable[supi]
	log.Println("Release Req, Subs: ", uuid)
	url := "/nsmf-pdusession/v1/sm-contexts/" + uuid + "/release"
	//req, err := http.NewRequest("POST", "http://smf:29502/nsmf-pdusession/v1/sm-contexts/urn:uuid:553162e0-3ecb-43b8-b74e-04c21ca6f12a/release", r)
	req, err := http.NewRequest("POST", "http://smf:29502"+url, r)
	if err != nil {
		log.Fatalln(err)
	}

	//Make header of http Req
	makeHeader(req, contentType)

	return req
}

func buildRelMsgBody() (*bytes.Buffer, string) {

	body := &bytes.Buffer{}

	//replace IMSI, PDUID as received
	request := buildSessReleaseData()

	contentType, err := openapi.MultipartEncode(request, body)
	if err != nil {
		log.Fatalln(err)
		return nil, ""
	}

	return body, contentType
}

func buildSessReleaseData() *models.ReleaseSmContextRequest {
	var ueDataStr string
	var request models.ReleaseSmContextRequest
	ueDataJson := `{
		"ueLocation":{
		   "nrLocation":{
			  "tai":{
				 "plmnId":{
					"mcc":"208",
					"mnc":"93"
				 },
				 "tac":"000001"
			  },
			  "ncgi":{
				 "plmnId":{
					"mcc":"208",
					"mnc":"93"
				 },
				 "nrCellId":"000004000"
			  },
			  "ueLocationTimestamp":"2021-06-25T09:43:56.26582706Z"
		   }
		},
		"upCnxState":"DEACTIVATED",
		"ngApCause":{
		   "group":3,
		   "value":0
		}
	 }`
	ueDataStr = strings.ReplaceAll(ueDataJson, " ", "")
	request.JsonData = new(models.SmContextReleaseData)
	json.Unmarshal([]byte(ueDataStr), request.JsonData)

	return &request
}
