package dosen

import (
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/model"
)

type IDosen interface {
	SelectDosenByPersonalID(ctx *context.Context, personalID int64) (res model.Dosen, err error)
}
