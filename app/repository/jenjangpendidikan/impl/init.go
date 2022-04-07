package impl

import (
	jenjangPendidikanrp "project_we/app/repository/jenjangpendidikan"

	"github.com/beego/beego/v2/adapter/orm"
)

type jenjangPendidikanRepository struct {
	orm orm.Ormer
}

// NewJenjangPendidikanRepository ...
func NewJenjangPendidikanRepository(orm orm.Ormer) jenjangPendidikanrp.IJenjangPendidikan {
	return &jenjangPendidikanRepository{
		orm: orm,
	}
}
