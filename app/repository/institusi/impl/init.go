package impl

import (
	"github.com/beego/beego/v2/adapter/orm"
	institusirp "project_we/app/repository/institusi"
)

type institusiRepository struct {
	orm orm.Ormer
}

// NewInstitusiRepository ...
func NewInstitusiRepository(orm orm.Ormer) institusirp.IInstitusi {
	return &institusiRepository{
		orm: orm,
	}
}
