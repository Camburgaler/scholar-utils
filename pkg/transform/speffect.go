package transform

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	emevdParser "github.com/Camburgaler/scholar-utils/pkg/data/emevd/parse"
	"github.com/Camburgaler/scholar-utils/pkg/output"
)

var (
	attributeNames = []string{
		"Strength",
		"Dexterity",
		"Intelligence",
		"Faith",
		"Vigor",
		"Endurance",
		"Vitality",
		"Adaptability",
		"Attunement",
	}

	statNames = []string{
		"Max Equip Load",
		"Attunement Slots",
		"Max HP",
		"Spell Cast Time",
		"Bleed Damage",
		"Poison Damage",
		"Max Spell Usages",
	}
)

// toTitleCase converts a string from PascalCase to Title Case
func toTitleCase(pascal string) string {
	// Except for the very first index, insert a space before each uppercase letter
	var title strings.Builder
	for i, r := range pascal {
		if i > 0 && r >= 'A' && r <= 'Z' && !(pascal[i-1] == 'H' && r == 'P') {
			title.WriteString(" ")
		}

		title.WriteRune(r)
	}

	return title.String()
}

func extractStatName(spEffectArg string) string {
	if !strings.Contains(spEffectArg, ".") {
		return ""
	}

	argBreakdown := strings.Split(spEffectArg, ".")

	// The stat name is in PascalCase, so add spaces before each word to convert to Title Case
	statName := toTitleCase(argBreakdown[1])

	if argBreakdown[0] == "BuffType" {
		statName += " Damage"
	}

	if statName == "Equip Load" {
		statName = "Max Equip Load"
	}

	return statName
}

func createArmorModifiers(spEffects []emevdParser.Statement) ([]output.Modifier, error) {
	modifiers := []output.Modifier{}

	for _, spEffect := range spEffects {
		var statName = ""
		var value = 0.0
		var description = ""
		var method output.ModifierMethod
		var targetType = output.ModifierTargetTypeSpecial

		// ignore IncreaseDamageMultiplicatively for now
		switch spEffect.Name {
		case "ModifyStatAdditively", "ModifyStatMultiplicatively":
			// Generally, each statement has an arg that indicates the stat at arg 0
			statName = extractStatName(spEffect.Args[0])

			// It looks like most statements have 2 numeric args: index 1 that's always 0 and index 2 that is the actual value
			if num, err := strconv.ParseFloat(spEffect.Args[2], 64); err == nil {
				if num != 0 {
					value = num
				}
			}

			verb := "Increase"

			switch spEffect.Name {
			case "ModifyStatMultiplicatively":
				// convert to percentage
				value = (value - 1) * 100

				// round to two decimal places
				value = math.Round(value*100) / 100

				// derive verb
				if value < 1 {
					verb = "Decrease"
				}

				// generate description and method
				description = fmt.Sprintf("%s %s by %.1f%%", verb, statName, math.Abs(value))
				method = output.ModifierMethodMultiplicative

				// convert to decimal for output
				value = (value / 100)
			case "ModifyStatAdditively":
				// derive verb
				if value < 0 {
					verb = "Decrease"
				}

				// generate description and method
				description = fmt.Sprintf("%s %s by %d", verb, statName, int(math.Abs(value)))
				method = output.ModifierMethodAdditive
			}
		case "ModifyAttributeBasedOnCurrentValue":
			// Generally, each statement has an arg that indicates the stat at arg 0
			statName = extractStatName(spEffect.Args[0])

			// ModifyAttributeBasedOnCurrentValue always just reduces the attribute by 1
			value = -1
			description = fmt.Sprintf("Decrease %s by 1", statName)
			method = output.ModifierMethodAdditive
		case "ModifyDamageFlatToBaseAr":
			// Armor with ModifyDamageFlatToBaseAr has the following effects:
			//     Weapons with innate affliction of the relevant damage type get 50 points of extra damage.
			//     Weapons infused with the relevant damage type get 25 points of extra damage.
			//     Weapons with innate relevant damage AND infused with the relevant damage type get 60 points of extra damage.

			// ModifyDamageFlatToBaseAr has the buff type at index 1
			statName = extractStatName(spEffect.Args[1])

			method = output.ModifierMethodAdditive

			for i, descStr := range []string{"Increase %s by %d when the weapon has innate %s", "Increase %s by %d when the weapon is infused with %s"} {
				if statName == "" {
					return nil, fmt.Errorf("no stat name in Statement: %+v", spEffect)
				}

				value = 50.0 / float64(i+1)
				description = fmt.Sprintf(descStr, statName, int(math.Abs(value)), statName)

				modifiers = append(modifiers, output.Modifier{
					Description: description,
					TargetType:  output.ModifierTargetTypeStat,
					Target:      statName,
					Method:      method,
					Value:       value,
				})
			}

			value = 60.0
			description = fmt.Sprintf("Increase %s by %d when the weapon has innate AND infused %s", statName, int(math.Abs(value)), statName)
		case "IncreaseDamageMultiplicatively":
			// Ignore
			continue
		default:
			return nil, fmt.Errorf("unhandled SpEffect name for Statement: %+v", spEffect)
		}

		if statName == "" {
			return nil, fmt.Errorf("no stat name in Statement: %+v", spEffect)
		}

		if slices.Contains(attributeNames, statName) {
			targetType = output.ModifierTargetTypeAttribute
		}
		if slices.Contains(statNames, statName) {
			targetType = output.ModifierTargetTypeStat
		}

		modifiers = append(modifiers, output.Modifier{
			Description: description,
			TargetType:  targetType,
			Target:      statName,
			Method:      method,
			Value:       value,
		})
	}

	return modifiers, nil
}
