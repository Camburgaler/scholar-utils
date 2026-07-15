package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

var ValidLevelUpStatusCalcIDs = []id.Range{
	{Start: 0, End: 8},
}

type (
	// LevelUpStatusCalcParam is a struct for storing data from LevelUpStatusCalcParam.csv
	//
	// LevelUpStatusCalcParam.csv contains information about what stats are affected by what attributes when leveling up
	LevelUpStatusCalc struct {
		ID                             id.ID
		Name                           string
		MaximumHP                      bool
		MaximumStamina                 bool
		MaximumEquipLoad               bool
		SpellSlotCount                 bool
		SpellCastingSpeed              bool
		PhysicalAttackPowerByStrength  bool
		PhysicalAttackPowerByDexterity bool
		AttackPowerMagic               bool
		AttackPowerFire                bool
		AttackPowerLightning           bool
		AttackPowerDark                bool
		AttackPowerPoison              bool
		AttackPowerBleed               bool
		Defense                        bool
		AbsorptionMagic                bool
		AbsorptionFire                 bool
		AbsorptionLightning            bool
		AbsorptionDark                 bool
		AbsorptionPoison               bool
		AbsorptionBleed                bool
		AbsorptionPetrify              bool
		AbsorptionCurse                bool
		Agility                        bool
		Poise                          bool
		LeftHandWeaponPrimary          bool
		LeftHandWeaponSecondary        bool
		LeftHandWeaponTertiary         bool
		RightHandWeaponPrimary         bool
		RightHandWeaponSecondary       bool
		RightHandWeaponTertiary        bool
		DefenseStrike                  bool
		DefenseSlash                   bool
		DefenseThrust                  bool
		DefensePoise                   bool
		dummy00                        []byte
		dummy01                        []byte
	}
)
