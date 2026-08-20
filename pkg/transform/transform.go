// Package transform contains functions for transforming data from DS2 params to Scholar JSON
package transform

import (
	"fmt"
	"math"
	"reflect"
	"strings"

	emevdParser "github.com/Camburgaler/scholar-utils/pkg/data/emevd/parse"
	"github.com/Camburgaler/scholar-utils/pkg/data/param"
	paramParser "github.com/Camburgaler/scholar-utils/pkg/data/param/parse"
	"github.com/Camburgaler/scholar-utils/pkg/output"
)

var (
	Infusions = []string{
		"Standard",
		"Lightning",
		"Magic",
		"Fire",
		"Dark",
		"Poison",
		"Bleed",
		"Raw",
		"Enchanted",
		"Mundane",
		"Boss",
		"Special",
	}

	levelUpStatusCalcParamIndexToAttribute = map[int]string{
		0: "Vigor",
		1: "Endurance",
		2: "Vitality",
		3: "Attunement",
		4: "Strength",
		5: "Dexterity",
		6: "Adaptability",
		7: "Intelligence",
		8: "Faith",
	}

	noRing = output.Ring{
		Equippable: output.Equippable{
			Name:       "No Ring",
			Modifiers:  []output.Modifier{},
			Weight:     0,
			Durability: 0,
			RepairCost: 0,
		},
		ItemDiscovery: 0,
	}
)

func createClasses(playerStatusParams []param.PlayerStatus) []output.Class {
	classes := []output.Class{}

	fmt.Println("\nCreating classes...")

	for _, playerStatusParam := range playerStatusParams {
		if playerStatusParam.IsClass() {
			classes = append(classes, output.Class{
				Name:  strings.TrimPrefix(playerStatusParam.Name, "[Class] "),
				Level: playerStatusParam.Level,
				Attributes: output.Attributes[int]{
					Vigor:        playerStatusParam.Vigor,
					Endurance:    playerStatusParam.Endurance,
					Vitality:     playerStatusParam.Vitality,
					Attunement:   playerStatusParam.Attunement,
					Adaptability: playerStatusParam.Adaptability,
					Strength:     playerStatusParam.Strength,
					Dexterity:    playerStatusParam.Dexterity,
					Intelligence: playerStatusParam.Intelligence,
					Faith:        playerStatusParam.Faith,
				},
			})
		}
	}

	fmt.Printf("Created %d classes\n", len(classes))

	return classes
}

func createAttributeToStatMap(levelUpStatusCalcParams []param.LevelUpStatusCalc) output.Attributes[output.Stats[bool]] {
	attributesToStatMap := output.Attributes[output.Stats[bool]]{}
	vAttributesToStatMap := reflect.ValueOf(&attributesToStatMap).Elem()
	fields := reflect.VisibleFields(vAttributesToStatMap.Type())

	fmt.Println("\nCreating attribute to stat map...")

	for i, levelUpStatusCalcParam := range levelUpStatusCalcParams {
		vLevelUpStatusCalcParam := reflect.ValueOf(&levelUpStatusCalcParam).Elem()

		for _, field := range fields {
			if field.Anonymous {
				continue
			}

			if levelUpStatusCalcParamIndexToAttribute[i] == field.Name {
				stats := output.Stats[bool]{}
				vStats := reflect.ValueOf(&stats).Elem()

				for field := range vStats.Fields() {
					vStats.FieldByName(field.Name).SetBool(vLevelUpStatusCalcParam.FieldByName(field.Name).Bool())
				}

				vAttributesToStatMap.FieldByName(field.Name).Set(reflect.ValueOf(stats))
			}
		}
	}

	fmt.Println("Created attribute to stat map")

	return attributesToStatMap
}

func createCovenants(vowParams []param.Vow) []string {
	fmt.Println("\nCreating covenants...")

	covenants := []string{}

	for _, vowParam := range vowParams {
		covenants = append(covenants, vowParam.Name)
	}

	fmt.Printf("Created %d covenants\n", len(covenants))

	return covenants
}

func createLevels(playerLevelUpSoulsParams []param.PlayerLevelUpSouls) []int {
	fmt.Println("\nCreating levels...")

	levels := []int{}

	for _, playerLevelUpSoulsParam := range playerLevelUpSoulsParams {
		levels = append(levels, playerLevelUpSoulsParam.NecessarySouls)
	}

	fmt.Printf("Created %d levels\n", len(levels))

	return levels
}

