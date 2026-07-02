// Package weapon contains structs for storing weapon data
package weapon

import "github.com/Camburgaler/scholar-utils/pkg/params/id"

type (
	// Param is a struct for storing data from WeaponParam.csv
	//
	// WeaponParam.csv contains information about equippable weapons
	Param struct {
		ID                             id.ID
		Name                           string
		ItemID                         id.ID
		WeaponModelID                  id.ID
		WeaponReinforceID              id.ID
		WeaponActionCategoryID         id.ID
		WeaponTypeID                   id.ID
		Unk00_1                        int64
		Unk00_2                        int64
		Unk00_3                        int64
		Unk00_4                        int64
		StrRequirement                 int64
		DexRequirement                 int64
		IntRequirement                 int64
		FthRequirement                 int64
		Weight                         float64
		RecoveryAnimation_weight       float64
		MaxDurability                  int64
		BaseRepairCost                 int64
		Unk01                          int64
		StaminaConsumption             int64
		StaminaDamage                  int64
		WeaponStaminaCostID            id.ID
		Unk02                          int64
		Unk03                          int64
		Unk04                          int64
		Unk05                          int64
		PlayerDamageIDBackstabSmall    int64
		PlayerDamageIDBackstabMedium   int64
		PlayerDamageIDBackstabLarge    int64
		PlayerDamageIDGuardbreakSmall  int64
		PlayerDamageIDGuardbreakMedium int64
		PlayerDamageIDGuardbreakLarge  int64
		PlayerDamageIDRiposteSmall     int64
		PlayerDamageIDRiposteMedium    int64
		PlayerDamageIDRiposteLarge     int64
		ParryFramesDuration            float64
		Unk06                          int64
		Unk07                          int64
		DamageMult                     float64
		StaminaDamageMult              float64
		DurabilityDamageMult           float64
		GuardBreakField                float64
		StatusEffectAmount             float64
		PostureDamageMult              float64
		HitboxRadius                   float64
		HitboxLength                   float64
		HitbackRadius                  float64
		HitbackLength                  float64
		DamageTypeMenu_1               int64
		DamageTypeMenu_2               int64
		PoiseDamageMenu                int64
		CounterDamageMenu              int64
		CastingSpeedMenu               int64
		PoiseDamagePVP                 float64
		PoiseDamagePVE                 float64
		HyperArmmorPoiseMult           float64
	}

	// ReinforceParam is a struct for storing data from WeaponReinforceParam.csv
	//
	// WeaponReinforceParam.csv contains information about weapon reinforcement
	ReinforceParam struct {
		ID                    id.ID
		Name                  string
		DamagePhysical        int64
		DamageMagic           int64
		DamageLightning       int64
		DamageFire            int64
		DamageDark            int64
		DamagePoison          int64
		DamageBleed           int64
		DamagePetrify         int64
		DamageCurse           int64
		MaxDamagePhysical     int64
		MaxDamageMagic        int64
		MaxDamageLightning    int64
		MaxDamageFire         int64
		MaxDamageDark         int64
		MaxDamagePoison       int64
		MaxDamageBleed        int64
		MaxDamagePetrify      int64
		MaxDamageCurse        int64
		MaxReinforcementLevel int64
		WeaponStatsAffectID   id.ID
		AbsorptionPhysical    float64
		AbsorptionMagic       float64
		AbsorptionFire        float64
		AbsorptionLightning   float64
		AbsorptionDark        float64
		AbsorptionPoison      float64
		AbsorptionBleed       float64
		AbsorptionPetrify     float64
		AbsorptionCurse       float64
		StabilityZero         float64
		StabilityOne          float64
		StabilityTwo          float64
		StabilityThree        float64
		StabilityFour         float64
		StabilityFive         float64
		StabilitySix          float64
		StabilitySeven        float64
		StabilityEight        float64
		StabilityNine         float64
		StabilityTen          float64
		DamageMultPhysical    float64
		DamageMultMagic       float64
		DamageMultLightning   float64
		DamageMultFire        float64
		DamageMultDark        float64
		DamageMultPoison      float64
		DamageMultBleed       float64
		DamageMultPetrify     float64
		DamageMultCurse       float64
		InfusionMultMagic     float64
		InfusionMultLightning float64
		InfusionMultFire      float64
		InfusionMultDark      float64
		InfusionMultPoison    float64
		InfusionMultBleed     float64
		InfusionMultRaw       float64
		InfusionMultEnchanted float64
		InfusionMultMundane   float64
		CustomAttrSpecParamID id.ID
		CustomAttrCostParamID id.ID
		ReinforceCostID       id.ID
	}

	// StatsAffectParam is a struct for storing data from WeaponStatsAffectParam.csv
	//
	// WeaponStatsAffectParam.csv contains information about the way weapon Damage scales with Player stats
	StatsAffectParam struct {
		ID                        id.ID
		Name                      string
		BaseDamageMult            float64
		Unk00                     float64
		Level00PhysicalScalingStr float64
		Level00PhysicalScalingDex float64
		Level00MagicScaling       float64
		Level00LightningScaling   float64
		Level00FireScaling        float64
		Level00DarkScaling        float64
		Level00PoisonScaling      float64
		Level00BleedScaling       float64
		Level00EnchantedScaling   float64
		Level01PhysicalScalingStr float64
		Level01PhysicalScalingDex float64
		Level01MagicScaling       float64
		Level01LightningScaling   float64
		Level01FireScaling        float64
		Level01DarkScaling        float64
		Level01PoisonScaling      float64
		Level01BleedScaling       float64
		Level01EnchantedScaling   float64
		Level02PhysicalScalingStr float64
		Level02PhysicalScalingDex float64
		Level02MagicScaling       float64
		Level02LightningScaling   float64
		Level02FireScaling        float64
		Level02DarkScaling        float64
		Level02PoisonScaling      float64
		Level02BleedScaling       float64
		Level02EnchantedScaling   float64
		Level03PhysicalScalingStr float64
		Level03PhysicalScalingDex float64
		Level03MagicScaling       float64
		Level03LightningScaling   float64
		Level03FireScaling        float64
		Level03DarkScaling        float64
		Level03PoisonScaling      float64
		Level03BleedScaling       float64
		Level03EnchantedScaling   float64
		Level04PhysicalScalingStr float64
		Level04PhysicalScalingDex float64
		Level04MagicScaling       float64
		Level04LightningScaling   float64
		Level04FireScaling        float64
		Level04DarkScaling        float64
		Level04PoisonScaling      float64
		Level04BleedScaling       float64
		Level04EnchantedScaling   float64
		Level05PhysicalScalingStr float64
		Level05PhysicalScalingDex float64
		Level05MagicScaling       float64
		Level05LightningScaling   float64
		Level05FireScaling        float64
		Level05DarkScaling        float64
		Level05PoisonScaling      float64
		Level05BleedScaling       float64
		Level05EnchantedScaling   float64
		Level06PhysicalScalingStr float64
		Level06PhysicalScalingDex float64
		Level06MagicScaling       float64
		Level06LightningScaling   float64
		Level06FireScaling        float64
		Level06DarkScaling        float64
		Level06PoisonScaling      float64
		Level06BleedScaling       float64
		Level06EnchantedScaling   float64
		Level07PhysicalScalingStr float64
		Level07PhysicalScalingDex float64
		Level07MagicScaling       float64
		Level07LightningScaling   float64
		Level07FireScaling        float64
		Level07DarkScaling        float64
		Level07PoisonScaling      float64
		Level07BleedScaling       float64
		Level07EnchantedScaling   float64
		Level08PhysicalScalingStr float64
		Level08PhysicalScalingDex float64
		Level08MagicScaling       float64
		Level08LightningScaling   float64
		Level08FireScaling        float64
		Level08DarkScaling        float64
		Level08PoisonScaling      float64
		Level08BleedScaling       float64
		Level08EnchantedScaling   float64
		Level09PhysicalScalingStr float64
		Level09PhysicalScalingDex float64
		Level09MagicScaling       float64
		Level09LightningScaling   float64
		Level09FireScaling        float64
		Level09DarkScaling        float64
		Level09PoisonScaling      float64
		Level09BleedScaling       float64
		Level09EnchantedScaling   float64
		Level10PhysicalScalingStr float64
		Level10PhysicalScalingDex float64
		Level10MagicScaling       float64
		Level10LightningScaling   float64
		Level10FireScaling        float64
		Level10DarkScaling        float64
		Level10PoisonScaling      float64
		Level10BleedScaling       float64
		Level10EnchantedScaling   float64
	}
)

var (
	ValidParamIDs = []id.Range{
		{Min: 1000000, Max: 3400000},
		{Min: 3406000, Max: 3406000},
		{Min: 3410000, Max: 5540000},
		{Min: 11000000, Max: 11840000},
	}
	ValidReinforceParamIDs = []id.Range{
		{Min: 1000, Max: 3400},
		{Min: 3406, Max: 3406},
		{Min: 3410, Max: 5540},
		{Min: 11000, Max: 11840},
	}
	ValidStatsAffectParamIDs = []id.Range{
		{Min: 1000000, Max: 1003159},
		{Min: 1004000, Max: 1005039},
		{Min: 1005070, Max: 1006019},
		{Min: 1006030, Max: 1103109},
		{Min: 1104020, Max: 1104039},
		{Min: 1105030, Max: 2105039},
		{Min: 2204040, Max: 2204049},
		{Min: 3003050, Max: 11026009},
		{Min: 11029000, Max: 11029009},
		{Min: 11033000, Max: 11108009},
		{Min: 11130000, Max: 11154009},
	}
)
