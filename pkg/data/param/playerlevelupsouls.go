package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

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
	PlayerLevelUpSoul struct {
		// The unique ID for this PlayerLevelUpSouls (same as the level)
		Level id.ID

		// The name of this PlayerLevelUpSouls ("Level 1", "Level 2", etc.)
		name string

		// The level to which this record corresponds
		//
		// Not necessary since it's always the same as the ID
		soulLevelThreshold int64
		dummy              []byte

		// The increase in souls from the threshold
		//
		// Not necessary since it's always 0
		increasedSoulsFromThreshold int64

		// The number of souls required to level up
		NecessarySouls int64
	}
)