func createSpells(spellParams []param.Spell) []output.Spell {
	fmt.Println("\nCreating spells...")

	spells := []output.Spell{}

	for _, spellParam := range spellParams {
		spells = append(spells, output.Spell{
			Name:                 spellParam.Name,
			RequiredIntelligence: spellParam.RequiredIntelligence,
			RequiredFaith:        spellParam.RequiredFaith,
			SpellSlotCost:        spellParam.SpellSlotCost,
			UsageCountCurve: []int{
				spellParam.UsageCountLV1,
				spellParam.UsageCountLV2,
				spellParam.UsageCountLV3,
				spellParam.UsageCountLV4,
				spellParam.UsageCountLV5,
				spellParam.UsageCountLV6,
				spellParam.UsageCountLV7,
				spellParam.UsageCountLV8,
				spellParam.UsageCountLV9,
				spellParam.UsageCountLV10,
			},
		})
	}

	fmt.Printf("Created %d spells\n", len(spells))

	return spells
}

func createArmor(armorParams []param.Armor, armorReinforceParams []param.ArmorReinforce, itemParams []param.Item, armorSpEffects emevdParser.Events) ([]output.Armor, error) {
	armor := []output.Armor{}

	for _, armorParam := range armorParams {
		name := armorParam.Name

		if strings.Contains(name, "[Body]") {
			name = "No Armor"
		}

		// Armor ID of 113601XX is the visible Aurous set
		if math.Floor(float64(armorParam.ID)/100) == 113601 {
			name = armorParam.Name + " (Visible)"
		}

		// Armor ID of 113611XX is the invisible Aurous set
		if math.Floor(float64(armorParam.ID)/100) == 113611 {
			name = armorParam.Name + " (Invisible)"
		}

		// Armor ID of 12270100 is the Prisoner's Hood (Mask Only)
		if armorParam.ID == 12270100 {
			name = armorParam.Name + " (Mask Only)"
		}

		// Armor ID of 12270101 is the Prisoner's Tatters (Master's Attire)
		if armorParam.ID == 12270101 {
			name = armorParam.Name + " (Master's Attire)"
		}

		armorReinforceParam := param.ArmorReinforce{}
		for _, param := range armorReinforceParams {
			if param.ID == armorParam.ArmorReinforceID {
				armorReinforceParam = param
				break
			}
		}

		maxReinforcementLevel := armorReinforceParam.MaxReinforcementLevel
		if name == "No Armor" {
			maxReinforcementLevel = 0
		}

		itemParam := param.Item{}
		for _, param := range itemParams {
			if param.ArmorParamID == armorParam.ID {
				itemParam = param
				break
			}
		}

		spEffects := armorSpEffects[itemParam.SpecialEffectID]
		modifiers, err := createModifiers(spEffects)
		if err != nil {
			err = fmt.Errorf("failed to create armor modifiers for %s: %w", armorParam.Name, err)
			return nil, err
		}

		armor = append(armor, output.Armor{
			Equippable: output.Equippable{
				Name:       name,
				Modifiers:  modifiers,
				Weight:     armorParam.Weight,
				Durability: armorParam.Durability,
				RepairCost: armorParam.RepairCost,
			},
			Defenses: output.Defenses{
				Standard:  calculateSlopeIntercept(Point{X: 0, Y: float64(armorReinforceParam.MinDefenseStandard)}, Point{X: float64(maxReinforcementLevel), Y: float64(armorReinforceParam.MaxDefenseStandard)}),
				Slash:     calculateSlopeIntercept(Point{X: 0, Y: float64(armorReinforceParam.MinDefenseSlash)}, Point{X: float64(maxReinforcementLevel), Y: float64(armorReinforceParam.MaxDefenseSlash)}),
				Thrust:    calculateSlopeIntercept(Point{X: 0, Y: float64(armorReinforceParam.MinDefenseThrust)}, Point{X: float64(maxReinforcementLevel), Y: float64(armorReinforceParam.MaxDefenseThrust)}),
				Strike:    calculateSlopeIntercept(Point{X: 0, Y: float64(armorReinforceParam.MinDefenseStrike)}, Point{X: float64(maxReinforcementLevel), Y: float64(armorReinforceParam.MaxDefenseStrike)}),
				Magic:     calculateSlopeIntercept(Point{X: 0, Y: float64(armorReinforceParam.MinAbsorptionMagic)}, Point{X: float64(maxReinforcementLevel), Y: float64(armorReinforceParam.MaxAbsorptionMagic)}),
				Lightning: calculateSlopeIntercept(Point{X: 0, Y: float64(armorReinforceParam.MinAbsorptionLightning)}, Point{X: float64(maxReinforcementLevel), Y: float64(armorReinforceParam.MaxAbsorptionLightning)}),
				Fire:      calculateSlopeIntercept(Point{X: 0, Y: float64(armorReinforceParam.MinAbsorptionFire)}, Point{X: float64(maxReinforcementLevel), Y: float64(armorReinforceParam.MaxAbsorptionFire)}),
				Dark:      calculateSlopeIntercept(Point{X: 0, Y: float64(armorReinforceParam.MinAbsorptionDark)}, Point{X: float64(maxReinforcementLevel), Y: float64(armorReinforceParam.MaxAbsorptionDark)}),
			},
			DefenseScalingPhysical: armorParam.DefenseStatEffectScaling,
			DefenseScalingSlash:    armorParam.DefenseScalingSlash,
			DefenseScalingThrust:   armorParam.DefenseScalingThrust,
			DefenseScalingStrike:   armorParam.DefenseScalingStrike,
			Resistances: output.Resistances{
				Poison:  calculateSlopeIntercept(Point{X: 0, Y: float64(armorReinforceParam.MinAbsorptionPoison)}, Point{X: float64(maxReinforcementLevel), Y: float64(armorReinforceParam.MaxAbsorptionPoison)}),
				Bleed:   calculateSlopeIntercept(Point{X: 0, Y: float64(armorReinforceParam.MinAbsorptionBleed)}, Point{X: float64(maxReinforcementLevel), Y: float64(armorReinforceParam.MaxAbsorptionBleed)}),
				Petrify: calculateSlopeIntercept(Point{X: 0, Y: float64(armorReinforceParam.MinAbsorptionPetrify)}, Point{X: float64(maxReinforcementLevel), Y: float64(armorReinforceParam.MaxAbsorptionPetrify)}),
				Curse:   calculateSlopeIntercept(Point{X: 0, Y: float64(armorReinforceParam.MinAbsorptionCurse)}, Point{X: float64(maxReinforcementLevel), Y: float64(armorReinforceParam.MaxAbsorptionCurse)}),
			},
			Poise: armorParam.Poise,
			Requirements: output.ScalingAttributes[int]{
				Strength:     armorParam.PrerequisiteStrength,
				Dexterity:    armorParam.PrerequisiteDexterity,
				Intelligence: armorParam.PrerequisiteIntelligence,
				Faith:        armorParam.PrerequisiteFaith,
			},
			ItemDiscovery:         armorParam.ItemDiscovery,
			MaxReinforcementLevel: maxReinforcementLevel,
		})
	}

	return armor, nil
}

