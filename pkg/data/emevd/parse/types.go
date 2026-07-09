package parse

import "github.com/Camburgaler/scholar-utils/pkg/data/params/id"

type (
	Identifier struct {
		Parts []string
	}

	Statement struct {
		Name string
		Args []string
	}

	Event struct {
		ID         id.ID
		Statements []Statement
	}
	DS2EMEVD struct {
		SpEffectArmor  []Event
		SpEffectRing   []Event
		SpEffectWeapon []Event
	}
)
