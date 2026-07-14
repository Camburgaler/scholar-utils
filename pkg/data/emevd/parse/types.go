package parse

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

type (
	Identifier struct {
		Parts []string
	}

	Statement struct {
		Name string
		Args []string
	}

	Events map[id.ID][]Statement

	DS2EMEVD struct {
		SpEffectArmor  Events
		SpEffectRing   Events
		SpEffectWeapon Events
	}
)
