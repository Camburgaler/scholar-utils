// Package transform contains functions for transforming data from DS2 params to Scholar JSON
package transform

import (
	emevdParser "github.com/Camburgaler/scholar-utils/pkg/data/emevd/parse"
	paramParser "github.com/Camburgaler/scholar-utils/pkg/data/params/parse"
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
	helmets := map[string]output.Armor{
		"no-helmet": noHelmet,
	}
	chestpieces := map[string]output.Armor{
		"no-chestpiece": noChestpiece,
	}
	gauntlets := map[string]output.Armor{
		"no-gauntlets": noGauntlets,
	}
	leggings := map[string]output.Armor{
		"no-leggings": noLeggings,
	}
	rings := map[string]output.Ring{
		"no-ring": noRing,
	}
	weapons := map[string]output.Weapon{}

	return output.ScholarData{
		Helmets:     helmets,
		Chestpieces: chestpieces,
		Gauntlets:   gauntlets,
		Leggings:    leggings,
		Rings:       rings,
		Weapons:     weapons,
	}, nil
}
