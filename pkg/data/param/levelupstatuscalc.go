package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

type (
	// LevelUpStatusCalcParam is a struct for storing data from LevelUpStatusCalcParam.csv
	//
	// LevelUpStatusCalcParam.csv contains information about what stats are affected by what attributes when leveling up
	LevelUpStatusCalcParam struct {
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
		Defence                        bool
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
		DefenceStrike                  bool
		DefenceSlash                   bool
		DefenceThrust                  bool
		DefencePoise                   bool
		dummy00                        []byte
		dummy01                        []byte
	}
)
