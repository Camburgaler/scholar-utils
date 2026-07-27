// Package transform contains functions for transforming data from DS2 params to Scholar JSON
package transform

import (
	"fmt"
	"math"
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
				Attributes: output.Attributes[int]{
					Vigor:        playerStatusParam.Vigor,
					Endurance:    playerStatusParam.Endurance,
					Vitality:     playerStatusParam.Vitality,
					Attunement:   playerStatusParam.Attunement,
					Adaptability: playerStatusParam.Adaptability,
					Strength:     playerStatusParam.Strength,
					Dexterity:    playerStatusParam.Dexterity,
					Intelligence: playerStatusParam.Intelligence,
					Faith:        playerStatusParam.Faith,
				},
			})
		}
	}

	fmt.Printf("Created %d classes\n", len(classes))

	return classes
}

func createAttributeToStatMap(levelUpStatusCalcParams []param.LevelUpStatusCalc) output.Attributes[output.Stats[bool]] {
	attributesToStatMap := output.Attributes[output.Stats[bool]]{}
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
				stats := output.Stats[bool]{}
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

func createCovenants(vowParams []param.Vow) []string {
	fmt.Println("\nCreating covenants...")

	covenants := []string{}

	for _, vowParam := range vowParams {
		covenants = append(covenants, vowParam.Name)
	}

	fmt.Printf("Created %d covenants\n", len(covenants))

	return covenants
}

func createLevels(playerLevelUpSoulsParams []param.PlayerLevelUpSouls) []int {
	fmt.Println("\nCreating levels...")

	levels := []int{}

	for _, playerLevelUpSoulsParam := range playerLevelUpSoulsParams {
		levels = append(levels, playerLevelUpSoulsParam.NecessarySouls)
	}

	fmt.Printf("Created %d levels\n", len(levels))

	return levels
}

func createSpells(spellParams []param.Spell) []output.Spell {
	fmt.Println("\nCreating spells...")

	spells := []output.Spell{}

	for _, spellParam := range spellParams {
		spells = append(spells, output.Spell{
			Name:                 spellParam.Name,
			RequiredIntelligence: spellParam.RequiredIntelligence,
			RequiredFaith:        spellParam.RequiredFaith,
			SpellSlotCost:        spellParam.SpellSlotCost,
			UsageCountCurve: []int{
				spellParam.UsageCountLV1,
				spellParam.UsageCountLV2,
				spellParam.UsageCountLV3,
				spellParam.UsageCountLV4,
				spellParam.UsageCountLV5,
				spellParam.UsageCountLV6,
				spellParam.UsageCountLV7,
				spellParam.UsageCountLV8,
				spellParam.UsageCountLV9,
				spellParam.UsageCountLV10,
			},
		})
	}

	fmt.Printf("Created %d spells\n", len(spells))

	return spells
}

func createArmor(armorParams []param.Armor) []output.Armor {
	fmt.Println("\nCreating armor...")

	armor := []output.Armor{}

	for _, armorParam := range armorParams {
		armorName := armorParam.Name

		if strings.Contains(armorName, "[Body]") {
			armorName = "No Armor"
		}

		// Armor ID of 113601XX is the visible Aurous set
		if math.Floor(float64(armorParam.ID)/100) == 113601 {
			armorName = armorParam.Name + " (Visible)"
		}

		// Armor ID of 113611XX is the invisible Aurous set
		if math.Floor(float64(armorParam.ID)/100) == 113611 {
			armorName = armorParam.Name + " (Invisible)"
		}

		// Armor ID of 12270100 is the Prisoner's Hood (Mask Only)
		if armorParam.ID == 12270100 {
			armorName = armorParam.Name + " (Mask Only)"
		}

		// Armor ID of 12270101 is the Prisoner's Tatters (Master's Attire)
		if armorParam.ID == 12270101 {
			armorName = armorParam.Name + " (Master's Attire)"
		}

		armor = append(armor, output.Armor{
			Equippable: output.Equippable{
				Name:                    armorName,
				AdditiveModifiers:       map[string]float64{},
				MultiplicativeModifiers: map[string]float64{},
				Weight:                  armorParam.Weight,
			},
			Defenses:    output.Defenses{},
			Resistances: output.Resistances{},
			Poise:       armorParam.Poise,
			Requirements: output.ScalingAttributes[int]{
				Strength:     armorParam.PrerequisiteStrength,
				Dexterity:    armorParam.PrerequisiteDexterity,
				Intelligence: armorParam.PrerequisiteIntelligence,
				Faith:        armorParam.PrerequisiteFaith,
			},
		})
	}

	fmt.Printf("Created %d armor pieces\n", len(armor))

	return armor
}

// Transform transforms data from DS2 params/EMEVDs to Scholar-friendly data
func Transform(paramData paramParser.DS2Params, emevdData emevdParser.DS2EMEVD) (output.ScholarData, error) {
	rings := []output.Ring{
		noRing,
	}
	helmets := []param.Armor{}
	chestpieces := []param.Armor{}
	gauntlets := []param.Armor{}
	leggings := []param.Armor{}

	for _, armorParam := range paramData.ArmorParam {
		switch armorParam.ArmorCategory {
		case param.ArmorCategoryHead:
			helmets = append(helmets, armorParam)
		case param.ArmorCategoryChest:
			chestpieces = append(chestpieces, armorParam)
		case param.ArmorCategoryArms:
			gauntlets = append(gauntlets, armorParam)
		case param.ArmorCategoryLegs:
			leggings = append(leggings, armorParam)
		}
	}

	return output.ScholarData{
		Classes:            createClasses(paramData.PlayerStatusParam),
		Chestpieces:        createArmor(chestpieces),
		Gauntlets:          createArmor(gauntlets),
		Helmets:            createArmor(helmets),
		Leggings:           createArmor(leggings),
		Weapons:            []output.Weapon{},
		Rings:              rings,
		Levels:             createLevels(paramData.PlayerLevelUpSoulsParam),
		Covenants:          createCovenants(paramData.VowParam),
		Spells:             createSpells(paramData.SpellParam),
		AttributeToStatMap: createAttributeToStatMap(paramData.LevelUpStatusCalcParam),
		BaseStats:          output.Stats[float64]{},
		StatCalculation:    output.StatCalculationDetails{},
	}, nil
}
