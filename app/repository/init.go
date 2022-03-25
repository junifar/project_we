package repository

import (
	"github.com/beego/beego/v2/adapter/cache"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"

	dosenrp "project_we/app/repository/dosen"
	dosenrpi "project_we/app/repository/dosen/impl"
	locationrp "project_we/app/repository/location"
	locationrpi "project_we/app/repository/location/impl"
	userrp "project_we/app/repository/user"
	userrpi "project_we/app/repository/user/impl"
)

type Dependencies struct {
	ORM   orm.Ormer
	Cache cache.Cache
}

type repositories struct {
	LocationRP locationrp.ILocation
	UserRP     userrp.IUser
	DosenRP    dosenrp.IDosen
}

func Init(dependencies Dependencies) (res repositories) {
	res.LocationRP = locationrpi.New(dependencies.ORM)
	logs.Info("initialize repository location")
	res.UserRP = userrpi.New(dependencies.ORM, dependencies.Cache)
	logs.Info("initialize repository user")
	res.DosenRP = dosenrpi.NewDosenRepository(dependencies.ORM)
	logs.Info("initialize repository dosen")
	return
}
