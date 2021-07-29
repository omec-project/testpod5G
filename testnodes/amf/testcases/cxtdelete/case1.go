package cxtdelete

import (
	"log"
	"testpod/testnodes/amf/testcases/cxtcreate"
	"time"
)

type Case1 struct {
}

func (c Case1) Execute() bool {
	var status bool
	for supi := range cxtcreate.SubsImsiToUuidTable {
		log.Printf("initiating UP cnx deactivate request for supi[%v]", supi)

		if status = SendPduSessUpCnxDeactivateRequest(supi); !status {
			return status
		}

		log.Printf("initiating PDU Sess release request for supi[%v]", supi)
		if status = SendPduSessReleaseRequest(supi); !status {
			return status
		}

		time.Sleep(2 * time.Second)
	}

	return true
}
