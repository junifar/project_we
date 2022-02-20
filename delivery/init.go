package delivery

import (
	"github.com/beego/beego/v2/server/web"
	locationuc "project_we/app/usecase/location"
)

// Deliveries is delivery dependencies
type Deliveries struct {
	web.Controller
	Location locationuc.ILocation
}

// New is delivery constructor
func New(location locationuc.ILocation) *Deliveries {
	return &Deliveries{
		Location: location,
	}
}