func createRings(ringParams []param.Ring, itemParams []param.Item, ringSpEffects emevdParser.Events) ([]output.Ring, error) {
	fmt.Println("\nCreating rings...")
	rings := []output.Ring{noRing}

	for _, ringParam := range ringParams {
		itemParam := param.Item{}
		for _, param := range itemParams {
			if param.ID == ringParam.ID {
				itemParam = param
				break
			}
		}

		spEffects := ringSpEffects[itemParam.SpecialEffectID]
		modifiers, err := createModifiers(spEffects)
		if err != nil {
			err = fmt.Errorf("error creating ring modifiers for %s: %w", ringParam.Name, err)
			return nil, err
		}

		rings = append(rings, output.Ring{
			Equippable: output.Equippable{
				Name:       ringParam.Name,
				Modifiers:  modifiers,
				Weight:     ringParam.Weight,
				Durability: ringParam.Durability,
				RepairCost: ringParam.RepairCost,
			},
			ItemDiscovery: ringParam.ItemDiscovery,
		})
	}

	fmt.Printf("Created %d rings\n", len(rings))
	return rings, nil
}

func createWeapons(weaponParams []param.Weapon, weaponTypeParams []param.WeaponType, weaponReinforceParams []param.WeaponReinforce, itemParams []param.Item, weaponSpEffects emevdParser.Events) ([]output.Weapon, error) {
	fmt.Println("\nCreating weapons...")
	weapons := []output.Weapon{}

	for _, weaponParam := range weaponParams {
		name := weaponParam.Name

		// Majestic Greatsword has different stats for left vs right hand
		if weaponParam.ID == 1830000 {
			name += " (Right Hand)"
		}
		if weaponParam.ID == 1831000 {
			name += " (Left Hand)"
		}

		weaponTypeParam := param.WeaponType{}
		for _, param := range weaponTypeParams {
			if param.ID == weaponParam.WeaponTypeID {
				weaponTypeParam = param
				break
			}
		}

		category, _, _ := strings.Cut(weaponTypeParam.Name, ":")

		weaponReinforce := param.WeaponReinforce{}
		for _, param := range weaponReinforceParams {
			if param.ID == weaponParam.WeaponReinforceID {
				weaponReinforce = param
				break
			}
		}

		// Fists need to have its max reinforcement level overriden to 0
		maxReinforcementLevel := weaponReinforce.MaxReinforcementLevel
		if weaponParam.ID == 3400000 || weaponParam.ID == 3406000 {
			maxReinforcementLevel = 0
		}

		itemParam := param.Item{}
		for _, param := range itemParams {
			if param.ID == weaponParam.ID {
				itemParam = param
				break
			}
		}

		spEffects := weaponSpEffects[itemParam.SpecialEffectID]
		modifiers, err := createModifiers(spEffects)
		if err != nil {
			err = fmt.Errorf("error creating ring modifiers for %s: %w", weaponParam.Name, err)
			return nil, err
		}

		weapons = append(weapons, output.Weapon{
			Equippable: output.Equippable{
				Name:       name,
				Modifiers:  modifiers,
				Weight:     weaponParam.Weight,
				Durability: weaponParam.Durability,
				RepairCost: weaponParam.RepairCost,
			},
			Requirements: output.ScalingAttributes[int]{
				Strength:     weaponParam.RequiredStrength,
				Dexterity:    weaponParam.RequiredDexterity,
				Intelligence: weaponParam.RequiredIntelligence,
				Faith:        weaponParam.RequiredFaith,
			},
			Category:              category,
			Paired:                weaponTypeParam.DualWieldingPermission != 0,
			Infusions:             map[string]output.Infusion{}, // TODO: create infusions
			MaxReinforcementLevel: maxReinforcementLevel,
		})
	}

	fmt.Printf("Created %d weapons\n", len(weapons))
	return weapons, nil
}

