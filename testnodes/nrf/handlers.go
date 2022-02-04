// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package nrf

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/free5gc/http_wrapper"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/gin-gonic/gin"
)

var (
	ChanSmfRegistered chan bool
)

func init() {
	ChanSmfRegistered = make(chan bool, 1)
}

type nrfPlmnList struct {
	PlmnList []models.PlmnId `json:"plmnList,omitempty" yaml:"plmnList" bson:"plmnList" mapstructure:"PlmnList"`
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Index")
}

func HTTPUpdateNFInstance(c *gin.Context) {
}

func HTTPRegisterNFInstance(c *gin.Context) {
	var nfprofile models.NfProfile
	var locationHeader string
	requestBody, _ := c.GetRawData()
	// step 2: convert requestBody to openapi models
	openapi.Deserialize(&nfprofile, requestBody, "application/json")

	plmnJson := `{"plmnList":[{"mcc":"208","mnc":"93"}]}`
	var plmnList nrfPlmnList
	err := json.Unmarshal([]byte(plmnJson), &plmnList)
	if err != nil {
		log.Println("PLMN JSON Decode Error: ", err.Error())
	}
	nfprofile.PlmnList = &plmnList.PlmnList

	// set nfprofile location
	locationHeader = "http://nrf:29510/nnrf-nfm/v1/nf-instances/" + nfprofile.NfInstanceId
	header := make(http.Header)
	header.Add("Location", locationHeader)
	httpResponse := http_wrapper.NewResponse(http.StatusCreated, header, nfprofile)
	for key, val := range httpResponse.Header {
		c.Header(key, val[0])
	}

	responseBody, _ := openapi.Serialize(httpResponse.Body, "application/json")
	c.Data(httpResponse.Status, "application/json", responseBody)

	//Send Trigger for AMF test suite start
	ChanSmfRegistered <- true
}

func HTTPDeregisterNFInstance(c *gin.Context) {
}

func HTTPGetNFInstance(c *gin.Context) {
}

func HTTPGetNFInstances(c *gin.Context) {
}

func HTTPSearchNFInstances(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Query = c.Request.URL.Query()
	queryParameters := req.Query
	if queryParameters["target-nf-type"] != nil && queryParameters["requester-nf-type"] != nil {
		var httpResponse *http_wrapper.Response
		if queryParameters["target-nf-type"][0] == "UDM" {
			httpResponse = getUdmDiscInfo()
		} else if queryParameters["target-nf-type"][0] == "PCF" {
			httpResponse = getPcfDiscInfo()
		} else if queryParameters["target-nf-type"][0] == "AMF" {
			httpResponse = getAmfDiscInfo()
		}
		responseBody, _ := openapi.Serialize(httpResponse.Body, "application/json")
		c.Data(httpResponse.Status, "application/json", responseBody)

	}
}

func getAmfDiscInfo() *http_wrapper.Response {
	strSearchResult := `{
		"validityPeriod":100,
		"nfInstances":[
		   {
			  "nfInstanceId":"5c599907-63d0-4fcc-90c9-923ac5deb9fd",
			  "nfType":"AMF",
			  "nfStatus":"REGISTERED",
			  "plmnList":[
				 {
					"mcc":"208",
					"mnc":"93"
				 }
			  ],
			  "sNssais":[
				 {
					"sst":1,
					"sd":"010203"
				 }
			  ],
			  "ipv4Addresses":[
				 "127.0.0.1"
			  ],
			  "amfInfo":{
				 "amfSetId":"046",
				 "amfRegionId":"f5",
				 "guamiList":[
					{
					   "plmnId":{
						  "mcc":"208",
						  "mnc":"93"
					   },
					   "amfId":"f511b2"
					}
				 ],
				 "taiList":[
					{
					   "plmnId":{
						  "mcc":"208",
						  "mnc":"93"
					   },
					   "tac":"000001"
					}
				 ]
			  },
			  "nfServices":[
				 {
					"serviceInstanceId":"1",
					"serviceName":"namf-evts",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29518
					   }
					],
					"apiPrefix":"http://amf:29518"
				 },
				 {
					"serviceInstanceId":"2",
					"serviceName":"namf-mt",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29518
					   }
					],
					"apiPrefix":"http://amf:29518"
				 },
				 {
					"serviceInstanceId":"3",
					"serviceName":"namf-loc",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29518
					   }
					],
					"apiPrefix":"http://amf:29518"
				 },
				 {
					"serviceInstanceId":"4",
					"serviceName":"namf-oam",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29518
					   }
					],
					"apiPrefix":"http://amf:29518"
				 },
				 {
					"serviceInstanceId":"0",
					"serviceName":"namf-comm",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29518
					   }
					],
					"apiPrefix":"http://amf:29518"
				 }
			  ]
		   }
		]
	 }`
	jsonSearchResult := strings.ReplaceAll(strSearchResult, " ", "")

	var searchResult models.SearchResult
	json.Unmarshal([]byte(jsonSearchResult), &searchResult)

	return http_wrapper.NewResponse(http.StatusOK, nil, searchResult)
}

