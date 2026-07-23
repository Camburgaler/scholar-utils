package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

// ValidVowIDs is a list of valid VowParam IDs
var ValidVowIDs = []id.Range{
	{Start: 0, End: 9},
}

type (
	// Vow is a struct for storing data from VowParam.csv
	//
	// VowParam.csv contains information about Vows
	Vow struct {
		// The unique ID for this vow
		ID id.ID

		// The name of this vow
		Name                    string
		requiredContributionLv1 int64
		requiredContributionLv2 int64
		requiredContributionLv3 int64
		padding0                []byte
	}
)
