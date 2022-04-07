package impl

import (
	programStudirp "project_we/app/repository/programstudi"

	"github.com/beego/beego/v2/adapter/orm"
)

type programStudiRepository struct {
	orm orm.Ormer
}

// NewProgramStudiRepository ...
func NewProgramStudiRepository(orm orm.Ormer) programStudirp.IProgramStudi {
	return &programStudiRepository{
		orm: orm,
	}
}
