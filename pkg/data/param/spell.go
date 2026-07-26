package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

var ValidSpellIDs = []id.Range{
	{Start: 31010000, End: 35310000},
}

type SpellCategory int

const (
	SpellCategorySorcery SpellCategory = iota
	SpellCategoryMiracle
	SpellCategoryPyromancy
	SpellCategoryStaffHex
	SpellCategoryChimeHex
)

type Spell struct {
	// The unique ID for this spell
	ID id.ID

	// The name of this spell
	Name string

	// The category of this spell
	spellCategory SpellCategory

	isDual2HandedSpellAllowed bool

	availableCondition int

	dummy0 []byte

	// The minimum required intelligence to use this spell
	RequiredIntelligence int

	// The minimum required faith to use this spell
	RequiredFaith int

	baseSpellBulletID id.ID

	baseSpellDamageID id.ID

	single1HNormalChantTimeScale float64

	single1HNormalHoldCastTime float64

	single1HNormalAnimIDSpellChanting id.ID

	single1HNormalAnimIDSpellCasting id.ID

	single1HNormalStaminaCostScale float64

	single1HStrongChantTimeScale float64

	single1HStrongHoldCastTime float64

	single1HStrongAnimIDSpellChanting id.ID

	single1HStrongAnimIDSpellCasting id.ID

	single1HStrongStaminaCostScale float64

	single2HNormalChantTimeScale float64

	single2HNormalHoldCastTime float64

	single2HNormalAnimIDSpellChanting id.ID

	single2HNormalAnimIDSpellCasting id.ID

	single2HNormalStaminaCostScale float64

	single2HStrongChantTimeScale float64

	single2HHoldCastTime float64

	single2HAnimIDSpellChanting id.ID

	single2HAnimIDSpellCasting id.ID

	single2HStaminaCostScale float64

	dual1HNormalChantTimeScale float64

	dual1HNormalHoldCastTime float64

	dual1HNormalAnimIDSpellChanting id.ID

	dual1HNormalAnimIDSpellCasting id.ID

	dual1HNormalStaminaCostScale float64

	dual1HStrongChantTimeScale float64

	dual1HStrongHoldCastTime float64

	dual1HStrongAnimIDSpellChanting id.ID

	dual1HStrongAnimIDSpellCasting id.ID

	dual1HStrongStaminaCostScale float64

	dual2HNormalChantTimeScale float64

	dual2HNormalHoldCastTime float64

	dual2HNormalAnimIDSpellChanting id.ID

	dual2HNormalAnimIDSpellCasting id.ID

	dual2HNormalStaminaCostScale float64

	dual2HStrongChantTimeScale float64

	dual2HStrongHoldCastTime float64

	dual2HStrongAnimIDSpellChanting id.ID

	dual2HStrongAnimIDSpellCasting id.ID

	dual2HStrongStaminaCostScale float64

	lockonAvailableDistanceScale float64

	canMove bool

	canWaitShoot bool

	canPredictShot bool

	dummy1 []byte

	sE0ID id.ID

	sFX0ID id.ID

	sFX0AttachTarget bool

	sFX0AttachDummyPolyIDRight id.ID

	sFX0AttachDummyPolyIDLeft id.ID

	sFX0DisableEraseFade bool

	sE1ID id.ID

	sFX1ID id.ID

	sFX1AttachTarget bool

	sFX1AttachDummyPolyIDRight id.ID

	sFX1AttachDummyPolyIDLeft id.ID

	sFX1DisableEraseFade bool

	decreaseSpeedRateParallelWalk float64

	decreaseSpeedRateWalk float64

	decreaseSpeedRateLockOnWalk float64

	startPlaySpeed float64

	endPlaySpeed float64

	toughnessPeriodScale float64

	shootSECategory int

	// How many spell slots this spell consumes
	SpellSlotCost int

	UsageCountLV1 int

	UsageCountLV2 int

	UsageCountLV3 int

	UsageCountLV4 int

	UsageCountLV5 int

	UsageCountLV6 int

	UsageCountLV7 int

	UsageCountLV8 int

	UsageCountLV9 int

	UsageCountLV10 int

	dummy2 []byte

	soulConsumeParamID id.ID

	baseLeftHandSpellBulletID id.ID

	baseLeftHandSpellDamageID id.ID
}
