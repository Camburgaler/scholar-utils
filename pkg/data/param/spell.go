package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

// ValidSpellIDs is the list of valid SpellParam IDs
var ValidSpellIDs = []id.Range{
	{Start: 31010000, End: 35310000},
}

// Spell is a struct for storing data from SpellParam.csv
//
// SpellParam.csv contains information about spells
type Spell struct {
	// The unique ID for this spell
	ID id.ID

	// The name of this spell
	Name string

	spellCategory             int    //nolint:unused
	isDual2HandedSpellAllowed bool   //nolint:unused
	availableCondition        int    //nolint:unused
	dummy0                    []byte //nolint:unused

	// The minimum required intelligence to use this spell
	RequiredIntelligence int

	// The minimum required faith to use this spell
	RequiredFaith int

	baseSpellBulletID                 id.ID   //nolint:unused
	baseSpellDamageID                 id.ID   //nolint:unused
	single1HNormalChantTimeScale      float64 //nolint:unused
	single1HNormalHoldCastTime        float64 //nolint:unused
	single1HNormalAnimIDSpellChanting id.ID   //nolint:unused
	single1HNormalAnimIDSpellCasting  id.ID   //nolint:unused
	single1HNormalStaminaCostScale    float64 //nolint:unused
	single1HStrongChantTimeScale      float64 //nolint:unused
	single1HStrongHoldCastTime        float64 //nolint:unused
	single1HStrongAnimIDSpellChanting id.ID   //nolint:unused
	single1HStrongAnimIDSpellCasting  id.ID   //nolint:unused
	single1HStrongStaminaCostScale    float64 //nolint:unused
	single2HNormalChantTimeScale      float64 //nolint:unused
	single2HNormalHoldCastTime        float64 //nolint:unused
	single2HNormalAnimIDSpellChanting id.ID   //nolint:unused
	single2HNormalAnimIDSpellCasting  id.ID   //nolint:unused
	single2HNormalStaminaCostScale    float64 //nolint:unused
	single2HStrongChantTimeScale      float64 //nolint:unused
	single2HHoldCastTime              float64 //nolint:unused
	single2HAnimIDSpellChanting       id.ID   //nolint:unused
	single2HAnimIDSpellCasting        id.ID   //nolint:unused
	single2HStaminaCostScale          float64 //nolint:unused
	dual1HNormalChantTimeScale        float64 //nolint:unused
	dual1HNormalHoldCastTime          float64 //nolint:unused
	dual1HNormalAnimIDSpellChanting   id.ID   //nolint:unused
	dual1HNormalAnimIDSpellCasting    id.ID   //nolint:unused
	dual1HNormalStaminaCostScale      float64 //nolint:unused
	dual1HStrongChantTimeScale        float64 //nolint:unused
	dual1HStrongHoldCastTime          float64 //nolint:unused
	dual1HStrongAnimIDSpellChanting   id.ID   //nolint:unused
	dual1HStrongAnimIDSpellCasting    id.ID   //nolint:unused
	dual1HStrongStaminaCostScale      float64 //nolint:unused
	dual2HNormalChantTimeScale        float64 //nolint:unused
	dual2HNormalHoldCastTime          float64 //nolint:unused
	dual2HNormalAnimIDSpellChanting   id.ID   //nolint:unused
	dual2HNormalAnimIDSpellCasting    id.ID   //nolint:unused
	dual2HNormalStaminaCostScale      float64 //nolint:unused
	dual2HStrongChantTimeScale        float64 //nolint:unused
	dual2HStrongHoldCastTime          float64 //nolint:unused
	dual2HStrongAnimIDSpellChanting   id.ID   //nolint:unused
	dual2HStrongAnimIDSpellCasting    id.ID   //nolint:unused
	dual2HStrongStaminaCostScale      float64 //nolint:unused
	lockonAvailableDistanceScale      float64 //nolint:unused
	canMove                           bool    //nolint:unused
	canWaitShoot                      bool    //nolint:unused
	canPredictShot                    bool    //nolint:unused
	dummy1                            []byte  //nolint:unused
	sE0ID                             id.ID   //nolint:unused
	sFX0ID                            id.ID   //nolint:unused
	sFX0AttachTarget                  bool    //nolint:unused
	sFX0AttachDummyPolyIDRight        id.ID   //nolint:unused
	sFX0AttachDummyPolyIDLeft         id.ID   //nolint:unused
	sFX0DisableEraseFade              bool    //nolint:unused
	sE1ID                             id.ID   //nolint:unused
	sFX1ID                            id.ID   //nolint:unused
	sFX1AttachTarget                  bool    //nolint:unused
	sFX1AttachDummyPolyIDRight        id.ID   //nolint:unused
	sFX1AttachDummyPolyIDLeft         id.ID   //nolint:unused
	sFX1DisableEraseFade              bool    //nolint:unused
	decreaseSpeedRateParallelWalk     float64 //nolint:unused
	decreaseSpeedRateWalk             float64 //nolint:unused
	decreaseSpeedRateLockOnWalk       float64 //nolint:unused
	startPlaySpeed                    float64 //nolint:unused
	endPlaySpeed                      float64 //nolint:unused
	toughnessPeriodScale              float64 //nolint:unused
	shootSECategory                   int     //nolint:unused

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

	dummy2                    []byte //nolint:unused
	soulConsumeParamID        id.ID  //nolint:unused
	baseLeftHandSpellBulletID id.ID  //nolint:unused
	baseLeftHandSpellDamageID id.ID  //nolint:unused
}
