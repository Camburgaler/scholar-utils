package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

var (
	WeaponTypes = map[string][]id.ID{
		"Fists": {
			10,
			30,
			600,
		},
		"Hands": {
			11,
		},
		"Claw": {
			20,
			21,
			22,
		},
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
		"Thrusting Sword": {
			60,
			61,
			62,
			63,
			630,
		},
		"Whip": {
			70,
		},
		"Twinblade": {
			90,
			91,
			92,
			500,
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
		"Straight Sword": {
			140,
			141,
			142,
			145,
			510,
		},
		"Scythe": {
			150,
			151,
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
		"Ultra Greatsword": {
			170,
			171,
			172,
			530,
			670,
		},
		"Curved Greatsword": {
			180,
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
		"Axe": {
			270,
			271,
			272,
			275,
			277,
		},
		"Greataxe": {
			281,
			285,
			286,
			287,
		},
		"Halberd": {
			290,
			680,
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
		"Torch": {
			470,
		},
	}

	WeaponCategories = map[id.ID]string{
		0: "Weapon",
		1: "Bow",
		2: "Crossbow",
		3: "Shield",
	}
)

type (
	// Weapon is a struct for storing data from WeaponParam.csv
	//
	// WeaponParam.csv contains information about equippable weapons
	Weapon struct {
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

		// Multiplier applied to the weapon's shield/defense bypass
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
)

var (
	// ValidWeaponIDs is the list of valid WeaponParam IDs.
	ValidWeaponIDs = []id.Range{
		{Start: 1000000, End: 3400000},
		{Start: 3406000, End: 3406000},
		{Start: 3410000, End: 5540000},
		{Start: 11000000, End: 11840000},
	}
)
