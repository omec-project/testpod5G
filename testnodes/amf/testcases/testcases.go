// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package testcases

import (
	"testpod/testnodes/amf/testcases/cxtcreate"
	"testpod/testnodes/amf/testcases/cxtdelete"
)

func Execute() bool {

	var status bool

	if status = cxtcreate.Execute(); !status {
		return status
	}
	/*
		if status = cxtupdate.Execute(); !status {
			return status
		}
	*/
	if status = cxtdelete.Execute(); !status {
		return status
	}

	return true
}
