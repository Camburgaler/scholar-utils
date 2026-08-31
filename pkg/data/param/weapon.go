package param

import (
	"github.com/Camburgaler/scholar-utils/pkg/data/param/id"
)

var WeaponCategories = map[id.ID]string{
	0: "Weapon",
	1: "Bow",
	2: "Crossbow",
	3: "Shield",
}

type (
	// Weapon is a struct for storing data from WeaponParam.csv
	//
	// WeaponParam.csv contains information about equippable weapons
	Weapon struct {
		// The unique ID for this weapon
		ID id.ID

		// The name of this weapon
		Name string

		weaponID      id.ID
		weaponModelID id.ID

		// The ID of the WeaponReinforceParam that describes the weapon's reinforcement
		WeaponReinforceID id.ID

		weaponActionCategoryID id.ID

		// The type of this weapon (see WeaponTypes)
		WeaponTypeID id.ID

		isAttackAutoHoming bool

		// The category of this weapon (see WeaponCategories)
		WeaponCategory int

		clothCollidableID id.ID

		// The minimum required strength to use this weapon
		RequiredStrength int

		// The minimum required dexterity to use this weapon
		RequiredDexterity int

		// The minimum required intelligence to use this weapon
		RequiredIntelligence int

		// The minimum required faith to use this weapon
		RequiredFaith int

		// The weight of this weapon
		Weight float64

		adjustBalanceWeight float64

		// The durability of this weapon
		Durability int

		// The cost to repair this weapon
		RepairCost int

		materialID                   id.ID
		staminaConsumptionMelee      int
		staminaConsumptionRanged     int
		staminaConsumptionID         id.ID
		lockonAvailableDistanceScale int
		grabMotionRate               int
		grabRotationSpeed            int
		grabRotationMaxSpeed         int
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
		brokenSoundType              int

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
		StunDamageScale float64

		hitboxRadius      float64
		hitboxLength      float64
		hitbackRadius     float64
		hitbackLength     float64
		damageTypeMenu_1  int
		damageTypeMenu_2  int
		poiseDamageMenu   int
		counterDamageMenu int
		castingSpeedMenu  int

		// Adjusts poise damage against players
		PoiseDamagePVP float64

		// Adjusts poise damage against enemies
		PoiseDamagePVE float64

		// Determines how often attacks cannot be interrupted by poise damage
		HyperArmorPoiseRate float64
	}
)

// ValidWeaponIDs is the list of valid WeaponParam IDs.
var ValidWeaponIDs = []id.Range{
	{Start: 1000000, End: 3400000},
	{Start: 3406000, End: 3406000},
	{Start: 3410000, End: 5540000},
	{Start: 11000000, End: 11840000},
}
