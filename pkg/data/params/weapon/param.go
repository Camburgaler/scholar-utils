// Package weapon contains structs for storing weapon data
package weapon

import "github.com/Camburgaler/scholar-utils/pkg/data/params/id"

var (
	WeaponTypes = map[string][]id.ID{
		"Dagger": {
			40,
			41,
			42,
			43,
			44,
			45,
			740,
			741,
			742,
			745,
		},
		"Straight Sword": {
			140,
			141,
			142,
			145,
			510,
		},
		"Thrusting Sword": {
			60,
			61,
			62,
			63,
			630,
		},
		"Curved Sword": {
			120,
			121,
			122,
		},
		"Katana": {
			130,
			132,
			133,
			590,
		},
		"Greatsword": {
			160,
			161,
			162,
			163,
			164,
			520,
			710,
		},
		"Axe": {
			270,
			271,
			272,
			275,
			277,
		},
		"Scythe": {
			150,
			151,
		},
		"Greataxe": {
			281,
			285,
			286,
			287,
		},
		"Hammer": {
			200,
			201,
			202,
			203,
			204,
			205,
			610,
		},
		"Great Hammer": {
			240,
			241,
			246,
			690,
		},
		"Spear": {
			100,
			101,
			620,
			641,
			700,
		},
		"Lance": {
			110,
		},
		"Halberd": {
			290,
			680,
		},
		"Fists": {
			10,
			30,
			600,
		},
		"Claw": {
			20,
			21,
			22,
		},
		"Whip": {
			70,
		},
		"Staff": {
			330,
			331,
			334,
			335,
			336,
			337,
			338,
			339,
			340,
			341,
			342,
			343,
			730,
		},
		"Chime": {
			350,
			351,
			352,
			353,
			354,
			355,
			356,
			357,
			358,
			650,
		},
		"Bow": {
			360,
			361,
			362,
			363,
			364,
			365,
		},
		"Greatbow": {
			370,
			540,
			550,
		},
		"Crossbow": {
			380,
			381,
			385,
			386,
			560,
			570,
			580,
		},
		"Curved Greatsword": {
			180,
		},
		"Ultra Greatsword": {
			170,
			171,
			172,
			530,
			670,
		},
		"Twinblade": {
			90,
			91,
			92,
			500,
		},
		"Hands": {
			11,
		},
		"Torch": {
			470,
		},
		"Small Shield": {
			410,
			420,
		},
		"Medium Shield": {
			430,
			431,
			440,
		},
		"Greatshield": {
			460,
			465,
			466,
			461,
		},
	}

	WeaponCategories = map[id.ID]string{
		0: "Weapon",
	}
)

type (
	// Param is a struct for storing data from WeaponParam.csv
	//
	// WeaponParam.csv contains information about equippable weapons
	Param struct {
		// The unique ID for this weapon
		ID id.ID

		// The name of this weapon
		Name          string
		weaponID      id.ID
		weaponModelID id.ID

		// The ID of the WeaponReinforceParam that describes the weapon's reinforcement
		WeaponReinforceID      id.ID
		weaponActionCategoryID id.ID

		// The type of this weapon (see WeaponTypes)
		WeaponTypeID       id.ID
		isAttackAutoHoming bool

		// The category of this weapon (see WeaponCategories)
		WeaponCategory    int64
		clothCollidableID id.ID

		// The minimum required strength to use this weapon
		RequiredStrength int64

		// The minimum required dexterity to use this weapon
		RequiredDexterity int64

		// The minimum required intelligence to use this weapon
		RequiredIntelligence int64

		// The minimum required faith to use this weapon
		RequiredFaith int64

		// The weight of this weapon
		Weight              float64
		adjustBalanceWeight float64

		// The durability of this weapon
		Durability int64

		// The cost to repair this weapon
		RepairCost                   int64
		materialID                   id.ID
		staminaConsumptionMelee      int64
		staminaConsumptionRanged     int64
		staminaConsumptionID         id.ID
		lockonAvailableDistanceScale int64
		grabMotionRate               int64
		grabRotationSpeed            int64
		grabRotationMaxSpeed         int64
		backstabDamageSmallID        id.ID
		backstabDamageMediumID       id.ID
		backstabDamageLargeID        id.ID
		guardbreakDamageSmallID      id.ID
		guardbreakDamageMediumID     id.ID
		guardbreakDamageLargeID      id.ID
		parryDamageSmallID           id.ID
		parryDamageMediumID          id.ID
		parryDamageLargeID           id.ID
		parryFramesDuration          float64
		brokenSFXID                  id.ID
		brokenSoundType              int64

		// Multiplier applied to all damage
		DamageScale float64

		// Multiplier applied to all stamina depletion inflicted
		StaminaDamageScale float64

		// Multiplier applied to all durability lost per hit
		DurabilityDamageScale float64

		// Multiplier applied to the weapon's ability to ignore poise
		IgnorePoiseRateScale float64

		// Multiplier applied to the weapon's shield/defence bypass
		IgnoreDamageCutRateScale float64

		// Multiplier applied to the weapon's ability to apply stun buildup
		StunDamageScale   float64
		hitboxRadius      float64
		hitboxLength      float64
		hitbackRadius     float64
		hitbackLength     float64
		damageTypeMenu_1  int64
		damageTypeMenu_2  int64
		poiseDamageMenu   int64
		counterDamageMenu int64
		castingSpeedMenu  int64

		// Adjusts poise damage against players
		PoiseDamagePVP float64

		// Adjusts poise damage against enemies
		PoiseDamagePVE float64

		// Determines how often attacks cannot be interrupted by poise damage
		HyperArmorPoiseRate float64
	}

	// ReinforceParam is a struct for storing data from WeaponReinforceParam.csv
	//
	// WeaponReinforceParam.csv contains information about weapon reinforcement
	ReinforceParam struct {
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

	// StatsAffectParam is a struct for storing data from WeaponStatsAffectParam.csv
	//
	// WeaponStatsAffectParam.csv contains information about the way weapon Damage scales with Player stats
	StatsAffectParam struct {
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

var (
	// ValidParamIDs is the list of valid WeaponParam IDs.
	ValidParamIDs = []id.Range{
		{Start: 1000000, End: 3400000},
		{Start: 3406000, End: 3406000},
		{Start: 3410000, End: 5540000},
		{Start: 11000000, End: 11840000},
	}

	// ValidReinforceParamIDs is the list of valid WeaponReinforceParam IDs
	ValidReinforceParamIDs = []id.Range{
		{Start: 1000, End: 3400},
		{Start: 3406, End: 3406},
		{Start: 3410, End: 5540},
		{Start: 11000, End: 11840},
	}

	// ValidStatsAffectParamIDs is the list of valid WeaponStatsAffectParam IDs
	ValidStatsAffectParamIDs = []id.Range{
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
