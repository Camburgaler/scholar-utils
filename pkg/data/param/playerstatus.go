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
		genderType int
		livingType int
		bodyType   int
		pad01      []byte

		// The level of this class
		Level int

		// The Vigor stat of this class
		Vigor      int
		totalSouls int

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
		Adaptability            int
		durability              int
		luck                    int
		facePresetNo            int
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
		beltItem01Num           int
		beltItem02Num           int
		beltItem03Num           int
		beltItem04Num           int
		beltItem05Num           int
		beltItem06Num           int
		beltItem07Num           int
		beltItem08Num           int
		beltItem09Num           int
		beltItem10Num           int
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
		rightWeapon01Attribute  int
		dummyR1                 []byte
		rightWeapon02Enhanced   bool
		rightWeapon02Attribute  int
		dummyR2                 []byte
		rightWeapon03Enhanced   bool
		rightWeapon03Attribute  int
		dummyR3                 []byte
		leftWeapon01Enhanced    bool
		leftWeapon01Attribute   int
		dummyL1                 []byte
		leftWeapon02Enhanced    bool
		leftWeapon02Attribute   int
		dummyL2                 []byte
		leftWeapon03Enhanced    bool
		leftWeapon03Attribute   int
		dummyL3                 []byte
		headArmorEnhanced       bool
		bodyArmorEnhanced       bool
		armArmorEnhanced        bool
		legArmorEnhanced        bool
		bagItemParam01          id.ID
		bagItemParam02          id.ID
		bagItemParam03          id.ID
		estusFlaskLvUsableCount int
		estusFlaskHealingPoint  int
		dummy                   []byte
		arrowSlot01             id.ID
		arrowSlot02             id.ID
		boltSlot01              id.ID
		boltSlot02              id.ID
		arrowSlot01Num          int
		arrowSlot02Num          int
		boltSlot01Num           int
		boltSlot02Num           int
		gestureSlot01           id.ID
		gestureSlot02           id.ID
		gestureSlot03           id.ID
		gestureSlot04           id.ID
		gestureSlot05           id.ID
		gestureSlot06           id.ID
		gestureSlot07           id.ID
		gestureSlot08           id.ID
		vowID                   id.ID
		vowLevel                int
		vowContribution         int
	}
)

func (p *PlayerStatus) IsClass() bool {
	return strings.Contains(p.Name, "[Class]")
}
