package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

var ValidMenuStatsIDs = []id.Range{
	{Start: 1, End: 99},
}

type (
	// MenuStats is a struct for storing data from MenuStats.csv
	//
	// MenuStats.csv contains information about displaying character stats on the menu
	MenuStats struct {
		// The unique ID for this menu stat (1, 2, etc.)
		ID id.ID

		// The name of this menu stat (Level 1, Level 2, etc.)
		Name string

		MaxHP int64

		MaxStamina int64

		Poise int64

		SpellUsesLevel int64

		SpellSlots int64

		CastingTime int64

		AttackPhysicalStrength int64

		AttackPhysicalDexterity int64

		AttackMagic int64

		AttackFire int64

		AttackLightning int64

		AttackDark int64

		AttackPoison int64

		AttackBleed int64

		AttackDurability int64

		Defense int64

		AbsorptionMagic int64

		AbsorptionFire int64

		AbsorptionLightning int64

		AbsorptionDark int64

		AbsorptionPoison int64

		AbsorptionBleed int64

		DefenseDurability int64

		AbsorptionPetrify int64

		AbsorptionCurse int64

		EvasionInvincibilityFrames int64

		ActionSpeed int64

		AttackPhysicalHollow int64
		dummy1               []byte

		MaxEquipLoad float64
	}
)
