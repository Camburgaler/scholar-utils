package param

import (
	"github.com/Camburgaler/scholar-utils/pkg/data/param/id"
)

type (
	// Weapon is a struct for storing data from WeaponParam.csv
	//
	// WeaponParam.csv contains information about equippable weapons
	Weapon struct {
		// The unique ID for this weapon
		ID id.ID

		// The name of this weapon
		Name string

		weaponID      id.ID //nolint:unused
		weaponModelID id.ID //nolint:unused

		// The ID of the WeaponReinforceParam that describes the weapon's reinforcement
		WeaponReinforceID id.ID

		weaponActionCategoryID id.ID //nolint:unused

		// The type of this weapon (see WeaponTypes)
		WeaponTypeID id.ID

		isAttackAutoHoming bool //nolint:unused

		// The category of this weapon (see WeaponCategories)
		WeaponCategory int

		clothCollidableID id.ID //nolint:unused

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

		adjustBalanceWeight float64 //nolint:unused

		// The durability of this weapon
		Durability int

		// The cost to repair this weapon
		RepairCost int

		materialID                   id.ID   //nolint:unused
		staminaConsumptionMelee      int     //nolint:unused
		staminaConsumptionRanged     int     //nolint:unused
		staminaConsumptionID         id.ID   //nolint:unused
		lockonAvailableDistanceScale int     //nolint:unused
		grabMotionRate               int     //nolint:unused
		grabRotationSpeed            int     //nolint:unused
		grabRotationMaxSpeed         int     //nolint:unused
		backstabDamageSmallID        id.ID   //nolint:unused
		backstabDamageMediumID       id.ID   //nolint:unused
		backstabDamageLargeID        id.ID   //nolint:unused
		guardbreakDamageSmallID      id.ID   //nolint:unused
		guardbreakDamageMediumID     id.ID   //nolint:unused
		guardbreakDamageLargeID      id.ID   //nolint:unused
		parryDamageSmallID           id.ID   //nolint:unused
		parryDamageMediumID          id.ID   //nolint:unused
		parryDamageLargeID           id.ID   //nolint:unused
		parryFramesDuration          float64 //nolint:unused
		brokenSFXID                  id.ID   //nolint:unused
		brokenSoundType              int     //nolint:unused

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

		hitboxRadius      float64 //nolint:unused
		hitboxLength      float64 //nolint:unused
		hitbackRadius     float64 //nolint:unused
		hitbackLength     float64 //nolint:unused
		damageTypeMenu1   int     //nolint:unused
		damageTypeMenu2   int     //nolint:unused
		poiseDamageMenu   int     //nolint:unused
		counterDamageMenu int     //nolint:unused
		castingSpeedMenu  int     //nolint:unused

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
