package cxtcreate

var (
	c1 Case1
	//c2 cxtcreate.Case2
	//u2 cxtupdate.Case1
)

var SubsImsiToUuidTable map[string]string

func init() {
	SubsImsiToUuidTable = make(map[string]string)
}

func Execute() bool {

	var status bool
	if status = c1.Execute(); !status {
		return status
	}
	return true
}
