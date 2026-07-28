package transform

import "github.com/Camburgaler/scholar-utils/pkg/output"

var (
	noRing = output.Ring{
		Equippable: output.Equippable{
			Name:       "No Ring",
			Modifiers:  []output.Modifier{},
			Weight:     0,
			Durability: 0,
			RepairCost: 0,
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
