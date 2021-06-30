package cxtdelete

import (
	"log"
	"testpod/testnodes/amf/testcases/cxtcreate"
)

type Case1 struct {
}

func (c Case1) Execute() bool {
	var status bool
	for supi := range cxtcreate.SubsImsiToUuidTable {
		log.Printf("initiating UP cnx deactivate request")

		if status = SendPduSessUpCnxDeactivateRequest(supi); !status {
			return status
		}

		log.Printf("initiating PDU Sess release request")
		if status = SendPduSessReleaseRequest(supi); !status {
			return status
		}
	}

	return true
}
