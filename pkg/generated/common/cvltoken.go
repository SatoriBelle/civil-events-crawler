// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at 2019-03-20 17:29:20.245579 +0000 UTC
package common

var eventTypesCVLTokenContract = []string{
	"Approval",
	"OwnershipRenounced",
	"OwnershipTransferred",
	"Transfer",
}

func EventTypesCVLTokenContract() []string {
	tmp := make([]string, len(eventTypesCVLTokenContract))
	copy(tmp, eventTypesCVLTokenContract)
	return tmp
}
