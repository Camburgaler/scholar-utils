package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

var (
	// ValidArmorIDs is a list of valid ArmorParam IDs
	ValidArmorIDs = []id.Range{
		{Start: 11001100, End: 11001103},
		{Start: 11010100, End: 11040103},
		{Start: 11050100, End: 11370103},
		{Start: 11390100, End: 13060101},
		{Start: 13061100, End: 13081101},
		{Start: 13120100, End: 17550103},
		{Start: 17680100, End: 17950103},
	}

	// Categories of armor
	ArmorCategories = map[int64]string{
		1: "Head",
		2: "Chest",
		3: "Arms",
		4: "Legs",
	}
)

type (
	// Armor is a struct for storing data from ArmorParam.csv
	//
	// ArmorParam.csv contains information about equippable armor
	Armor struct {
		// The unique ID for this armor
		ID id.ID

		// The name of this armor
		Name string

		// The common ID for this armor set
		ArmorSetID id.ID
		armorType  int64

		// The category of this armor (see ArmorCategories)
		ArmorCategory     int64
		dummy0            []byte
		modelID           id.ID
		hasGenderedArmor  bool
		dummy1            []byte
		clothCollidableID id.ID
		interfereID       id.ID
		iconID            id.ID

		// Reference to reinforcement data in ArmorReinforceParam
		ArmorReinforceID id.ID

		// Coefficient for defence scaling
		DefenceScalingStatusEffect float64

		// Coefficient for slash defence scaling
		DefenceScalingSlash float64

		// Coefficient for thrust defence scaling
		DefenseScalingThrust float64

		// Coefficient for strike defence scaling
		DefenseScalingStrike float64

		// Minimum strength stat required to equip this armor
		PrerequisiteStrength int64

		// Minimum dexterity stat required to equip this armor
		PrerequisiteDexterity int64

		// Minimum intelligence stat required to equip this armor
		PrerequisiteIntelligence int64

		// Minimum faith stat required to equip this armor
		PrerequisiteFaith int64

		// Weight of this armor
		Weight float64

		// Durability of this armor
		Durability int64

		// Cost to repair this armor
		RepairCost int64

		// Poise of this armor
		Poise         float64
		handsUpWeight int64
		handsUpFLevel int64

		// Modifier to item discovery provided by this armor
		ItemDiscovery   int64
		material        int64
		brokenSFXID     id.ID
		brokenSoundType int64
	}
)
