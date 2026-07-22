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
		ScalingAttributes[T]
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

	Stats struct {
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
		Covenants   []Covenant

		// AttributeToStatMap is a map of Attributes to Stats to booleans (i.e. a 2D array)
		//
		// Chose a map for this field because it will be accessed by the Scholar UI in terms of the keys, and not necessarily iterated over.
		AttributeToStatMap Attributes[Stats]
	}
)
