// Package armor contains structs for storing armor data
package armor

import (
	"github.com/Camburgaler/scholar-utils/pkg/params/id"
)

type (
	// Param is a struct for storing data from ArmorParam.csv
	//
	// ArmorParam.csv contains information about equippable armor
	Param struct {
		ID                             id.ID
		Name                           string
		ArmorSetID                     id.ID
		Unk00                          int64
		SlotCategory                   int64
		Unk06                          int64
		Unk07                          int64
		ModelID                        id.ID
		HasGenderedArmor               bool
		Unk08                          int64
		Unk09                          int64
		EquipInterfereID               id.ID
		Unk10                          int64
		ArmorReinforceID               id.ID
		PhysicalDefenceScaling         float64
		Unk11                          float64
		Unk12                          float64
		Unk13                          float64
		StrengthRequirement            int64
		DexterityRequirement           int64
		IntelligenceRequirement        int64
		FaithRequirement               int64
		Weight                         float64
		Durability                     int64
		RepairCost                     int64
		Poise                          float64
		ControlParametersBlendWeight1H int64
		ControlParametersBlendWeight2H int64
		ItemDiscovery                  int64
		Unk15                          int64
		Unk16                          int64
		Unk17                          int64
	}

	// ReinforceParam is a struct for storing data from ArmorReinforceParam.csv
	//
	// ArmorReinforceParam.csv contains information about armor reinforcement
	ReinforceParam struct {
		ID                     id.ID
		Name                   string
		DefenceSlash           int64
		DefenceThrust          int64
		DefenceStrike          int64
		DefenceStandard        int64
		AbsorptionMagic        float64
		AbsorptionLightning    float64
		AbsorptionFire         float64
		AbsorptionDark         float64
		AbsorptionPoison       float64
		AbsorptionBleed        float64
		AbsorptionPetrify      float64
		AbsorptionCurse        float64
		MaxDefenceSlash        int64
		MaxDefenceThrust       int64
		MaxDefenceStrike       int64
		MaxDefenceStandard     int64
		MaxAbsorptionMagic     float64
		MaxAbsorptionLightning float64
		MaxAbsorptionFire      float64
		MaxAbsorptionDark      float64
		MaxAbsorptionPoison    float64
		MaxAbsorptionBleed     float64
		MaxAbsorptionPetrify   float64
		MaxAbsorptionCurse     float64
		MaxReinforcementLevel  int64
		Unk64                  int64
		ReinforceMaterial      id.ID
		Unk6C                  int64
		ReinforceCostID        id.ID
	}
)

var (
	ValidParamIDs = []id.Range{
		{Min: 11010100, Max: 11040103},
		{Min: 11050100, Max: 11370103},
		{Min: 11390100, Max: 13060101},
		{Min: 13061100, Max: 13081101},
		{Min: 13120100, Max: 17550103},
		{Min: 17680100, Max: 17950103},
	}
	ValidReinforceParamIDs = []id.Range{
		{Min: 11010100, Max: 12230103},
		{Min: 12240100, Max: 13060101},
		{Min: 13061100, Max: 13081101},
		{Min: 13120100, Max: 17550103},
		{Min: 17680100, Max: 17830101},
		{Min: 17950100, Max: 17950103},
	}
)

// SlotCategory constants
const (
	SlotCategoryHead  int64 = 1
	SlotCategoryChest int64 = 2
	SlotCategoryHands int64 = 3
	SlotCategoryLegs  int64 = 4
)
