// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package pcf

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/omec-project/http_wrapper"
	"github.com/omec-project/openapi"
	"github.com/omec-project/openapi/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Index")
}

func HTTPSmPoliciesPost(c *gin.Context) {
	var smPolicyContextData models.SmPolicyContextData
	// step 1: retrieve http request body
	requestBody, _ := c.GetRawData()
	openapi.Deserialize(&smPolicyContextData, requestBody, "application/json")
	req := http_wrapper.NewRequest(c.Request, smPolicyContextData)
	request := req.Body.(models.SmPolicyContextData)
	if request.Supi == "" || request.SliceInfo == nil || len(request.SliceInfo.Sd) != 6 {
		//TODO: Throw some error
		panic("Error in PCC Req")
	}

	strPolicyRsp := `{    
		"sessRules":{
		   "SessRuleId-5":{
			  "authSessAmbr":{
				 "uplink":"200 Mbps",
				 "downlink":"100 Mbps"
			  },
			  "authDefQos":{
				 "5qi":9,
				 "arp":{
					"priorityLevel":8,
					"preemptCap":"",
					"preemptVuln":""
				 },
				 "priorityLevel":8
			  },
			  "sessRuleId":"SessRuleId-5"
		   }
		},
		"policyCtrlReqTriggers":[
		   "PLMN_CH",
		   "RES_MO_RE",
		   "AC_TY_CH",
		   "UE_IP_CH",
		   "PS_DA_OFF",
		   "DEF_QOS_CH",
		   "SE_AMBR_CH",
		   "QOS_NOTIF",
		   "RAT_TY_CH"
		],
		"suppFeat":"000f"
	 }`

	jsonPolicyRsp := strings.ReplaceAll(strPolicyRsp, " ", "")

	var policyRsp *models.SmPolicyDecision
	json.Unmarshal([]byte(jsonPolicyRsp), &policyRsp)

	policyRsp.SessRules["SessRuleId-5"].AuthSessAmbr.Downlink = "100 Mbps"
	policyRsp.SessRules["SessRuleId-5"].AuthSessAmbr.Uplink = "200 Mbps"

	httpResponse := http_wrapper.NewResponse(http.StatusCreated, nil, policyRsp)

	responseBody, _ := openapi.Serialize(httpResponse.Body, "application/json")
	c.Data(httpResponse.Status, "application/json", responseBody)

}

func HTTPSmPoliciesSmPolicyIdDeletePost(c *gin.Context) {

}
func HTTPSmPoliciesSmPolicyIDGet(c *gin.Context) {

}
func HTTPSmPoliciesSmPolicyIdUpdatePost(c *gin.Context) {

}
