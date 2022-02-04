// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package cxtcreate

// 10 PDU Session Create
import (
	"strconv"
	"time"
)

type Case1 struct {
}

func (c Case1) Execute() bool {
	var status bool
	for count := 1; count <= 10; count++ {
		if status = SendPduSessCreateRequest("imsi-20893000000000"+strconv.Itoa(count), 5); !status {
			return status
		}

		time.Sleep(2 * time.Second)
	}

	return true
}
