// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package cxtupdate

import (
	"bytes"
	"encoding/hex"
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

type Case1 struct {
}

var ueDataStr string
var request models.UpdateSmContextRequest

//Get from Yaml file
func init() {
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
			  "ueLocationTimestamp":"2021-04-21T11:43:42.647807417Z"
		   }
		},
		"n2SmInfo":{
		   "contentId":"N2SmInfo"
		},
		"n2SmInfoType":"PDU_RES_SETUP_RSP"
	 }`
	ueDataStr = strings.ReplaceAll(ueDataJson, " ", "")
	request.JsonData = new(models.SmContextUpdateData)
	json.Unmarshal([]byte(ueDataStr), request.JsonData)
	n2SmMsg := "0003e0c0a8fbfc000000010009"
	decoded, err := hex.DecodeString(n2SmMsg)
	if err != nil {
		log.Fatal(err)
	}
	request.BinaryDataN1SmMessage = nil
	request.BinaryDataN2SmInformation = []byte(decoded)
}

func (c Case1) makeHeader(req *http.Request, contentType string) {
	req.Header.Set("Accept", "application/json")
	//req.Header.Set("Content-Type", "application/json")
	fmt.Println("Content-Type: ", contentType)
	req.Header.Set("Content-Type", contentType)
}

func (c Case1) makeBody() (*bytes.Buffer, string) {

	body := &bytes.Buffer{}
	contentType, err := openapi.MultipartEncode(&request, body)
	if err != nil {
		log.Fatalln(err)
		return nil, ""
	}

	return body, contentType

	/*
		byteData, err := json.Marshal(request.JsonData)
		if err != nil {
			fmt.Println("Marshal Error: ", err.Error())
		}

		fmt.Println("Json body:", string(byteData))
		return bytes.NewReader(byteData)
	*/
}

func (c Case1) Execute(supi string) bool {
	r, contentType := c.makeBody()
	uuid := cxtcreate.SubsImsiToUuidTable[supi]
	fmt.Println("Modify Req, Subs: ", uuid)
	url := "/nsmf-pdusession/v1/sm-contexts/" + uuid + "/modify"
	//req, err := http.NewRequest("POST", "http://smf:29502/nsmf-pdusession/v1/sm-contexts/urn:uuid:553162e0-3ecb-43b8-b74e-04c21ca6f12a/modify", r)
	req, err := http.NewRequest("POST", "http://smf:29502"+url, r)
	if err != nil {
		log.Fatalln(err)
	}

	c.makeHeader(req, contentType)

	resp := network.Send(req)
	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(string(bytes))
	return true
}
