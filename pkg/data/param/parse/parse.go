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
		MenuStatsParam          []param.MenuStats
		PlayerLevelUpSoulsParam []param.PlayerLevelUpSouls
		PlayerStatusParam       []param.PlayerStatus
		ItemParam               []param.Item
		RingParam               []param.Ring
		VowParam                []param.Vow
		WeaponParam             []param.Weapon
		WeaponReinforceParam    []param.WeaponReinforce
		WeaponStatsAffectParam  []param.WeaponStatsAffect
	}

	// ParamFile is an enum for the param CSV files
	ParamFile int

	// ParamFileMetadata is a struct for storing metadata about a param CSV file
	ParamFileMetadata struct {
		Name     string
		DataType reflect.Type
		ValidIDs []id.Range
	}
)

// Param file enum
const (
	ParamFileArmor ParamFile = iota
	ParamFileArmorReinforce
	ParamFileCustomAttrSpec
	ParamFileItem
	ParamFileLevelUpStatusCalc
	ParamFileMenuStats
	ParamFilePlayerLevelUpSouls
	ParamFilePlayerStatus
	ParamFileRing
	ParamFileVow
	ParamFileWeapon
	ParamFileWeaponReinforce
	ParamFileWeaponStatsAffect
	ParamFileCount // Must always be last
)

var (
	ParamFiles = map[ParamFile]ParamFileMetadata{
		ParamFileArmor: {
			Name:     "ArmorParam",
			DataType: reflect.TypeFor[param.Armor](),
			ValidIDs: param.ValidArmorIDs,
		},
		ParamFileArmorReinforce: {
			Name:     "ArmorReinforceParam",
			DataType: reflect.TypeFor[param.ArmorReinforce](),
			ValidIDs: param.ValidArmorReinforceIDs,
		},
		ParamFileCustomAttrSpec: {
			Name:     "CustomAttrSpecParam",
			DataType: reflect.TypeFor[param.CustomAttrSpec](),
			ValidIDs: param.ValidCustomAttrSpecIDs,
		},
		ParamFileItem: {
			Name:     "ItemParam",
			DataType: reflect.TypeFor[param.Item](),
			ValidIDs: param.ValidItemIDs,
		},
		ParamFileLevelUpStatusCalc: {
			Name:     "LevelUpStatusCalcParam",
			DataType: reflect.TypeFor[param.LevelUpStatusCalc](),
			ValidIDs: param.ValidLevelUpStatusCalcIDs,
		},
		ParamFileMenuStats: {
			Name:     "MenuStatsParam",
			DataType: reflect.TypeFor[param.MenuStats](),
			ValidIDs: param.ValidMenuStatsIDs,
		},
		ParamFilePlayerLevelUpSouls: {
			Name:     "PlayerLevelUpSoulsParam",
			DataType: reflect.TypeFor[param.PlayerLevelUpSouls](),
			ValidIDs: param.ValidPlayerLevelUpSoulsIDs,
		},
		ParamFilePlayerStatus: {
			Name:     "PlayerStatusParam",
			DataType: reflect.TypeFor[param.PlayerStatus](),
			ValidIDs: param.ValidPlayerStatusIDs,
		},
		ParamFileRing: {
			Name:     "RingParam",
			DataType: reflect.TypeFor[param.Ring](),
			ValidIDs: param.ValidRingIDs,
		},
		ParamFileVow: {
			Name:     "VowParam",
			DataType: reflect.TypeFor[param.Vow](),
			ValidIDs: param.ValidVowIDs,
		},
		ParamFileWeapon: {
			Name:     "WeaponParam",
			DataType: reflect.TypeFor[param.Weapon](),
			ValidIDs: param.ValidWeaponIDs,
		},
		ParamFileWeaponReinforce: {
			Name:     "WeaponReinforceParam",
			DataType: reflect.TypeFor[param.WeaponReinforce](),
			ValidIDs: param.ValidWeaponReinforceIDs,
		},
		ParamFileWeaponStatsAffect: {
			Name:     "WeaponStatsAffectParam",
			DataType: reflect.TypeFor[param.WeaponStatsAffect](),
			ValidIDs: param.ValidWeaponStatsAffectIDs,
		},
	}
)

// parseRow parses a single row from a CSV file
//
// @param row - The row to parse
//
// @param dataType - The reflect.Type of the struct field
func parseRow(row []string, dataType reflect.Type) any {
	v := reflect.New(dataType).Elem()

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

	return v.Interface()
}

// parseFile parses a single Param CSV file
//
// @param v - The reflect.Value of the struct field to populate
//
// @param file - The name of the file to parse
//
// @param dataType - The reflect.Type of the struct field
//
// @returns an error
func (p *DS2Params) parseFile(paramFile ParamFile) error {
	// Get metadata
	metadata, ok := ParamFiles[paramFile]
	if !ok {
		return fmt.Errorf("unknown param file: %d", paramFile)
	}

	v := reflect.ValueOf(p).Elem().FieldByName(metadata.Name)

	// Open the file
	f, err := os.Open(fmt.Sprintf("inputs/%s.csv", metadata.Name))
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
		for _, r := range metadata.ValidIDs {
			if r.Contains(rowID) {
				validID = true
				break
			}
		}

		if !validID {
			continue
		}

		// Parse
		record := reflect.ValueOf(parseRow(row, metadata.DataType))
		v.Set(reflect.Append(v, record))
	}

	fmt.Printf("Parsed %d rows\n", v.Len())

	return nil
}

// Parse parses the Param CSV files
//
// @param name - The name of the struct field to populate
//
// @param dataType - The reflect.Type of the struct field
//
// @returns an error
func Parse() (DS2Params, error) {
	result := DS2Params{}

	for i := range ParamFileCount {
		metadata, ok := ParamFiles[i]
		if !ok {
			return DS2Params{}, fmt.Errorf("unknown param file: %d", i)
		}

		fmt.Printf("\nParsing inputs/%s.csv...\n", metadata.Name)
		fmt.Println("Data Type:", metadata.DataType.String())

		// Parse
		err := result.parseFile(i)
		if err != nil {
			return DS2Params{}, err
		}

	}

	return result, nil
}
