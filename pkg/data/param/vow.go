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
		Name string

		requiredContributionLv1 int    //nolint:unused
		requiredContributionLv2 int    //nolint:unused
		requiredContributionLv3 int    //nolint:unused
		padding0                []byte //nolint:unused
	}
)
