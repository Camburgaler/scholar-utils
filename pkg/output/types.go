package output

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
		Level      int64
		Attributes Attributes[int64]
	}

	// Equippable is a struct for the common fields of equippable items
	Equippable struct {
		Name                    string //pk
		AdditiveModifiers       map[string]float64
		MultiplicativeModifiers map[string]float64
		Weight                  float64
	}

	// Defenses is a struct for the defenses of an armor
	Defenses struct {
		Slash    int64
		Thrust   int64
		Strike   int64
		Standard int64
	}

	// Absorptions is a struct for the absorptions of an armor
	Absorptions struct {
		Magic     float64
		Lightning float64
		Fire      float64
		Dark      float64
		Poison    float64
		Bleed     float64
		Petrify   float64
		Curse     float64
	}

	// Armor is a struct for equippable armor
	Armor struct {
		Equippable
		Defenses    Defenses
		Absorptions Absorptions
		Poise       float64
	}

	// SlopeIntercept is a struct for the slope and intercept of a line
	SlopeIntercept struct {
		Slope     float64
		Intercept float64
	}

	// Infusion is a struct for a weapon infusion
	Infusion struct {
		Name              string
		DamageUpgradeRate map[string]SlopeIntercept
		StatScalingRate   ScalingAttributes[float64]
	}

	// Weapon is a struct for a weapon
	Weapon struct {
		Name         string
		Requirements ScalingAttributes[int64]
		Category     string
		Paired       bool
		Infusions    map[string]Infusion
	}

	// Ring is a struct for equippable rings
	Ring struct {
		Equippable
	}

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
		Level                  int64
		SoulsRequiredToLevelUp int64
	}

	// Covenant is a struct for a covenant
	Covenant struct {
		Name string
	}

	// // damage is a struct for the damage of a weapon
	// damage struct {
	// 	Physical  int64
	// 	Magic     int64
	// 	Lightning int64
	// 	Fire      int64
	// 	Dark      int64
	// 	Poison    int64
	// 	Bleed     int64
	// 	Petrify   int64
	// 	Curse     int64
	// }

	// ScholarData is a struct for Scholar-friendly data
	ScholarData struct {
		Classes     []Class
		Chestpieces []Armor
		Gauntlets   []Armor
		Helmets     []Armor
		Leggings    []Armor
		Weapons     []Weapon
		Rings       []Ring
		Levels      []Level
		Covenants   []string

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
