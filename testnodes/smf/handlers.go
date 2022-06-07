// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package smf

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

	"github.com/omec-project/openapi"
	"github.com/omec-project/openapi/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Index")
}

func HTTPPostSmContexts(c *gin.Context) {
	var request models.PostSmContextsRequest
	request.JsonData = new(models.SmContextCreateData)

	s := strings.Split(c.GetHeader("Content-Type"), ";")
	fmt.Println("Received Content Type: ", s)
	var err error
	switch s[0] {
	case "application/json":
		err = c.ShouldBindJSON(request.JsonData)
	case "multipart/related":
		err = c.ShouldBindWith(&request, openapi.MultipartRelatedBinding{})
	}

	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}

		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	jsonBody, _ := json.Marshal(request.JsonData)
	fmt.Println("Received JSON body: ", string(jsonBody))

	///Prepare Response
	var response models.PostSmContextsResponse
	createdData := new(models.SmContextCreatedData)
	createdData.SNssai = &models.Snssai{
		Sst: 1,
		Sd:  "010203"}
	response.JsonData = createdData

	c.Header("Location", "urn:uuid:553162e0-3ecb-43b8-b74e-04c21ca6f12a")

	c.Render(http.StatusCreated, openapi.MultipartRelatedRender{Data: response})

	//Trigger N1N2 Msg towards AMF
	go sendN1N2Req()
}

func makeBody(request *models.N1N2MessageTransferRequest) (*bytes.Buffer, string) {
	ueDataJson := `{
		"n1MessageContainer":{
		   "n1MessageClass":"SM",
		   "n1MessageContent":{
			  "contentId":"GSM_NAS"
		   }
		},
		"n2InfoContainer":{
		   "n2InformationClass":"SM",
		   "smInfo":{
			  "pduSessionId":5,
			  "n2InfoContent":{
				 "ngapIeType":"PDU_RES_SETUP_REQ",
				 "ngapData":{
					"contentId":"N2SmInformation"
				 }
			  },
			  "sNssai":{
				 "sst":1,
				 "sd":"010203"
			  }
		   }
		},
		"pduSessionId":5
	 }`
	ueDataStr := strings.ReplaceAll(ueDataJson, " ", "")
	request.JsonData = new(models.N1N2MessageTransferReqData)
	json.Unmarshal([]byte(ueDataStr), request.JsonData)

	//NAS DATA
	smNasBuf := `2e0501c211000901000631310101ff09060600640600c82905010afa0003220401010203790006092041010109250908696e7465726e6574`
	decodedNas, err := hex.DecodeString(smNasBuf)
	if err != nil {
		log.Fatal(err)
	}
	request.BinaryDataN1Message = []byte(decodedNas)

	//NGAP DATA
	n2Pdu := `0000040082000a0c05f5e100300bebc200008b000a01f00a31d7e30000000500860001000088000700090000093800`
	decodedNgap, err := hex.DecodeString(n2Pdu)
	if err != nil {
		log.Fatal(err)
	}
	request.BinaryDataN2Information = []byte(decodedNgap)

	body := &bytes.Buffer{}
	contentType, err := openapi.MultipartEncode(request, body)
	if err != nil {
		log.Fatalln(err)
		return nil, ""
	}

	return body, contentType
}

func makeHeader(req *http.Request, contentType string) {
	req.Header.Set("Accept", "application/json")
	//req.Header.Set("Content-Type", "application/json")
	fmt.Println("Content-Type: ", contentType)
	req.Header.Set("Content-Type", contentType)
}

func sendN1N2Req() bool {
	n1n2Request := models.N1N2MessageTransferRequest{}
	r, contentType := makeBody(&n1n2Request)
	req, err := http.NewRequest("POST", "http://amf:29518/namf-comm/v1/ue-contexts/imsi-208930000000003/n1-n2-messages", r)
	if err != nil {
		log.Fatalln(err)
	}

	makeHeader(req, contentType)

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

func HTTPUpdateSmContext(c *gin.Context) {
	var request models.UpdateSmContextRequest
	request.JsonData = new(models.SmContextUpdateData)

	s := strings.Split(c.GetHeader("Content-Type"), ";")
	var err error
	switch s[0] {
	case "application/json":
		err = c.ShouldBindJSON(request.JsonData)
	case "multipart/related":
		err = c.ShouldBindWith(&request, openapi.MultipartRelatedBinding{})
	}
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		c.JSON(http.StatusBadRequest, rsp)
		log.Print(err)
		return
	}

	jsonBody, _ := json.Marshal(request.JsonData)
	fmt.Println("Received JSON body: ", string(jsonBody))

	//Prepare Response
	var response models.UpdateSmContextResponse
	response.JsonData = new(models.SmContextUpdatedData)

	c.Render(http.StatusOK, openapi.MultipartRelatedRender{Data: response})
	//c.JSON(http.StatusOK, gin.H{"message": "HTTPUpdateSmContext"})
}

func HTTPReleaseSmContext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "HTTPReleaseSmContext"})
}

func RetrieveSmContext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "RetrieveSmContext"})
}

// PostPduSessions - Create
func PostPduSessions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PostPduSessions"})
}

// UpdatePduSession - Update (initiated by V-SMF)
func UpdatePduSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdatePduSession"})
}

func ReleasePduSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ReleasePduSession"})
}
