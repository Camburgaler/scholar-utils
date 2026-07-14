package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

// Item types
//
// Comes from ItemTypeParam
var ItemTypes = map[int64]string{
	0:  "None",
	10: "Weapon",
	20: "Shield",
	30: "Armor",
	40: "Ring",
}

type (
	// Item is a struct for storing data from ItemParam.csv
	//
	// ItemParam.csv contains information about all items, equippable and otherwise
	Item struct {
		// The unique ID for this item
		ID id.ID

		// The name of this item
		Name   string
		iconID id.ID

		// SpEffect info is stored in the EMEVD files, which are accessed using DarkScript3
		SpecialEffectID id.ID
		startSFXID      id.ID
		actionSFXID     id.ID
		effectSFXID     id.ID

		// The ID of the weapon that this item data describes (if applicable)
		WeaponParamID id.ID

		// The ID of the armor that this item data describes (if applicable)
		ArmorParamID id.ID

		// The ID of the ring that this item data describes (if applicable)
		RingParamID     id.ID
		spellParamID    id.ID
		gestureID       id.ID
		sortID          id.ID
		price           int64
		salePrice       int64
		motionSpeedRate float64

		// The weight of this item
		Weight float64

		// The ID of the item type (see ItemTypes)
		ItemTypeParamID      int64
		itemUsageParamID     int64
		getFrag              bool
		maxNum               int64
		spEffectTargetRegion int64
		spEffectDuplication  int64
		spEffectSFXPriority  int64
		menuItemCategory     int64

		// Determines if this item can be enchanted
		EnchantType            bool
		completeTrophy         bool
		isOnlyOne              bool
		isOnlyOneUtilGameClear bool
		isTakeOver             bool
		canGetWhenGetAllItem   bool
		padding                []byte
	}
)

// ValidItemIDs is a list of valid ItemParam IDs
var ValidItemIDs = []id.Range{
	{Start: 1000000, End: 3370000},
	{Start: 3410000, End: 11840000},
	{Start: 21010100, End: 21502100},
	{Start: 21600100, End: 21600103},
	{Start: 21610100, End: 21610103},
	{Start: 21630100, End: 21630100},
	{Start: 21640100, End: 21640100},
	{Start: 21650100, End: 21650100},
	{Start: 21660100, End: 21660100},
	{Start: 21670100, End: 21670100},
	{Start: 21680100, End: 21680100},
	{Start: 21690103, End: 21690103},
	{Start: 21700100, End: 21700100},
	{Start: 21710100, End: 26260103},
	{Start: 26510100, End: 26510103},
	{Start: 26590100, End: 26700100},
	{Start: 26750100, End: 26750103},
	{Start: 26770101, End: 26770101},
	{Start: 26800100, End: 26800103},
	{Start: 26880100, End: 26880103},
	{Start: 26890100, End: 26890103},
	{Start: 26900100, End: 26900103},
	{Start: 26940100, End: 27520103},
	{Start: 27521100, End: 27950103},
	{Start: 40010000, End: 42000000},
}
