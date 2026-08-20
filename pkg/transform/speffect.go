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
		"Magic Defense",
		"Lightning Defense",
		"Fire Defense",
		"Dark Defense",
		"Poison Defense",
		"Bleed Defense",
		"Curse Defense",
		"Petrify Defense",
	}
)

// toTitleCase converts a string from PascalCase to Title Case
func toTitleCase(pascal string) string {
	// Except for the very first index, insert a space before each uppercase letter
	var title strings.Builder
	for i, r := range pascal {
		if i > 0 &&
			r >= 'A' &&
			r <= 'Z' &&
			!(pascal[i-1] == 'H' &&
				r == 'P') &&
			!(i > 1 &&
				pascal[i-2] == 'P' &&
				pascal[i-1] == 'v' &&
				r == 'E') {
			title.WriteString(" ")
		}

		title.WriteRune(r)
	}

	return title.String()
}

// extractStatName returns the stat name from a special effect arg by parsing it into Title Case
func extractStatName(spEffectArg string) string {
	argBreakdown := []string{"", spEffectArg}
	if strings.Contains(spEffectArg, ".") {
		argBreakdown = strings.Split(spEffectArg, ".")
	}

	// The stat name is in PascalCase, so add spaces before each word to convert to Title Case
	statName := toTitleCase(argBreakdown[1])

	if argBreakdown[0] == "BuffType" {
		statName += " Damage"
	}

	if statName == "Equip Load" {
		statName = "Max Equip Load"
	}

	if strings.Contains(statName, "Degredation") {
		statName = strings.ReplaceAll(statName, "Degredation", "Degradation")
	}

	if spEffectArg == "MultiplicativeStatType.Unknown29" {
		statName = "Hollowing HP Reduction"
	}

	if spEffectArg == "MultiplicativeStatType.Unknown61" {
		statName = "Physical Damage From Behind"
	}

	if spEffectArg == "MultiplicativeStatType.Unknown62" {
		statName = "Elemental Damage From Behind"
	}

	if statName == "Anti Stability Modifier" {
		statName = "Stamina Damage"
	}

	return statName
}

func parseModifyStatMultiplicatively(arg0, arg2 string) output.Modifier {
	var (
		statName string
		value    float64
	)

	// Generally, each statement has an arg that indicates the stat at arg 0
	statName = extractStatName(arg0)

	// It looks like most statements have 2 numeric args: index 1 that's always 0 and index 2 that is the actual value
	if num, err := strconv.ParseFloat(arg2, 64); err == nil {
		if num != 0 {
			value = num
		}
	}

	// convert to percentage
	value = (value - 1) * 100

	// round to two decimal places
	value = math.Round(value*100) / 100

	// derive verb
	verb := "Increase"
	if value < 1 {
		verb = "Decrease"
	}

	return output.Modifier{
		Description: fmt.Sprintf("%s %s by %.1f%%", verb, statName, math.Abs(value)),
		Target:      statName,
		Method:      output.ModifierMethodMultiplicative,
		Value:       (value / 100),
	}
}

func parseModifyStatAdditively(arg0, arg2 string) output.Modifier {
	var (
		statName string
		value    float64
	)

	// Generally, each statement has an arg that indicates the stat at arg 0
	statName = extractStatName(arg0)

	// It looks like most statements have 2 numeric args: index 1 that's always 0 and index 2 that is the actual value
	if num, err := strconv.ParseFloat(arg2, 64); err == nil {
		if num != 0 {
			value = num
		}
	}

	if statName == "Current HP Ratio" || statName == "Current HP Ratio2" {
		return output.Modifier{
			Description: fmt.Sprintf("Triggers an auxiliary effect while below %d%% HP", int(math.Abs(value))),
			Target:      statName,
			Method:      output.ModifierMethodAdditive,
			Value:       value,
		}
	}

	// derive verb
	verb := "Increase"
	if value < 0 {
		verb = "Decrease"
	}

	return output.Modifier{
		Description: fmt.Sprintf("%s %s by %d", verb, statName, int(math.Abs(value))),
		Target:      statName,
		Method:      output.ModifierMethodAdditive,
		Value:       value,
	}
}

func parseModifyAttributeBasedOnCurrentValue(arg0 string) output.Modifier {
	// Generally, each statement has an arg that indicates the stat at arg 0
	statName := extractStatName(arg0)

	// ModifyAttributeBasedOnCurrentValue always just reduces the attribute by 1
	return output.Modifier{
		Description: fmt.Sprintf("Decrease %s by 1", statName),
		Target:      statName,
		Method:      output.ModifierMethodAdditive,
		Value:       -1,
	}
}

