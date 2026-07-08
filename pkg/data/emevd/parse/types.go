package parse

import "github.com/Camburgaler/scholar-utils/pkg/data/params/id"

type (
	Statement struct {
		Name string
		Args []string
	}

	Event struct {
		ID         id.ID
		Statements []Statement
	}
)
