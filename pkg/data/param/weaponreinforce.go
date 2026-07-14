package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

var (
	// ValidWeaponReinforceIDs is the list of valid WeaponReinforceParam IDs
	ValidWeaponReinforceIDs = []id.Range{
		{Start: 1000, End: 3400},
		{Start: 3406, End: 3406},
		{Start: 3410, End: 5540},
		{Start: 11000, End: 11840},
	}
)

type (
	// WeaponReinforce is a struct for storing data from WeaponReinforceParam.csv
	//
	// WeaponReinforceParam.csv contains information about weapon reinforcement
	WeaponReinforce struct {
		// The unique ID for this reinforcement
		ID id.ID

		// The name of this reinforcement
		Name string

		// The minimum (non-reinforced) physical damage provided by the weapon using this data
		MinPhysical int64

		// The minimum (non-reinforced) magic damage provided by the weapon using this data
		MinMagic int64

		// The minimum (non-reinforced) lightning damage provided by the weapon using this data
		MinLightning int64

		// The minimum (non-reinforced) fire damage provided by the weapon using this data
		MinFire int64

		// The minimum (non-reinforced) dark damage provided by the weapon using this data
		MinDark int64

		// The minimum (non-reinforced) poison damage provided by the weapon using this data
		MinPoison int64

		// The minimum (non-reinforced) bleed damage provided by the weapon using this data
		MinBleed int64

		// The minimum (non-reinforced) petrify damage provided by the weapon using this data
		MinPetrify int64

		// The minimum (non-reinforced) curse damage provided by the weapon using this data
		MinCurse int64

		// The maximum (fully reinforced) physical damage provided by the weapon using this data
		MaxPhysical int64

		// The maximum (fully reinforced) magic damage provided by the weapon using this data
		MaxMagic int64

		// The maximum (fully reinforced) lightning damage provided by the weapon using this data
		MaxLightning int64

		// The maximum (fully reinforced) fire damage provided by the weapon using this data
		MaxFire int64

		// The maximum (fully reinforced) dark damage provided by the weapon using this data
		MaxDark int64

		// The maximum (fully reinforced) poison damage provided by the weapon using this data
		MaxPoison int64

		// The maximum (fully reinforced) bleed damage provided by the weapon using this data
		MaxBleed int64

		// The maximum (fully reinforced) petrify damage provided by the weapon using this data
		MaxPetrify int64

		// The maximum (fully reinforced) curse damage provided by the weapon using this data
		MaxCurse int64

		// The maximum reinforcement level
		MaxReinforcementLevel int64

		dummy []byte

		// The ID of the WeaponStatsAffectParam that defines the stats affect
		WeaponStatsAffectID id.ID

		// Physical absorption
		AbsorptionPhysical float64

		// Magic absorption
		AbsorptionMagic float64

		// Fire absorption
		AbsorptionFire float64

		// Lightning absorption
		AbsorptionLightning float64

		// Dark absorption
		AbsorptionDark float64

		// Poison absorption
		AbsorptionPoison float64

		// Bleed absorption
		AbsorptionBleed float64

		// Petrify absorption
		AbsorptionPetrify float64

		// Curse absorption
		AbsorptionCurse float64

		// Guard stability value at reinforcement level 0
		StabilityRate00 float64

		// Guard stability value at reinforcement level 1
		StabilityRate01 float64

		// Guard stability value at reinforcement level 2
		StabilityRate02 float64

		// Guard stability value at reinforcement level 3
		StabilityRate03 float64

		// Guard stability value at reinforcement level 4
		StabilityRate04 float64

		// Guard stability value at reinforcement level 5
		StabilityRate05 float64

		// Guard stability value at reinforcement level 6
		StabilityRate06 float64

		// Guard stability value at reinforcement level 7
		StabilityRate07 float64

		// Guard stability value at reinforcement level 8
		StabilityRate08 float64

		// Guard stability value at reinforcement level 9
		StabilityRate09 float64

		// Guard stability value at reinforcement level 10
		StabilityRate10 float64

		// Base physical attribute multiplier applied befor elemental modifiers
		PhysicalRate float64

		// Base magic attribute multiplier applied befor elemental modifiers
		MagicRate float64

		// Base lightning attribute multiplier applied befor elemental modifiers
		LightningRate float64

		// Base fire attribute multiplier applied befor elemental modifiers
		FireRate float64

		// Base dark attribute multiplier applied befor elemental modifiers
		DarkRate float64

		// Base poison attribute multiplier applied befor elemental modifiers
		PoisonRate float64

		// Base bleed attribute multiplier applied befor elemental modifiers
		BleedRate float64

		// Base petrify attribute multiplier applied befor elemental modifiers
		PetrifyRate float64

		// Base curse attribute multiplier applied befor elemental modifiers
		CurseRate float64

		// Attribute multiplier applied by magic infusion
		InfusionMagicRate float64

		// Attribute multiplier applied by lightning infusion
		InfusionLightningRate float64

		// Attribute multiplier applied by fire infusion
		InfusionFireRate float64

		// Attribute multiplier applied by dark infusion
		InfusionDarkRate float64

		// Attribute multiplier applied by poison infusion
		InfusionPoisonRate float64

		// Attribute multiplier applied by bleed infusion
		InfusionBleedRate float64

		// Attribute multiplier applied by raw infusion
		InfusionRawRate float64

		// Attribute multiplier applied by enchanted infusion
		InfusionEnchantedRate float64

		// Attribute multiplier applied by mundane infusion
		InfusionMundaneRate float64

		// The ID of the CustomAttributeSpecParam that defines the custom attributes
		CustomAttrSpecParamID id.ID
		customAttrCostParamID id.ID
		reinforceCostID       id.ID
	}
)
