package param

import (
	"strings"

	"github.com/Camburgaler/scholar-utils/pkg/data/param/id"
)

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
		Name       string
		genderType int64
		livingType int64
		bodyType   int64
		pad01      []byte

		// The level of this class
		Level int64

		// The Vigor stat of this class
		Vigor      int64
		totalSouls int64

		// The Endurance stat of this class
		Endurance int64

		// The Attunement stat of this class
		Attunement int64

		// The Vitality stat of this class
		Vitality int64

		// The Strength stat of this class
		Strength int64

		// The Dexterity stat of this class
		Dexterity int64

		// The Intelligence stat of this class
		Intelligence int64

		// The Faith stat of this class
		Faith int64

		// The Adaptability stat of this class
		Adaptability            int64
		durability              int64
		luck                    int64
		facePresetNo            int64
		pad00                   []byte
		beltItem01              id.ID
		beltItem02              id.ID
		beltItem03              id.ID
		beltItem04              id.ID
		beltItem05              id.ID
		beltItem06              id.ID
		beltItem07              id.ID
		beltItem08              id.ID
		beltItem09              id.ID
		beltItem10              id.ID
		beltItem01Num           int64
		beltItem02Num           int64
		beltItem03Num           int64
		beltItem04Num           int64
		beltItem05Num           int64
		beltItem06Num           int64
		beltItem07Num           int64
		beltItem08Num           int64
		beltItem09Num           int64
		beltItem10Num           int64
		spellSlotItem01         id.ID
		spellSlotItem02         id.ID
		spellSlotItem03         id.ID
		spellSlotItem04         id.ID
		spellSlotItem05         id.ID
		spellSlotItem06         id.ID
		spellSlotItem07         id.ID
		spellSlotItem08         id.ID
		spellSlotItem09         id.ID
		spellSlotItem10         id.ID
		rightWeapon01           id.ID
		rightWeapon02           id.ID
		rightWeapon03           id.ID
		leftWeapon01            id.ID
		leftWeapon02            id.ID
		leftWeapon03            id.ID
		headArmor               id.ID
		bodyArmor               id.ID
		armArmor                id.ID
		legArmor                id.ID
		ringSlot01              id.ID
		ringSlot02              id.ID
		ringSlot03              id.ID
		ringSlot04              id.ID
		rightWeapon01Enhanced   bool
		rightWeapon01Attribute  int64
		dummyR1                 []byte
		rightWeapon02Enhanced   bool
		rightWeapon02Attribute  int64
		dummyR2                 []byte
		rightWeapon03Enhanced   bool
		rightWeapon03Attribute  int64
		dummyR3                 []byte
		leftWeapon01Enhanced    bool
		leftWeapon01Attribute   int64
		dummyL1                 []byte
		leftWeapon02Enhanced    bool
		leftWeapon02Attribute   int64
		dummyL2                 []byte
		leftWeapon03Enhanced    bool
		leftWeapon03Attribute   int64
		dummyL3                 []byte
		headArmorEnhanced       bool
		bodyArmorEnhanced       bool
		armArmorEnhanced        bool
		legArmorEnhanced        bool
		bagItemParam01          id.ID
		bagItemParam02          id.ID
		bagItemParam03          id.ID
		estusFlaskLvUsableCount int64
		estusFlaskHealingPoint  int64
		dummy                   []byte
		arrowSlot01             id.ID
		arrowSlot02             id.ID
		boltSlot01              id.ID
		boltSlot02              id.ID
		arrowSlot01Num          int64
		arrowSlot02Num          int64
		boltSlot01Num           int64
		boltSlot02Num           int64
		gestureSlot01           id.ID
		gestureSlot02           id.ID
		gestureSlot03           id.ID
		gestureSlot04           id.ID
		gestureSlot05           id.ID
		gestureSlot06           id.ID
		gestureSlot07           id.ID
		gestureSlot08           id.ID
		vowID                   id.ID
		vowLevel                int64
		vowContribution         int64
	}
)

func (p *PlayerStatus) IsClass() bool {
	return strings.Contains(p.Name, "[Class]")
}
