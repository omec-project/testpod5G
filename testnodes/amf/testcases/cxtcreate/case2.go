package cxtcreate

//Context Replacements of 5 calls
import (
	"strconv"
	"time"
)

type Case2 struct {
}

func (c Case2) Execute() bool {
	var status bool
	for count := 1; count <= 5; count++ {
		if status = SendPduSessCreateRequest("imsi-20893000000001"+strconv.Itoa(count), 5); !status {
			return status
		}

		time.Sleep(2 * time.Second)
	}

	for count := 1; count <= 5; count++ {
		if status = SendPduSessCreateRequest("imsi-20893000000001"+strconv.Itoa(count), 5); !status {
			return status
		}

		time.Sleep(2 * time.Second)
	}
	return true
}
