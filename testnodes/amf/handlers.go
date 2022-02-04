// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package amf

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testpod/testnodes/amf/testcases/cxtupdate"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/gin-gonic/gin"
)

func HTTPN1N2MessageTransfer(c *gin.Context) {

	var n1n2MessageTransferRequest models.N1N2MessageTransferRequest
	n1n2MessageTransferRequest.JsonData = new(models.N1N2MessageTransferReqData)

	requestBody, err := c.GetRawData()
	if err != nil {
		panic("N1N2 decode error, " + err.Error())
	}
	contentType := c.GetHeader("Content-Type")
	s := strings.Split(contentType, ";")
	switch s[0] {
	case "application/json":
		err = fmt.Errorf("n1 and n2 datas are both Empty in N1N2MessgeTransfer")
	case "multipart/related":
		err = openapi.Deserialize(&n1n2MessageTransferRequest, requestBody, contentType)
	default:
		err = fmt.Errorf("wrong content type")
	}

	if err != nil {
		panic("N1N2 content error, " + err.Error())
	}

	ueCxtId := c.Params.ByName("ueContextId")
	log.Println("N1N2 RequestURI", c.Request.RequestURI)

	var u2 cxtupdate.Case1

	c.JSON(http.StatusOK, gin.H{"message": "HTTPN1N2MessageTransfer"})
	go u2.Execute(ueCxtId)
}
