package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

type (
	// CustomAttrSpec is a struct for storing data from CustomAttributeSpecParam.csv
	//
	// CustomAttributeSpecParam.csv contains information about what types of weapons are compatible with each infusion
	CustomAttrSpec struct {
		// The unique ID for this infusion
		ID id.ID

		// The name of this infusion
		Name string

		// A flag for whether the Physical infusion is compatible (should always be true)
		Physical bool

		// A flag for whether the Fire infusion is compatible
		Fire bool

		// A flag for whether the Magic infusion is compatible
		Magic bool

		// A flag for whether the Lightning infusion is compatible
		Lightning bool

		// A flag for whether the Dark infusion is compatible
		Dark bool

		// A flag for whether the Poison infusion is compatible
		Poison bool

		// A flag for whether the Bleed infusion is compatible
		Bleed bool

		// A flag for whether the Raw infusion is compatible
		Raw bool

		// A flag for whether the Enchanted infusion is compatible
		Enchanted bool

		// A flag for whether the Mundane infusion is compatible
		Mundane bool
		padding []byte
	}
)

// ValidParamIDs is a list of valid CustomAttrSpecParam IDs
var ValidCustomAttrSpecIDs = []id.Range{
	{Start: 0, End: 200100},
}
