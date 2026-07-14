package output

type (
	Defenses struct {
		Slash    int
		Thrust   int
		Strike   int
		Standard int
	}
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
	Equippable struct {
		Name                    string //pk
		AdditiveModifiers       map[string]float64
		MultiplicativeModifiers map[string]float64
	}
	Armor struct {
		Equippable
		Defenses    Defenses
		Absorptions Absorptions
		Poise       float64
		Weight      float64
	}
	Ring struct {
		Equippable
	}
	stats struct {
		Vigor        int
		Endurance    int
		Vitality     int
		Adaptability int
		Strength     int
		Dexterity    int
		Intelligence int
		Faith        int
		Attunement   int
	}
	Damage struct {
		Physical  int64
		Magic     int64
		Lightning int64
		Fire      int64
		Dark      int64
		Poison    int64
		Bleed     int64
		Petrify   int64
		Curse     int64
	}
	ScalingStats struct {
		Strength     float64
		Dexterity    float64
		Intelligence float64
		Faith        float64
	}
	SlopeIntercept struct {
		Slope     float64
		Intercept float64
	}
	Infusion struct {
		Name              string
		DamageUpgradeRate map[string]SlopeIntercept
		StatScalingRate   ScalingStats
	}
	Weapon struct {
		Name         string
		Requirements ScalingStats
		Category     string
		Paired       bool
		Infusions    map[string]Infusion
	}
	class struct {
		Name  string
		Level int
		Stats stats
	}
	ScholarData struct {
		Classes     []class
		Helmets     []Armor
		Chestpieces []Armor
		Gauntlets   []Armor
		Leggings    []Armor
		Rings       []Ring
		Weapons     []Weapon
	}
)
