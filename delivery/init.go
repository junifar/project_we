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

// Usecase is usecase dependencies
type Usecase struct {
	LocationUC locationuc.ILocation
	UserUC     useruc.IUser
}

// New is delivery constructor
func New(usecase Usecase) *Deliveries {
	return &Deliveries{
		Location: usecase.LocationUC,
		User:     usecase.UserUC,
	}
}
