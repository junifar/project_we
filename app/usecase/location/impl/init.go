package impl

import (
	locationrp "project_we/app/repository/location"
	locationuc "project_we/app/usecase/location"
)

// LocationImpl is location dependencies
type LocationImpl struct {
	Location locationrp.ILocation
}

// New is location constructor
func New(location locationrp.ILocation) locationuc.ILocation {
	return &LocationImpl{
		Location: location,
	}
}
