// Package param contains structs for storing data from the param CSV files
package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

type armorCategory int

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
	// ArmorCategoryHead is the head armor category
	ArmorCategoryHead armorCategory = iota + 1
	// ArmorCategoryChest is the chest armor category
	ArmorCategoryChest
	// ArmorCategoryArms is the arms armor category
	ArmorCategoryArms
	// ArmorCategoryLegs is the legs armor category
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

		armorType int //nolint:unused

		// The category of this armor (see ArmorCategories)
		ArmorCategory armorCategory

		dummy0            []byte //nolint:unused
		modelID           id.ID  //nolint:unused
		hasGenderedArmor  bool   //nolint:unused
		dummy1            []byte //nolint:unused
		clothCollidableID id.ID  //nolint:unused
		interfereID       id.ID  //nolint:unused
		iconID            id.ID  //nolint:unused

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
		Poise float64

		handsUpWeight int //nolint:unused
		handsUpFLevel int //nolint:unused

		// Modifier to item discovery provided by this armor
		ItemDiscovery int

		material        int   //nolint:unused
		brokenSFXID     id.ID //nolint:unused
		brokenSoundType int   //nolint:unused
	}
)