func getPcfDiscInfo() *http_wrapper.Response {
	strSearchResult := `{
		"validityPeriod":100,
		"nfInstances":[
		   {
			  "nfInstanceId":"2d3cebcf-f9ad-4b75-8f80-94245109011b",
			  "nfType":"PCF",
			  "nfStatus":"REGISTERED",
			  "plmnList":[
				 {
					"mcc":"208",
					"mnc":"93"
				 }
			  ],
			  "ipv4Addresses":[
				 "127.0.0.1"
			  ],
			  "pcfInfo":{
				 "dnnList":[
					"free5gc",
					"internet"
				 ]
			  },
			  "nfServices":[
				 {
					"serviceInstanceId":"2",
					"serviceName":"npcf-bdtpolicycontrol",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29507
					   }
					],
					"apiPrefix":"http://pcf:29507"
				 },
				 {
					"serviceInstanceId":"3",
					"serviceName":"npcf-policyauthorization",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29507
					   }
					],
					"apiPrefix":"http://pcf:29507",
					"supportedFeatures":"3"
				 },
				 {
					"serviceInstanceId":"4",
					"serviceName":"npcf-eventexposure",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29507
					   }
					],
					"apiPrefix":"http://pcf:29507"
				 },
				 {
					"serviceInstanceId":"5",
					"serviceName":"npcf-ue-policy-control",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29507
					   }
					],
					"apiPrefix":"http://pcf:29507"
				 },
				 {
					"serviceInstanceId":"0",
					"serviceName":"npcf-am-policy-control",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29507
					   }
					],
					"apiPrefix":"http://pcf:29507"
				 },
				 {
					"serviceInstanceId":"1",
					"serviceName":"npcf-smpolicycontrol",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29507
					   }
					],
					"apiPrefix":"http://pcf:29507",
					"supportedFeatures":"3fff"
				 }
			  ]
		   }
		]
	 }`
	jsonSearchResult := strings.ReplaceAll(strSearchResult, " ", "")

	var searchResult models.SearchResult
	json.Unmarshal([]byte(jsonSearchResult), &searchResult)

	return http_wrapper.NewResponse(http.StatusOK, nil, searchResult)
}

func getUdmDiscInfo() *http_wrapper.Response {
	strSearchResult := `{
		"validityPeriod":100,
		"nfInstances":[
		   {
			  "nfInstanceId":"64f21163-c693-4ce6-8983-ec0a056ec2b0",
			  "nfType":"UDM",
			  "nfStatus":"REGISTERED",
			  "plmnList":[
				 {
					"mcc":"208",
					"mnc":"93"
				 }
			  ],
			  "ipv4Addresses":[
				 "127.0.0.1"
			  ],
			  "udmInfo":{
				 
			  },
			  "nfServices":[
				 {
					"serviceInstanceId":"0",
					"serviceName":"nudm-sdm",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29503
					   }
					],
					"apiPrefix":"http://udm:29503"
				 },
				 {
					"serviceInstanceId":"1",
					"serviceName":"nudm-uecm",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29503
					   }
					],
					"apiPrefix":"http://udm:29503"
				 },
				 {
					"serviceInstanceId":"2",
					"serviceName":"nudm-ueau",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29503
					   }
					],
					"apiPrefix":"http://udm:29503"
				 },
				 {
					"serviceInstanceId":"3",
					"serviceName":"nudm-ee",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29503
					   }
					],
					"apiPrefix":"http://udm:29503"
				 },
				 {
					"serviceInstanceId":"4",
					"serviceName":"nudm-pp",
					"versions":[
					   {
						  "apiVersionInUri":"v1",
						  "apiFullVersion":"1.0.0"
					   }
					],
					"scheme":"http",
					"nfServiceStatus":"REGISTERED",
					"ipEndPoints":[
					   {
						  "ipv4Address":"127.0.0.1",
						  "transport":"TCP",
						  "port":29503
					   }
					],
					"apiPrefix":"http://udm:29503"
				 }
			  ]
		   }
		]
	 }`

	jsonSearchResult := strings.ReplaceAll(strSearchResult, " ", "")

	var searchResult models.SearchResult
	json.Unmarshal([]byte(jsonSearchResult), &searchResult)

	return http_wrapper.NewResponse(http.StatusOK, nil, searchResult)

}
