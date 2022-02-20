package impl

import (
	"github.com/beego/beego/v2/adapter/orm"
	locationrp "project_we/app/repository/location"
)

// LocationImpl location dependencies
type LocationImpl struct {
	orm orm.Ormer
}

// New location contructor
func New(orm orm.Ormer) locationrp.ILocation {
	return &LocationImpl{orm: orm}
}
