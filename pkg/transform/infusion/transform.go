package infusion

import (
	"fmt"
	"reflect"

	"github.com/Camburgaler/scholar-utils/pkg/data/params/weapon"
	"github.com/Camburgaler/scholar-utils/pkg/output"
)

var (
	scales = []string{
		"PhysicalScalingStr",
		"PhysicalScalingDex",
		"MagicScaling",
		"LightningScaling",
		"FireScaling",
		"DarkScaling",
		"PoisonScaling",
		"BleedScaling",
		"EnchantedScaling",
	}
)

const (
	MaxUpgrade        = 10
	MaxSpecialUpgrade = 5
)

func leastSquares(x, y []float64) (float64, float64) {
	// least-squares sum regression

	n := len(x)
	nf64 := float64(n)

	xy := 0.0
	xsq := 0.0
	sumX := 0.0
	sumY := 0.0

	for i := range n {
		xy += x[i] * y[i]
		xsq += x[i] * x[i]
		sumX += x[i]
		sumY += y[i]
	}

	slope := (nf64*xy - sumX*sumY) / (nf64*xsq - sumX*sumX)
	intercept := y[0]

	return slope, intercept
}

// Special scales stop scaling after level 5, so to check if it's special, compare levels 5 and 6
func isSpecial(row weapon.StatsAffectParam) bool {
	special := true
	rowV := reflect.ValueOf(&row).Elem()

	for _, scale := range scales {
		lvlFiveFieldName := fmt.Sprintf("Level05%s", scale)
		lvlSixFieldName := fmt.Sprintf("Level06%s", scale)
		if rowV.FieldByName(lvlFiveFieldName).Float() == rowV.FieldByName(lvlSixFieldName).Float() {
			special = false
		}
	}

	return special
}

func ExtractInfusions(rows []weapon.StatsAffectParam) []output.Infusion {
	infusions := []output.Infusion{}

	for _, row := range rows {
		// Initialize scale
		scale := output.Infusion{
			Name:              row.Name,
			DamageUpgradeRate: map[string]output.SlopeIntercept{},
			StatScalingRate:   output.ScalingStats{},
		}

		// Determine upgrade range
		upgRange := MaxUpgrade
		if isSpecial(row) {
			upgRange = MaxSpecialUpgrade
		}

		// Initialize container for graph data
		scales := map[string][]float64{
			"PhysicalScalingStr": {},
			"PhysicalScalingDex": {},
			"MagicScaling":       {},
			"LightningScaling":   {},
			"FireScaling":        {},
			"DarkScaling":        {},
			"PoisonScaling":      {},
			"BleedScaling":       {},
			"EnchantedScaling":   {},
		}
		vScales := reflect.ValueOf(&scales).Elem()
		vRow := reflect.ValueOf(&row).Elem()

		// Populate graph data
		for i := range upgRange {
			for scaleName := range scales {
				rowFieldName := fmt.Sprintf("Level%02d%s", i, scaleName)
				vScales.FieldByName(scaleName).Set(reflect.Append(vScales.FieldByName(scaleName), reflect.ValueOf(vRow.FieldByName(rowFieldName).Float())))
			}
		}

		// Calculate formulae
		for scaleName := range scales {
			x := []float64{}
			for i := range upgRange {
				x = append(x, float64(i))
			}
			slope, intercept := leastSquares(x, scales[scaleName])

			scale.DamageUpgradeRate[scaleName] = output.SlopeIntercept{
				Slope:     slope,
				Intercept: intercept,
			}
		}

		infusions = append(infusions, scale)
	}

	return infusions
}

//         # damage & upgrade
//         physical = [float(relevant[i]["physicsAtkRate"]) for i in range(0, upgRange)]
//         magic = [float(relevant[i]["magicAtkRate"]) for i in range(0, upgRange)]
//         fire = [float(relevant[i]["fireAtkRate"]) for i in range(0, upgRange)]
//         lightning = [float(relevant[i]["thunderAtkRate"]) for i in range(0, upgRange)]
//         holy = [float(relevant[i]["darkAtkRate"]) for i in range(0, upgRange)]

//         physical_upg, physical_dmg = regression(xs, physical)
//         magic_upg, magic_dmg = regression(xs, magic)
//         fire_upg, fire_dmg = regression(xs, fire)
//         lightning_upg, lightning_dmg = regression(xs, lightning)
//         holy_upg, holy_dmg = regression(xs, holy)

//         infusion["damageUpgradeRate"] = {
//             "physical": {
//                 "slope": physical_upg,
//                 "intercept": physical_dmg,
//             },
//             "magic": {
//                 "slope": magic_upg,
//                 "intercept": magic_dmg,
//             },
//             "fire": {
//                 "slope": fire_upg,
//                 "intercept": fire_dmg,
//             },
//             "lightning": {
//                 "slope": lightning_upg,
//                 "intercept": lightning_dmg,
//             },
//             "holy": {
//                 "slope": holy_upg,
//                 "intercept": holy_dmg,
//             },
//         }

//         # scaling
//         strength = [float(relevant[i]["correctStrengthRate"]) for i in range(0, upgRange)]
//         dexterity = [float(relevant[i]["correctAgilityRate"]) for i in range(0, upgRange)]
//         intelligence = [float(relevant[i]["correctMagicRate"]) for i in range(0, upgRange)]
//         faith = [float(relevant[i]["correctFaithRate"]) for i in range(0, upgRange)]
//         arcane = [float(relevant[i]["correctLuckRate"]) for i in range(0, upgRange)]

//         infusion["statScalingRate"] = {
//             "STR": strength,
//             "DEX": dexterity,
//             "INT": intelligence,
//             "FTH": faith,
//             "ARC": arcane,
//         }

//         infusions[id] = infusion
