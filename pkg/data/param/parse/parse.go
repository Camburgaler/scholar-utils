// Package parse contains functions for parsing CSV files
package parse

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Camburgaler/scholar-utils/pkg/data/param"
	"github.com/Camburgaler/scholar-utils/pkg/data/param/id"
)

type (
	// DS2Params is a struct for storing data from the param CSV files
	DS2Params struct {
		ArmorParam              []param.Armor
		ArmorReinforceParam     []param.ArmorReinforce
		CustomAttrSpecParam     []param.CustomAttrSpec
		LevelUpStatusCalcParam  []param.LevelUpStatusCalc
		PlayerLevelUpSoulsParam []param.PlayerLevelUpSoul
		ItemParam               []param.Item
		RingParam               []param.Ring
		VowParam                []param.Vow
		WeaponParam             []param.Weapon
		WeaponReinforceParam    []param.WeaponReinforce
		WeaponStatsAffectParam  []param.WeaponStatsAffect
	}
)

// File names
const (
	ParamFileArmor              = "ArmorParam"
	ParamFileArmorReinforce     = "ArmorReinforceParam"
	ParamFileCustomAttrSpec     = "CustomAttrSpecParam"
	ParamFileItem               = "ItemParam"
	ParamFileLevelUpStatusCalc  = "LevelUpStatusCalcParam"
	ParamFilePlayerLevelUpSouls = "PlayerLevelUpSoulsParam"
	ParamFileRing               = "RingParam"
	ParamFileVow                = "VowParam"
	ParamFileWeapon             = "WeaponParam"
	ParamFileWeaponReinforce    = "WeaponReinforceParam"
	ParamFileWeaponStatsAffect  = "WeaponStatsAffectParam"
)

var (
	// File names in a slice
	ParamFiles = []string{
		ParamFileArmor,
		ParamFileArmorReinforce,
		ParamFileCustomAttrSpec,
		ParamFileItem,
		ParamFileLevelUpStatusCalc,
		ParamFilePlayerLevelUpSouls,
		ParamFileRing,
		ParamFileVow,
		ParamFileWeapon,
		ParamFileWeaponReinforce,
		ParamFileWeaponStatsAffect,
	}
	validParamIDs = map[string][]id.Range{
		ParamFileArmor:              param.ValidArmorIDs,
		ParamFileArmorReinforce:     param.ValidArmorReinforceIDs,
		ParamFileCustomAttrSpec:     param.ValidCustomAttrSpecIDs,
		ParamFileItem:               param.ValidItemIDs,
		ParamFileLevelUpStatusCalc:  param.ValidLevelUpStatusCalcIDs,
		ParamFilePlayerLevelUpSouls: param.ValidPlayerLevelUpSoulsIDs,
		ParamFileRing:               param.ValidRingIDs,
		ParamFileVow:                param.ValidVowIDs,
		ParamFileWeapon:             param.ValidWeaponIDs,
		ParamFileWeaponReinforce:    param.ValidWeaponReinforceIDs,
		ParamFileWeaponStatsAffect:  param.ValidWeaponStatsAffectIDs,
	}
)

func parseRow[T any](row []string) T {
	var t T
	v := reflect.ValueOf(&t).Elem()

	// For each field on the struct, set the value from the row
	for i := 0; i < v.NumField(); i++ {
		// Skip unexported fields
		if v.Field(i).CanSet() == false {
			continue
		}

		if v.Field(i).Kind() == reflect.String {
			v.Field(i).SetString(row[i])
		} else if v.Field(i).Kind() == reflect.Int64 {
			num, err := strconv.ParseInt(row[i], 10, 64)
			if err != nil {
				panic(err)
			}
			v.Field(i).SetInt(int64(num))
		} else if v.Field(i).Kind() == reflect.Float64 {
			num, err := strconv.ParseFloat(row[i], 64)
			if err != nil {
				panic(err)
			}
			v.Field(i).SetFloat(num)
		} else if v.Field(i).Kind() == reflect.Bool {
			v.Field(i).SetBool(row[i] == "1")
		} else {
			panic(fmt.Sprintf("unknown type: %s\nvalue: %s", v.Field(i).Kind().String(), row[i]))
		}
	}

	return t
}

func (p *DS2Params) parseFile(file string) error {
	// Open the file
	f, err := os.Open(fmt.Sprintf("inputs/%s.csv", file))
	if err != nil {
		return err
	}
	defer f.Close()

	// Create scanner
	s := bufio.NewScanner(f)

	fieldsPerRecord := 0

	for s.Scan() {
		line, err := s.Text(), s.Err()
		if err != nil {
			fmt.Println(err)
			break
		}

		// Split line into fields
		row := strings.Split(line, ",")

		// Skip header
		if row[0] == "ID" {
			// Check if the last field is empty (indicates an unnecessary trailing comma)
			if row[len(row)-1] == "" {
				fieldsPerRecord = len(row) - 1
			} else {
				fieldsPerRecord = len(row)
			}
			fmt.Printf("Header has %d fields\n", fieldsPerRecord)
			continue
		}

		// Check if the row has the correct number of fields
		if len(row) != fieldsPerRecord {
			fmt.Printf("Encountered a row with %d fields:\n%s\n", len(row), row)
			break
		}

		// Check if the ID is valid
		rowAtZeroI, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			return err
		}
		rowID := id.ID(rowAtZeroI)

		validID := false
		for _, r := range validParamIDs[file] {
			if r.Contains(rowID) {
				validID = true
				break
			}
		}

		if !validID {
			continue
		}

		// Parse
		switch file {
		case ParamFileArmor:
			p.ArmorParam = append(p.ArmorParam, parseRow[param.Armor](row))
		case ParamFileArmorReinforce:
			p.ArmorReinforceParam = append(p.ArmorReinforceParam, parseRow[param.ArmorReinforce](row))
		case ParamFileCustomAttrSpec:
			p.CustomAttrSpecParam = append(p.CustomAttrSpecParam, parseRow[param.CustomAttrSpec](row))
		case ParamFileItem:
			p.ItemParam = append(p.ItemParam, parseRow[param.Item](row))
		case ParamFileLevelUpStatusCalc:
			p.LevelUpStatusCalcParam = append(p.LevelUpStatusCalcParam, parseRow[param.LevelUpStatusCalc](row))
		case ParamFilePlayerLevelUpSouls:
			p.PlayerLevelUpSoulsParam = append(p.PlayerLevelUpSoulsParam, parseRow[param.PlayerLevelUpSoul](row))
		case ParamFileRing:
			p.RingParam = append(p.RingParam, parseRow[param.Ring](row))
		case ParamFileVow:
			p.VowParam = append(p.VowParam, parseRow[param.Vow](row))
		case ParamFileWeapon:
			p.WeaponParam = append(p.WeaponParam, parseRow[param.Weapon](row))
		case ParamFileWeaponReinforce:
			p.WeaponReinforceParam = append(p.WeaponReinforceParam, parseRow[param.WeaponReinforce](row))
		case ParamFileWeaponStatsAffect:
			p.WeaponStatsAffectParam = append(p.WeaponStatsAffectParam, parseRow[param.WeaponStatsAffect](row))
		}
	}

	return nil
}

// Parse parses the Param CSV files
//
// Returns a DS2Params struct
func Parse() (DS2Params, error) {
	result := DS2Params{}

	// For each file, populate the struct
	for _, file := range ParamFiles {
		fmt.Printf("\nParsing %s.csv...\n", file)

		// Parse
		err := result.parseFile(file)
		if err != nil {
			return DS2Params{}, err
		}
	}

	return result, nil
}
