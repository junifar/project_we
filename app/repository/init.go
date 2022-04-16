package repository

import (
	"github.com/beego/beego/v2/adapter/cache"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	"project_we/app/pkg/curl"
	jabatanfungsionalrp "project_we/app/repository/jabatanfungsional"
	jabatanfungsionalrpi "project_we/app/repository/jabatanfungsional/impl"
	sintarp "project_we/app/repository/sinta"
	sintarpi "project_we/app/repository/sinta/impl"

	dosenrp "project_we/app/repository/dosen"
	dosenrpi "project_we/app/repository/dosen/impl"
	institusirp "project_we/app/repository/institusi"
	institusirpi "project_we/app/repository/institusi/impl"
	jenjangpendidikanrp "project_we/app/repository/jenjangpendidikan"
	jenjangpendidikanrpi "project_we/app/repository/jenjangpendidikan/impl"
	locationrp "project_we/app/repository/location"
	locationrpi "project_we/app/repository/location/impl"
	programstudirp "project_we/app/repository/programstudi"
	programstudirpi "project_we/app/repository/programstudi/impl"
	userrp "project_we/app/repository/user"
	userrpi "project_we/app/repository/user/impl"
)

type Dependencies struct {
	ORM   orm.Ormer
	Cache cache.Cache
	Curl  curl.IHttpRequestor
}

type repositories struct {
	LocationRP          locationrp.ILocation
	UserRP              userrp.IUser
	DosenRP             dosenrp.IDosen
	InstitusiRP         institusirp.IInstitusi
	JenjangPendidikanRP jenjangpendidikanrp.IJenjangPendidikan
	ProgramStudiRP      programstudirp.IProgramStudi
	JabatanFungsionalRP jabatanfungsionalrp.IJabatanFungsional
	SintaRP             sintarp.ISinta
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
	res.ProgramStudiRP = programstudirpi.NewProgramStudiRepository(dependencies.ORM)
	logs.Info("initialize repository program studi")
	res.JabatanFungsionalRP = jabatanfungsionalrpi.NewJabatanFungsionalRepository(dependencies.ORM)
	logs.Info("initialize repository jabatan fungsional")
	res.SintaRP = sintarpi.NewSintaRepository(dependencies.ORM, dependencies.Curl)
	logs.Info("initialize repository sinta")
	return
}
