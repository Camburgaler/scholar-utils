package param

import (
	"strings"

	"github.com/Camburgaler/scholar-utils/pkg/data/param/id"
)

// ValidPlayerStatusIDs is the list of valid PlayerStatusParam IDs
var ValidPlayerStatusIDs = []id.Range{
	{Start: 20, End: 110},
}

type (
	// PlayerStatus is a struct for storing data from PlayerStatusParam.csv
	//
	// PlayerStatusParam.csv contains information about the starting classes
	PlayerStatus struct {
		// The unique ID for this player status
		ID id.ID

		// The name of this player status
		Name string

		genderType int    //nolint:unused
		livingType int    //nolint:unused
		bodyType   int    //nolint:unused
		pad01      []byte //nolint:unused

		// The level of this class
		Level int

		// The Vigor stat of this class
		Vigor int

		totalSouls int //nolint:unused

		// The Endurance stat of this class
		Endurance int

		// The Attunement stat of this class
		Attunement int

		// The Vitality stat of this class
		Vitality int

		// The Strength stat of this class
		Strength int

		// The Dexterity stat of this class
		Dexterity int

		// The Intelligence stat of this class
		Intelligence int

		// The Faith stat of this class
		Faith int

		// The Adaptability stat of this class
		Adaptability int

		durability              int    //nolint:unused
		luck                    int    //nolint:unused
		facePresetNo            int    //nolint:unused
		pad00                   []byte //nolint:unused
		beltItem01              id.ID  //nolint:unused
		beltItem02              id.ID  //nolint:unused
		beltItem03              id.ID  //nolint:unused
		beltItem04              id.ID  //nolint:unused
		beltItem05              id.ID  //nolint:unused
		beltItem06              id.ID  //nolint:unused
		beltItem07              id.ID  //nolint:unused
		beltItem08              id.ID  //nolint:unused
		beltItem09              id.ID  //nolint:unused
		beltItem10              id.ID  //nolint:unused
		beltItem01Num           int    //nolint:unused
		beltItem02Num           int    //nolint:unused
		beltItem03Num           int    //nolint:unused
		beltItem04Num           int    //nolint:unused
		beltItem05Num           int    //nolint:unused
		beltItem06Num           int    //nolint:unused
		beltItem07Num           int    //nolint:unused
		beltItem08Num           int    //nolint:unused
		beltItem09Num           int    //nolint:unused
		beltItem10Num           int    //nolint:unused
		spellSlotItem01         id.ID  //nolint:unused
		spellSlotItem02         id.ID  //nolint:unused
		spellSlotItem03         id.ID  //nolint:unused
		spellSlotItem04         id.ID  //nolint:unused
		spellSlotItem05         id.ID  //nolint:unused
		spellSlotItem06         id.ID  //nolint:unused
		spellSlotItem07         id.ID  //nolint:unused
		spellSlotItem08         id.ID  //nolint:unused
		spellSlotItem09         id.ID  //nolint:unused
		spellSlotItem10         id.ID  //nolint:unused
		rightWeapon01           id.ID  //nolint:unused
		rightWeapon02           id.ID  //nolint:unused
		rightWeapon03           id.ID  //nolint:unused
		leftWeapon01            id.ID  //nolint:unused
		leftWeapon02            id.ID  //nolint:unused
		leftWeapon03            id.ID  //nolint:unused
		headArmor               id.ID  //nolint:unused
		bodyArmor               id.ID  //nolint:unused
		armArmor                id.ID  //nolint:unused
		legArmor                id.ID  //nolint:unused
		ringSlot01              id.ID  //nolint:unused
		ringSlot02              id.ID  //nolint:unused
		ringSlot03              id.ID  //nolint:unused
		ringSlot04              id.ID  //nolint:unused
		rightWeapon01Enhanced   bool   //nolint:unused
		rightWeapon01Attribute  int    //nolint:unused
		dummyR1                 []byte //nolint:unused
		rightWeapon02Enhanced   bool   //nolint:unused
		rightWeapon02Attribute  int    //nolint:unused
		dummyR2                 []byte //nolint:unused
		rightWeapon03Enhanced   bool   //nolint:unused
		rightWeapon03Attribute  int    //nolint:unused
		dummyR3                 []byte //nolint:unused
		leftWeapon01Enhanced    bool   //nolint:unused
		leftWeapon01Attribute   int    //nolint:unused
		dummyL1                 []byte //nolint:unused
		leftWeapon02Enhanced    bool   //nolint:unused
		leftWeapon02Attribute   int    //nolint:unused
		dummyL2                 []byte //nolint:unused
		leftWeapon03Enhanced    bool   //nolint:unused
		leftWeapon03Attribute   int    //nolint:unused
		dummyL3                 []byte //nolint:unused
		headArmorEnhanced       bool   //nolint:unused
		bodyArmorEnhanced       bool   //nolint:unused
		armArmorEnhanced        bool   //nolint:unused
		legArmorEnhanced        bool   //nolint:unused
		bagItemParam01          id.ID  //nolint:unused
		bagItemParam02          id.ID  //nolint:unused
		bagItemParam03          id.ID  //nolint:unused
		estusFlaskLvUsableCount int    //nolint:unused
		estusFlaskHealingPoint  int    //nolint:unused
		dummy                   []byte //nolint:unused
		arrowSlot01             id.ID  //nolint:unused
		arrowSlot02             id.ID  //nolint:unused
		boltSlot01              id.ID  //nolint:unused
		boltSlot02              id.ID  //nolint:unused
		arrowSlot01Num          int    //nolint:unused
		arrowSlot02Num          int    //nolint:unused
		boltSlot01Num           int    //nolint:unused
		boltSlot02Num           int    //nolint:unused
		gestureSlot01           id.ID  //nolint:unused
		gestureSlot02           id.ID  //nolint:unused
		gestureSlot03           id.ID  //nolint:unused
		gestureSlot04           id.ID  //nolint:unused
		gestureSlot05           id.ID  //nolint:unused
		gestureSlot06           id.ID  //nolint:unused
		gestureSlot07           id.ID  //nolint:unused
		gestureSlot08           id.ID  //nolint:unused
		vowID                   id.ID  //nolint:unused
		vowLevel                int    //nolint:unused
		vowContribution         int    //nolint:unused
	}
)

// IsClass returns true if this player status is a class
func (p *PlayerStatus) IsClass() bool {
	return strings.Contains(p.Name, "[Class]")
}
