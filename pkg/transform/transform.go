// Package transform contains functions for transforming data from DS2 params to Scholar JSON
package transform

import (
	"strings"

	emevdParser "github.com/Camburgaler/scholar-utils/pkg/data/emevd/parse"
	"github.com/Camburgaler/scholar-utils/pkg/data/param"
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

func transformClasses(playerStatusParams []param.PlayerStatus) []output.Class {
	classes := []output.Class{}

	for _, playerStatusParam := range playerStatusParams {
		if playerStatusParam.IsClass() {
			classes = append(classes, output.Class{
				Name:  strings.TrimPrefix(playerStatusParam.Name, "[Class] "),
				Level: playerStatusParam.Level,
				Stats: output.Attributes[int64]{
					Vigor:        playerStatusParam.Vigor,
					Endurance:    playerStatusParam.Endurance,
					Vitality:     playerStatusParam.Vitality,
					Attunement:   playerStatusParam.Attunement,
					Adaptability: playerStatusParam.Adaptability,
					ScalingAttributes: output.ScalingAttributes[int64]{
						Strength:     playerStatusParam.Strength,
						Dexterity:    playerStatusParam.Dexterity,
						Intelligence: playerStatusParam.Intelligence,
						Faith:        playerStatusParam.Faith,
					},
				},
			})
		}
	}

	return classes
}

// Transform transforms data from DS2 params/EMEVDs to Scholar-friendly data
func Transform(paramData paramParser.DS2Params, emevdData emevdParser.DS2EMEVD) (output.ScholarData, error) {
	classes := transformClasses(paramData.PlayerStatusParam)

	return output.ScholarData{
		Classes: classes,
	}, nil
}
