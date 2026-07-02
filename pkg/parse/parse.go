// Package parse contains functions for parsing CSV files
package parse

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/Camburgaler/scholar-utils/pkg/params/armor"
	"github.com/Camburgaler/scholar-utils/pkg/params/id"
	"github.com/Camburgaler/scholar-utils/pkg/params/ring"
	"github.com/Camburgaler/scholar-utils/pkg/params/weapon"
)

type (
	// DS2Params is a struct for storing data from the param CSV files
	DS2Params struct {
		ArmorParam             []armor.Param
		ArmorReinforceParam    []armor.ReinforceParam
		RingParam              []ring.Param
		WeaponParam            []weapon.Param
		WeaponReinforceParam   []weapon.ReinforceParam
		WeaponStatsAffectParam []weapon.StatsAffectParam
	}
)

// File names
const (
	FileArmorParam             = "ArmorParam"
	FileArmorReinforceParam    = "ArmorReinforceParam"
	FileRingParam              = "RingParam"
	FileWeaponParam            = "WeaponParam"
	FileWeaponReinforceParam   = "WeaponReinforceParam"
	FileWeaponStatsAffectParam = "WeaponStatsAffectParam"
)

var (
	// File names in a slice
	Files = []string{
		FileArmorParam,
		FileArmorReinforceParam,
		FileRingParam,
		FileWeaponParam,
		FileWeaponReinforceParam,
		FileWeaponStatsAffectParam,
	}
	validParamIDs = map[string][]id.Range{
		FileArmorParam:             armor.ValidParamIDs,
		FileArmorReinforceParam:    armor.ValidReinforceParamIDs,
		FileRingParam:              ring.ValidParamIDs,
		FileWeaponParam:            weapon.ValidParamIDs,
		FileWeaponReinforceParam:   weapon.ValidReinforceParamIDs,
		FileWeaponStatsAffectParam: weapon.ValidStatsAffectParamIDs,
	}
)

func parseRow[T any](row []string) T {
	var t T
	v := reflect.ValueOf(&t).Elem()

	// For each field on the struct, set the value from the row
	for i := 0; i < v.NumField(); i++ {
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

// Parse parses the CSV files
//
// Returns a DS2Params struct
func Parse() (DS2Params, error) {
	armorParam := []armor.Param{}
	armorReinforceParam := []armor.ReinforceParam{}
	ringParam := []ring.Param{}
	weaponParam := []weapon.Param{}
	weaponReinforceParam := []weapon.ReinforceParam{}
	weaponStatsAffectParam := []weapon.StatsAffectParam{}

	// For each file, populate the struct
	for _, file := range Files {
		fmt.Printf("Parsing %s.csv...\n", file)

		// Open the file
		f, err := os.Open(fmt.Sprintf("inputs/%s.csv", file))
		if err != nil {
			return DS2Params{}, err
		}
		defer f.Close()

		csvReader := csv.NewReader(f)

		for {
			row, err := csvReader.Read()

			// End of file
			if row == nil {
				break
			}

			// Skip header
			if row[0] == "ID" {
				fmt.Printf("Header has %d fields\n", len(row))
				continue
			}

			// Check if the row has the correct number of fields
			if len(row) != csvReader.FieldsPerRecord {
				fmt.Printf("Encountered a row with %d fields\n", len(row))
				break
			}

			if err != nil {
				fmt.Println(err)
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
				if range_.Min <= rowID && range_.Max >= rowID {
					break
				}
				validID = false
			}

			if !validID {
				continue
			}

			// Parse
			switch file {
			case FileArmorParam:
				armorParam = append(armorParam, parseRow[armor.Param](row))
			case FileArmorReinforceParam:
				armorReinforceParam = append(armorReinforceParam, parseRow[armor.ReinforceParam](row))
			case FileRingParam:
				ringParam = append(ringParam, parseRow[ring.Param](row))
			case FileWeaponParam:
				weaponParam = append(weaponParam, parseRow[weapon.Param](row))
			case FileWeaponReinforceParam:
				weaponReinforceParam = append(weaponReinforceParam, parseRow[weapon.ReinforceParam](row))
			case FileWeaponStatsAffectParam:
				weaponStatsAffectParam = append(weaponStatsAffectParam, parseRow[weapon.StatsAffectParam](row))
			}
		}

		fmt.Println("Parsed!")
	}

	return DS2Params{armorParam, armorReinforceParam, ringParam, weaponParam, weaponReinforceParam, weaponStatsAffectParam}, nil
}
