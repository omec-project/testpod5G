// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package amf

import (
	"log"
	"testpod/testnodes/amf/testcases"
	"testpod/testnodes/nrf"
	"time"
)

func StartTestSuite() {
	log.Println("waiting for SMF Registration...")
	<-nrf.ChanSmfRegistered
	log.Println("SMF registered, starting AMF test suite")
	time.Sleep(time.Second * 2)
	status := true
	if status = testcases.Execute(); !status {
		log.Println("TEST case failed")
	}
}
