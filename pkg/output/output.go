// Package output contains functions for outputting data
package output

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	emevdParser "github.com/Camburgaler/scholar-utils/pkg/data/emevd/parse"
	paramParser "github.com/Camburgaler/scholar-utils/pkg/data/params/parse"
)

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
	Armor struct {
		Name        string
		Defenses    Defenses
		Absorptions Absorptions
		Poise       float64
		Weight      float64
	}
	Ring struct {
		Name string
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
		Classes     map[string]class
		Helmets     map[string]Armor
		Chestpieces map[string]Armor
		Gauntlets   map[string]Armor
		Leggings    map[string]Armor
		Rings       map[string]Ring
		Weapons     map[string]Weapon
	}
)

var (
	classes = map[string]class{
		"warrior": {
			Name:  "Warrior",
			Level: 12,
			Stats: stats{
				Vigor:        7,
				Endurance:    6,
				Vitality:     6,
				Attunement:   5,
				Strength:     15,
				Dexterity:    11,
				Adaptability: 5,
				Intelligence: 5,
				Faith:        5,
			},
		},
		"knight": {
			Name:  "Knight",
			Level: 13,
			Stats: stats{
				Vigor:        12,
				Endurance:    6,
				Vitality:     7,
				Attunement:   4,
				Strength:     11,
				Dexterity:    8,
				Adaptability: 9,
				Intelligence: 3,
				Faith:        6,
			},
		},
		"swordsman": {
			Name:  "Swordsman",
			Level: 12,
			Stats: stats{
				Vigor:        4,
				Endurance:    8,
				Vitality:     4,
				Attunement:   6,
				Strength:     9,
				Dexterity:    16,
				Adaptability: 6,
				Intelligence: 7,
				Faith:        5,
			},
		},
		"bandit": {
			Name:  "Bandit",
			Level: 11,
			Stats: stats{
				Vigor:        9,
				Endurance:    7,
				Vitality:     11,
				Attunement:   2,
				Strength:     9,
				Dexterity:    14,
				Adaptability: 3,
				Intelligence: 1,
				Faith:        8,
			},
		},
		"cleric": {
			Name:  "Cleric",
			Level: 14,
			Stats: stats{
				Vigor:        10,
				Endurance:    3,
				Vitality:     8,
				Attunement:   10,
				Strength:     11,
				Dexterity:    5,
				Adaptability: 4,
				Intelligence: 4,
				Faith:        12,
			},
		},
		"sorcerer": {
			Name:  "Sorcerer",
			Level: 11,
			Stats: stats{
				Vigor:        5,
				Endurance:    6,
				Vitality:     5,
				Attunement:   12,
				Strength:     3,
				Dexterity:    7,
				Adaptability: 8,
				Intelligence: 14,
				Faith:        4,
			},
		},
		"explorer": {
			Name:  "Explorer",
			Level: 10,
			Stats: stats{
				Vigor:        7,
				Endurance:    6,
				Vitality:     9,
				Attunement:   7,
				Strength:     6,
				Dexterity:    6,
				Adaptability: 12,
				Intelligence: 5,
				Faith:        5,
			},
		},
		"deprived": {
			Name:  "Deprived",
			Level: 1,
			Stats: stats{
				Vigor:        6,
				Endurance:    6,
				Vitality:     6,
				Attunement:   6,
				Strength:     6,
				Dexterity:    6,
				Adaptability: 6,
				Intelligence: 6,
				Faith:        6,
			},
		},
	}
)

func writeData(file string, data any) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	j, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	_, err = f.Write(j)
	if err != nil {
		return err
	}

	return nil
}

// Output outputs the data to JSON files
//
// @param data ScholarData - the parsed, transformed data to output
//
// @param paramData DS2Params - temporary, just for development
//
// @param emevdData DS2EMEVD - temporary, just for development
//
// @return error
func Output(data ScholarData, paramData paramParser.DS2Params, emevdData emevdParser.DS2EMEVD) error {
	os.Mkdir("outputs", 0755)

	// TODO: Remove before release
	paramFileValueMap := map[string]any{
		paramParser.ParamFileArmor:             paramData.ArmorParam,
		paramParser.ParamFileArmorReinforce:    paramData.ArmorReinforceParam,
		paramParser.ParamFileCustomAttrSpec:    paramData.CustomAttrSpecParam,
		paramParser.ParamFileItem:              paramData.ItemParam,
		paramParser.ParamFileRing:              paramData.RingParam,
		paramParser.ParamFileWeapon:            paramData.WeaponParam,
		paramParser.ParamFileWeaponReinforce:   paramData.WeaponReinforceParam,
		paramParser.ParamFileWeaponStatsAffect: paramData.WeaponStatsAffectParam,
	}

	// TODO: Remove before release
	emevdFileMap := map[string]emevdParser.Events{
		emevdParser.EMEVDFileArmor:  emevdData.SpEffectArmor,
		emevdParser.EMEVDFileRing:   emevdData.SpEffectRing,
		emevdParser.EMEVDFileWeapon: emevdData.SpEffectWeapon,
	}

	// TODO: Remove before release
	for _, paramFileName := range paramParser.ParamFiles {
		outputFileName := fmt.Sprintf("outputs/%s.json", paramFileName)
		err := writeData(outputFileName, paramFileValueMap[paramFileName])
		if err != nil {
			return err
		}
	}

	// TODO: Remove before release
	for _, emevdFileName := range emevdParser.EMEVDFiles {
		outputFileName := fmt.Sprintf("outputs/%s.json", emevdFileName)
		err := writeData(outputFileName, emevdFileMap[emevdFileName])
		if err != nil {
			return err
		}
	}

	data.Classes = classes

	// Output the data using reflection
	v := reflect.ValueOf(&data).Elem()
	for i := 0; i < v.NumField(); i++ {
		outputFileName := fmt.Sprintf("outputs/%s.json", v.Type().Field(i).Name)
		err := writeData(outputFileName, v.Field(i).Interface())
		if err != nil {
			return err
		}
	}

	return nil
}
