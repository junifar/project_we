package usecase

import (
	"github.com/beego/beego/v2/core/logs"
	useruc "project_we/app/usecase/user"

	locationrp "project_we/app/repository/location"
	userrp "project_we/app/repository/user"
	locationuc "project_we/app/usecase/location"
	locationuci "project_we/app/usecase/location/impl"
	_ "project_we/app/usecase/user"
	useruci "project_we/app/usecase/user/impl"
)

type Usecase struct {
	LocationUC locationuc.ILocation
	UserUC     useruc.IUser
}

type Repository struct {
	LocationRP locationrp.ILocation
	UserRP     userrp.IUser
}

func Init(repository Repository) (res Usecase) {
	res.LocationUC = locationuci.New(repository.LocationRP)
	logs.Info("initialize usecase location")
	res.UserUC = useruci.New(repository.UserRP)
	logs.Info("initialize usecase user")
	return
}
