package cxtcreate

// 10 PDU Session Create
import (
	"strconv"
	"time"
)

type Case3 struct {
}

func (c Case3) Execute() bool {
	var status bool
	for count := 1; count <= 10; count++ {
		if status = SendPduSessCreateRequest("imsi-20893000000001"+strconv.Itoa(count), 5); !status {
			return status
		}
		time.Sleep(2 * time.Second)
	}

	return true
}