func parseModifyDamageFlatToBaseAr(arg1, arg2 string) ([]output.Modifier, error) {
	// ModifyDamageFlatToBaseAr has 2 args:
	// Arg0 is always 0
	// Arg1 is the damage type
	// Arg2 is the modifier value
	// ModifyDamageFlatToBaseAr has the following effects:
	//     Weapons with innate affliction of the relevant damage type get Arg2 points of extra damage.
	//     Weapons infused with the relevant damage type get (Arg2 / 2) points of extra damage.
	//     Weapons with innate relevant damage AND infused with the relevant damage type get (Arg2 * 1.2) points of extra damage.
	var (
		damageType   string
		value        float64
		newModifiers = []output.Modifier{}
	)

	// ModifyDamageFlatToBaseAr has the buff type at index 1
	damageType = extractStatName(arg1)
	infusion := extractStatName(arg1)
	if infusion == "All Physical Damage" {
		infusion = "Standard"
	}

	value, err := strconv.ParseFloat(arg2, 64)
	if err != nil {
		err = fmt.Errorf("failed to convert value to float: %w", err)
		return nil, err
	}

	newModifiers = append(newModifiers, output.Modifier{
		Description: fmt.Sprintf("Increase %s by %d when the weapon is infused with %s", damageType, int(math.Abs(value/2)), infusion),
		Target:      damageType,
		Method:      output.ModifierMethodAdditive,
		Value:       value / 2,
	})

	newModifiers = append(newModifiers, output.Modifier{
		Description: fmt.Sprintf("Increase %s by %d when the weapon has innate %s", damageType, int(math.Abs(value)), infusion),
		Target:      damageType,
		Method:      output.ModifierMethodAdditive,
		Value:       value,
	})

	newModifiers = append(newModifiers, output.Modifier{
		Description: fmt.Sprintf("Increase %s by %d when the weapon has innate AND infused %s", damageType, int(math.Abs(value*1.2)), infusion),
		Target:      damageType,
		Method:      output.ModifierMethodAdditive,
		Value:       value * 1.2,
	})

	return newModifiers, nil
}

