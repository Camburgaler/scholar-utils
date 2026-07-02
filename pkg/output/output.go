// Package output contains functions for outputting data
package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Camburgaler/scholar-utils/pkg/parse"
)

func Output(data parse.DS2Params) error {
	os.Mkdir("outputs", 0755)

	for _, file := range parse.Files {
		var j []byte
		var err error

		switch file {
		case parse.FileArmorParam:
			j, err = json.Marshal(data.ArmorParam)
			if err != nil {
				return err
			}
		case parse.FileArmorReinforceParam:
			j, err = json.Marshal(data.ArmorReinforceParam)
			if err != nil {
				return err
			}
		case parse.FileRingParam:
			j, err = json.Marshal(data.RingParam)
			if err != nil {
				return err
			}
		case parse.FileWeaponParam:
			j, err = json.Marshal(data.WeaponParam)
			if err != nil {
				return err
			}
		case parse.FileWeaponReinforceParam:
			j, err = json.Marshal(data.WeaponReinforceParam)
			if err != nil {
				return err
			}
		case parse.FileWeaponStatsAffectParam:
			j, err = json.Marshal(data.WeaponStatsAffectParam)
			if err != nil {
				return err
			}
		}

		err = os.WriteFile(fmt.Sprintf("outputs/%s.json", file), j, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
