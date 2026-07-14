package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

var (
	// ValidWeaponStatsAffectIDs is the list of valid WeaponStatsAffectParam IDs
	ValidWeaponStatsAffectIDs = []id.Range{
		{Start: 1000000, End: 1003159},
		{Start: 1004000, End: 1005039},
		{Start: 1005070, End: 1006019},
		{Start: 1006030, End: 1103109},
		{Start: 1104020, End: 1104039},
		{Start: 1105030, End: 2105039},
		{Start: 2204040, End: 2204049},
		{Start: 3003050, End: 11026009},
		{Start: 11029000, End: 11029009},
		{Start: 11033000, End: 11108009},
		{Start: 11130000, End: 11154009},
	}
)

type (
	// WeaponStatsAffect is a struct for storing data from WeaponStatsAffectParam.csv
	//
	// WeaponStatsAffectParam.csv contains information about the way weapon Damage scales with Player stats
	WeaponStatsAffect struct {
		// The ID of the WeaponStatsAffectParam
		ID id.ID

		// The name of the WeaponStatsAffectParam
		Name string

		// Multiplier applied to the weapon's base attack value during reinforcement
		BaseDamageScaling    float64
		menuBaseValueScaling float64

		// Multiplier applied to the weapon's physical damage scaling from Strength. Tier 0.
		Level00PhysicalByStrength float64

		// Multiplier applied to the weapon's physical damage scaling from Dexterity. Tier 0.
		Level00PhysicalByDexterity float64

		// Multiplier applied to the weapon's magic damage scaling. Tier 0.
		Level00Magic float64

		// Multiplier applied to the weapon's lightning damage scaling. Tier 0.
		Level00Lightning float64

		// Multiplier applied to the weapon's fire damage scaling. Tier 0.
		Level00Fire float64

		// Multiplier applied to the weapon's dark damage scaling. Tier 0.
		Level00Dark float64

		// Multiplier applied to the weapon's poison damage scaling. Tier 0.
		Level00Poison float64

		// Multiplier applied to the weapon's bleed damage scaling. Tier 0.
		Level00Bleed float64

		// Multiplier applied to the weapon's physical damage scaling from the Enchanted infusion. Tier 0.
		Level00PhysicalByEnchant float64

		// Multiplier applied to the weapon's physical damage scaling from Strength. Tier 1.
		Level01PhysicalByStrength float64

		// Multiplier applied to the weapon's physical damage scaling from Dexterity. Tier 1.
		Level01PhysicalByDexterity float64

		// Multiplier applied to the weapon's magic damage scaling. Tier 1.
		Level01Magic float64

		// Multiplier applied to the weapon's lightning damage scaling. Tier 1.
		Level01Lightning float64

		// Multiplier applied to the weapon's fire damage scaling. Tier 1.
		Level01Fire float64

		// Multiplier applied to the weapon's dark damage scaling. Tier 1.
		Level01Dark float64

		// Multiplier applied to the weapon's poison damage scaling. Tier 1.
		Level01Poison float64

		// Multiplier applied to the weapon's bleed damage scaling. Tier 1.
		Level01Bleed float64

		// Multiplier applied to the weapon's physical damage scaling from the Enchanted infusion. Tier 1.
		Level01PhysicalByEnchant float64

		// Multiplier applied to the weapon's physical damage scaling from Strength. Tier 2.
		Level02PhysicalByStrength float64

		// Multiplier applied to the weapon's physical damage scaling from Dexterity. Tier 2.
		Level02PhysicalByDexterity float64

		// Multiplier applied to the weapon's magic damage scaling. Tier 2.
		Level02Magic float64

		// Multiplier applied to the weapon's lightning damage scaling. Tier 2.
		Level02Lightning float64

		// Multiplier applied to the weapon's fire damage scaling. Tier 2.
		Level02Fire float64

		// Multiplier applied to the weapon's dark damage scaling. Tier 2.
		Level02Dark float64

		// Multiplier applied to the weapon's poison damage scaling. Tier 2.
		Level02Poison float64

		// Multiplier applied to the weapon's bleed damage scaling. Tier 2.
		Level02Bleed float64

		// Multiplier applied to the weapon's physical damage scaling from the Enchanted infusion. Tier 2.
		Level02PhysicalByEnchant float64

		// Multiplier applied to the weapon's physical damage scaling from Strength. Tier 3.
		Level03PhysicalByStrength float64

		// Multiplier applied to the weapon's physical damage scaling from Dexterity. Tier 3.
		Level03PhysicalByDexterity float64

		// Multiplier applied to the weapon's magic damage scaling. Tier 3.
		Level03Magic float64

		// Multiplier applied to the weapon's lightning damage scaling. Tier 3.
		Level03Lightning float64

		// Multiplier applied to the weapon's fire damage scaling. Tier 3.
		Level03Fire float64

		// Multiplier applied to the weapon's dark damage scaling. Tier 3.
		Level03Dark float64

		// Multiplier applied to the weapon's poison damage scaling. Tier 3.
		Level03Poison float64

		// Multiplier applied to the weapon's bleed damage scaling. Tier 3.
		Level03Bleed float64

		// Multiplier applied to the weapon's physical damage scaling from the Enchanted infusion. Tier 3.
		Level03PhysicalByEnchant float64

		// Multiplier applied to the weapon's physical damage scaling from Strength. Tier 4.
		Level04PhysicalByStrength float64

		// Multiplier applied to the weapon's physical damage scaling from Dexterity. Tier 4.
		Level04PhysicalByDexterity float64

		// Multiplier applied to the weapon's magic damage scaling. Tier 4.
		Level04Magic float64

		// Multiplier applied to the weapon's lightning damage scaling. Tier 4.
		Level04Lightning float64

		// Multiplier applied to the weapon's fire damage scaling. Tier 4.
		Level04Fire float64

		// Multiplier applied to the weapon's dark damage scaling. Tier 4.
		Level04Dark float64

		// Multiplier applied to the weapon's poison damage scaling. Tier 4.
		Level04Poison float64

		// Multiplier applied to the weapon's bleed damage scaling. Tier 4.
		Level04Bleed float64

		// Multiplier applied to the weapon's physical damage scaling from the Enchanted infusion. Tier 4.
		Level04PhysicalByEnchant float64

		// Multiplier applied to the weapon's physical damage scaling from Strength. Tier 5.
		Level05PhysicalByStrength float64

		// Multiplier applied to the weapon's physical damage scaling from Dexterity. Tier 5.
		Level05PhysicalByDexterity float64

		// Multiplier applied to the weapon's magic damage scaling. Tier 5.
		Level05Magic float64

		// Multiplier applied to the weapon's lightning damage scaling. Tier 5.
		Level05Lightning float64

		// Multiplier applied to the weapon's fire damage scaling. Tier 5.
		Level05Fire float64

		// Multiplier applied to the weapon's dark damage scaling. Tier 5.
		Level05Dark float64

		// Multiplier applied to the weapon's poison damage scaling. Tier 5.
		Level05Poison float64

		// Multiplier applied to the weapon's bleed damage scaling. Tier 5.
		Level05Bleed float64

		// Multiplier applied to the weapon's physical damage scaling from the Enchanted infusion. Tier 5.
		Level05PhysicalByEnchant float64

		// Multiplier applied to the weapon's physical damage scaling from Strength. Tier 6.
		Level06PhysicalByStrength float64

		// Multiplier applied to the weapon's physical damage scaling from Dexterity. Tier 6.
		Level06PhysicalByDexterity float64

		// Multiplier applied to the weapon's magic damage scaling. Tier 6.
		Level06Magic float64

		// Multiplier applied to the weapon's lightning damage scaling. Tier 6.
		Level06Lightning float64

		// Multiplier applied to the weapon's fire damage scaling. Tier 6.
		Level06Fire float64

		// Multiplier applied to the weapon's dark damage scaling. Tier 6.
		Level06Dark float64

		// Multiplier applied to the weapon's poison damage scaling. Tier 6.
		Level06Poison float64

		// Multiplier applied to the weapon's bleed damage scaling. Tier 6.
		Level06Bleed float64

		// Multiplier applied to the weapon's physical damage scaling from the Enchanted infusion. Tier 6.
		Level06PhysicalByEnchant float64

		// Multiplier applied to the weapon's physical damage scaling from Strength. Tier 7.
		Level07PhysicalByStrength float64

		// Multiplier applied to the weapon's physical damage scaling from Dexterity. Tier 7.
		Level07PhysicalByDexterity float64

		// Multiplier applied to the weapon's magic damage scaling. Tier 7.
		Level07Magic float64

		// Multiplier applied to the weapon's lightning damage scaling. Tier 7.
		Level07Lightning float64

		// Multiplier applied to the weapon's fire damage scaling. Tier 7.
		Level07Fire float64

		// Multiplier applied to the weapon's dark damage scaling. Tier 7.
		Level07Dark float64

		// Multiplier applied to the weapon's poison damage scaling. Tier 7.
		Level07Poison float64

		// Multiplier applied to the weapon's bleed damage scaling. Tier 7.
		Level07Bleed float64

		// Multiplier applied to the weapon's physical damage scaling from the Enchanted infusion. Tier 7.
		Level07PhysicalByEnchant float64

		// Multiplier applied to the weapon's physical damage scaling from Strength. Tier 8.
		Level08PhysicalByStrength float64

		// Multiplier applied to the weapon's physical damage scaling from Dexterity. Tier 8.
		Level08PhysicalByDexterity float64

		// Multiplier applied to the weapon's magic damage scaling. Tier 8.
		Level08Magic float64

		// Multiplier applied to the weapon's lightning damage scaling. Tier 8.
		Level08Lightning float64

		// Multiplier applied to the weapon's fire damage scaling. Tier 8.
		Level08Fire float64

		// Multiplier applied to the weapon's dark damage scaling. Tier 8.
		Level08Dark float64

		// Multiplier applied to the weapon's poison damage scaling. Tier 8.
		Level08Poison float64

		// Multiplier applied to the weapon's bleed damage scaling. Tier 8.
		Level08Bleed float64

		// Multiplier applied to the weapon's physical damage scaling from the Enchanted infusion. Tier 8.
		Level08PhysicalByEnchant float64

		// Multiplier applied to the weapon's physical damage scaling from Strength. Tier 9.
		Level09PhysicalByStrength float64

		// Multiplier applied to the weapon's physical damage scaling from Dexterity. Tier 9.
		Level09PhysicalByDexterity float64

		// Multiplier applied to the weapon's magic damage scaling. Tier 9.
		Level09Magic float64

		// Multiplier applied to the weapon's lightning damage scaling. Tier 9.
		Level09Lightning float64

		// Multiplier applied to the weapon's fire damage scaling. Tier 9.
		Level09Fire float64

		// Multiplier applied to the weapon's dark damage scaling. Tier 9.
		Level09Dark float64

		// Multiplier applied to the weapon's poison damage scaling. Tier 9.
		Level09Poison float64

		// Multiplier applied to the weapon's bleed damage scaling. Tier 9.
		Level09Bleed float64

		// Multiplier applied to the weapon's physical damage scaling from the Enchanted infusion. Tier 9.
		Level09PhysicalByEnchant float64

		// Multiplier applied to the weapon's physical damage scaling from Strength. Tier 10.
		Level10PhysicalByStrength float64

		// Multiplier applied to the weapon's physical damage scaling from Dexterity. Tier 10.
		Level10PhysicalByDexterity float64

		// Multiplier applied to the weapon's magic damage scaling. Tier 10.
		Level10Magic float64

		// Multiplier applied to the weapon's lightning damage scaling. Tier 10.
		Level10Lightning float64

		// Multiplier applied to the weapon's fire damage scaling. Tier 10.
		Level10Fire float64

		// Multiplier applied to the weapon's dark damage scaling. Tier 10.
		Level10Dark float64

		// Multiplier applied to the weapon's poison damage scaling. Tier 10.
		Level10Poison float64

		// Multiplier applied to the weapon's bleed damage scaling. Tier 10.
		Level10Bleed float64

		// Multiplier applied to the weapon's physical damage scaling from the Enchanted infusion. Tier 10.
		Level10PhysicalByEnchant float64
	}
)