func parseSetStateInfo(arg0 string) (output.Modifier, error) {
	// SetStateInfo will always be special cases (flag flips)

	switch arg0 {
	case "StateInfoType.DisableBackstab":
		// StateInfoType.DisableBackstab is a flag for whether backstab is enabled
		return output.Modifier{
			Description: "Cannot be backstabbed",
			Target:      "Backstab",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.DisableFootstepSound":
		// StateInfoType.DisableFootstepSound is a flag for whether footstep sound is enabled
		return output.Modifier{
			Description: "Disable Footstep Sound",
			Target:      "Footstep Sound",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.DisableSoulLoss":
		// DisableSoulLoss is a flag for whether soul loss is enabled
		return output.Modifier{
			Description: "Disable Soul Loss On Death",
			Target:      "Soul Loss On Death",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.DisableHollowingAndSoulLoss":
		// DisableHollowingAndSoulLoss is a flag for whether hollowing and soul loss is enabled
		return output.Modifier{
			Description: "Disable Hollowing and Soul Loss On Death",
			Target:      "Hollowing and Soul Loss On Death",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.ReduceBackDamage":
		// ReduceBackDamage is a flag for whether back damage is reduced
		return output.Modifier{
			Description: "Reduce Damage From Behind",
			Target:      "Damage From Behind",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.HideWeaponLeft":
		// HideWeaponLeft is a flag for whether the left weapon is hidden
		return output.Modifier{
			Description: "Hide Left Hand Weapon",
			Target:      "Left Weapon",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.HideWeaponRight":
		// HideWeaponRight is a flag for whether the right weapon is hidden
		return output.Modifier{
			Description: "Hide Right Hand Weapon",
			Target:      "Right Weapon",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.DisableChrPhantomColor":
		// DisableChrPhantomColor is a flag for whether phantom color is enabled
		return output.Modifier{
			Description: "Disable Phantom Color",
			Target:      "Phantom Color",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.DisableSoulAbsorption":
		// DisableSoulAbsorption is a flag for whether one can gain souls
		return output.Modifier{
			Description: "Disable Soul Absorption",
			Target:      "Soul Absorption",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.Unknown0":
		// StateInfoType.Unknown0 is a flag for whether headshot stagger is enabled
		return output.Modifier{
			Description: "Disable Headshot Stagger",
			Target:      "Headshot Stagger",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.Unknown4":
		// StateInfoType.Unknown4 enables a 5% crit rate wherein crits deal 50% more damage
		return output.Modifier{
			Description: "Enable 5% chance for crits, which deal 50% more damage",
			Target:      "Critical Damage",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.Unknown6":
		// StateInfoType.Unknown6 is a flag for whether one takes on the appearance of a human even when hollow
		return output.Modifier{
			Description: "Force human appearance even when hollow",
			Target:      "Human Appearance",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.Unknown19":
		// StateInfoType.Unknown19 is a flag for increasing the wearer's item discovery by 50
		return output.Modifier{
			Description: "Increase Item Discovery by 50",
			Target:      "Item Discovery",
			Method:      output.ModifierMethodAdditive,
		}, nil
	case "StateInfoType.Unknown29":
		// StateInfoType.Unknown29 is a flag for whether headshot damage is enabled
		return output.Modifier{
			Description: "Disable Headshot Damage",
			Target:      "Headshot Damage",
			Method:      output.ModifierMethodToggle,
		}, nil
	case "StateInfoType.Unknown31":
		// StateInfoType.Unknown31 is a flag for making the wearer invisible while rolling
		return output.Modifier{
			Description: "Invisible while rolling",
			Target:      "Invisible While Rolling",
			Method:      output.ModifierMethodToggle,
		}, nil
	default:
		return output.Modifier{}, fmt.Errorf("unhandled flag: %s", arg0)
	}
}

func parseModifySpellEffectLength(arg0 string) (output.Modifier, error) {
	// ModifySpellEffectLength has the buff type at index 0
	value, err := strconv.ParseFloat(arg0, 64)
	if err != nil {
		return output.Modifier{}, err
	}

	// Derive verb
	verb := "Increase"
	if value < 0 {
		verb = "Decrease"
	}

	statName := "Spell Effect Length"

	// Convert to percent
	value -= 1.0
	value *= 100

	return output.Modifier{
		Description: fmt.Sprintf("%s %s by %.1f%%", verb, statName, math.Abs(value)),
		Target:      statName,
		Method:      output.ModifierMethodMultiplicative,
	}, nil
}

func parseModifyBulletParam(arg2, arg3 string) (output.Modifier, error) {
	// ModifyBulletParam extends the range of bows
	// Arg0 seems to always be 0
	// Arg1 seems to always refer to regular bows
	// Arg2 is always 0, 1, or 2 (0 is regualr bows, 1 is greatbows, 2 is crossbows)
	// Arg3 is the modifier value
	// Arg4 is always 1
	// Arg5 is also the modifier value
	// Arg6 is also the modifier value
	// Arg7 is always 1

	var bowType string
	switch arg2 {
	case "0":
		bowType = "Bow"
	case "1":
		bowType = "Greatbow"
	case "2":
		bowType = "Crossbow"
	default:
		return output.Modifier{}, fmt.Errorf("unhandled bow type: %s", arg2)
	}

	value, err := strconv.ParseFloat(arg3, 64)
	if err != nil {
		return output.Modifier{}, err
	}

	// Convert to percent
	value -= 1.0
	value *= 100

	statName := fmt.Sprintf("%s Range", bowType)

	// Derive verb
	verb := "Increase"
	if value < 0 {
		verb = "Decrease"
	}

	return output.Modifier{
		Description: fmt.Sprintf("%s %s by %.1f%%", verb, statName, math.Abs(value)),
		Target:      statName,
		Method:      output.ModifierMethodAdditive,
	}, nil
}

func parseModifyProperty(arg0, arg2 string) (output.Modifier, error) {
	// ModifyProperty is another special case with flag flipping and such
	// Arg0 is the flag
	// Arg1 is always 0
	// Arg2 is an optional modifier value

	switch arg0 {
	case "PlayerPropertyType.Appearence":
		// PlayerPropertyType.Appearence is a flag for the player appearing as a white phantom
		return output.Modifier{
			Description: "Force white phantom appearance",

			Target: "White Phantom Appearance",
			Method: output.ModifierMethodToggle,
		}, nil
	case "PlayerPropertyType.HexSelfDamageStaff":
		// PlayerPropertyType.HexSelfDamageStaff is a flag for the player being damaged by casting hexes
		damage, err := strconv.Atoi(arg2)
		if err != nil {
			err = fmt.Errorf("failed to convert arg2 to int for PlayerPropertyType.HexSelfDamageStaff: %w", err)
			return output.Modifier{}, err
		}
		return output.Modifier{
			Description: fmt.Sprintf("Take %d damage every time the player casts a hex from a staff", damage),

			Target: "Hex Self Damage Staff",
			Method: output.ModifierMethodToggle,
		}, nil
	case "PlayerPropertyType.HexSelfDamageChime":
		// PlayerPropertyType.HexSelfDamageChime is a flag for the player being damaged by casting hexes
		damage, err := strconv.Atoi(arg2)
		if err != nil {
			err = fmt.Errorf("failed to convert arg2 to int for PlayerPropertyType.HexSelfDamageChime: %w", err)
			return output.Modifier{}, err
		}

		return output.Modifier{
			Description: fmt.Sprintf("Take %d damage every time the player casts a hex from a chime", damage),

			Target: "Hex Self Damage Chime",
			Method: output.ModifierMethodToggle,
		}, nil
	case "PlayerPropertyType.Unknown21":
		// PlayerPropertyType.Unknown21 is a flag for the player being wet
		return output.Modifier{
			Description: "Soaks the wearer",
			Target:      "Soak",
			Method:      output.ModifierMethodToggle,
		}, nil
	default:
		return output.Modifier{}, fmt.Errorf("unhandled flag: %s", arg0)
	}
}

func parseModifyStaminaRecovery(arg1 string) (output.Modifier, error) {
	// ModifyStaminaRecovery has the buff value at index 1
	value, err := strconv.ParseFloat(arg1, 64)
	if err != nil {
		err = fmt.Errorf("failed to convert value to float: %w", err)
		return output.Modifier{}, err
	}

	// Convert to percent
	value -= 1
	value *= 100

	return output.Modifier{
		Description: fmt.Sprintf("Increase Stamina Recovery by %.1f%%", value),
		Target:      "Stamina Recovery",
		Method:      output.ModifierMethodMultiplicative,
		Value:       value,
	}, nil
}

func parseBuffDefense(arg1, arg2 string) (output.Modifier, error) {
	// BuffDefense has the stat at arg 1 and the buff value at arg 2
	statName := extractStatName(arg1)
	statName = strings.ReplaceAll(statName, " Damage", " Defense")

	value, err := strconv.Atoi(arg2)
	if err != nil {
		err = fmt.Errorf("failed to convert value to int: %w", err)
		return output.Modifier{}, err
	}

	return output.Modifier{
		Description: fmt.Sprintf("Increase %s by %d", statName, value),
		Target:      statName,
		Method:      output.ModifierMethodAdditive,
		Value:       float64(value),
	}, nil
}

func parseUnknownCommand10004002(arg0, arg1 string) (output.Modifier, error) {
	// UnknownCommand10004002 is for modifying HP on a time interval
	// Arg0 is the number of seconds in the interval
	// Arg1 is the modifier value

	interval, err := strconv.Atoi(arg0)
	if err != nil {
		err = fmt.Errorf("failed to convert interval to int: %w", err)
		return output.Modifier{}, err
	}

	value, err := strconv.Atoi(arg1)
	if err != nil {
		err = fmt.Errorf("failed to convert value to int: %w", err)
		return output.Modifier{}, err
	}

	return output.Modifier{
		Description: fmt.Sprintf("Increase HP by %d every %d second(s)", value, interval),
		Target:      "HP Regen",
		Method:      output.ModifierMethodAdditive,
		Value:       float64(value),
	}, nil
}

func parseRandomlySpawnBulletAfterTakingDamageDurabilitybasedRandomness(arg0, arg1 string) (output.Modifier, error) {
	// RandomlySpawnBulletAfterTakingDamageDurabilitybasedRandomness is too long of a name (/_\;)
	// This has 5 args and is seemingly only used by the Old Sun Ring
	// Arg0 is the max probability of triggering the effect
	// Arg1 is (probably) the min probability of triggering the effect
	// Arg2 could also be the min probability of triggering the effect, not sure
	// Arg3 is an ID, presumably the effect being triggered
	// Arg4 is also an ID. For the Old Sun Ring, it's the same ID as Arg3. Not sure what it does

	maxProbability, err := strconv.Atoi(arg0)
	if err != nil {
		err = fmt.Errorf("failed to convert max probability to int: %w", err)
		return output.Modifier{}, err
	}

	minProbability, err := strconv.Atoi(arg1)
	if err != nil {
		err = fmt.Errorf("failed to convert min probability to int: %w", err)
		return output.Modifier{}, err
	}

	return output.Modifier{
		Description: fmt.Sprintf("Chance to trigger an explosion on taking damage with a probability between %d%% and %d%%, scaling with remaining durability", minProbability, maxProbability),
		Target:      "Trigger Explosion",
		Method:      output.ModifierMethodToggle,
	}, nil
}

func parseRandomlySpawnBulletAfterTakingDamageFlatRandomness(arg0 string) (output.Modifier, error) {
	// RandomlySpawnBulletAfterTakingDamageFlatRandomness has a chance to spawn thorns after taking damage
	// Arg0 is the probability of triggering the effect
	// Arg1 is a number. E.g. 10, 9, 8. Not sure what it does
	// Arg2 is always 50
	// Arg3 is a number. E.g. 500, 400, 300. Not sure what it does
	// Arg4 is an ID. Presumably the effect
	// Arg5 is also an ID. For the Ring of Thorns, it's the same ID as Arg4

	probability, err := strconv.Atoi(arg0)
	if err != nil {
		err = fmt.Errorf("failed to convert probability to int: %w", err)
		return output.Modifier{}, err
	}

	return output.Modifier{
		Description: fmt.Sprintf("%d%% chance to counterattack on taking damage", probability),
		Target:      "Counterattack",
		Method:      output.ModifierMethodToggle,
	}, nil
}

func parseInitiatePassiveServermediatedMultiplayerItem(arg0 string) (output.Modifier, error) {
	// InitiatePassiveServermediatedMultiplayerItem facilitates covenant invasions
	// Arg0 is the covenant to which one belongs

	arg0Breakdown := strings.Split(arg0, ".")

	if arg0Breakdown[0] != "PassiveServerItemType" {
		return output.Modifier{}, fmt.Errorf("unhandled arg0: %s", arg0)
	}

	var covenant string

	switch arg0Breakdown[1] {
	case "CrestoftheRat":
		covenant = "Rat King"
	case "BellKeepersSoul":
		covenant = "Bell Keepers"
	case "GuardiansSeal":
		covenant = "Blue Seninels"
	default:
		return output.Modifier{}, fmt.Errorf("unhandled PassiveServerItemType: %s", arg0)
	}

	return output.Modifier{
		Description: fmt.Sprintf("Covenant Invasion: %s", covenant),
		Target:      "Covenant Invasion",
		Method:      output.ModifierMethodToggle,
	}, nil
}

func parseNerfDefense(arg1, arg2 string) (output.Modifier, error) {
	// NerfDefense has 3 args
	// Arg0 is always 0
	// Arg1 is the defense stat to nerf
	// Arg2 is the nerf value

	statName := extractStatName(arg1)
	statName = strings.ReplaceAll(statName, " Damage", " Defense")

	value, err := strconv.Atoi(arg2)
	if err != nil {
		err = fmt.Errorf("failed to convert value to int: %w", err)
		return output.Modifier{}, err
	}

	return output.Modifier{
		Description: fmt.Sprintf("Decrease %s by %d", statName, value),
		Target:      statName,
		Method:      output.ModifierMethodAdditive,
		Value:       float64(value * -1),
	}, nil
}

func parseApplySoulScalingToWeapon(arg1, arg2, arg3, arg4 string) (output.Modifier, error) {
	// ApplySoulScalingToWeapon scales weapon damage based on the number of souls held
	// This modifier has 6 args
	// Arg0 is always 0
	// Arg1 is the minimum number of souls to apply the scaling
	// Arg2 is the maximum number of souls to apply the scaling
	// Arg3 is the minimum scaling value
	// Arg4 is the maximum scaling value
	// Arg5 is always 0

	minSouls, err := strconv.Atoi(arg1)
	if err != nil {
		err = fmt.Errorf("failed to convert min souls to int: %w", err)
		return output.Modifier{}, err
	}

	maxSouls, err := strconv.Atoi(arg2)
	if err != nil {
		err = fmt.Errorf("failed to convert max souls to int: %w", err)
		return output.Modifier{}, err
	}

	minScaling, err := strconv.ParseFloat(arg3, 64)
	if err != nil {
		err = fmt.Errorf("failed to convert min scaling to float: %w", err)
		return output.Modifier{}, err
	}

	// convert to percentage
	minScaling *= 100

	maxScaling, err := strconv.ParseFloat(arg4, 64)
	if err != nil {
		err = fmt.Errorf("failed to convert max scaling to float: %w", err)
		return output.Modifier{}, err
	}

	// convert to percentage
	maxScaling *= 100

	return output.Modifier{
		Description: fmt.Sprintf("Scale weapon damage based on number of souls held. Min souls: %d, Max souls: %d, Min scaling: %.1f%%, Max scaling: %.1f%%", minSouls, maxSouls, minScaling, maxScaling),
		Target:      "Weapon Damage Scaling",
		Method:      output.ModifierMethodMultiplicative,
	}, nil
}

func parseApplySpecialScalingToWeapon(arg0, arg1, arg2, arg3 string) (output.Modifier, error) {
	// ApplySpecialScalingToWeapon scales weapon damage based on a non-standard scaling stat
	// This modifier has 5 args
	// Arg0 is the type of scaling to apply
	// Arg1 is the max value of the scaling stat for scaling purposes (assume the min is 1?)
	// Arg2 is the min scaling rate
	// Arg3 is the max scaling rate
	// Arg4 is always 0

	statName := extractStatName(arg0)

	maxStatValue, err := strconv.Atoi(arg1)
	if err != nil {
		err = fmt.Errorf("failed to convert max stat value to int: %w", err)
		return output.Modifier{}, err
	}

	minScaling, err := strconv.ParseFloat(arg2, 64)
	if err != nil {
		err = fmt.Errorf("failed to convert min scaling to float: %w", err)
		return output.Modifier{}, err
	}

	maxScaling, err := strconv.ParseFloat(arg3, 64)
	if err != nil {
		err = fmt.Errorf("failed to convert max scaling to float: %w", err)
		return output.Modifier{}, err
	}

	// convert to percentage
	if minScaling < maxScaling {
		minScaling += 1
		maxScaling += 1
	}
	minScaling *= 100
	maxScaling *= 100

	return output.Modifier{
		Description: fmt.Sprintf("Scale weapon damage based on %s. Min value: 1, Max value: %d, Min scaling: %.1f%%, Max scaling: %.1f%%", statName, maxStatValue, minScaling, maxScaling),
		Target:      "Weapon Damage Scaling",
		Method:      output.ModifierMethodMultiplicative,
	}, nil
}

func parseApplyNgScalingToWeapon(arg0, arg1, arg2, arg3 string) (output.Modifier, error) {
	// ApplyNgScalingToWeapon adds weapon damage based on the number of the current NG cycle
	// This modifier has 5 args
	// Arg0 is the minimum NG cycle to apply the scaling (i.e. the first "Journey" upon which the scaling is applied)
	// Arg1 is the maximum NG cycle to apply the scaling
	// Arg2 is the minimum attack value to add
	// Arg3 is the maximum attack value to add
	// Arg4 is not always 0, but idk what it does

	minCycle, err := strconv.Atoi(arg0)
	if err != nil {
		err = fmt.Errorf("failed to convert min cycle to int: %w", err)
		return output.Modifier{}, err
	}

	maxCycle, err := strconv.Atoi(arg1)
	if err != nil {
		err = fmt.Errorf("failed to convert max cycle to int: %w", err)
		return output.Modifier{}, err
	}

	minAttack, err := strconv.Atoi(arg2)
	if err != nil {
		err = fmt.Errorf("failed to convert min attack to int: %w", err)
		return output.Modifier{}, err
	}

	maxAttack, err := strconv.Atoi(arg3)
	if err != nil {
		err = fmt.Errorf("failed to convert max attack to int: %w", err)
		return output.Modifier{}, err
	}

	return output.Modifier{
		Description: fmt.Sprintf("Scale weapon damage based on NG cycle. Min cycle: %d, Max cycle: %d, Min attack bonus: %d, Max attack bonus: %d", minCycle, maxCycle, minAttack, maxAttack),
		Target:      "Weapon Damage",
		Method:      output.ModifierMethodAdditive,
	}, nil
}

func createModifiers(spEffects []emevdParser.Statement) ([]output.Modifier, error) {
	modifiers := []output.Modifier{}

	for _, spEffect := range spEffects {
		var newModifiers []output.Modifier

		switch spEffect.Name {
		case "ModifyStatMultiplicatively":
			if spEffect.Args[0] == "MultiplicativeStatType.Unknown46" {
				continue
			}
			newModifiers = []output.Modifier{parseModifyStatMultiplicatively(spEffect.Args[0], spEffect.Args[2])}

		case "ModifyStatAdditively":
			// Not sure what this does, but it is used by the Ivory Warrior Ring
			if spEffect.Args[0] == "AdditiveStatType.Unknown22" {
				continue
			}
			newModifiers = []output.Modifier{parseModifyStatAdditively(spEffect.Args[0], spEffect.Args[2])}

		case "ModifyAttributeBasedOnCurrentValue":
			newModifiers = []output.Modifier{parseModifyAttributeBasedOnCurrentValue(spEffect.Args[0])}

		case "ModifyDamageFlatToBaseAr":
			var err error
			newModifiers, err = parseModifyDamageFlatToBaseAr(spEffect.Args[1], spEffect.Args[2])
			if err != nil {
				err = fmt.Errorf("error parsing ModifyDamageFlatToBaseAr: %w", err)
				return nil, err
			}

		case "SetStateInfo":
			// Ignore, not sure what this does
			if spEffect.Args[0] == "StateInfoType.Unknown23" {
				continue
			}
			modifier, err := parseSetStateInfo(spEffect.Args[0])
			if err != nil {
				err = fmt.Errorf("error parsing SetStateInfo: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "ModifySpellEffectLength":
			modifier, err := parseModifySpellEffectLength(spEffect.Args[0])
			if err != nil {
				err = fmt.Errorf("error parsing ModifySpellEffectLength: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "ModifyBulletParam":
			// A modifier of x1.0 does nothing
			if spEffect.Args[3] == "1" {
				continue
			}
			modifier, err := parseModifyBulletParam(spEffect.Args[2], spEffect.Args[3])
			if err != nil {
				err = fmt.Errorf("error parsing ModifyBulletParam: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "ModifyProperty":
			// Ignore these modifiers
			if spEffect.Args[0] == "PlayerPropertyType.SfxIDOnKill" ||
				spEffect.Args[0] == "PlayerPropertyType.SfxOnKillOrigin" {
				continue
			}
			modifier, err := parseModifyProperty(spEffect.Args[0], spEffect.Args[2])
			if err != nil {
				err = fmt.Errorf("error parsing ModifyProperty: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "ModifyStaminaRecovery":
			modifier, err := parseModifyStaminaRecovery(spEffect.Args[1])
			if err != nil {
				err = fmt.Errorf("error parsing ModifyStaminaRecovery: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "BuffDefense":
			if spEffect.Args[1] == "BuffType.Unknown7" {
				continue
			}

			modifier, err := parseBuffDefense(spEffect.Args[1], spEffect.Args[2])
			if err != nil {
				err = fmt.Errorf("error parsing BuffDefense: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "UnknownCommand10004002":
			modifier, err := parseUnknownCommand10004002(spEffect.Args[0], spEffect.Args[1])
			if err != nil {
				err = fmt.Errorf("error parsing UnknownCommand10004002: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "ModifyEnemyAggroToPlayer":
			// ModifyEnemyAggroToPlayer has a single arg, but I don't know what it does
			newModifiers = []output.Modifier{
				{
					Description: "Increase enemy aggro",
					Target:      "Aggro",
					Method:      output.ModifierMethodToggle,
				},
			}

		case "RandomlySpawnBulletAfterTakingDamageDurabilitybasedRandomness":
			modifier, err := parseRandomlySpawnBulletAfterTakingDamageDurabilitybasedRandomness(spEffect.Args[0], spEffect.Args[1])
			if err != nil {
				err = fmt.Errorf("error parsing RandomlySpawnBulletAfterTakingDamageDurabilitybasedRandomness: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "RandomlySpawnBulletAfterTakingDamageFlatRandomness":
			modifier, err := parseRandomlySpawnBulletAfterTakingDamageFlatRandomness(spEffect.Args[0])
			if err != nil {
				err = fmt.Errorf("error parsing RandomlySpawnBulletAfterTakingDamageFlatRandomness: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "ResetInvasionCooldownTimer":
			// ResetInvasionCooldownTimer has no args
			newModifiers = []output.Modifier{
				{
					Description: "Disable invasion cooldown timer",
					Target:      "Invasion Cooldown Timer",
					Method:      output.ModifierMethodToggle,
				},
			}

		case "InitiatePassiveServermediatedMultiplayerItem":
			modifier, err := parseInitiatePassiveServermediatedMultiplayerItem(spEffect.Args[0])
			if err != nil {
				err = fmt.Errorf("error parsing InitiatePassiveServermediatedMultiplayerItem: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "ReplaceFistsWithAlternateWeapon":
			// ReplaceFistsWithAlternateWeapon has 4 args, none of which are consequential
			newModifiers = []output.Modifier{
				// Increase base damage with bare fists
				{
					Description: "Increase base damage of bare fists to 200",
					Target:      "Bare Fist Base Damage",
					Method:      output.ModifierMethodToggle,
				},
				// Increase damage scaling with bare fists
				{
					Description: "Increase damage scaling of bare fists",
					Target:      "Bare Fist Damage Scaling",
					Method:      output.ModifierMethodToggle,
				},
				// Increase damage on ladders
				{
					Description: "Increase damage on ladders",
					Target:      "Ladder Damage",
					Method:      output.ModifierMethodToggle,
				},
				// Enable powerstancing bare fists
				{
					Description: "Enable powerstancing bare fists",
					Target:      "Powerstancing Bare Fist",
					Method:      output.ModifierMethodToggle,
				},
			}

		case "ModifyEstusUsage":
			// ModifyEstusUsage has two args, neither of which are consequential
			// Arg0 specifies what kind of buff this is, but it's only used once by the Ancient Dragon Seal
			// Arg1 is always 0
			newModifiers = []output.Modifier{
				{
					Description: "Increase estus healing by 50",
					Target:      "Estus Healing",
					Method:      output.ModifierMethodToggle,
				},
			}

		case "NerfDefense":
			modifier, err := parseNerfDefense(spEffect.Args[1], spEffect.Args[2])
			if err != nil {
				err = fmt.Errorf("error parsing NerfDefense: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "ApplyEquipLoadbasedFlatDamageToWeapon":
			// ApplyEquipLoadbasedFlatDamageToWeapon has 7 args, none of which are consequential since it's only used once for Flynn's Ring
			newModifiers = []output.Modifier{
				{
					Description: "Increase Physical Attack Rating by 50 at or less than 60.4 Maximum Equip Load, inversely scaling linearly to 0 at 85.4 Maximum Equip Load",
					Target:      "Physical Attack Rating",
					Method:      output.ModifierMethodToggle,
				},
			}

		case "AddSpellDeflectChance":
			// AddSpellDeflectChance has 2 args, neither of which are consequential since it's only used once for Yorgh's Ring
			newModifiers = []output.Modifier{
				{
					Description: "50% chance to deflect spells",
					Target:      "Spell Deflection Chance",
					Method:      output.ModifierMethodToggle,
				},
				{
					Description: "Enables spell deflection for shield parries",
					Target:      "Shiled Parry Spell Deflection",
					Method:      output.ModifierMethodToggle,
				},
			}

		case "ApplySoulScalingToWeapon":
			modifier, err := parseApplySoulScalingToWeapon(spEffect.Args[1], spEffect.Args[2], spEffect.Args[3], spEffect.Args[4])
			if err != nil {
				err = fmt.Errorf("error parsing ApplySoulScalingToWeapon: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "ApplySpecialScalingToWeapon":
			modifier, err := parseApplySpecialScalingToWeapon(spEffect.Args[0], spEffect.Args[1], spEffect.Args[2], spEffect.Args[3])
			if err != nil {
				err = fmt.Errorf("error parsing ApplySpecialScalingToWeapon: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		case "ApplyNgScalingToWeapon":
			modifier, err := parseApplyNgScalingToWeapon(spEffect.Args[0], spEffect.Args[1], spEffect.Args[2], spEffect.Args[3])
			if err != nil {
				err = fmt.Errorf("error parsing ApplyNgScalingToWeapon: %w", err)
				return nil, err
			}
			newModifiers = []output.Modifier{modifier}

		default:
			return nil, fmt.Errorf("unhandled SpEffect name for Statement: %+v", spEffect)
		}

		for _, newModifier := range newModifiers {
			if newModifier.Target == "" {
				return nil, fmt.Errorf("no stat name in Statement: %+v", spEffect)
			}

			newModifier.TargetType = output.ModifierTargetTypeSpecial
			if slices.Contains(attributeNames, newModifier.Target) {
				newModifier.TargetType = output.ModifierTargetTypeAttribute
			}
			if slices.Contains(statNames, newModifier.Target) {
				newModifier.TargetType = output.ModifierTargetTypeStat
			}

			modifiers = append(modifiers, newModifier)
		}
	}

	return modifiers, nil
}
