package param

import (
	"github.com/Camburgaler/scholar-utils/pkg/data/param/id"
)

type (
	// ArmorReinforce is a struct for storing data from ArmorReinforceParam.csv
	//
	// ArmorReinforceParam.csv contains information about armor reinforcement
	ArmorReinforce struct {
		// The unique ID for this reinforcement data
		ID id.ID

		// The name of this reinforcement data
		Name string

		// The minimum (non-reinforced) slash defence provided by the armor using this data
		MinDefenceSlash int64

		// The minimum (non-reinforced) thrust defence provided by the armor using this data
		MinDefenceThrust int64

		// The minimum (non-reinforced) strike defence provided by the armor using this data
		MinDefenceStrike int64

		// The minimum (non-reinforced) standard defence provided by the armor using this data
		MinDefenceStandard int64

		// The minimum (non-reinforced) magic absorption provided by the armor using this data
		MinAbsorptionMagic float64

		// The minimum (non-reinforced) lightning absorption provided by the armor using this data
		MinAbsorptionLightning float64

		// The minimum (non-reinforced) fire absorption provided by the armor using this data
		MinAbsorptionFire float64

		// The minimum (non-reinforced) dark absorption provided by the armor using this data
		MinAbsorptionDark float64

		// The minimum (non-reinforced) poison absorption provided by the armor using this data
		MinAbsorptionPoison float64

		// The minimum (non-reinforced) bleed absorption provided by the armor using this data
		MinAbsorptionBleed float64

		// The minimum (non-reinforced) petrify absorption provided by the armor using this data
		MinAbsorptionPetrify float64

		// The minimum (non-reinforced) curse absorption provided by the armor using this data
		MinAbsorptionCurse float64

		// The maximum (fully reinforced) slash defence provided by the armor using this data
		MaxDefenceSlash int64

		// The maximum (fully reinforced) thrust defence provided by the armor using this data
		MaxDefenceThrust int64

		// The maximum (fully reinforced) strike defence provided by the armor using this data
		MaxDefenceStrike int64

		// The maximum (fully reinforced) standard defence provided by the armor using this data
		MaxDefenceStandard int64

		// The maximum (fully reinforced) magic absorption provided by the armor using this data
		MaxAbsorptionMagic float64

		// The maximum (fully reinforced) lightning absorption provided by the armor using this data
		MaxAbsorptionLightning float64

		// The maximum (fully reinforced) fire absorption provided by the armor using this data
		MaxAbsorptionFire float64

		// The maximum (fully reinforced) dark absorption provided by the armor using this data
		MaxAbsorptionDark float64

		// The maximum (fully reinforced) poison absorption provided by the armor using this data
		MaxAbsorptionPoison float64

		// The maximum (fully reinforced) bleed absorption provided by the armor using this data
		MaxAbsorptionBleed float64

		// The maximum (fully reinforced) petrify absorption provided by the armor using this data
		MaxAbsorptionPetrify float64

		// The maximum (fully reinforced) curse absorption provided by the armor using this data
		MaxAbsorptionCurse float64

		// The maximum reinforcement level for this armor
		MaxReinforcementLevel int64
		dummy                 []byte

		// The defence power correction scale to use within the menu. Display only.
		MenuResistanceScale float64

		// ID for the CustomAttrSpecParam that defines how this armor can be infused
		CustomAttrSpecID id.ID
		customAttrCost   int64
		reinforceCostID  id.ID
	}
)

var (
	// ValidArmorReinforceIDs is a list of valid ArmorReinforceParam IDs
	ValidArmorReinforceIDs = []id.Range{
		{Start: 11001100, End: 11001103},
		{Start: 11010100, End: 12230103},
		{Start: 12240100, End: 13060101},
		{Start: 13061100, End: 13081101},
		{Start: 13120100, End: 17550103},
		{Start: 17680100, End: 17830101},
		{Start: 17950100, End: 17950103},
	}
)
