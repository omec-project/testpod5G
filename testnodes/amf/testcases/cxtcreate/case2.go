package cxtcreate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"testpod/network"
)

type Case2 struct {
}

type clientData1 struct {
	ReqType string `json:"reqtype"`
	ReqMsg  string `json:"reqmsg"`
}

func (c Case2) makeHeader(req *http.Request) {
	req.Header.Set("Accept", "application/json")
}

func (c Case2) makeBody() *bytes.Reader {
	dataBody := clientData1{
		ReqType: "ReqType_2",
		ReqMsg:  "ReqData_2"}

	byteData, err := json.Marshal(dataBody)
	if err != nil {
		fmt.Println("Marshal Error: ", err.Error())
	}

	fmt.Println("Json body:", string(byteData))
	return bytes.NewReader(byteData)
}

func (c Case2) Execute() bool {
	r := c.makeBody()
	req, err := http.NewRequest("GET", "http://dummysmfsvc:6000/nsmf-pdusession/v1/", r)
	if err != nil {
		log.Fatalln(err)
	}

	c.makeHeader(req)

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
