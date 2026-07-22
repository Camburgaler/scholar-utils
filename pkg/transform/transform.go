// Package transform contains functions for transforming data from DS2 params to Scholar JSON
package transform

import (
	"fmt"
	"reflect"
	"strings"

	emevdParser "github.com/Camburgaler/scholar-utils/pkg/data/emevd/parse"
	"github.com/Camburgaler/scholar-utils/pkg/data/param"
	paramParser "github.com/Camburgaler/scholar-utils/pkg/data/param/parse"
	"github.com/Camburgaler/scholar-utils/pkg/output"
)

var (
	Infusions = []string{
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

	levelUpStatusCalcParamIndexToAttribute = map[int]string{
		0: "Vigor",
		1: "Endurance",
		2: "Vitality",
		3: "Attunement",
		4: "Strength",
		5: "Dexterity",
		6: "Adaptability",
		7: "Intelligence",
		8: "Faith",
	}
)

func createClasses(playerStatusParams []param.PlayerStatus) []output.Class {
	classes := []output.Class{}

	fmt.Println("\nCreating classes...")

	for _, playerStatusParam := range playerStatusParams {
		if playerStatusParam.IsClass() {
			classes = append(classes, output.Class{
				Name:  strings.TrimPrefix(playerStatusParam.Name, "[Class] "),
				Level: playerStatusParam.Level,
				Attributes: output.Attributes[int64]{
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

	fmt.Printf("Created %d classes\n", len(classes))

	return classes
}

func createAttributeToStatMap(levelUpStatusCalcParams []param.LevelUpStatusCalc) output.Attributes[output.Stats] {
	attributesToStatMap := output.Attributes[output.Stats]{}
	vAttributesToStatMap := reflect.ValueOf(&attributesToStatMap).Elem()
	fields := reflect.VisibleFields(vAttributesToStatMap.Type())

	fmt.Println("\nCreating attribute to stat map...")

	for i, levelUpStatusCalcParam := range levelUpStatusCalcParams {
		vLevelUpStatusCalcParam := reflect.ValueOf(&levelUpStatusCalcParam).Elem()

		for _, field := range fields {
			if field.Anonymous {
				continue
			}

			if levelUpStatusCalcParamIndexToAttribute[i] == field.Name {
				stats := output.Stats{}
				vStats := reflect.ValueOf(&stats).Elem()

				for field := range vStats.Fields() {
					vStats.FieldByName(field.Name).SetBool(vLevelUpStatusCalcParam.FieldByName(field.Name).Bool())
				}

				vAttributesToStatMap.FieldByName(field.Name).Set(reflect.ValueOf(stats))
			}
		}
	}

	fmt.Println("Created attribute to stat map")

	return attributesToStatMap
}

// Transform transforms data from DS2 params/EMEVDs to Scholar-friendly data
func Transform(paramData paramParser.DS2Params, emevdData emevdParser.DS2EMEVD) (output.ScholarData, error) {
	classes := createClasses(paramData.PlayerStatusParam)
	attributesToStatMap := createAttributeToStatMap(paramData.LevelUpStatusCalcParam)
	rings := []output.Ring{
		noRing,
	}
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

	return output.ScholarData{
		Classes:            classes,
		AttributeToStatMap: attributesToStatMap,
		Rings:              rings,
		Helmets:            helmets,
		Chestpieces:        chestpieces,
		Gauntlets:          gauntlets,
		Leggings:           leggings,
	}, nil
}
