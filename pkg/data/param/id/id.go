// Package id contains structs for storing data about param IDs
package id

type (
	// ID is an alias for int
	ID int

	// Range is a struct for storing a range of IDs
	//
	// The range is inclusive
	Range struct {
		// Beginning of the range (inclusive)
		Start ID

		// End of the range (inclusive)
		End ID
	}
)

func (r Range) Contains(id ID) bool {
	return id >= r.Start && id <= r.End
}

func (i ID) Plus(num int) ID {
	return ID(int(i) + num)
}
