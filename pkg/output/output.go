// Package output contains functions for outputting data
package output

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	emevdParser "github.com/Camburgaler/scholar-utils/pkg/data/emevd/parse"
	paramParser "github.com/Camburgaler/scholar-utils/pkg/data/param/parse"
)

func writeData(file string, data any) error {
	fmt.Printf("\nWriting %s...\n", file)

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

	data.BaseStats = BaseStats
	data.StatCalculation = DS2StatCalculationDetails

	// TODO: Remove before release
	paramFileValueMap := map[paramParser.ParamFile]any{
		paramParser.ParamFileArmor:              paramData.ArmorParam,
		paramParser.ParamFileArmorReinforce:     paramData.ArmorReinforceParam,
		paramParser.ParamFileCustomAttrSpec:     paramData.CustomAttrSpecParam,
		paramParser.ParamFileItem:               paramData.ItemParam,
		paramParser.ParamFileMenuStats:          paramData.MenuStatsParam,
		paramParser.ParamFilePlayerLevelUpSouls: paramData.PlayerLevelUpSoulsParam,
		paramParser.ParamFileRing:               paramData.RingParam,
		paramParser.ParamFileVow:                paramData.VowParam,
		paramParser.ParamFileWeapon:             paramData.WeaponParam,
		paramParser.ParamFileWeaponReinforce:    paramData.WeaponReinforceParam,
		paramParser.ParamFileWeaponStatsAffect:  paramData.WeaponStatsAffectParam,
	}

	// TODO: Remove before release
	emevdFileMap := map[string]emevdParser.Events{
		emevdParser.EMEVDFileArmor:  emevdData.SpEffectArmor,
		emevdParser.EMEVDFileRing:   emevdData.SpEffectRing,
		emevdParser.EMEVDFileWeapon: emevdData.SpEffectWeapon,
	}

	// TODO: Remove before release
	for i := range paramParser.ParamFileCount {
		metadata := paramParser.ParamFiles[i]

		outputFileName := fmt.Sprintf("outputs/_%s.json", metadata.Name)
		err := writeData(outputFileName, paramFileValueMap[i])
		if err != nil {
			return err
		}
	}

	// TODO: Remove before release
	for _, emevdFileName := range emevdParser.EMEVDFiles {
		outputFileName := fmt.Sprintf("outputs/_%s.json", emevdFileName)
		err := writeData(outputFileName, emevdFileMap[emevdFileName])
		if err != nil {
			return err
		}
	}

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
