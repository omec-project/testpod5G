package cxtcreate

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

		time.Sleep(5 * time.Second)
	}
	time.Sleep(10 * time.Second)
	return true
}
