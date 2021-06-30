package amf

import (
	"fmt"
	"testpod/testnodes/amf/testcases"
	"time"
)

func StartTestSuite() {
	time.Sleep(time.Second * 5)
	status := true
	if status = testcases.Execute(); !status {
		fmt.Println("TEST case failed")
	}
}

/*
func main() {
	//Start Test cases
	go startTestSuite()

	//Start all dummy interfaces
	//hostServices()
}
*/
/*
type clientData struct {
	ReqType string `json:"reqtype"`
	ReqMsg  string `json:"reqmsg"`
}

func main() {
	dataBody := clientData{ReqType: "testType", ReqMsg: "TestData"}
	fmt.Println("Data body:", dataBody)
	byteData, err := json.Marshal(dataBody)
	if err != nil {
		fmt.Println("Marshal Error: ", err.Error())
	}
	fmt.Println("Json body:", string(byteData))
	r := bytes.NewReader(byteData)
	req, err := http.NewRequest("GET", "http://myserversvc:6000/api/test", r)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bytes))
}


client := http.Client{
    Transport: &http2.Transport{
        // So http2.Transport doesn't complain the URL scheme isn't 'https'
        AllowHTTP: true,
        // Pretend we are dialing a TLS endpoint.
        // Note, we ignore the passed tls.Config
        DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
            return net.Dial(network, addr)
        },
    },
}

resp, _ := client.Get(url)
fmt.Printf("Client Proto: %d\n", resp.ProtoMajor)
*/
