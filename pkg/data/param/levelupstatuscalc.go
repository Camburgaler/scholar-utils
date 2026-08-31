package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

// ValidLevelUpStatusCalcIDs is the list of valid LevelUpStatusCalcParam IDs
var ValidLevelUpStatusCalcIDs = []id.Range{
	{Start: 0, End: 8},
}

type (
	// LevelUpStatusCalc is a struct for storing data from LevelUpStatusCalcParam.csv
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
		ResistancePoison               bool
		ResistanceBleed                bool
		ResistancePetrify              bool
		ResistanceCurse                bool
		Agility                        bool
		Poise                          bool

		leftHandWeaponPrimary    bool   //nolint:unused
		leftHandWeaponSecondary  bool   //nolint:unused
		leftHandWeaponTertiary   bool   //nolint:unused
		rightHandWeaponPrimary   bool   //nolint:unused
		rightHandWeaponSecondary bool   //nolint:unused
		rightHandWeaponTertiary  bool   //nolint:unused
		defenseStrike            bool   //nolint:unused
		defenseSlash             bool   //nolint:unused
		defenseThrust            bool   //nolint:unused
		defensePoise             bool   //nolint:unused
		dummy00                  []byte //nolint:unused
		dummy01                  []byte //nolint:unused
	}
)
