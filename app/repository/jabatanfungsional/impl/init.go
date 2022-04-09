package impl

import (
	"github.com/beego/beego/v2/adapter/orm"

	jabatanfungsionalrp "project_we/app/repository/jabatanfungsional"
)

type jabatanFungsionalRepository struct {
	orm orm.Ormer
}

// NewJabatanFungsionalRepository ...
func NewJabatanFungsionalRepository(orm orm.Ormer) jabatanfungsionalrp.IJabatanFungsional {
	return &jabatanFungsionalRepository{
		orm: orm,
	}
}
