package output

type (
	modifierTargetType string
	modifierMethod     string
)

// Modifier target types
const (
	ModifierTargetTypeAttribute modifierTargetType = "attribute"
	ModifierTargetTypeStat      modifierTargetType = "stat"
	ModifierTargetTypeSpecial   modifierTargetType = "special"
)

// Modifier methods
const (
	ModifierMethodAdditive       modifierMethod = "additive"
	ModifierMethodMultiplicative modifierMethod = "multiplicative"
	ModifierMethodToggle         modifierMethod = "toggle"
)

type (
	// ScalingAttributes is a struct for the attributes that affect the damage of a weapon
	ScalingAttributes[T any] struct {
		Strength     T
		Dexterity    T
		Intelligence T
		Faith        T
	}

	// Attributes is a struct for character Attributes
	Attributes[T any] struct {
		Strength     T
		Dexterity    T
		Intelligence T
		Faith        T
		Vigor        T
		Endurance    T
		Vitality     T
		Adaptability T
		Attunement   T
	}

	// Class is a struct for a starting class
	Class struct {
		Name       string
		Level      int
		Attributes Attributes[int]
	}

	// Modifier is a struct for an item's modifier, also called special effects
	Modifier struct {
		Description string
		TargetType  modifierTargetType
		Target      string
		Method      modifierMethod
		Value       float64
	}

	// Equippable is a struct for the common fields of equippable items
	Equippable struct {
		// Name of the equippable item
		Name string

		// Special effects of the equippable item
		Modifiers []Modifier

		// Weight of the equippable item
		Weight float64

		// Durability of the equippable item
		Durability int

		// Cost to repair the equippable item once durability is depleted
		RepairCost int
	}

	// Defenses is a struct for the defenses of an armor
	Defenses struct {
		Slash     SlopeIntercept
		Thrust    SlopeIntercept
		Strike    SlopeIntercept
		Standard  SlopeIntercept
		Magic     SlopeIntercept
		Lightning SlopeIntercept
		Fire      SlopeIntercept
		Dark      SlopeIntercept
	}

	// Resistances is a struct for the absorptions of an armor
	Resistances struct {
		Poison  SlopeIntercept
		Bleed   SlopeIntercept
		Petrify SlopeIntercept
		Curse   SlopeIntercept
	}

	// Armor is a struct for equippable armor
	Armor struct {
		Equippable

		// The slope-intercept formulae that represent the armor's defenses for each damage type at each reinforcement level
		Defenses               Defenses
		DefenseScalingPhysical float64
		DefenseScalingSlash    float64
		DefenseScalingThrust   float64
		DefenseScalingStrike   float64
		Resistances            Resistances
		Poise                  float64
		Requirements           ScalingAttributes[int]
		ItemDiscovery          int
		MaxReinforcementLevel  int
	}

	// SlopeIntercept is a struct for the slope and intercept of a line
	SlopeIntercept struct {
		Slope     float64
		Intercept float64
	}

	// Damages is a struct for the damages of a weapon
	Damages[T any] struct {
		Physical  T
		Magic     T
		Lightning T
		Fire      T
		Dark      T
		Poison    T
		Bleed     T
		Curse     T
		Petrify   T
	}

	// ScalingFactors is a struct for the scaling factors of a weapon
	ScalingFactors struct {
		PhysicalByStrength  float64
		PhysicalByDexterity float64
		Magic               float64
		Lightning           float64
		Fire                float64
		Dark                float64
		Poison              float64
		Bleed               float64
		PhysicalByEnchant   float64
	}

	// Scaling is a struct for the scaling of a weapon at each reinforcement level
	Scaling struct {
		Level00 ScalingFactors
		Level01 ScalingFactors
		Level02 ScalingFactors
		Level03 ScalingFactors
		Level04 ScalingFactors
		Level05 ScalingFactors
		Level06 ScalingFactors
		Level07 ScalingFactors
		Level08 ScalingFactors
		Level09 ScalingFactors
		Level10 ScalingFactors
	}

	// Infusion is a struct for a weapon infusion
	Infusion struct {
		// The name of the infusion
		Name string

		// The slope-intercept formulae that represent the infusion's base damage for each damage type at each reinforcement level
		Damages Damages[SlopeIntercept]

		// The scaling values for each damage stat at each reinforcement level
		Scaling           Scaling
		DamageRates       Damages[float64]
		BaseDamageScaling float64
	}

	// Weapon is a struct for a weapon
	Weapon struct {
		Equippable

		// The minimum required attributes to wiled this weapon effectively
		Requirements ScalingAttributes[int]

		// The category of this weapon (see WeaponCategories)
		Category string

		// Whether this weapon will be paired when two-handed
		Paired bool

		// Data for each valid infusion on this weapon (includes default infusion)
		Infusions []Infusion

		// The number of times this weapon can be reinforced to increase its stats
		MaxReinforcementLevel int
	}

	// Ring is a struct for equippable rings
	Ring struct {
		Equippable
		ItemDiscovery int
	}

	// Stats is a struct for the stats of a character
	Stats[T any] struct {
		MaximumHP                      T
		MaximumStamina                 T
		MaximumEquipLoad               T
		SpellSlotCount                 T
		SpellCastingSpeed              T
		PhysicalAttackPowerByStrength  T
		PhysicalAttackPowerByDexterity T
		AttackPowerMagic               T
		AttackPowerFire                T
		AttackPowerLightning           T
		AttackPowerDark                T
		AttackPowerPoison              T
		AttackPowerBleed               T
		Defense                        T
		AbsorptionMagic                T
		AbsorptionFire                 T
		AbsorptionLightning            T
		AbsorptionDark                 T
		ResistancePoison               T
		ResistanceBleed                T
		ResistancePetrify              T
		ResistanceCurse                T
		Agility                        T
		Poise                          T
	}

	// Level is a struct for a level
	Level struct {
		Level                  int
		SoulsRequiredToLevelUp int
	}

	// Covenant is a struct for a covenant
	Covenant struct {
		Name string
	}

	// Spell is a struct for spell data
	Spell struct {
		Name                 string
		RequiredIntelligence int
		RequiredFaith        int
		SpellSlotCost        int

		// At certain Attunement breakpoints, spells gain more usages
		// The breakpoints are 15, 26, 32, 38, 43, 49, 58, 79, and 94
		UsageCountCurve []int
	}

	// ScholarData is a struct for Scholar-friendly data
	ScholarData struct {
		Classes     []Class
		Chestpieces []Armor
		Gauntlets   []Armor
		Helmets     []Armor
		Leggings    []Armor
		Weapons     []Weapon
		Rings       []Ring
		Levels      []int
		Covenants   []string
		Spells      []Spell

		// AttributeToStatMap is a map of Attributes to Stats to booleans (i.e. a 2D object)
		//
		// Chose a map for this field because it will be accessed by the Scholar UI in terms of the keys, and not necessarily iterated over.
		AttributeToStatMap Attributes[Stats[bool]]

		// BaseStats is a map of Stats to floats
		//
		// Chose a map for this field because it will be accessed by the Scholar UI in terms of the keys, and not necessarily iterated over.
		BaseStats Stats[float64]

		// StatCalculation is a map of Attributes to Stats to StatCurves (i.e. a 3D object)
		//
		// Chose a map for this field because it will be accessed by the Scholar UI in terms of the keys, and not necessarily iterated over.
		StatCalculation Attributes[Stats[StatCurve]]
	}
)
