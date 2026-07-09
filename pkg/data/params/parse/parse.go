// Package parse contains functions for parsing CSV files
package parse

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Camburgaler/scholar-utils/pkg/data/params/armor"
	"github.com/Camburgaler/scholar-utils/pkg/data/params/customattrspec"
	"github.com/Camburgaler/scholar-utils/pkg/data/params/id"
	"github.com/Camburgaler/scholar-utils/pkg/data/params/item"
	"github.com/Camburgaler/scholar-utils/pkg/data/params/ring"
	"github.com/Camburgaler/scholar-utils/pkg/data/params/weapon"
)

type (
	// DS2Params is a struct for storing data from the param CSV files
	DS2Params struct {
		ArmorParam             []armor.Param
		ArmorReinforceParam    []armor.ReinforceParam
		CustomAttrSpecParam    []customattrspec.Param
		ItemParam              []item.Param
		RingParam              []ring.Param
		WeaponParam            []weapon.Param
		WeaponReinforceParam   []weapon.ReinforceParam
		WeaponStatsAffectParam []weapon.StatsAffectParam
	}
)

// File names
const (
	ParamFileArmor             = "ArmorParam"
	ParamFileArmorReinforce    = "ArmorReinforceParam"
	ParamFileCustomAttrSpec    = "CustomAttrSpecParam"
	ParamFileItem              = "ItemParam"
	ParamFileRing              = "RingParam"
	ParamFileWeapon            = "WeaponParam"
	ParamFileWeaponReinforce   = "WeaponReinforceParam"
	ParamFileWeaponStatsAffect = "WeaponStatsAffectParam"
)

var (
	// File names in a slice
	ParamFiles = []string{
		ParamFileArmor,
		ParamFileArmorReinforce,
		ParamFileCustomAttrSpec,
		ParamFileItem,
		ParamFileRing,
		ParamFileWeapon,
		ParamFileWeaponReinforce,
		ParamFileWeaponStatsAffect,
	}
	validParamIDs = map[string][]id.Range{
		ParamFileArmor:             armor.ValidParamIDs,
		ParamFileArmorReinforce:    armor.ValidReinforceParamIDs,
		ParamFileCustomAttrSpec:    customattrspec.ValidParamIDs,
		ParamFileItem:              item.ValidParamIDs,
		ParamFileRing:              ring.ValidParamIDs,
		ParamFileWeapon:            weapon.ValidParamIDs,
		ParamFileWeaponReinforce:   weapon.ValidReinforceParamIDs,
		ParamFileWeaponStatsAffect: weapon.ValidStatsAffectParamIDs,
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

func parseFile(s *bufio.Scanner, file string, params DS2Params) (DS2Params, error) {
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
		rowZero, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			return DS2Params{}, err
		}
		rowID := id.ID(rowZero)

		validID := true
		for _, range_ := range validParamIDs[file] {
			if range_.Start <= rowID && range_.End >= rowID {
				break
			}
			validID = false
		}

		if !validID {
			continue
		}

		// Parse
		switch file {
		case ParamFileArmor:
			params.ArmorParam = append(params.ArmorParam, parseRow[armor.Param](row))
		case ParamFileArmorReinforce:
			params.ArmorReinforceParam = append(params.ArmorReinforceParam, parseRow[armor.ReinforceParam](row))
		case ParamFileCustomAttrSpec:
			params.CustomAttrSpecParam = append(params.CustomAttrSpecParam, parseRow[customattrspec.Param](row))
		case ParamFileItem:
			params.ItemParam = append(params.ItemParam, parseRow[item.Param](row))
		case ParamFileRing:
			params.RingParam = append(params.RingParam, parseRow[ring.Param](row))
		case ParamFileWeapon:
			params.WeaponParam = append(params.WeaponParam, parseRow[weapon.Param](row))
		case ParamFileWeaponReinforce:
			params.WeaponReinforceParam = append(params.WeaponReinforceParam, parseRow[weapon.ReinforceParam](row))
		case ParamFileWeaponStatsAffect:
			params.WeaponStatsAffectParam = append(params.WeaponStatsAffectParam, parseRow[weapon.StatsAffectParam](row))
		}
	}

	return params, nil
}

// Parse parses the Param CSV files
//
// Returns a DS2Params struct
func Parse() (DS2Params, error) {
	result := DS2Params{}

	// For each file, populate the struct
	for _, file := range ParamFiles {
		fmt.Printf("\nParsing %s.csv...\n", file)

		// Open the file
		f, err := os.Open(fmt.Sprintf("inputs/%s.csv", file))
		if err != nil {
			return DS2Params{}, err
		}

		// Create scanner
		s := bufio.NewScanner(f)

		// Parse
		result, err = parseFile(s, file, result)
		if err != nil {
			return DS2Params{}, err
		}

		// Close the file
		f.Close()
	}

	return result, nil
}
