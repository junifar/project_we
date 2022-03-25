package impl

import (
	"github.com/beego/beego/v2/adapter/orm"
	dosenrp "project_we/app/repository/dosen"
)

type dosenRepository struct {
	orm orm.Ormer
}

// NewDosenRepository ...
func NewDosenRepository(orm orm.Ormer) dosenrp.IDosen {
	return &dosenRepository{
		orm: orm,
	}
}
