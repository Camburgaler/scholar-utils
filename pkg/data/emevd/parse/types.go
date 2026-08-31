package parse

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

type (
	// Statement is a struct for storing a function call from an EMEVD file
	Statement struct {
		Name string
		Args []string
	}

	// Events is a struct for storing a list of statements from an EMEVD file
	Events map[id.ID][]Statement

	// DS2EMEVD is a struct for storing the EMEVD file data
	DS2EMEVD struct {
		SpEffectArmor  Events
		SpEffectRing   Events
		SpEffectWeapon Events
	}
)
