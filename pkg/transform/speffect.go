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

// extractStatName returns the stat name from a special effect arg by parsing it into Title Case
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
		case "SetStateInfo":
			// SetStateInfo will always be special cases (mostly flag flips)
			method = output.ModifierMethodToggle

			switch spEffect.Args[0] {
			case "StateInfoType.DisableBackstab":
				// StateInfoType.DisableBackstab is a flag for whether backstab is enabled
				description = "Cannot be backstabbed"
				statName = "Backstab"
			case "StateInfoType.DisableFootstepSound":
				// StateInfoType.DisableFootstepSound is a flag for whether footstep sound is enabled
				description = "Disable Footstep Sound"
				statName = "Footstep Sound"
			case "StateInfoType.Unknown0":
				// StateInfoType.Unknown0 is a flag for whether headshot stagger is enabled
				description = "Disable Headshot Stagger"
				statName = "Headshot Stagger"
			case "StateInfoType.Unknown4":
				// StateInfoType.Unknown4 enables a 5% crit rate wherein crits deal 50% more damage
				description = "Enable 5% chance for crits, which deal 50% more damage"
				statName = "Critical Damage"
			case "StateInfoType.Unknown29":
				// StateInfoType.Unknown29 is a flag for whether headshot damage is enabled
				description = "Disable Headshot Damage"
				statName = "Headshot Damage"
			default:
				return nil, fmt.Errorf("unhandled flag for Statement: %+v", spEffect)
			}
		case "ModifySpellEffectLength":
			// ModifySpellEffectLength has the buff type at index 0
			value, err := strconv.ParseFloat(spEffect.Args[0], 64)
			if err != nil {
				return nil, err
			}
			verb := "Increase"
			if value < 0 {
				verb = "Decrease"
			}
			statName = strings.Replace(toTitleCase(spEffect.Name), "Modify", verb, 1)

			// Convert to percent
			value -= 1.0
			value *= 100

			method = output.ModifierMethodMultiplicative
			description = fmt.Sprintf("%s by %.2f%%", statName, math.Abs(value))

			// Convert back to decimal
			value /= 100
		case "ModifyBulletParam":
			// ModifyBulletParam extends the range of bows
			// Arg0 seems to always be 0
			// Arg1 seems to always refer to regular bows
			// Arg2 is always 0, 1, or 2 (not sure what it does)
			// Arg3 is the modifier value
			// Arg4 is always 1
			// Arg5 is also the modifier value
			// Arg6 is also the modifier value
			// Arg7 is always 1

			if spEffect.Args[2] != "0" {
				continue
			}

			method = output.ModifierMethodAdditive
			value, err := strconv.ParseFloat(spEffect.Args[3], 64)
			if err != nil {
				return nil, err
			}

			// Convert to percent
			value -= 1.0
			value *= 100

			statName = "Bow Range"
			description = fmt.Sprintf("Increase %s by %.1f%%", statName, value)
		case "ModifyProperty":
			// ModifyProperty is another special case with flag flipping and such
			// Arg0 is the flag

			switch spEffect.Args[0] {
			case "PlayerPropertyType.Unknown21":
				// PlayerPropertyType.Unknown21 is a flag for the player being wet
				description = "Soaks the wearer"
				statName = "Soak"
				method = output.ModifierMethodToggle
			default:
				return nil, fmt.Errorf("unhandled flag for Statement: %+v", spEffect)
			}
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

func createRingModifiers(spEffects []emevdParser.Statement) ([]output.Modifier, error) {
	modifiers := []output.Modifier{}

	for _, spEffect := range spEffects {
		modifiers = append(modifiers, output.Modifier{
			Description: spEffect.Name,
			TargetType:  output.ModifierTargetTypeStat,
			Target:      extractStatName(spEffect.Args[0]),
			Method:      output.ModifierMethodAdditive,
			Value:       1,
		})
	}

	return modifiers, nil
}
