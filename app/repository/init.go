package repository

import (
	"github.com/beego/beego/v2/adapter/cache"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"

	dosenrp "project_we/app/repository/dosen"
	dosenrpi "project_we/app/repository/dosen/impl"
	institusirp "project_we/app/repository/institusi"
	institusirpi "project_we/app/repository/institusi/impl"
	jenjangpendidikanrp "project_we/app/repository/jenjangpendidikan"
	jenjangpendidikanrpi "project_we/app/repository/jenjangpendidikan/impl"
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
	LocationRP          locationrp.ILocation
	UserRP              userrp.IUser
	DosenRP             dosenrp.IDosen
	InstitusiRP         institusirp.IInstitusi
	JenjangPendidikanRP jenjangpendidikanrp.IJenjangPendidikan
}

func Init(dependencies Dependencies) (res repositories) {
	res.LocationRP = locationrpi.New(dependencies.ORM)
	logs.Info("initialize repository location")
	res.UserRP = userrpi.New(dependencies.ORM, dependencies.Cache)
	logs.Info("initialize repository user")
	res.DosenRP = dosenrpi.NewDosenRepository(dependencies.ORM)
	logs.Info("initialize repository dosen")
	res.InstitusiRP = institusirpi.NewInstitusiRepository(dependencies.ORM)
	logs.Info("initialize repository institusi")
	res.JenjangPendidikanRP = jenjangpendidikanrpi.NewJenjangPendidikanRepository(dependencies.ORM)
	logs.Info("initialize repository jenjang pendidikan")
	return
}
