package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

// ValidPlayerLevelUpSoulsIDs is the list of valid PlayerLevelUpSoulsParam IDs
var ValidPlayerLevelUpSoulsIDs = []id.Range{
	{
		Start: 0,
		End:   850,
	},
}

type (
	// PlayerLevelUpSouls is a struct for storing data from PlayerLevelUpSoulsParam.csv
	//
	// PlayerLevelUpSoulsParam.csv contains information about the souls required to level up
	PlayerLevelUpSouls struct {
		// The unique ID for this PlayerLevelUpSouls (same as the level)
		Level id.ID

		name                        string //nolint:unused
		soulLevelThreshold          int    //nolint:unused
		dummy                       []byte //nolint:unused
		increasedSoulsFromThreshold int    //nolint:unused

		// The number of souls required to level up
		NecessarySouls int
	}
)
