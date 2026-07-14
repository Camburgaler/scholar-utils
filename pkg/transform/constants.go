package transform

import "github.com/Camburgaler/scholar-utils/pkg/output"

var (
	noHelmet = output.Armor{
		Equippable: output.Equippable{
			Name:                    "No Helmet",
			AdditiveModifiers:       map[string]float64{},
			MultiplicativeModifiers: map[string]float64{},
		},
		Defenses: output.Defenses{
			Slash:    0,
			Thrust:   0,
			Strike:   0,
			Standard: 0,
		},
		Absorptions: output.Absorptions{
			Magic:     0,
			Lightning: 0,
			Fire:      0,
			Dark:      0,
			Poison:    0,
			Bleed:     0,
			Petrify:   0,
			Curse:     0,
		},
		Poise:  0,
		Weight: 0,
	}
	noChestpiece = output.Armor{
		Equippable: output.Equippable{
			Name:                    "No Chestpiece",
			AdditiveModifiers:       map[string]float64{},
			MultiplicativeModifiers: map[string]float64{},
		},
		Defenses: output.Defenses{
			Slash:    0,
			Thrust:   0,
			Strike:   0,
			Standard: 0,
		},
		Absorptions: output.Absorptions{
			Magic:     0,
			Lightning: 0,
			Fire:      0,
			Dark:      0,
			Poison:    0,
			Bleed:     0,
			Petrify:   0,
			Curse:     0,
		},
		Poise:  0,
		Weight: 0,
	}
	noGauntlets = output.Armor{
		Equippable: output.Equippable{
			Name:                    "No Gauntlets",
			AdditiveModifiers:       map[string]float64{},
			MultiplicativeModifiers: map[string]float64{},
		},
		Defenses: output.Defenses{
			Slash:    0,
			Thrust:   0,
			Strike:   0,
			Standard: 0,
		},
		Absorptions: output.Absorptions{
			Magic:     0,
			Lightning: 0,
			Fire:      0,
			Dark:      0,
			Poison:    0,
			Bleed:     0,
			Petrify:   0,
			Curse:     0,
		},
		Poise:  0,
		Weight: 0,
	}
	noLeggings = output.Armor{
		Equippable: output.Equippable{
			Name:                    "No Leggings",
			AdditiveModifiers:       map[string]float64{},
			MultiplicativeModifiers: map[string]float64{},
		},
		Defenses: output.Defenses{
			Slash:    0,
			Thrust:   0,
			Strike:   0,
			Standard: 0,
		},
		Absorptions: output.Absorptions{
			Magic:     0,
			Lightning: 0,
			Fire:      0,
			Dark:      0,
			Poison:    0,
			Bleed:     0,
			Petrify:   0,
			Curse:     0,
		},
		Poise:  0,
		Weight: 0,
	}
	noRing = output.Ring{
		Equippable: output.Equippable{
			Name:                    "No Ring",
			AdditiveModifiers:       map[string]float64{},
			MultiplicativeModifiers: map[string]float64{},
		},
	}
	// fists = output.Weapon{
	// 	Name: "Fists",
	// 	Requirements: output.ScalingStats{
	// 		Strength:     0,
	// 		Dexterity:    0,
	// 		Intelligence: 0,
	// 		Faith:        0,
	// 	},
	// 	Category: "fists",
	// 	Paired:   false,
	// 	Infusions: map[string]output.Infusion{
	// 		"standard": {
	// 			Name: "Standard",
	// 			BaseDamage: output.Damage{
	// 				Physical:  10,
	// 				Magic:     10,
	// 				Lightning: 10,
	// 				Fire:      10,
	// 				Dark:      10,
	// 				Poison:    800,
	// 				Bleed:     800,
	// 				Petrify:   0,
	// 				Curse:     0,
	// 			},
	// 			Scaling: output.ScalingStats{
	// 				// TODO
	// 				Strength:     0,
	// 				Dexterity:    0,
	// 				Intelligence: 0,
	// 				Faith:        0,
	// 			},
	// 		},
	// 	},
	// }
)
