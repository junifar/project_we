package delivery

import (
	"github.com/beego/beego/v2/server/web"

	locationuc "project_we/app/usecase/location"
	useruc "project_we/app/usecase/user"
)

// Deliveries is delivery dependencies
type Deliveries struct {
	web.Controller
	Location locationuc.ILocation
	User     useruc.IUser
}

// New is delivery constructor
func New(
	location locationuc.ILocation,
	User useruc.IUser,
) *Deliveries {
	return &Deliveries{
		Location: location,
		User:     User,
	}
}
