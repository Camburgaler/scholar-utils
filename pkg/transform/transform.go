// Package transform contains functions for transforming data from DS2 params to Scholar JSON
package transform

import (
	emevdParser "github.com/Camburgaler/scholar-utils/pkg/data/emevd/parse"
	paramParser "github.com/Camburgaler/scholar-utils/pkg/data/params/parse"
	"github.com/Camburgaler/scholar-utils/pkg/output"
)

// Weapon infusions
const (
	infusionStandard = iota
	infusionLightning
	infusionMagic
	infusionFire
	infusionDark
	infusionPoison
	infusionBleed
	infusionRaw
	infusionEnchanted
	infusionMundane
	infusionBoss
	infusionSpecial
)

var Infusions = []string{
	"Standard",
	"Lightning",
	"Magic",
	"Fire",
	"Dark",
	"Poison",
	"Bleed",
	"Raw",
	"Enchanted",
	"Mundane",
	"Boss",
	"Special",
}

// Weapon categories
const (
	dagger = iota
	straightSword
	greatsword
	ultraGreatsword
	curvedSword
	katana
	curvedGreatsword
	piercingSword
	axe
	greatAxe
	hammer
	greatHammer
	fist
	claw
	spear
	halberd
	lance
	reaper
	twinblade
	whip
	bow
	greatbow
	crossbow
	staff
	pyromancyFlame
	chime
)

func Transform(paramData paramParser.DS2Params, emevdData []emevdParser.Event) (output.ScholarData, error) {
	helmets := map[string]output.Armor{
		"no-helmet": {
			Name: "No Helmet",
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
		},
	}
	chestpieces := map[string]output.Armor{
		"no-chestpiece": {
			Name: "No Chestpiece",
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
		},
	}
	gauntlets := map[string]output.Armor{
		"no-gauntlets": {
			Name: "No Gauntlets",
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
		},
	}
	leggings := map[string]output.Armor{
		"no-leggings": {
			Name: "No Leggings",
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
		},
	}
	rings := map[string]output.Ring{
		"no-ring": {
			Name: "No Ring",
		},
	}
	weapons := map[string]output.Weapon{}
	// weapons := map[string]output.Weapon{
	// 	"fists": {
	// 		Name: "Fists",
	// 		Requirements: output.ScalingStats{
	// 			Strength:     0,
	// 			Dexterity:    0,
	// 			Intelligence: 0,
	// 			Faith:        0,
	// 		},
	// 		Category: "fists",
	// 		Paired:   false,
	// 		Infusions: map[string]output.Infusion{
	// 			"standard": {
	// 				Name: "Standard",
	// 				BaseDamage: output.Damage{
	// 					Physical:  10,
	// 					Magic:     10,
	// 					Lightning: 10,
	// 					Fire:      10,
	// 					Dark:      10,
	// 					Poison:    800,
	// 					Bleed:     800,
	// 					Petrify:   0,
	// 					Curse:     0,
	// 				},
	// 				Scaling: output.ScalingStats{
	// 					// TODO
	// 					Strength:     0,
	// 					Dexterity:    0,
	// 					Intelligence: 0,
	// 					Faith:        0,
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	return output.ScholarData{
		Helmets:     helmets,
		Chestpieces: chestpieces,
		Gauntlets:   gauntlets,
		Leggings:    leggings,
		Rings:       rings,
		Weapons:     weapons,
	}, nil
}
