// Package id contains structs for storing data about param IDs
package id

type (
	// ID is an alias for int64
	ID int64

	// Range is a struct for storing a range of IDs
	//
	// The range is inclusive
	Range struct {
		Min ID
		Max ID
	}
)
