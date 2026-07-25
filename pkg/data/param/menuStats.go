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

		MaxHP int

		MaxStamina int

		Poise int

		SpellUsesLevel int

		SpellSlots int

		CastingTime int

		AttackPhysicalStrength int

		AttackPhysicalDexterity int

		AttackMagic int

		AttackFire int

		AttackLightning int

		AttackDark int

		AttackPoison int

		AttackBleed int

		AttackDurability int

		Defense int

		AbsorptionMagic int

		AbsorptionFire int

		AbsorptionLightning int

		AbsorptionDark int

		AbsorptionPoison int

		AbsorptionBleed int

		DefenseDurability int

		AbsorptionPetrify int

		AbsorptionCurse int

		EvasionInvincibilityFrames int

		ActionSpeed int

		AttackPhysicalHollow int
		dummy1               []byte

		MaxEquipLoad float64
	}
)
