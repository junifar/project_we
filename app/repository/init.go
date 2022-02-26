package repository

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"

	locationrp "project_we/app/repository/location"
	locationrpi "project_we/app/repository/location/impl"
	userrp "project_we/app/repository/user"
	userrpi "project_we/app/repository/user/impl"
)

type Dependencies struct {
	ORM orm.Ormer
}

type repositories struct {
	LocationRP locationrp.ILocation
	UserRP     userrp.IUser
}

func Init(dependencies Dependencies) (res repositories) {
	res.LocationRP = locationrpi.New(dependencies.ORM)
	logs.Info("initialize repository location")
	res.UserRP = userrpi.New(dependencies.ORM)
	logs.Info("initialize repository user")
	return
}
