package output

type (
	// StatCurve is a struct for storing a cumulative sequence of stat values
	//
	// The value at a given index is the how much one should add to the relevant stat given an attribute value as the index
	//
	// An empty StatCurve indicates that the attribute does not affect the stat
	//
	// A value of [-1] indicates that the attribute's contibution to the stat is "special" (not sure yet what that means)
	StatCurve []float64

	// A map of Attributes to Stats to StatCurves
	StatCalculationDetails = Attributes[Stats[StatCurve]]
)

var BaseStats = Stats[float64]{
	MaximumHP:                      500,
	MaximumStamina:                 80,
	MaximumEquipLoad:               38.5,
	SpellSlotCount:                 0,
	SpellCastingSpeed:              35,
	PhysicalAttackPowerByStrength:  49,
	PhysicalAttackPowerByDexterity: 49,
	AttackPowerMagic:               49,
	AttackPowerFire:                50,
	AttackPowerLightning:           49,
	AttackPowerDark:                46,
	AttackPowerPoison:              50,
	AttackPowerBleed:               50,
	Defense:                        50,
	AbsorptionMagic:                0,
	AbsorptionFire:                 0,
	AbsorptionLightning:            0,
	AbsorptionDark:                 0,
	ResistancePoison:               0,
	ResistanceBleed:                0,
	ResistancePetrify:              0,
	ResistanceCurse:                0,
	Agility:                        85,
	Poise:                          0,
}

