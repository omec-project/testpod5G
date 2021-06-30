package udm

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/free5gc/http_wrapper"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Index")
}

func HTTPGetSmData(c *gin.Context) {

	strSubData := `[
	{
	   "singleNssai":{
		  "sst":1,
		  "sd":"010203"
	   },
	   "dnnConfigurations":{
		  "internet":{
			 "pduSessionTypes":{
				"defaultSessionType":"IPV4",
				"allowedSessionTypes":[
				   "IPV4"
				]
			 },
			 "sscModes":{
				"defaultSscMode":"SSC_MODE_1",
				"allowedSscModes":[
				   "SSC_MODE_2",
				   "SSC_MODE_3"
				]
			 },
			 "5gQosProfile":{
				"5qi":9,
				"arp":{
				   "priorityLevel":8,
				   "preemptCap":"",
				   "preemptVuln":""
				},
				"priorityLevel":8
			 },
			 "sessionAmbr":{
				"uplink":"200 Mbps",
				"downlink":"100 Mbps"
			 }
		  },
		  "internet2":{
			 "pduSessionTypes":{
				"defaultSessionType":"IPV4",
				"allowedSessionTypes":[
				   "IPV4"
				]
			 },
			 "sscModes":{
				"defaultSscMode":"SSC_MODE_1",
				"allowedSscModes":[
				   "SSC_MODE_2",
				   "SSC_MODE_3"
				]
			 },
			 "5gQosProfile":{
				"5qi":9,
				"arp":{
				   "priorityLevel":8,
				   "preemptCap":"",
				   "preemptVuln":""
				},
				"priorityLevel":8
			 },
			 "sessionAmbr":{
				"uplink":"200 Mbps",
				"downlink":"100 Mbps"
			 }
		  }
	   }
	}
 ]`
	jsonSubData := strings.ReplaceAll(strSubData, " ", "")

	rspSMSubDataList := make([]models.SessionManagementSubscriptionData, 0, 1)
	json.Unmarshal([]byte(jsonSubData), &rspSMSubDataList)

	rspSMSubDataList[0].DnnConfigurations["internet"].SessionAmbr.Downlink = "100 Mbps"
	rspSMSubDataList[0].DnnConfigurations["internet2"].SessionAmbr.Downlink = "100 Mbps"
	rspSMSubDataList[0].DnnConfigurations["internet"].SessionAmbr.Uplink = "300 Mbps"
	rspSMSubDataList[0].DnnConfigurations["internet2"].SessionAmbr.Uplink = "300 Mbps"

	httpResponse := http_wrapper.NewResponse(http.StatusOK, nil, rspSMSubDataList)
	responseBody, _ := openapi.Serialize(httpResponse.Body, "application/json")
	c.Data(httpResponse.Status, "application/json", responseBody)
}
