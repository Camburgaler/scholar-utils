// Package ring contains structs for storing ring data
package ring

import "github.com/Camburgaler/scholar-utils/pkg/data/params/id"

type (
	// Param is a struct for storing data from RingParam.csv
	//
	// RingParam.csv contains information about equippable rings
	Param struct {
		// The unique ID for this ring
		ID id.ID

		// The name of this ring
		Name string

		// The weight of this ring
		Weight float64

		// The durability of this ring
		Durability int64

		// The cost to repair this ring
		RepairCost int64

		// The modifier to item discovery provided by this ring
		ItemDiscovery int64
		dummy0        []byte
	}
)

// ValidParamIDs is a list of valid RingParam IDs
var ValidParamIDs = []id.Range{
	{Start: 40010000, End: 42000000},
}