// StatCurve presets
var (
	specialStatCurve = StatCurve{-1}

	commonMaximumHPStatCurve = StatCurve{
		// Attribute of 0 is 0
		0,
		// 1-20 increase in increments of 2
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		// 21-50 increase in increments of 1
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
		51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
		61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
		// 51-99 increase in increments of 0
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
		70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
		70, 70, 70, 70, 70, 70, 70, 70, 70,
	}

	commonAttackPowerStatCurve = StatCurve{
		// Attribute of 0 is 0
		0,
		// 1-4 increase in increments of 1
		1, 2, 3, 4,
		// 5 increases by 0
		4,
		// 6-8 increase in increments of 1
		5, 6, 7,
		// 9 increases by 0
		7,
		// 10 increases by 1
		8,
		// 11-13 increase in increments of 2
		10, 12, 14,
		// 14 increases by 3
		17,
		// 15-16 increase in increments of 2
		19, 21,
		// 17 increases by 3
		24,
		// 18-19 increase in increments of 2
		26, 28,
		// 20 increases by 3
		31,
		// 21-23 increase in increments of 2
		33, 35, 37,
		// 24 increases by 3
		40,
		// 25-27 increase in increments of 2
		42, 44, 46,
		// 28 increases by 3
		49,
		// 29-30 increase in increments of 2
		51, 53,
		// 31 increases by 3
		56,
		// 32-35 increase in increments of 4
		60, 64, 68, 72,
		// 36 increases by 3
		75,
		// 37-40 increase in increments of 4
		79, 83, 87, 91,
		// 41-50 alternate between 1 and 2
		92, 94, 95, 97, 98, 100, 101, 103, 104, 106,
		// 50-70 increment in a repeating pattern of 0, 1, 1, 1
		106, 107, 108, 109, 109, 110, 111, 112, 112, 113,
		114, 115, 115, 116, 117, 118, 118, 119, 120, 121,
		// 71-80 alternate between 1 and 2
		122, 124, 125, 127, 128, 130, 131, 133, 134, 136,
		// 81-90 increment in a repeating pattern of 0, 1, 1, 1
		136, 137, 138, 139, 139, 140, 141, 142, 142, 143,
		// 91 increases by 0
		143,
		// 92-99 increase in increments of 1
		144, 145, 146, 147, 148, 149, 150, 151,
	}

	commonAbsorptionStatCurve = StatCurve{
		// Attribute of 0 is 0
		0,
		// 1-10 increase in increments of 6
		6, 12, 18, 24, 30, 36, 42, 48, 54, 60,
		// 11-20 increase in increments of 8
		68, 76, 84, 92, 100, 108, 116, 124, 132, 140,
		// 21-60 increase in increments of 1
		141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
		151, 152, 153, 154, 155, 156, 157, 158, 159, 160,
		161, 162, 163, 164, 165, 166, 167, 168, 169, 170,
		171, 172, 173, 174, 175, 176, 177, 178, 179, 180,
		// 61-98 alternate between 0 and 1
		180, 181, 181, 182, 182, 183, 183, 184, 184, 185,
		185, 186, 186, 187, 187, 188, 188, 189, 189, 190,
		190, 191, 191, 192, 192, 193, 193, 194, 194, 195,
		195, 196, 196, 197, 197, 198, 198, 199,
		// 99 increases by 1
		200,
	}

	vigorMaximumHPStatCurve = StatCurve{
		// Attribute of 0 is 0
		0,
		// 1-20 increase in increments of 30
		30, 60, 90, 120, 150, 180, 210, 240, 270, 300,
		330, 360, 390, 420, 450, 480, 510, 540, 570, 600,
		// 21-50 increase in increments of 20
		620, 640, 660, 680, 700, 720, 740, 760, 780, 800,
		820, 840, 860, 880, 900, 920, 940, 960, 980, 1000,
		1020, 1040, 1060, 1080, 1100, 1120, 1140, 1160, 1180, 1200,
		// 51-99 increase in increments of 5
		1205, 1210, 1215, 1220, 1225, 1230, 1235, 1240, 1245, 1250,
		1255, 1260, 1265, 1270, 1275, 1280, 1285, 1290, 1295, 1300,
		1305, 1310, 1315, 1320, 1325, 1330, 1335, 1340, 1345, 1350,
		1355, 1360, 1365, 1370, 1375, 1380, 1385, 1390, 1395, 1400,
		1405, 1410, 1415, 1420, 1425, 1430, 1435, 1440, 1445,
	}

	enduranceMaximumStaminaStatCurve = StatCurve{
		// Attribute of 0 is 0
		0,
		// 1-20 increase in increments of 2
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		// 21-98 increase in increments of 1
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
		51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
		61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
		81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
		91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
		101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
		111, 112, 113, 114, 115, 116, 117, 118,
		// 99 increase by 2
		120,
	}

	vitalityMaximumEquipLoadStatCurve = StatCurve{
		// Attribute of 0 is 0
		0,
		// 1-29 increase in increments of 1.5
		1.5, 3, 4.5, 6, 7.5, 9, 10.5, 12, 13.5, 15,
		16.5, 18, 19.5, 21, 22.5, 24, 25.5, 27, 28.5, 30,
		31.5, 33, 34.5, 36, 37.5, 39, 40.5, 42, 43.5,
		// 30-49 increase in increments of 1
		44.5,
		45.5, 46.5, 47.5, 48.5, 49.5, 50.5, 51.5, 52.5, 53.5, 54.5,
		55.5, 56.5, 57.5, 58.5, 59.5, 60.5, 61.5, 62.5, 63.5,
		// 50-70 increase in increments of 0.5
		64,
		64.5, 65, 65.5, 66, 66.5, 67, 67.5, 68, 68.5, 69,
		69.5, 70, 70.5, 71, 71.5, 72, 72.5, 73, 73.5, 74,
		// 70-98 alternate between 0 and 0.5
		74, 74.5, 74.5, 75, 75, 75.5, 75.5, 76, 76, 76.5,
		76.5, 77, 77, 77.5, 77.5, 78, 78, 78.5, 78.5, 79,
		79, 79.5, 79.5, 80, 80, 80.5, 80.5, 81,
		// 99 increases by 0.5
		81.5,
	}

	attunementSpellSlotCountStatCurve = StatCurve{
		// Attribute of 0 is 0
		0,
		// 1-9 increase in increments of 0
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		// 10 increases by 1
		1,
		// 11-12 increase in increments of 0
		1, 1,
		// 13 increases by 1
		2,
		// 14-15 increase in increments of 0
		2, 2,
		// 16 increases by 1
		3,
		// 17-19 increase in increments of 0
		3, 3, 3,
		// 20 increases by 1
		4,
		// 21-24 increase in increments of 0
		4, 4, 4, 4,
		// 25 increases by 1
		5,
		// 26-29 increase in increments of 0
		5, 5, 5, 5,
		// 30 increases by 1
		6,
		// 31-39 increase in increments of 0
		6, 6, 6, 6, 6, 6, 6, 6, 6,
		// 40 increases by 1
		7,
		// 41-49 increase in increments of 0
		7, 7, 7, 7, 7, 7, 7, 7, 7,
		// 50 increases by 1
		8,
		// 51-59 increase in increments of 0
		8, 8, 8, 8, 8, 8, 8, 8, 8,
		// 60 increases by 1
		9,
		// 61-74 increase in increments of 0
		9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
		9, 9, 9, 9,
		// 75 increases by 1
		10,
		// 76-99 increase in increments of 0
		10, 10, 10, 10, 10,
		10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
		10, 10, 10, 10, 10, 10, 10, 10, 10,
	}

	intelligenceAttackPowerMagicStatCurve = StatCurve{
		// Attribute of 0 is 0
		0,
		// 1-10 increment in a repeating pattern of 1, 0, 1, 0, 0
		1, 1, 2, 2, 2, 3, 3, 4, 4, 4,
		// 11-13 increase in increments of 5
		9, 14, 19,
		// 14-20 increment in a repeating pattern of 6, 5, 5
		25, 30, 35, 41, 46, 51, 57,
		// 21-30 increment in a repeating pattern of 0, 0, 1, 0, 1
		57, 57, 58, 58, 59, 59, 59, 60, 60, 61,
		// 31-40 increase in increments of 3
		64, 67, 70, 73, 76, 79, 82, 85, 88, 91,
		// 41-50 alternate between 1 and 2
		92, 94, 95, 97, 98, 100, 101, 103, 104, 106,
		// 51-60 increment in a repeating pattern of 0, 1, 1, 1
		106, 107, 108, 109, 109, 110, 111, 112, 112, 113,
		// 61-70 increment in a repeating pattern of 0, 1, 1, 1, 1
		113, 114, 115, 116, 117, 117, 118, 119, 120, 121,
		// 71-80 alternate between 1 and 2
		122, 124, 125, 127, 128, 130, 131, 133, 134, 136,
		// 81-90 increment in a repeating pattern of 0, 1, 1, 1
		136, 137, 138, 139, 139, 140, 141, 142, 142, 143,
		// 91 increases by 0
		143,
		// 92-99 increase in increments of 1
		144, 145, 146, 147, 148, 149, 150, 151,
	}
)

