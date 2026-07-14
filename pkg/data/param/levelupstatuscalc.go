package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

type (
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
		ResistanceMagic                bool
		ResistanceFire                 bool
		ResistanceLightning            bool
		ResistanceDark                 bool
		ResistancePoison               bool
		ResistanceBleed                bool
		ResistancePetrify              bool
		ResistanceCurse                bool
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
