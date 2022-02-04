// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package cxtcreate

import "log"

var (
	c1 Case1
	c2 Case2
	c3 Case3
	//u2 cxtupdate.Case1
)

var SubsImsiToUuidTable map[string]string

var testcaseTable []bool

func init() {
	SubsImsiToUuidTable = make(map[string]string)
	testcaseTable = []bool{false, true, false, false}
}

func Execute() bool {

	var status bool

	if testcaseTable[1] {
		log.Println("starting create testcase 1")
		status = c1.Execute()
	}

	if testcaseTable[2] {
		log.Println("starting create testcase 2")
		status = c2.Execute()
	}

	if testcaseTable[3] {
		log.Println("starting create testcase 3")
		status = c3.Execute()
	}

	return status
}