var DS2StatCalculationDetails = StatCalculationDetails{
	Vigor: Stats[StatCurve]{
		MaximumHP:         vigorMaximumHPStatCurve,
		ResistancePetrify: specialStatCurve,
	},
	Endurance: Stats[StatCurve]{
		MaximumHP:      commonMaximumHPStatCurve,
		MaximumStamina: enduranceMaximumStaminaStatCurve,
		Poise:          specialStatCurve,
		Defense:        specialStatCurve,
	},
	Vitality: Stats[StatCurve]{
		MaximumHP:        commonMaximumHPStatCurve,
		MaximumEquipLoad: vitalityMaximumEquipLoadStatCurve,
		Defense:          specialStatCurve,
		ResistancePoison: specialStatCurve,
	},
	Attunement: Stats[StatCurve]{
		MaximumHP:         commonMaximumHPStatCurve,
		SpellSlotCount:    attunementSpellSlotCountStatCurve,
		SpellCastingSpeed: specialStatCurve,
		ResistanceCurse:   specialStatCurve,
		Agility:           specialStatCurve,
	},
	Strength: Stats[StatCurve]{
		MaximumHP:                     commonMaximumHPStatCurve,
		PhysicalAttackPowerByStrength: commonAttackPowerStatCurve,
		Defense:                       specialStatCurve,
	},
	Dexterity: Stats[StatCurve]{
		MaximumHP:                      commonMaximumHPStatCurve,
		PhysicalAttackPowerByDexterity: commonAttackPowerStatCurve,
		AttackPowerPoison:              specialStatCurve,
		AttackPowerBleed:               specialStatCurve,
		Defense:                        specialStatCurve,
	},
	Adaptability: Stats[StatCurve]{
		MaximumHP:         commonMaximumHPStatCurve,
		Agility:           specialStatCurve,
		Poise:             specialStatCurve,
		AttackPowerPoison: specialStatCurve,
		ResistancePetrify: specialStatCurve,
		ResistanceCurse:   specialStatCurve,
		ResistancePoison:  specialStatCurve,
		ResistanceBleed:   specialStatCurve,
	},
	Intelligence: Stats[StatCurve]{
		MaximumHP:         commonMaximumHPStatCurve,
		SpellCastingSpeed: specialStatCurve,
		AttackPowerMagic:  intelligenceAttackPowerMagicStatCurve,
		AttackPowerFire:   specialStatCurve,
		AttackPowerDark:   specialStatCurve,
		AttackPowerPoison: specialStatCurve,
		AbsorptionMagic:   commonAbsorptionStatCurve,
		AbsorptionFire:    specialStatCurve,
		AbsorptionDark:    specialStatCurve,
	},
	Faith: Stats[StatCurve]{
		MaximumHP:            commonMaximumHPStatCurve,
		SpellCastingSpeed:    specialStatCurve,
		AttackPowerFire:      specialStatCurve,
		AttackPowerLightning: commonAttackPowerStatCurve,
		AttackPowerDark:      specialStatCurve,
		AttackPowerBleed:     specialStatCurve,
		AbsorptionFire:       specialStatCurve,
		AbsorptionLightning:  commonAbsorptionStatCurve,
		AbsorptionDark:       specialStatCurve,
		ResistanceBleed:      specialStatCurve,
	},
}
