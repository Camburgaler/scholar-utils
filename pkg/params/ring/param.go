// Package ring contains structs for storing ring data
package ring

import "github.com/Camburgaler/scholar-utils/pkg/params/id"

type (
	// Param is a struct for storing data from RingParam.csv
	//
	// RingParam.csv contains information about equippable rings
	Param struct {
		ID            id.ID
		Name          string
		Weight        float64
		Durability    int64
		RepairCost    int64
		ItemDiscovery int64
	}
)

var (
	ValidParamIDs = []id.Range{
		{Min: 40010000, Max: 42000000},
	}
)
