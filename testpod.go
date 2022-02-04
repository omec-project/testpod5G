// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
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
		log.Println("Trigger SMF")
		//host SMF test pod Service
		smf.HostServices(":29502")

	} else if os.Args[1] == "amf" {
		log.Println("Trigger AMF")

		//Start test-suite(make it blocking)
		go amf.StartTestSuite()

		//host AMF Service
		amf.HostServices(":29518")
	}
}
