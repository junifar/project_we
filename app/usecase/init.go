package usecase

import (
	"github.com/beego/beego/v2/adapter/cache"
	"github.com/beego/beego/v2/core/logs"

	dosenrp "project_we/app/repository/dosen"
	locationrp "project_we/app/repository/location"
	userrp "project_we/app/repository/user"
	locationuc "project_we/app/usecase/location"
	locationuci "project_we/app/usecase/location/impl"
	_ "project_we/app/usecase/user"
	useruc "project_we/app/usecase/user"
	useruci "project_we/app/usecase/user/impl"
)

type Usecase struct {
	LocationUC locationuc.ILocation
	UserUC     useruc.IUser
	Cache      cache.Cache
}

type Repository struct {
	LocationRP locationrp.ILocation
	UserRP     userrp.IUser
	DosenRP    dosenrp.IDosen
}

func Init(repository Repository) (res Usecase) {
	res.LocationUC = locationuci.New(repository.LocationRP)
	logs.Info("initialize usecase location")
	res.UserUC = useruci.New(repository.UserRP, repository.DosenRP)
	logs.Info("initialize usecase user")
	return
}
