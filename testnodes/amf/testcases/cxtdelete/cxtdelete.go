// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

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
