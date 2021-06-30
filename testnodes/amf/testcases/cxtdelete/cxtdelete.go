package cxtdelete

var (
	c1 Case1
)

func Execute() bool {

	var status bool

	if status = c1.Execute(); !status {
		return status
	}

	return true
}
