// Package transform contains functions for transforming data from DS2 params to Scholar JSON
package transform

import (
	emevdParser "github.com/Camburgaler/scholar-utils/pkg/data/emevd/parse"
	paramParser "github.com/Camburgaler/scholar-utils/pkg/data/param/parse"
	"github.com/Camburgaler/scholar-utils/pkg/output"
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

func Transform(paramData paramParser.DS2Params, emevdData emevdParser.DS2EMEVD) (output.ScholarData, error) {
	helmets := []output.Armor{
		noHelmet,
	}
	chestpieces := []output.Armor{
		noChestpiece,
	}
	gauntlets := []output.Armor{
		noGauntlets,
	}
	leggings := []output.Armor{
		noLeggings,
	}
	rings := []output.Ring{
		noRing,
	}
	weapons := []output.Weapon{}

	return output.ScholarData{
		Helmets:     helmets,
		Chestpieces: chestpieces,
		Gauntlets:   gauntlets,
		Leggings:    leggings,
		Rings:       rings,
		Weapons:     weapons,
	}, nil
}