// Transform transforms data from DS2 params/EMEVDs to Scholar-friendly data
func Transform(paramData paramParser.DS2Params, emevdData emevdParser.DS2EMEVD) (output.ScholarData, error) {
	helmetParams := []param.Armor{}
	chestpieceParams := []param.Armor{}
	gauntletParams := []param.Armor{}
	leggingParams := []param.Armor{}

	for _, armorParam := range paramData.ArmorParam {
		switch armorParam.ArmorCategory {
		case param.ArmorCategoryHead:
			helmetParams = append(helmetParams, armorParam)
		case param.ArmorCategoryChest:
			chestpieceParams = append(chestpieceParams, armorParam)
		case param.ArmorCategoryArms:
			gauntletParams = append(gauntletParams, armorParam)
		case param.ArmorCategoryLegs:
			leggingParams = append(leggingParams, armorParam)
		}
	}

	rings, err := createRings(paramData.RingParam, paramData.ItemParam, emevdData.SpEffectRing)
	if err != nil {
		return output.ScholarData{}, err
	}

	fmt.Println("\nCreating helmets...")
	helmets, err := createArmor(helmetParams, paramData.ArmorReinforceParam, paramData.ItemParam, emevdData.SpEffectArmor)
	if err != nil {
		return output.ScholarData{}, err
	}
	fmt.Printf("Created %d helmets\n", len(helmets))

	fmt.Println("\nCreating chestpieces...")
	chestpieces, err := createArmor(chestpieceParams, paramData.ArmorReinforceParam, paramData.ItemParam, emevdData.SpEffectArmor)
	if err != nil {
		return output.ScholarData{}, err
	}
	fmt.Printf("Created %d chestpieces\n", len(chestpieces))

	fmt.Println("\nCreating gauntlets...")
	gauntlets, err := createArmor(gauntletParams, paramData.ArmorReinforceParam, paramData.ItemParam, emevdData.SpEffectArmor)
	if err != nil {
		return output.ScholarData{}, err
	}
	fmt.Printf("Created %d gauntlets\n", len(gauntlets))

	fmt.Println("\nCreating leggings...")
	leggings, err := createArmor(leggingParams, paramData.ArmorReinforceParam, paramData.ItemParam, emevdData.SpEffectArmor)
	if err != nil {
		return output.ScholarData{}, err
	}
	fmt.Printf("Created %d leggings\n", len(leggings))

	weapons, err := createWeapons(paramData.WeaponParam, paramData.WeaponTypeParam, paramData.WeaponReinforceParam, paramData.ItemParam, emevdData.SpEffectWeapon)
	if err != nil {
		return output.ScholarData{}, err
	}

	return output.ScholarData{
		Classes:            createClasses(paramData.PlayerStatusParam),
		Chestpieces:        chestpieces,
		Gauntlets:          gauntlets,
		Helmets:            helmets,
		Leggings:           leggings,
		Weapons:            weapons,
		Rings:              rings,
		Levels:             createLevels(paramData.PlayerLevelUpSoulsParam),
		Covenants:          createCovenants(paramData.VowParam),
		Spells:             createSpells(paramData.SpellParam),
		AttributeToStatMap: createAttributeToStatMap(paramData.LevelUpStatusCalcParam),
	}, nil
}
