package main

import (
	"fmt"
	"os"
	"testpod/testnodes/amf"
	"testpod/testnodes/nrf"
	"testpod/testnodes/pcf"
	"testpod/testnodes/smf"
	"testpod/testnodes/udm"
)

//var Router *gin.Engine

//Initialize basic framework
func init() {

	//NRF
	go nrf.HostServices(":29510")

	//PCF
	go pcf.HostServices(":29507")

	//UDM
	go udm.HostServices(":29503")

}

//Start desired nod as testpod
func main() {

	if os.Args[1] == "smf" {
		fmt.Println("Trigger SMF")
		//host SMF test pod Service
		smf.HostServices(":29502")

	} else if os.Args[1] == "amf" {
		fmt.Println("Trigger AMF")

		//Start test-suite(make it blocking)
		go amf.StartTestSuite()

		//host AMF Service
		amf.HostServices(":29518")
	}
}
