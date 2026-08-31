package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

type ArmorCategory int

// ValidArmorIDs is a list of valid ArmorParam IDs
var ValidArmorIDs = []id.Range{
	{Start: 11001100, End: 11001103},
	{Start: 11010100, End: 11040103},
	{Start: 11050100, End: 11370103},
	{Start: 11390100, End: 13060101},
	{Start: 13061100, End: 13081101},
	{Start: 13120100, End: 17550103},
	{Start: 17680100, End: 17950103},
}

const (
	ArmorCategoryHead ArmorCategory = iota + 1
	ArmorCategoryChest
	ArmorCategoryArms
	ArmorCategoryLegs
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
		armorType  int

		// The category of this armor (see ArmorCategories)
		ArmorCategory ArmorCategory

		dummy0            []byte
		modelID           id.ID
		hasGenderedArmor  bool
		dummy1            []byte
		clothCollidableID id.ID
		interfereID       id.ID
		iconID            id.ID

		// Reference to reinforcement data in ArmorReinforceParam
		ArmorReinforceID id.ID

		// The influence of the character's Physical DEF stat on the effectiveness of armor worn.
		DefenseStatEffectScaling float64

		// Coefficient for slash defense scaling
		DefenseScalingSlash float64

		// Coefficient for thrust defense scaling
		DefenseScalingThrust float64

		// Coefficient for strike defense scaling
		DefenseScalingStrike float64

		// Minimum strength stat required to equip this armor
		PrerequisiteStrength int

		// Minimum dexterity stat required to equip this armor
		PrerequisiteDexterity int

		// Minimum intelligence stat required to equip this armor
		PrerequisiteIntelligence int

		// Minimum faith stat required to equip this armor
		PrerequisiteFaith int

		// Weight of this armor
		Weight float64

		// Durability of this armor
		Durability int

		// Cost to repair this armor
		RepairCost int

		// Poise of this armor
		Poise         float64
		handsUpWeight int
		handsUpFLevel int

		// Modifier to item discovery provided by this armor
		ItemDiscovery   int
		material        int
		brokenSFXID     id.ID
		brokenSoundType int
	}
)
